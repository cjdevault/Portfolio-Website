// Chrisopther De Vault
// CSC 372
// Project Source Code
// main.go
//
// A simple arithmetic expression evaluator in Go.
// This program demonstrates how to:
//  1. Tokenize an input string into meaningful symbols (lexer).
//  2. Parse tokens using a recursive-descent parser according to grammar rules.
//  3. Build an Abstract Syntax Tree (AST) of Nodes.
//  4. Evaluate the AST to compute the numeric result.
//  5. Provide a simple Read–Eval–Print Loop (REPL) for interactive use.
//  6. Utilize a Go-specific goroutine to showcase a concurrency feature.
//
// Usage:
//   go run main.go
// Then type expressions like:
//   3 + 4 * (2 - 1)
// Type "exit" or press Enter on an empty line to quit.

package main

import (
	"bufio"   // for buffered input from stdin
	"fmt"     // for formatted I/O
	"os"      // to access OS features like Stdin
	"strconv" // for converting string to integer
	"strings" // for trimming and manipulating input strings
	"unicode" // for checking character properties (e.g., IsDigit, IsSpace)
)

// --- Token definitions --------------------------------

// TokenType enumerates the different kinds of tokens our lexer can produce.
type TokenType int

const (
	EOF    TokenType = iota // End of input
	PLUS                    // '+'
	MINUS                   // '-'
	MUL                     // '*'
	DIV                     // '/'
	MOD                     // '%'
	LPAREN                  // '('
	RPAREN                  // ')'
	NUMBER                  // sequence of digits, e.g. "123"
)

// Token holds a type and the actual string value for NUMBER tokens.
type Token struct {
	Type  TokenType
	Value string
}

// --- Lexer -------------------------------------------

// Lexer reads the input text character by character and groups them into tokens.
type Lexer struct {
	text        string // the full input string
	pos         int    // current position in text
	currentChar rune   // character at current pos (0 if end)
}

// NewLexer initializes a Lexer for the given input string.
func NewLexer(input string) *Lexer {
	l := &Lexer{text: input, pos: 0}
	if len(input) > 0 {
		l.currentChar = rune(input[0])
	}
	return l
}

// advance moves the lexer one character forward.
func (l *Lexer) advance() {
	l.pos++
	if l.pos >= len(l.text) {
		l.currentChar = 0 // using 0 to signal end of input
	} else {
		l.currentChar = rune(l.text[l.pos])
	}
}

// skipWhitespace keeps advancing as long as currentChar is whitespace.
func (l *Lexer) skipWhitespace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) {
		l.advance()
	}
}

// number reads a sequence of digits and returns a NUMBER token.
func (l *Lexer) number() Token {
	var sb strings.Builder
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) {
		sb.WriteRune(l.currentChar)
		l.advance()
	}
	return Token{Type: NUMBER, Value: sb.String()}
}

// GetNextToken identifies the next token in the input and returns it.
func (l *Lexer) GetNextToken() Token {
	for l.currentChar != 0 {
		// Skip any whitespace
		if unicode.IsSpace(l.currentChar) {
			l.skipWhitespace()
			continue
		}
		// Single-character tokens
		switch l.currentChar {
		case '+':
			l.advance()
			return Token{Type: PLUS, Value: "+"}
		case '-':
			l.advance()
			return Token{Type: MINUS, Value: "-"}
		case '*':
			l.advance()
			return Token{Type: MUL, Value: "*"}
		case '/':
			l.advance()
			return Token{Type: DIV, Value: "/"}
		case '%':
			l.advance()
			return Token{Type: MOD, Value: "%"}
		case '(':
			l.advance()
			return Token{Type: LPAREN, Value: "("}
		case ')':
			l.advance()
			return Token{Type: RPAREN, Value: ")"}
		default:
			// If it's a digit, read a full number
			if unicode.IsDigit(l.currentChar) {
				return l.number()
			}
			// Unknown character -> error
			panic(fmt.Sprintf("Lexer error: unknown character %q", l.currentChar))
		}
	}
	// No more input: return EOF token
	return Token{Type: EOF, Value: ""}
}

// --- AST Nodes ---------------------------------------

// Node is the interface for all nodes in the Abstract Syntax Tree.
// Each node must implement Eval(), which computes its integer value.
type Node interface {
	Eval() int
}

// NumberNode represents a literal integer in the AST.
type NumberNode struct {
	Value int
}

// Eval simply returns the stored integer value.
func (n *NumberNode) Eval() int {
	return n.Value
}

// UnaryNode represents a unary operation (+ or -) applied to one operand.
type UnaryNode struct {
	Op   TokenType // PLUS or MINUS
	Node Node      // the operand
}

