# 🚀 go-reloaded

[![Go Version](https://img.shields.io/badge/Go-1.22+-blue?logo=go)](https://go.dev)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](#)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/tests-✓_passing-success)](#)
[![Made with ❤️ in Go](https://img.shields.io/badge/made%20with-%E2%9D%A4%20in%20Go-00ADD8?logo=go)](#)
[![Docs](https://img.shields.io/badge/docs-agents.md-blue?logo=readme)](agents.md)
[![CLI Usage](https://img.shields.io/badge/CLI-Make_Help-orange?logo=gnu-bash)](#makefile-commands)

> 📘 See [**agents.md**](agents.md) for a detailed technical specification of all pipeline agents.  
> 💻 Run `make help` for a list of available CLI commands.
> 🧩 Read the [**Technical Specification (agents.md)**](agents.md) for internal logic and agent details.


---

## 🧩 Overview

**go-reloaded** is a command-line text transformation tool written in Go using a **pipeline architecture**.

It reads an input text file, applies a sequence of smart transformations (“agents”), and writes the corrected text to an output file.

This project was designed for educational auditing: it demonstrates clean modular Go design, testability, and readable rule-based text processing.

---

## 🏗️ Architecture

Each transformation is an **agent** — an independent module implementing:

```go
type Agent interface {
    Name() string
    Process(input string) string
}
The agents form a pipeline that processes text in order:
+-------------------+     +------------------+     +---------------------+     +----------------+
|   HexBinAgent     | --> |  CaseConvAgent   | --> |  PunctuationAgent   | --> |  ArticleAgent  |
| N (hex/bin) → dec |     | (up/low/cap[,N]) |     | spaces, ..., !?,' ' |     | a/A → an/An    |
+-------------------+     +------------------+     +---------------------+     +----------------+
This design keeps each rule isolated, testable, and auditable.
See detailed specs in agents.md.

⚙️ Installation
Clone the repository
git clone https://github.com/<yourusername>/go-reloaded.git
cd go-reloaded
Install Go (v1.22 or higher)
go version
# should print go version go1.22+ ...
Download dependencies
go mod tidy
🚦 Usage
🖥️ Basic command
go run ./cmd/go-reloaded <input_file> <output_file>
Example:
go run ./cmd/go-reloaded testdata/punctuation.txt result.txt
cat result.txt
🧠 Example
Input (testdata/punctuation.txt)
Punctuation tests are ... kinda boring ,what do you think ?
As Elton John said: ' I am the most well-known homosexual in the world '
I am exactly how they describe me: ' awesome '
BAMM !!  Are you serious !?  yes...right now
Output
Punctuation tests are... kinda boring, what do you think?
As Elton John said: 'I am the most well-known homosexual in the world'
I am exactly how they describe me: 'awesome'
BAMM!! Are you serious!? yes... right now

## 🧩 Agents Summary

| Agent | Purpose | Example |
|-------|----------|----------|
| 🧮 **HexBinAgent** | Converts numbers marked `(hex)` or `(bin)` to decimal | `1E (hex)` → `30` |
| 🔠 **CaseConvAgent** | Changes case with `(up)`, `(low)`, `(cap[,N])` | `this is so exciting (up,2)` → `this is SO EXCITING` |
| ✍️ **PunctuationAgent** | Fixes punctuation spacing & quotes | `,what ?!` → `, what?!` |
| 📰 **ArticleAgent** | Changes `a/A` → `an/An` before vowels or ‘h’ | `a apple` → `an apple` |

Full agent descriptions and internal specs: [`agents.md`](agents.md).

---

## 📄 Understanding `agents.md` (Technical Specification)

The file [`agents.md`](agents.md) serves as the **blueprint** of the project — a detailed technical and functional guide to every transformation stage (*agent*) in the pipeline.

It is not just documentation — it’s a **living specification** that defines what the code *should* do.  
This makes the project easy to **audit, test, and extend**.

### 🧭 Why it matters

| Audience | Purpose |
|-----------|----------|
| 🧑‍🏫 **Auditors / Reviewers** | Verify that the tool behaves as described in each agent’s section — without reading Go code. |
| 👩‍💻 **Developers / Contributors** | Understand or modify a specific agent using its description, prompt, and examples as a contract. |
| 🧪 **Testers** | Use the example input/output pairs to design test cases for each agent. |
| 🧠 **You (the Maintainer)** | Quickly recall how and why each rule exists, and ensure consistency when updating logic. |

---

### 🧩 What It Contains

Each agent section in `agents.md` includes:
- **Name & File** — where it’s implemented  
- **Order in Pipeline** — so you know when it runs  
- **Responsibility** — what transformation it performs  
- **Prompt (Spec)** — plain English rule that defines its job  
- **Examples** — input/output illustrations for clarity  
- **Notes** — special cases or assumptions  

Example snippet from `agents.md`:
```markdown
## 📰 ArticleAgent
**File:** `internal/transformations/article.go`  
**Purpose:** Replace `a/A` with `an/An` when the next word begins with a vowel or 'h'.

### Prompt
> If the article is `a` (or `A`) and the next token’s first letter is in `[a e i o u h]`, change it to `an` (or `An`).

### Examples
| Input | Output |
|--------|---------|
| a apple | an apple |
| A honest person | An honest person |

🧪 Testing
The project includes both integration and unit tests.
Run all tests:

make test
# or
go test ./...
Force tests to re-run (ignore cache):
go test ./... -count=1
Show coverage:
go test ./... -cover
🧰 Makefile Commands
Command	Description
make help	Show help menu with all commands
make run	Run pipeline on sample data (testdata/cases.txt)
make build	Build binary into bin/go-reloaded
make test	Run all tests
make fmt	Format code
make vet	Run static analysis
make clean	Remove build artifacts
make clear-cache	(Optional) Wipe Go build/test caches
Run the help menu any time:
make help
🧹 Clearing Cache (optional)
To clear Go’s internal caches manually:
go clean -cache -modcache -testcache
Or use the Makefile shortcut:
make clear-cache
🧱 Project Structure
go-reloaded/
├── cmd/
│   └── go-reloaded/        # main CLI entrypoint
│       └── main.go
│
├── internal/
│   └── transformations/    # pipeline & agents
│       ├── pipeline.go
│       ├── hexbin.go
│       ├── caseconv.go
│       ├── punctuation.go
│       ├── article.go
│       └── *_test.go
│
├── testdata/               # sample test files
│   ├── sample.txt
│   ├── punctuation.txt
│   └── cases.txt
│
├── docs/                   # GitHub Pages site
│   └── index.html
│
├── agents.md               # detailed agent specifications
├── main_test.go            # integration tests
├── Makefile                # development commands
├── go.mod / go.sum
├── README.md               # you're here :)
└── LICENSE
🧑‍💻 Contributing
Pull requests and audits are welcome.
Please follow these guidelines:
Run make fmt and make vet before committing.
Add or update tests for any changes.
Keep agents single-purpose; don’t mix rules in one agent.
Update agents.md if your change affects behavior.
🧾 License
Distributed under the MIT License.
See LICENSE for details.
❤️ Acknowledgments
This project is part of a Go learning exercise focused on:
file system I/O
text manipulation and regex
modular design and testing
pipeline architecture
