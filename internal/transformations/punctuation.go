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

		// 1. Normalize ellipses first (protect them from dot rules)
		l = regexp.MustCompile(`\.\s*\.\s*\.`).ReplaceAllString(l, `...`)
		l = regexp.MustCompile(`\.{4,}`).ReplaceAllString(l, `...`)
		l = strings.ReplaceAll(l, "...", "ELLIPSIS_PLACEHOLDER")

		// 2. Fix exclamation/question marks
		for {
			newL := regexp.MustCompile(`([!?]+)\s+([!?]+)`).ReplaceAllString(l, `$1$2`)
			if newL == l {
				break
			}
			l = newL
		}
		l = regexp.MustCompile(`\s+([!?]+)`).ReplaceAllString(l, `$1`)
		l = regexp.MustCompile(`([!?]+)([^\s!?])`).ReplaceAllString(l, `$1 $2`)

		// 3. Fix simple punctuation: , . : ;
		l = regexp.MustCompile(`\s+([,.:;])`).ReplaceAllString(l, `$1`)
		l = regexp.MustCompile(`([,.:;])(\S)`).ReplaceAllString(l, `$1 $2`)

		// 4. Restore ellipses and fix their spacing
		l = strings.ReplaceAll(l, "ELLIPSIS_PLACEHOLDER", "...")
		l = regexp.MustCompile(`\s+\.{3}`).ReplaceAllString(l, `...`)
		l = regexp.MustCompile(`\.{3}(\S)`).ReplaceAllString(l, `... $1`)

		// 5. Clean up multiple spaces
		l = regexp.MustCompile(` {2,}`).ReplaceAllString(l, ` `)

		// 6. Fix quotes: remove spaces inside ' ... ' and " ... " (do this last)
		// Handle nested quotes: first fix double quotes, then single quotes
		l = regexp.MustCompile(`"\s*([^"]*?)\s*"`).ReplaceAllString(l, `"$1"`)
		l = regexp.MustCompile(`'\s*([^']*?)\s*'`).ReplaceAllString(l, `'$1'`)

		// 7. Final cleanup - trim trailing spaces
		l = strings.TrimRight(l, " ")

		out = append(out, l)
	}
	return strings.Join(out, "\n")
}
