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
	for j := 0; j < n; j++ {
		idx := len(res) - 1 - j
		if idx < 0 {
			break // not enough previous words
		}
		switch command {
		case "up":
			res[idx] = strings.ToUpper(res[idx])
		case "low":
			res[idx] = strings.ToLower(res[idx])
		case "cap":
			res[idx] = capitalize(res[idx])
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
