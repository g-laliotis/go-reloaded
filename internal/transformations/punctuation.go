// punctuation.go implements PunctuationAgent.
// It enforces spacing rules for punctuation (. , ! ? : ;).
//  - No space before punctuation; exactly one space after (unless end-of-line).
//  - Preserve grouped punctuations: "..." and mixes of ! and ? (e.g., "!?").
//  - Trim spaces immediately inside paired single quotes: ' awesome ' -> 'awesome'.
// Works line-by-line to keep line breaks.

package transformations

import (
	"regexp"
	"strings"
)

type PunctuationAgent struct{}

func (p PunctuationAgent) Name() string { return "PunctuationAgent" }

func (p PunctuationAgent) Process(input string) string {
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		l := line

		// 1. Fix quotes: remove spaces inside ' ... '
		l = regexp.MustCompile(`'\s*([^']*?)\s*'`).ReplaceAllString(l, `'$1'`)

		// 2. Normalize ellipses first (protect them from dot rules)
		l = regexp.MustCompile(`\.\s*\.\s*\.`).ReplaceAllString(l, `...`)
		l = regexp.MustCompile(`\.{4,}`).ReplaceAllString(l, `...`)
		l = strings.ReplaceAll(l, "...", "ELLIPSIS_PLACEHOLDER")

		// 3. Fix exclamation/question marks
		for {
			newL := regexp.MustCompile(`([!?]+)\s+([!?]+)`).ReplaceAllString(l, `$1$2`)
			if newL == l {
				break
			}
			l = newL
		}
		l = regexp.MustCompile(`\s+([!?]+)`).ReplaceAllString(l, `$1`)
		l = regexp.MustCompile(`([!?]+)([^\s!?])`).ReplaceAllString(l, `$1 $2`)

		// 4. Fix simple punctuation: , . : ;
		l = regexp.MustCompile(`\s+([,.:;])`).ReplaceAllString(l, `$1`)
		l = regexp.MustCompile(`([,.:;])(\S)`).ReplaceAllString(l, `$1 $2`)

		// 5. Restore ellipses and fix their spacing
		l = strings.ReplaceAll(l, "ELLIPSIS_PLACEHOLDER", "...")
		l = regexp.MustCompile(`\s+\.{3}`).ReplaceAllString(l, `...`)
		l = regexp.MustCompile(`\.{3}(\S)`).ReplaceAllString(l, `... $1`)

		// 6. Clean up multiple spaces and trailing spaces
		l = regexp.MustCompile(` {2,}`).ReplaceAllString(l, ` `)
		l = strings.TrimRight(l, " ")

		out = append(out, l)
	}
	return strings.Join(out, "\n")
}
