# AGENTS SPECIFICATION

## PIPELINE_ORDER
1. HexBinAgent
2. CaseConvAgent  
3. PunctuationAgent
4. ArticleAgent

## AGENT_INTERFACE
```go
type Agent interface {
    Name() string
    Process(input string) string
}
```

## HEXBIN_AGENT
FILE: internal/transformations/hexbin.go
ORDER: 1
FUNCTION: Convert hex/binary numbers to decimal
PATTERNS:
- INPUT: NUMBER (hex) -> OUTPUT: DECIMAL
- INPUT: NUMBER (bin) -> OUTPUT: DECIMAL
REGEX_HEX: [0-9A-Fa-f]+
REGEX_BIN: [01]+
EXAMPLES:
- "1E (hex)" -> "30"
- "10 (bin)" -> "2"
- "Simply add 1E (hex) and 10 (bin)." -> "Simply add 30 and 2."
EDGE_CASES:
- Invalid hex/bin formats remain unchanged
- Preserves line breaks and spacing

## CASECONV_AGENT
FILE: internal/transformations/caseconv.go
ORDER: 2
FUNCTION: Apply case transformations with markers
MARKERS:
- (up) -> UPPERCASE previous word
- (low) -> lowercase previous word  
- (cap) -> Capitalize previous word
- (up,N) -> UPPERCASE previous N words
- (low,N) -> lowercase previous N words
- (cap,N) -> Capitalize previous N words
EXAMPLES:
- "Ready set go (up)" -> "Ready set GO"
- "this is so exciting (up, 2)" -> "this is SO EXCITING"
- "STOP SHOUTING (low)" -> "STOP shouting"
- "welcome to the brooklyn bridge (cap, 2)" -> "welcome to the Brooklyn Bridge"
BEHAVIOR:
- Markers are removed after processing
- Count defaults to 1 if not specified
- Supports both compact and spaced formats

## PUNCTUATION_AGENT
FILE: internal/transformations/punctuation.go
ORDER: 3
FUNCTION: Normalize punctuation spacing and quotes
RULES:
- BEFORE_PUNCT: Remove spaces before . , ! ? : ;
- AFTER_PUNCT: Add single space after punctuation if followed by text
- ELLIPSIS: Convert ". . ." or " ... " to "..."
- MULTI_PUNCT: Compact "! !" to "!!" and "? !" to "?!"
- QUOTES: Remove spaces inside single quotes ' text ' -> 'text'
EXAMPLES:
- "I was thinking ... You were right" -> "I was thinking... You were right"
- "over there ,and then BAMM !!" -> "over there, and then BAMM!!"
- "Are you serious !? yes...right now" -> "Are you serious!? yes... right now"
- "As Elton John said: ' awesome '" -> "As Elton John said: 'awesome'"
PRESERVATION:
- Line breaks preserved
- Multiple punctuation groups maintained

## ARTICLE_AGENT
FILE: internal/transformations/article.go
ORDER: 4
FUNCTION: Convert a/A to an/An before vowels and h
TRIGGER_LETTERS: a e i o u h
CONVERSIONS:
- "a" + vowel/h -> "an"
- "A" + vowel/h -> "An"
EXAMPLES:
- "a apple" -> "an apple"
- "A honest person" -> "An honest person"
- "a 'idea' a (hour) a [umbrella]" -> "an 'idea' an (hour) an [umbrella]"
CONSTRAINTS:
- Only applies to alphabetic words
- Ignores leading punctuation/quotes
- Preserves original case
- Numbers and symbols unchanged

## TESTING_FRAMEWORK
UNIT_TESTS: internal/transformations/agents_test.go
INTEGRATION_TESTS: main_test.go, edge_cases_test.go
COMPREHENSIVE_TESTS: testcases_test.go (24 cases)
PERFORMANCE_TESTS: benchmark_test.go
TEST_DATA: testcases/*.txt files
COMMANDS:
- make test
- make testcases  
- make bench
- make clear-cache

## IMPLEMENTATION_REQUIREMENTS
- Process text line-by-line
- Preserve newlines and line structure
- Sequential pipeline execution
- Deterministic output
- No side effects between agents
- UTF-8 text support

## EXTENSION_PROTOCOL
1. Create agent file in internal/transformations/
2. Implement Agent interface
3. Add to pipeline in pipeline.go
4. Add unit tests
5. Update this specification