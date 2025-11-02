// hexbin.go implements HexBinAgent.
// It finds numbers followed by (hex) or (bin) and replaces them with
// their decimal value. We process line-by-line to preserve newlines.

package transformations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type HexBinAgent struct{}

func (h HexBinAgent) Name() string { return "HexBinAgent" }

func (h HexBinAgent) Process(input string) string {
	lines := strings.Split(input, "\n")
	out := make([]string, 0, len(lines))

	// More efficient: separate regex for hex and bin
	reHex := regexp.MustCompile(`\b([0-9A-Fa-f]+)\s*\(hex\)`)
	reBin := regexp.MustCompile(`\b([01]+)\s*\(bin\)`)

	for _, line := range lines {
		// Process hex conversions
		line = reHex.ReplaceAllStringFunc(line, func(m string) string {
			num := reHex.FindStringSubmatch(m)[1]
			if v, err := strconv.ParseInt(num, 16, 64); err == nil {
				return fmt.Sprintf("%d", v)
			}
			return m
		})

		// Process binary conversions
		line = reBin.ReplaceAllStringFunc(line, func(m string) string {
			num := reBin.FindStringSubmatch(m)[1]
			if v, err := strconv.ParseInt(num, 2, 64); err == nil {
				return fmt.Sprintf("%d", v)
			}
			return m
		})

		out = append(out, line)
	}
	return strings.Join(out, "\n")
}
