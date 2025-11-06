package main

import (
	"strings"
	"testing"

	"go-reloaded/internal/transformations"
)

func BenchmarkPipeline(b *testing.B) {
	input := "Simply add 1E (hex) and 10 (bin) and make it exciting (up, 2) ... what do you think ?"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transformations.RunPipeline(input)
	}
}

func TestLargeInput(t *testing.T) {
	// Test with larger input to ensure no performance issues
	lines := make([]string, 100)
	for i := range lines {
		lines[i] = "test line with 1A (hex) and punctuation ... ok ?"
	}
	input := strings.Join(lines, "\n")

	result := transformations.RunPipeline(input)

	// Should contain converted hex and fixed punctuation
	if !strings.Contains(result, "26") || !strings.Contains(result, "ok?") {
		t.Error("Large input processing failed")
	}
}