// Eval applies the unary operator to the operand's Eval result.
func (u *UnaryNode) Eval() int {
	v := u.Node.Eval()
	if u.Op == PLUS {
		return +v
	}
	return -v
}

// BinaryNode represents a binary operation (e.g., +, -, *, /, %) between two operands.
type BinaryNode struct {
	Left  Node      // left operand
	Op    TokenType // operator token
	Right Node      // right operand
}

// Eval computes Left Eval, Right Eval, then applies the operator.
func (b *BinaryNode) Eval() int {
	l := b.Left.Eval()
	r := b.Right.Eval()
	switch b.Op {
	case PLUS:
		return l + r
	case MINUS:
		return l - r
	case MUL:
		return l * r
	case DIV:
		return l / r
	case MOD:
		return l % r
	}
	return 0 // should not reach here
}

// --- Parser ------------------------------------------

// Parser orchestrates consuming tokens and building the AST following grammar.
//
// Grammar:
//
//	expr   := term (("+"|"-") term)*
//	term   := unary (("*"|"/"|"%") unary)*
//	unary  := ("+"|"-") unary | factor
//	factor := NUMBER | "(" expr ")"
type Parser struct {
	lexer        *Lexer
	currentToken Token
}

// NewParser constructs a Parser given a Lexer, and reads the first token.
func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.currentToken = l.GetNextToken()
	return p
}

// eat checks that the current token matches the expected type, then reads the next.
func (p *Parser) eat(t TokenType) {
	if p.currentToken.Type != t {
		panic(fmt.Sprintf("Parser error: expected %v, got %v", t, p.currentToken.Type))
	}
	p.currentToken = p.lexer.GetNextToken()
}

// factor handles NUMBER or parenthesized subexpression.
func (p *Parser) factor() Node {
	tok := p.currentToken
	if tok.Type == NUMBER {
		p.eat(NUMBER)
		v, err := strconv.Atoi(tok.Value)
		if err != nil {
			// on overflow or non‐numeric, treat as zero
			v = 0
		}
		return &NumberNode{Value: v}
	}
	if tok.Type == LPAREN {
		p.eat(LPAREN)
		node := p.expr() // parse sub-expression
		p.eat(RPAREN)
		return node
	}
	panic("Parser error: expected number or '('")
}

// unary handles unary plus/minus or delegates to factor.
func (p *Parser) unary() Node {
	tok := p.currentToken
	if tok.Type == PLUS || tok.Type == MINUS {
		p.eat(tok.Type)
		node := p.unary() // allow chained unary ops
		return &UnaryNode{Op: tok.Type, Node: node}
	}
	return p.factor()
}

// term handles multiplication, division, and modulus with left associativity.
func (p *Parser) term() Node {
	node := p.unary()
	for p.currentToken.Type == MUL || p.currentToken.Type == DIV || p.currentToken.Type == MOD {
		op := p.currentToken.Type
		p.eat(op)
		right := p.unary()
		node = &BinaryNode{Left: node, Op: op, Right: right}
	}
	return node
}

// expr handles addition and subtraction with left associativity.
func (p *Parser) expr() Node {
	node := p.term()
	for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
		op := p.currentToken.Type
		p.eat(op)
		right := p.term()
		node = &BinaryNode{Left: node, Op: op, Right: right}
	}
	return node
}

// Parse is a convenience method to start parsing at expr.
func (p *Parser) Parse() Node {
	return p.expr()
}

// evaluateAsync runs node.Eval() in a goroutine and sends the result on ch.
func evaluateAsync(node Node, ch chan int) {
	// This runs in a separate goroutine:
	ch <- node.Eval()
}

// --- Main / REPL ------------------------------------

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Go Arithmetic Evaluator (with goroutine!)")
	fmt.Println("Enter an expression, or 'exit' to quit.")

	for {
		fmt.Print(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Input error:", err)
			return
		}
		text := strings.TrimSpace(line)
		if text == "" || strings.EqualFold(text, "exit") {
			fmt.Println("Goodbye!")
			return
		}

		// Recover from any panics during parse/Eval
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
			}
		}()

		// Lex and parse
		lexer := NewLexer(text)
		parser := NewParser(lexer)
		tree := parser.Parse()

		// evaluate in a goroutine
		resultCh := make(chan int)       // create a channel for int
		go evaluateAsync(tree, resultCh) // start evaluation concurrently
		result := <-resultCh             // wait for the result

		// Print the result once the goroutine sends it
		fmt.Printf("= %d\n\n", result)
	}
}
