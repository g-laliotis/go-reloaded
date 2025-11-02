[← Back to README](README.md)

---

# 🧩 Agents Technical Specification & Reference

> 📘 This document describes every text transformation agent in **`go-reloaded`**.  
> Each agent is a self-contained pipeline stage, with its own rules, examples, and validation notes.

---

## 🧭 Pipeline Overview

```text
+-------------------+     +------------------+     +---------------------+     +----------------+
|   HexBinAgent     | --> |  CaseConvAgent   | --> |  PunctuationAgent   | --> |  ArticleAgent  |
| N (hex/bin) → dec |     | (up/low/cap[,N]) |     | spaces, ..., !?,' ' |     | a/A → an/An    |
+-------------------+     +------------------+     +---------------------+     +----------------+
```

**Legend:**
- 🟩 Input: raw text (line-preserving)
- 🟦 Output: processed text passed to the next agent
- 🧩 Order matters: numbers → case → punctuation/quotes → articles

All agents implement the shared interface:
```go
type Agent interface {
    Name() string
    Process(input string) string
}
```

📁 **Defined in:** `internal/transformations/`

---

## ⚙️ Design Principles

| Principle | Description |
|-----------|-------------|
| 🧱 **Single Responsibility** | Each agent performs one transformation. |
| 🔁 **Composability** | Agents can be reordered, extended, or replaced. |
| 🧪 **Testability** | Each agent has dedicated unit tests (`*_test.go`). |
| 🔍 **Transparency** | Examples double as test and audit references. |
| ⚖️ **Determinism** | Same input → same output; no randomness. |

---

## 🧮 1. HexBinAgent

**📄 File:** `internal/transformations/hexbin.go`  
**🔢 Order:** 1  
**🎯 Purpose:** Convert numbers followed by `(hex)` or `(bin)` to decimal.

### 🧠 Prompt (Spec)

Replace numbers followed by `(hex)` or `(bin)` with their decimal equivalents.  
Keep punctuation and spacing intact.

**Supported formats:**
- Hexadecimal: `[0-9A-Fa-f]+`
- Binary: `[01]+`

### 💬 Examples

| Input | Output |
|-------|--------|
| `1E (hex)` | `30` |
| `10 (bin)` | `2` |
| `Simply add 1E (hex) and 10 (bin).` | `Simply add 30 and 2.` |

### 📝 Notes

- Invalid formats (e.g., `G1 (hex)`) remain unchanged.
- Works line-by-line and preserves newlines.

---

## 🔠 2. CaseConvAgent

**📄 File:** `internal/transformations/caseconv.go`  
**🔢 Order:** 2  
**🎯 Purpose:** Apply uppercase, lowercase, or capitalization rules using `(up)`, `(low)`, `(cap)` markers.  
Supports counts like `(up, 3)`.

### 🧠 Prompt (Spec)

When encountering `(up)`, `(low)`, or `(cap)`, apply the transformation to the previous word(s).  
If followed by `(cmd, N)` → apply to the previous N words.  
Then remove the marker(s).

### 💬 Examples

| Input | Output |
|-------|--------|
| `Ready set go (up)` | `Ready set GO` |
| `this is so exciting (up, 2)` | `this is SO EXCITING` |
| `STOP SHOUTING (low)` | `STOP shouting` |
| `welcome to the brooklyn bridge (cap, 2)` | `welcome to the Brooklyn Bridge` |

### 📝 Notes

- `(cap)` capitalizes first letter, lowers others.
- Supports both compact `(cmd, N)` and split forms.
- Markers are stripped after processing.

---

## ✍️ 3. PunctuationAgent

**📄 File:** `internal/transformations/punctuation.go`  
**🔢 Order:** 3  
**🎯 Purpose:** Normalize punctuation spacing for `. , ! ? : ; ...` and fix quotes `' '`.

### 🧠 Prompt (Spec)

