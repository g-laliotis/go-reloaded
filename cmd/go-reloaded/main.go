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
	// Handle help flag
	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		showHelp()
		return
	}

	// Handle version flag
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println(version.GetBuildInfo())
		return
	}

	// Expect exactly 2 arguments: input and output paths.
	if len(os.Args) != 3 {
		showUsage()
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

// showHelp displays comprehensive help information
func showHelp() {
	fmt.Printf("%s\n\n", version.Short())
	fmt.Println("DESCRIPTION:")
	fmt.Println("  A command-line text transformation tool using pipeline architecture.")
	fmt.Println("  Applies smart transformations through independent agents:")
	fmt.Println("    • HexBinAgent: Convert hex/binary numbers to decimal")
	fmt.Println("    • CaseConvAgent: Apply case transformations (up/low/cap)")
	fmt.Println("    • PunctuationAgent: Fix punctuation spacing and quotes")
	fmt.Println("    • ArticleAgent: Convert a/A to an/An before vowels")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Printf("  %s <input_file> <output_file>\n", os.Args[0])
	fmt.Printf("  %s [OPTIONS]\n", os.Args[0])
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Printf("  %s input.txt output.txt\n", os.Args[0])
	fmt.Printf("  %s sample.txt result.txt\n", os.Args[0])
	fmt.Println()
	fmt.Println("TRANSFORMATION EXAMPLES:")
	fmt.Println("  Input:  \"I have 1E (hex) apples and a orange (up)!\"")
	fmt.Println("  Output: \"I have 30 apples and an ORANGE!\"")
	fmt.Println()
	fmt.Println("For more information, visit: https://github.com/g-laliotis/go-reloaded")
}

// showUsage displays brief usage information
func showUsage() {
	fmt.Printf("Usage: %s <input_file> <output_file>\n", os.Args[0])
	fmt.Printf("       %s --help\n", os.Args[0])
	fmt.Printf("       %s --version\n", os.Args[0])
}
