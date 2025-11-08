# 🚀 go-reloaded

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue?logo=go)](https://go.dev)
[![Docker](https://img.shields.io/badge/Docker-supported-blue?logo=docker)](https://www.docker.com/)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-✓_passing-success)](#)
[![Made with ❤️ in Go](https://img.shields.io/badge/made%20with-%E2%9D%A4%20in%20Go-00ADD8?logo=go)](#)
[![Docs](https://img.shields.io/badge/docs-AGENTS.md-blue?logo=readme)](AGENTS.md)
[![CLI Usage](https://img.shields.io/badge/CLI-Make_Help-orange?logo=gnu-bash)](#makefile-commands)
[![Security](https://img.shields.io/badge/security-policy-red?logo=shield)](SECURITY.md)
[![Code of Conduct](https://img.shields.io/badge/code%20of-conduct-ff69b4?logo=handshake)](CODE_OF_CONDUCT.md)
[![Contributing](https://img.shields.io/badge/contributing-guidelines-brightgreen?logo=github)](CONTRIBUTING.md)
[![GitHub Pages](https://img.shields.io/badge/GitHub%20Pages-Live%20Demo-brightgreen?logo=github)](https://g-laliotis.github.io/go-reloaded/)

> 📘 See [**AGENTS.md**](AGENTS.md) for detailed technical specification of all pipeline agents.  
> 💻 Run `make help` for a list of available CLI commands.

---

## 🧩 TL;DR Overview

**go-reloaded** is a command-line text transformation tool written in Go using a **pipeline architecture**.

It reads an input text file, applies a sequence of smart transformations ("agents"), and writes the corrected text to an output file.

This project demonstrates clean modular Go design, testability, and readable rule-based text processing.

**Professional Features:**
- 🔒 [**Security Policy**](SECURITY.md) – clear vulnerability reporting process
- 🤝 [**Code of Conduct**](CODE_OF_CONDUCT.md) – community & collaboration rules
- 📋 **Issue Templates** – structured bug/feature reports
- 🔄 **CI/CD** – automated tests on Go 1.22 & 1.23
- 📚 [**Comprehensive Docs**](AGENTS.md) – agents spec, [contribution guide](CONTRIBUTING.md), [presentation guide](PRESENTATION_GUIDE.md)

**Ideal for:** developers who want a reference implementation of a modular Go CLI with full tests, CI, Docker, and security best practices.

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

**Check version**
```bash
go run ./cmd/go-reloaded --version
# go-reloaded v1.1.0 (built with go1.22+ for darwin/amd64)
```

---

## 🐳 Docker Usage

### For Non-Technical Users

**Step 1: Install Docker**
- Download Docker Desktop from https://www.docker.com/products/docker-desktop/
- Install and start Docker Desktop

**Step 2: Create a test file**
- Create a text file called `input.txt` on your Desktop
- Add some test text like: `I have 1E (hex) apples and a orange (up)!`

**Step 3: Open Terminal/Command Prompt**
- **Windows**: Press `Win + R`, type `cmd`, press Enter
- **Mac**: Press `Cmd + Space`, type `terminal`, press Enter

**Step 4: Navigate to your Desktop**
```bash
# Windows
cd Desktop

# Mac
cd Desktop
```

**Step 5: Run the transformation**
```bash
# Windows
docker run --rm -v "%cd%:/data" ghcr.io/g-laliotis/go-reloaded:latest /data/input.txt /data/output.txt

# Mac/Linux
docker run --rm -v "$(pwd):/data" ghcr.io/g-laliotis/go-reloaded:latest /data/input.txt /data/output.txt
```

**Step 6: Check the result**
- Look for `output.txt` on your Desktop
- Open it to see the transformed text

### For Developers

**Using pre-built image (recommended):**
```bash
# Pull and run directly from GitHub Container Registry
docker run --rm -v "$(pwd):/data" ghcr.io/g-laliotis/go-reloaded:latest /data/input.txt /data/output.txt

# Example with sample file
docker run --rm -v "$(pwd):/data" ghcr.io/g-laliotis/go-reloaded:latest /data/testdata/samples/sample.txt /data/result.txt
```

**Building locally (for development):**
```bash
make docker-build
make docker-run INPUT=testdata/samples/sample.txt OUTPUT=result.txt

# Or manually
docker build -t go-reloaded .
docker run --rm -v "$(pwd):/data" go-reloaded /data/input.txt /data/output.txt
```

---

## 🚦 Usage

**Basic command**
```bash
go run ./cmd/go-reloaded <input_file> <output_file>
```

**Example:**
```bash
go run ./cmd/go-reloaded testdata/samples/sample.txt result.txt
cat result.txt
```

**Example transformation:**

*Input (testdata/samples/punctuation.txt)*
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

Full agent descriptions and internal specs: [`AGENTS.md`](AGENTS.md).

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
make testcases
# or
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
- **Comprehensive cases**: `testcases_test.go` (24 comprehensive cases)
- **Performance tests**: `benchmark_test.go`

---

## 🧰 Makefile Commands

| Command | Description |
|---------|-------------|
| `make help` | Show help menu with all commands |
| `make run` | Run pipeline on sample data (testdata/samples/) |
| `make build` | Build binary into bin/go-reloaded |
| `make test` | Run all tests |
| `make testcases` | Run comprehensive test cases (24 cases) |
| `make bench` | Run benchmark tests with memory stats |
| `make fmt` | Format code |
| `make vet` | Run static analysis |
| `make clean` | Remove build artifacts |
| `make clear-cache` | Clear Go build and test cache |
| `make docker-build` | Build Docker image locally |
| `make docker-run` | Run with Docker (specify INPUT/OUTPUT) |

**Docker Registry:**
- **GHCR**: `ghcr.io/g-laliotis/go-reloaded:latest` (recommended)
- **Local build**: Use `make docker-build` for development

Run the help menu any time:
```bash
make help
```

---

## 🧱 Project Structure

```
go-reloaded/
├── .github/
│   ├── ISSUE_TEMPLATE/
│   │   ├── bug_report.yml      # Bug report form
│   │   └── feature_request.yml # Feature request form
│   ├── workflows/
│   │   ├── docker.yml          # Docker build and publish to GHCR
│   │   ├── pages.yml           # GitHub Pages deployment
│   │   └── test.yml            # CI/CD testing workflow
│   └── PULL_REQUEST_TEMPLATE.md # Pull request template
├── cmd/
│   └── go-reloaded/
│       └── main.go         # CLI entrypoint with version support
├── docs/
│   ├── .nojekyll           # bypass Jekyll processing
│   ├── index.html          # GitHub Pages site
│   └── README.md           # docs folder readme
├── internal/
│   ├── transformations/
│   │   ├── agents_test.go  # unit tests
│   │   ├── article.go      # article correction agent
│   │   ├── caseconv.go     # case conversion agent
│   │   ├── hexbin.go       # hex/binary conversion agent
│   │   ├── pipeline.go     # pipeline orchestration
│   │   ├── punctuation.go  # punctuation fixing agent
│   │   └── utils.go        # shared utilities
│   └── version/
│       └── version.go      # version and build info
├── testdata/
│   ├── comprehensive/      # comprehensive test cases (24 cases)
│   │   ├── case*_input.txt     # test inputs
│   │   └── case*_expected_output.txt # expected outputs
│   └── samples/            # sample files for demos
│       ├── cases.txt       # basic test cases
│       ├── punctuation.txt # punctuation examples
│       └── sample.txt      # sample input
├── .dockerignore           # Docker build exclusions
├── .gitignore
├── AGENTS.md               # technical specification
├── benchmark_test.go       # performance tests
├── CHANGELOG.md            # version history
├── CODE_OF_CONDUCT.md      # community guidelines
├── CONTRIBUTING.md         # contribution guidelines
├── Dockerfile              # Docker containerization
├── edge_cases_test.go      # edge case tests
├── go.mod
├── LICENSE
├── main_test.go            # integration tests
├── Makefile                # build automation
├── PRESENTATION_GUIDE.md   # stakeholder briefing document
├── README.md               # project documentation
├── SECURITY.md             # security policy
└── testcases_test.go       # comprehensive test runner
```

---

## 🧑💻 Contributing

Pull requests and audits are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

**Quick start:**
- Run `make fmt` and `make vet` before committing
- Add or update tests for any changes
- Keep agents single-purpose; don't mix rules in one agent
- Update `AGENTS.md` if your change affects behavior

**Version history:** See [CHANGELOG.md](CHANGELOG.md) for release notes and changes.

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