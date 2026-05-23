# Go Arithmetic Expression Parser & Evaluator

An interactive **lexer → parser → AST → evaluator** pipeline written in Go, with a REPL and goroutine-based evaluation. Built as a from-scratch demonstration of recursive-descent parsing and interpreter design.

---

## Overview

This project implements a small **arithmetic expression language** in a single Go module. Given input such as `3 + 4 * (2 - 1)`, the program:

1. **Tokenizes** the string (lexer)
2. **Parses** tokens into an abstract syntax tree (recursive-descent parser)
3. **Evaluates** the AST to an integer result
4. Runs evaluation in a **goroutine** and returns the result through a channel (concurrency requirement)

A read–eval–print loop (**REPL**) accepts expressions from stdin until you type `exit` or submit an empty line.

---

## Features

| Component | Description |
|-----------|-------------|
| **Lexer** | Emits tokens for `+`, `-`, `*`, `/`, `%`, `(`, `)`, integers, and `EOF`; skips whitespace |
| **Parser** | Recursive-descent grammar with correct operator precedence |
| **AST** | `NumberNode`, `UnaryNode`, `BinaryNode`; all implement `Eval() int` |
| **Operators** | `+`, `-`, `*`, `/`, `%`; unary `+` / `-`; parentheses |
| **REPL** | Interactive `>>>` prompt with panic recovery on bad input |
| **Concurrency** | `evaluateAsync` runs `Eval()` in a goroutine; main receives via `chan int` |
| **Tests** | Lexer, parser/evaluator, invalid input, async eval, REPL integration, benchmarks, fuzz |

---

## Grammar

```text
expr   := term (("+" | "-") term)*
term   := unary (("*" | "/" | "%") unary)*
unary  := ("+" | "-") unary | factor
factor := NUMBER | "(" expr ")"
```

Precedence (low → high): addition/subtraction → multiplication/division/modulus → unary → parentheses and literals.

---

## Requirements

- **Go 1.24+** (see `code/go.mod`; module path `github.com/cjdevault/project`)

---

## Build & run

From the **`code/`** directory (canonical source with `go.mod`):

```bash
cd code
go run .
```

Or build a binary:

```bash
cd code
go build -o gocalc .
./gocalc
```

### REPL usage

```text
Go Arithmetic Evaluator (with goroutine!)
Enter an expression, or 'exit' to quit.
>>> 3 + 4 * (2 - 1)
= 7

>>> 10 - (3 * 2)
= 4

>>> exit
Goodbye!
```

**Examples**

| Input | Result |
|-------|--------|
| `1+2` | 3 |
| `2*(3+4)` | 14 |
| `-1+2` | 1 |
| `5%2` | 1 |
| `(1+2)*3-4/2` | 7 |
| `--5` | 5 |

Invalid input (unknown characters, malformed parentheses, trailing garbage) triggers a recovered panic and prints `Error: ...`.

---

## Testing

```bash
cd code
go test -v ./...
```

With coverage:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

Test coverage includes:

- Lexer token streams (operators, whitespace, parentheses)
- Parser + evaluator correctness and precedence
- Invalid expressions (empty, `()`, `a+1`, etc.)
- Goroutine evaluation (`evaluateAsync`)
- REPL integration (piped stdin/stdout)
- Benchmarks and fuzz harness
- Edge cases (oversized integers default to 0, stress expressions)

---

## Project structure

```text
Go Parser/
├── README.md                          # This file
│
├── code/                              # Primary source (use this to build & test)
│   ├── go.mod                         # Go module definition
│   ├── main.go                        # Lexer, parser, AST, REPL, goroutine eval
│   ├── main_test.go                   # Unit, integration, benchmark, fuzz tests
│   └── coverage.out                   # Test coverage output (generated)
│
├── GO_interpreter_project/            # Submission copy of source (same as code/)
│   ├── main.go
│   └── main_test.go
├── GO_interpreter_project.zip         # Archived submission bundle
│
├── coverage.out                       # Coverage artifact (root copy)
└── Project_Presentation.mp4         # Recorded project presentation
```

**Note:** Prefer **`code/`** for development (`go.mod` lives there). `GO_interpreter_project/` mirrors the same `main.go` / `main_test.go` for course submission.

---

## Architecture

```text
  Input string
       │
       ▼
    Lexer ──► []Token
       │
       ▼
    Parser ──► AST (Node)
       │
       ▼
  evaluateAsync (goroutine)
       │
       ▼
    chan int ──► REPL prints result
```

**Lexer** — character-by-character scan; `GetNextToken()` returns the next `Token` until `EOF`.

**Parser** — `Parser` holds current token; `eat()` advances; `expr()` / `term()` / `unary()` / `factor()` build the tree.

**Evaluator** — tree walk via `Eval()` on each node type; integer division and modulus follow Go semantics.

---

## Design notes

- **Single file:** All components live in `main.go` for clarity in a course interpreter project.
- **Error handling:** Lexer/parser use `panic` + `recover` in the REPL rather than explicit error returns.
- **Large literals:** `strconv.Atoi` failures map to `0` instead of crashing.
- **Goroutine:** Evaluation is cheap; the goroutine demonstrates channels and concurrency as required by the assignment.

---

## License

Coursework submission. All rights reserved by the author unless your course specifies otherwise.