Correct spacing and punctuation to ensure grammatical consistency:
- `. , : ;` → no space before, one after
- `...` → compact form, one space after if followed by text
- `!?, ?!, !!` → compact form, one space after if followed by text
- `' text '` → `'text'` (remove inner spaces)

Preserve multi-punctuation groups and newlines.

### 💬 Examples

| Input | Output |
|-------|--------|
| `I was thinking ... You were right` | `I was thinking... You were right` |
| `over there ,and then BAMM !!` | `over there, and then BAMM!!` |
| `Are you serious !? yes...right now` | `Are you serious!? yes... right now` |
| `As Elton John said: ' awesome '` | `As Elton John said: 'awesome'` |

### 📝 Notes

- Converts `. . .` → `...`
- Preserves original newlines.
- Prevents spaces splitting `!?` or `!!`.

---

## 📰 4. ArticleAgent

**📄 File:** `internal/transformations/article.go`  
**🔢 Order:** 4  
**🎯 Purpose:** Replace `a` → `an` and `A` → `An` when next word begins with a vowel (`a, e, i, o, u`) or `h`.

### 🧠 Prompt (Spec)

If the article `a` (or `A`) precedes a word beginning with a vowel or `h`,  
replace it with `an` (or `An`).  
Ignore leading quotes or brackets before the word.

### 💬 Examples

| Input | Output |
|-------|--------|
| `a apple` | `an apple` |
| `A honest person` | `An honest person` |
| `a 'idea' a (hour) a [umbrella]` | `an 'idea' an (hour) an [umbrella]` |

### 📝 Notes

- Only affects alphabetic words.
- Skips numbers and symbols (`a 42` → unchanged).
- Preserves case (`A` → `An`).

---

## 🧪 Testing Reference

| Type | Files | Description |
|------|-------|-------------|
| 🔹 **Unit tests** | `internal/transformations/agents_test.go` | Individual agent testing |
| 🔹 **Integration tests** | `main_test.go`, `edge_cases_test.go` | End-to-end pipeline testing |
| 🔹 **Comprehensive cases** | `testcases_test.go` | 12 edge cases with tricky scenarios |
| 🔹 **Performance tests** | `benchmark_test.go` | Performance monitoring |
| 🔹 **Test data** | `testcases/case*_input.txt` | Input files for comprehensive testing |
| 🔹 **Expected outputs** | `testcases/case*_expected_output.txt` | Expected results for validation |

**Run Tests:**
```bash
make test                    # All tests
go test -v -run TestCases    # Comprehensive edge cases
go test ./... -count=1       # Force re-run
```

✅ All tests pass and confirm agents follow this specification exactly.

---

## 🧰 Implementation Notes

- All agents process text line-by-line (split on `\n`, then rejoined).
- The pipeline runs sequentially; output of one feeds the next.
- Extend the pipeline easily in `pipeline.go`:

```go
agents := []Agent{
    HexBinAgent{},
    CaseConvAgent{},
    PunctuationAgent{},
    ArticleAgent{},
    // Add more agents here
}
```

---

## 🧩 Extending the Pipeline

To add a new agent:

1. Create a file in `internal/transformations/`, e.g. `newagent.go`.
2. Implement the interface:
   ```go
   type NewAgent struct{}

   func (n NewAgent) Name() string { return "NewAgent" }

   func (n NewAgent) Process(input string) string {
       // transformation logic
       return input
   }
   ```
3. Register it in `RunPipeline()` inside `pipeline.go`.
4. Document it here with name, purpose, spec, and examples.

---

## 🧱 Document History

| Date | Change |
|------|--------|
| 2025-10-29 | Initial finalized version with all four core agents |
| 2025-10-30 | Added ASCII diagram, test refs, and extension guide |

---

This specification ensures consistent behavior across all transformations and serves as the authoritative reference for testing, auditing, and extending the go-reloaded pipeline.

