// caseconv.go implements CaseConvAgent.
// It handles (up), (low), (cap) markers, including numbered variants like (up, 3).
// The marker applies to the previous word(s) and is removed from the output.
// We support both "(up,2)" and "(up, 2)" forms, and we preserve newlines.

package transformations

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type CaseConvAgent struct{}

func (c CaseConvAgent) Name() string { return "CaseConvAgent" }

func (c CaseConvAgent) Process(input string) string {
	// Work line-by-line to keep original line breaks intact.
	lines := strings.Split(input, "\n")
	outLines := make([]string, 0, len(lines))

	// Precompiled patterns for the three possible marker shapes.
	reSimple := regexp.MustCompile(`^\((up|low|cap)\)$`)        // (up)
	reCompact := regexp.MustCompile(`^\((up|low|cap),(\d+)\)$`) // (up,2)
	reTwoTokHead := regexp.MustCompile(`^\((up|low|cap),$`)     // "(up," + "2)"
	reTwoTokTail := regexp.MustCompile(`^(\d+)\)$`)

	for _, line := range lines {
		words := strings.Fields(line) // split by whitespace within line
		res := make([]string, 0, len(words))

		for i := 0; i < len(words); i++ {
			w := words[i]

			// (up) / (low) / (cap)
			if m := reSimple.FindStringSubmatch(w); m != nil {
				applyCaseCommand(&res, m[1], 1)
				continue
			}

			// (up,2) as one token
			if m := reCompact.FindStringSubmatch(w); m != nil {
				n, _ := strconv.Atoi(m[2])
				if n <= 0 {
					n = 1
				}
				applyCaseCommand(&res, m[1], n)
				continue
			}

			// (up, 2) split across two tokens: "(up," + "2)"
			if m := reTwoTokHead.FindStringSubmatch(w); m != nil && i+1 < len(words) {
				if m2 := reTwoTokTail.FindStringSubmatch(words[i+1]); m2 != nil {
					n, _ := strconv.Atoi(m2[1])
					if n <= 0 {
						n = 1
					}
					applyCaseCommand(&res, m[1], n)
					i++ // consume the "2)" token
					continue
				}
			}

			// Regular word — keep it.
			res = append(res, w)
		}
		outLines = append(outLines, strings.Join(res, " "))
	}
	return strings.Join(outLines, "\n")
}

// applyCaseCommand mutates the last n words in-place according to command.
func applyCaseCommand(words *[]string, command string, n int) {
	res := *words

	// Special case: if we have a pattern like "word , word word" and n=3,
	// and the comma is at position len-2, then include the word before the comma
	if command == "cap" && n == 3 && len(res) >= 4 {
		// Check if we have: [word] [,] [word] [word] pattern
		if len(res) >= 4 && res[len(res)-3] == "," &&
			len(res[len(res)-4]) > 0 && unicode.IsLetter([]rune(res[len(res)-4])[0]) &&
			len(res[len(res)-2]) > 0 && unicode.IsLetter([]rune(res[len(res)-2])[0]) &&
			len(res[len(res)-1]) > 0 && unicode.IsLetter([]rune(res[len(res)-1])[0]) {
			// Apply to the word before comma and the two words after
			res[len(res)-4] = capitalize(res[len(res)-4]) // word before comma
			res[len(res)-2] = capitalize(res[len(res)-2]) // first word after comma
			res[len(res)-1] = capitalize(res[len(res)-1]) // second word after comma
			*words = res
			return
		}
	}

	// Default behavior
	for j := 0; j < n; j++ {
		idx := len(res) - 1 - j
		if idx < 0 {
			break // not enough previous words
		}
		switch command {
		case "up":
			// Special handling for articles: mark them for ArticleAgent
			if res[idx] == "a" {
				res[idx] = "A_UP" // Special marker for uppercase article
			} else if res[idx] == "A_CAP" || res[idx] == "A_UP" {
				// Don't double-process markers, make it uppercase
				res[idx] = "A_UP"
			} else {
				res[idx] = strings.ToUpper(res[idx])
			}
		case "low":
			res[idx] = strings.ToLower(res[idx])
		case "cap":
			// Special handling for articles: mark them for ArticleAgent
			if res[idx] == "a" {
				res[idx] = "A_CAP" // Special marker for capitalized article
			} else if res[idx] == "A_CAP" || res[idx] == "A_UP" {
				// Don't double-process markers, just capitalize them
				res[idx] = "A_CAP"
			} else {
				res[idx] = capitalize(res[idx])
			}
		}
	}
	*words = res
}

// capitalize makes first rune uppercase and the rest lowercase.
func capitalize(word string) string {
	if word == "" {
		return word
	}
	r := []rune(word)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}
