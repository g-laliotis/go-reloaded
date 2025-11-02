package transformations

import "strings"

// NormalizeSpaces removes redundant spaces and ensures consistent spacing.
func NormalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
