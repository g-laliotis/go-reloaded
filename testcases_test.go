package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-reloaded/internal/transformations"
)

func TestCases(t *testing.T) {
	testDir := "testcases"
	
	// Find all input files
	inputFiles, err := filepath.Glob(filepath.Join(testDir, "*_input.txt"))
	if err != nil {
		t.Fatalf("Failed to find input files: %v", err)
	}

	for _, inputFile := range inputFiles {
		// Extract case name (e.g., "case01_basic" from "case01_basic_input.txt")
		baseName := strings.TrimSuffix(filepath.Base(inputFile), "_input.txt")
		expectedFile := filepath.Join(testDir, baseName+"_expected_output.txt")

		t.Run(baseName, func(t *testing.T) {
			// Read input
			input, err := os.ReadFile(inputFile)
			if err != nil {
				t.Fatalf("Failed to read input file %s: %v", inputFile, err)
			}

			// Read expected output
			expected, err := os.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("Failed to read expected file %s: %v", expectedFile, err)
			}

			// Run pipeline
			actual := transformations.RunPipeline(string(input))

			// Compare
			if actual != string(expected) {
				t.Errorf("Case %s failed:\nInput:\n%s\nExpected:\n%s\nActual:\n%s", 
					baseName, string(input), string(expected), actual)
			}
		})
	}
}