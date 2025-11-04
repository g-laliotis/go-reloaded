// main.go is the entry point of the CLI app.
// It:
//  1) reads an input text file,
//  2) runs the text through the transformation pipeline,
//  3) writes the transformed result to an output file.
//
// Usage:
//   go run ./cmd/go-reloaded <input_file> <output_file>

package main

import (
	"fmt"
	"os"

	"go-reloaded/internal/transformations"
	"go-reloaded/internal/version"
)

func main() {
	// Handle version flag
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println(version.GetBuildInfo())
		return
	}

	// Expect exactly 2 arguments: input and output paths.
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <input_file> <output_file>\n", os.Args[0])
		fmt.Printf("       %s --version\n", os.Args[0])
		os.Exit(1)
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	// Read the entire input file into memory.
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		os.Exit(1)
	}

	// Run our pipeline (sequence of agents) over the text.
	result := transformations.RunPipeline(string(data))

	// Write the transformed text to the output file.
	if err := os.WriteFile(outputPath, []byte(result), 0o644); err != nil {
		fmt.Printf("error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("OK →", outputPath)
}
