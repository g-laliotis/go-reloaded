package main

import (
	"testing"

	"go-reloaded/internal/transformations"
)

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		// Empty/whitespace
		{"empty", "", ""},
		{"whitespace only", "   \n  \n  ", ""},

		// Multi-line combinations
		{"multiline combo", "1A (hex) test (up)\na apple ... wow", "26 TEST\nan apple... wow"},

		// Invalid formats should remain unchanged
		{"invalid hex", "GZ (hex)", "GZ (hex)"},
		{"invalid bin", "123 (bin)", "123 (bin)"},

		// Complex punctuation
		{"complex punct", "wow !! ... really !? yes", "wow!!... really!? yes"},

		// Article edge cases
		{"article with numbers", "a 8-hour shift and a honest person", "a 8-hour shift and an honest person"},
		{"article with symbols", "a @symbol and a #hashtag", "a @symbol and a #hashtag"},

		// Case conversion edge cases
		{"case with punct", "hello world (up, 2) !", "HELLO WORLD!"},
		{"case beyond bounds", "word (up, 5)", "WORD"}, // should handle gracefully
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := transformations.RunPipeline(test.input)
			if result != test.expected {
				t.Errorf("got %q, want %q", result, test.expected)
			}
		})
	}
}

func TestLinePreservation(t *testing.T) {
	input := "line1 (up)\n\nline3 1A (hex)\n"
	expected := "LINE1\n\nline3 26"
	result := transformations.RunPipeline(input)
	if result != expected {
		t.Errorf("Line preservation failed:\ngot:  %q\nwant: %q", result, expected)
	}
}
