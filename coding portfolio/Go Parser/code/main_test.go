// Christopher De Vault
// CSC 372
// Project test code
// main_test.go
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// --- Helper Functions --------------------------------

// equalTokens compares two slices of Token for equality.
func equalTokens(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Type != b[i].Type || a[i].Value != b[i].Value {
			return false
		}
	}
	return true
}

// safeEval runs parsing + evaluation, recovers from panic, returns result and panicked flag.
func safeEval(input string) (result int, panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	lex := NewLexer(input)
	parser := NewParser(lex)
	tree := parser.Parse()
	result = tree.Eval()
	return
}

// --- Unit Tests for Lexer --------------------------------

func TestLexer(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Token
	}{
		{"single digit", "7", []Token{{NUMBER, "7"}, {EOF, ""}}},
		{"simple plus", "3+4", []Token{
			{NUMBER, "3"}, {PLUS, "+"}, {NUMBER, "4"}, {EOF, ""},
		}},
		{"with whitespace", " 12 \t -5 ",
			[]Token{{NUMBER, "12"}, {MINUS, "-"}, {NUMBER, "5"}, {EOF, ""}},
		},
		{"all operators", "1+2-3*4/5%6", []Token{
			{NUMBER, "1"}, {PLUS, "+"}, {NUMBER, "2"},
			{MINUS, "-"}, {NUMBER, "3"}, {MUL, "*"},
			{NUMBER, "4"}, {DIV, "/"}, {NUMBER, "5"},
			{MOD, "%"}, {NUMBER, "6"}, {EOF, ""},
		}},
		{"parentheses", "(1+(2*3))", []Token{
			{LPAREN, "("}, {NUMBER, "1"}, {PLUS, "+"},
			{LPAREN, "("}, {NUMBER, "2"}, {MUL, "*"},
			{NUMBER, "3"}, {RPAREN, ")"}, {RPAREN, ")"},
			{EOF, ""},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := NewLexer(tt.input)
			var got []Token
			for tok := lex.GetNextToken(); tok.Type != EOF; tok = lex.GetNextToken() {
				got = append(got, tok)
			}
			got = append(got, Token{Type: EOF, Value: ""})
			if !equalTokens(got, tt.want) {
				t.Errorf("Lexer(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// --- Unit Tests for Parser + Evaluator -------------------

func TestParserEval(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"single add", "1+2", 3},
		{"mixed ops", "2*3+4", 10},
		{"parentheses", "2*(3+4)", 14},
		{"unary minus", "-1+2", 1},
		{"modulus", "5%2", 1},
		{"complex", "(1+2)*3-4/2", 7},
		{"nested unary", "--5", 5},
		{"whitespace", " 10 - ( 3 * 2 ) ", 4},
		{"double unary plus", "1++2", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, panicked := safeEval(tt.input)
			if panicked {
				t.Fatalf("Eval(%q) panicked", tt.input)
			}
			if got != tt.want {
				t.Errorf("Eval(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestInvalidExpressions(t *testing.T) {
	invalid := []string{
		"", "+", "()", "(1+2", "a+1", "1/0",
	}
	for _, input := range invalid {
		t.Run(input, func(t *testing.T) {
			_, panicked := safeEval(input)
			if !panicked {
				t.Errorf("Expected panic for %q, got none", input)
			}
		})
	}
}

// --- Unit Test for Concurrency Feature -------------------

func TestEvaluateAsync(t *testing.T) {
	// Parse a known expression
	lex := NewLexer("7*6")
	parser := NewParser(lex)
	node := parser.Parse()

	ch := make(chan int)
	go evaluateAsync(node, ch) // run in a goroutine

	result := <-ch
	if result != 42 {
		t.Errorf("evaluateAsync: got %d, want 42", result)
	}
}

// --- Integration Test for REPL ----------------------------

func TestMainIntegration(t *testing.T) {
	// Backup original stdin/stdout
	oldStdin, oldStdout := os.Stdin, os.Stdout
	defer func() {
		os.Stdin, os.Stdout = oldStdin, oldStdout
	}()

	// Create pipes for stdin and stdout
	rIn, wIn, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rOut, wOut, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// Write simulated user input
	go func() {
		_, _ = io.WriteString(wIn, "1+2\nexit\n")
		wIn.Close()
	}()

	os.Stdin = rIn
	os.Stdout = wOut

	// Run the REPL (main)
	main()

	// Close writer to allow reader to finish
	wOut.Close()

	// Capture and verify output
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, rOut)
	out := buf.String()

	if !strings.Contains(out, "= 3") {
		t.Errorf("Integration output missing '= 3', got:\n%s", out)
	}
}

// --- Benchmarks -------------------------------------------

func BenchmarkEvalSimple(b *testing.B) {
	expr := "1+2*3-4/2"
	for i := 0; i < b.N; i++ {
		lex := NewLexer(expr)
		parser := NewParser(lex)
		tree := parser.Parse()
		_ = tree.Eval()
	}
}

func BenchmarkEvalComplex(b *testing.B) {
	expr := "((1+2)*3 - (4%3 + 5/2)) * (6-7+8) % 9"
	for i := 0; i < b.N; i++ {
		lex := NewLexer(expr)
		parser := NewParser(lex)
		tree := parser.Parse()
		_ = tree.Eval()
	}
}

// --- Fuzz Test -------------------------------

func FuzzEval(f *testing.F) {
	f.Add("1+1")
	f.Add("2*(3+4)")
	f.Fuzz(func(t *testing.T, input string) {
		defer func() { recover() }()
		lex := NewLexer(input)
		parser := NewParser(lex)
		tree := parser.Parse()
		_ = tree.Eval()
	})
}

// TestVeryLargeNumber ensures that extremely large integer literals
// do not panic and default to 0 (strconv.Atoi error ignored).
func TestVeryLargeNumber(t *testing.T) {
	input := "999999999999999999999999999999999999"
	got, panicked := safeEval(input)
	if panicked {
		t.Fatalf("Expected no panic for oversized number %q, got panic", input)
	}
	if got != 0 {
		t.Errorf("Expected result 0 for oversized input %q, got %d", input, got)
	}
}

// TestLeadingTrailingGarbage ensures inputs with extraneous characters
// before or after a valid expression panic.
func TestLeadingTrailingGarbage(t *testing.T) {
	invalid := []string{
		"foo1+2",
		"1+2bar",
		"xyz 3+4",
		"3+4 xyz",
	}
	for _, input := range invalid {
		t.Run(input, func(t *testing.T) {
			_, panicked := safeEval(input)
			if !panicked {
				t.Errorf("Expected panic for input with garbage %q, got none", input)
			}
		})
	}
}

// TestStressLargeExpression builds a long expression (1000 terms)
// and ensures the evaluator can handle it correctly.
func TestStressLargeExpression(t *testing.T) {
	const n = 1000
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			sb.WriteByte('+')
		}
		sb.WriteString("1")
	}
	expr := sb.String()

	got, panicked := safeEval(expr)
	if panicked {
		t.Fatalf("Stress test panicked on expression of length %d", len(expr))
	}
	if got != n {
		t.Errorf("Stress test result = %d, want %d", got, n)
	}
}
