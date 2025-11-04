// article.go implements ArticleAgent.
// It changes "a" → "an" (and "A" → "An") when the NEXT word starts
// with a vowel (a,e,i,o,u) or 'h'. We keep capitalization and
// handle leading wrappers on the next token (quotes/brackets).

package transformations

import (
	"strings"
	"unicode"
)

type ArticleAgent struct{}

func (a ArticleAgent) Name() string { return "ArticleAgent" }

func (a ArticleAgent) Process(input string) string {
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		words := strings.Fields(line)
		for i := 0; i < len(words)-1; i++ {
			current := words[i]
			if current == "a" || current == "A" || current == "A_UP" || current == "A_CAP" {
				next := trimLeadingWrappers(words[i+1])
				first := firstLetterRune(next)
				// Only change if next word starts with a letter (not number) and is vowel/h
				if first != 0 && isVowelOrH(first) {
					switch current {
					case "a":
						words[i] = "an"
					case "A":
						words[i] = "An" // Default A becomes An
					case "A_UP":
						words[i] = "AN" // (up) marker becomes AN
					case "A_CAP":
						words[i] = "An" // (cap) marker becomes An
					}
				} else {
					// If not followed by vowel, clean up markers
					if current == "A_UP" {
						words[i] = "A"
					} else if current == "A_CAP" {
						words[i] = "A"
					}
				}
			}
		}
		// Clean up any remaining markers that weren't processed
		for j := range words {
			if words[j] == "A_UP" || words[j] == "A_CAP" {
				words[j] = "A"
			}
		}
		out = append(out, strings.Join(words, " "))
	}
	return strings.Join(out, "\n")
}

// Remove leading quotes/brackets from the next token: "'apple" -> "apple"
func trimLeadingWrappers(s string) string {
	for len(s) > 0 {
		r := []rune(s)[0]
		switch r {
		case '\'', '"', '“', '”', '‘', '’', '(', '[', '{':
			s = string([]rune(s)[1:])
		default:
			return s
		}
	}
	return s
}

// firstLetterRune returns the first rune in lowercase if it's a letter; 0 if first rune is not a letter.
func firstLetterRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	first := []rune(s)[0]
	if unicode.IsLetter(first) {
		return unicode.ToLower(first)
	}
	return 0 // first character is not a letter (could be number, symbol, etc.)
}

func isVowelOrH(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false
	}
}
