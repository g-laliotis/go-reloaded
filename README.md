# 🚀 go-reloaded

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue?logo=go)](https://go.dev)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-✓_passing-success)](#)
[![Made with ❤️ in Go](https://img.shields.io/badge/made%20with-%E2%9D%A4%20in%20Go-00ADD8?logo=go)](#)
[![Docs](https://img.shields.io/badge/docs-agents.md-blue?logo=readme)](agents.md)
[![CLI Usage](https://img.shields.io/badge/CLI-Make_Help-orange?logo=gnu-bash)](#makefile-commands)

> 📘 See [**agents.md**](agents.md) for detailed technical specification of all pipeline agents.  
> 💻 Run `make help` for a list of available CLI commands.

---

## 🧩 Overview

**go-reloaded** is a command-line text transformation tool written in Go using a **pipeline architecture**.

It reads an input text file, applies a sequence of smart transformations ("agents"), and writes the corrected text to an output file.

This project demonstrates clean modular Go design, testability, and readable rule-based text processing.

---

## 🏗️ Architecture

Each transformation is an **agent** — an independent module implementing:

```go
type Agent interface {
    Name() string
    Process(input string) string
}
```

The agents form a pipeline that processes text in order:
```
+-------------------+     +------------------+     +---------------------+     +----------------+
|   HexBinAgent     | --> |  CaseConvAgent   | --> |  PunctuationAgent   | --> |  ArticleAgent  |
| N (hex/bin) → dec |     | (up/low/cap[,N]) |     | spaces, ..., !?,' ' |     | a/A → an/An    |
+-------------------+     +------------------+     +---------------------+     +----------------+
```

This design keeps each rule isolated, testable, and auditable.

---

## ⚙️ Installation

**Clone the repository**
```bash
git clone https://github.com/g-laliotis/go-reloaded.git
cd go-reloaded
```

**Install Go (v1.22 or higher)**
```bash
go version
# should print go version go1.22+ ...
```

**Download dependencies**
```bash
go mod tidy
```

---

## 🚦 Usage

**Basic command**
```bash
go run ./cmd/go-reloaded <input_file> <output_file>
```

**Example:**
```bash
go run ./cmd/go-reloaded testdata/sample.txt result.txt
cat result.txt
```

**Example transformation:**

*Input (testdata/punctuation.txt)*
```
Punctuation tests are ... kinda boring ,what do you think ?
As Elton John said: ' I am the most well-known homosexual in the world '
I am exactly how they describe me: ' awesome '
BAMM !!  Are you serious !?  yes...right now
```

*Output*
```
Punctuation tests are... kinda boring, what do you think?
As Elton John said: 'I am the most well-known homosexual in the world'
I am exactly how they describe me: 'awesome'
BAMM!! Are you serious!? yes... right now
```

---

## 🧩 Agents Summary

| Agent | Purpose | Example |
|-------|----------|----------|
| 🧮 **HexBinAgent** | Converts numbers marked `(hex)` or `(bin)` to decimal | `1E (hex)` → `30` |
| 🔠 **CaseConvAgent** | Changes case with `(up)`, `(low)`, `(cap[,N])` | `this is so exciting (up,2)` → `this is SO EXCITING` |
| ✍️ **PunctuationAgent** | Fixes punctuation spacing & quotes | `,what ?!` → `, what?!` |
| 📰 **ArticleAgent** | Changes `a/A` → `an/An` before vowels or 'h' | `a apple` → `an apple` |

Full agent descriptions and internal specs: [`agents.md`](agents.md).

---

## 🧪 Testing

The project includes comprehensive unit, integration, benchmark, and edge case tests.

**Run all tests:**
```bash
make test
# or
go test ./...
```

**Run comprehensive test cases:**
```bash
go test -v -run TestCases
```

**Force tests to re-run (ignore cache):**
```bash
go test ./... -count=1
```

**Show coverage:**
```bash
go test ./... -cover
```

**Test Types:**
- **Unit tests**: `internal/transformations/agents_test.go`
- **Integration tests**: `main_test.go`, `edge_cases_test.go`
- **Comprehensive cases**: `testcases_test.go` (12 edge cases)
- **Performance tests**: `benchmark_test.go`

---

## 🧰 Makefile Commands

| Command | Description |
|---------|-------------|
| `make help` | Show help menu with all commands |
| `make run` | Run pipeline on sample data (testdata/cases.txt) |
| `make build` | Build binary into bin/go-reloaded |
| `make test` | Run all tests |
| `make fmt` | Format code |
| `make vet` | Run static analysis |
| `make clean` | Remove build artifacts |
| `make clear-cache` | (Optional) Wipe Go build/test caches |

Run the help menu any time:
```bash
make help
```

---

## 🧱 Project Structure

```
go-reloaded/
├── .github/
│   └── workflows/
│       └── pages.yml       # GitHub Pages deployment
├── cmd/
│   └── go-reloaded/
│       └── main.go         # CLI entrypoint
├── internal/
│   └── transformations/
│       ├── agents_test.go  # unit tests
│       ├── article.go      # article correction agent
│       ├── caseconv.go     # case conversion agent
│       ├── hexbin.go       # hex/binary conversion agent
│       ├── pipeline.go     # pipeline orchestration
│       ├── punctuation.go  # punctuation fixing agent
│       └── utils.go        # shared utilities
├── testcases/
│   ├── case*_input.txt     # comprehensive test inputs
│   └── case*_expected_output.txt # expected outputs
├── testdata/
│   ├── cases.txt           # basic test cases
│   ├── punctuation.txt     # punctuation examples
│   └── sample.txt          # sample input
├── .gitignore
├── LICENSE
├── Makefile                # build automation
├── README.md               # project documentation
├── agents.md               # technical specification
├── benchmark_test.go       # performance tests
├── edge_cases_test.go      # edge case tests
├── go.mod
├── main_test.go            # integration tests
└── testcases_test.go       # comprehensive test runner
```

---

## 🧑💻 Contributing

Pull requests and audits are welcome.

Please follow these guidelines:
- Run `make fmt` and `make vet` before committing
- Add or update tests for any changes
- Keep agents single-purpose; don't mix rules in one agent
- Update `agents.md` if your change affects behavior

---

## 🧾 License

Distributed under the MIT License.  
See [LICENSE](LICENSE) for details.

---

## ❤️ Acknowledgments

This project is part of a Go learning exercise focused on:
- File system I/O
- Text manipulation and regex
- Modular design and testing
- Pipeline architecture