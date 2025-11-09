package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"go-reloaded/internal/transformations"
	"go-reloaded/internal/version"
)

func formatOutput(result string, format string) {
	switch format {
	case "json":
		data := struct {
			Result string `json:"result"`
			Status string `json:"status"`
		}{
			Result: result,
			Status: "success",
		}
		jsonOutput, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Printf("error generating JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonOutput))
	default:
		fmt.Println(result)
	}
}

func main() {
	// Define and parse flags first
	format := flag.String("format", "text", "Output format: text or json")
	flag.Parse()

	// Handle help flag
	if len(flag.Args()) == 0 || flag.Arg(0) == "--help" || flag.Arg(0) == "-h" {
		showHelp()
		return
	}

	// Handle version flag
	if flag.Arg(0) == "--version" || flag.Arg(0) == "-v" {
		fmt.Println(version.GetBuildInfo())
		return
	}

	// Get non-flag arguments
	args := flag.Args()
	if len(args) != 2 {
		showUsage()
		os.Exit(1)
	}
	inputPath := args[0]
	outputPath := args[1]

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

	formatOutput(result, *format)
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
	fmt.Printf("  %s [OPTIONS] <input_file> <output_file>\n", os.Args[0])
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -format string")
	fmt.Println("        Output format: text or json (default \"text\")")
	fmt.Println("  -h, --help")
	fmt.Println("        Show this help message")
	fmt.Println("  -v, --version")
	fmt.Println("        Show version information")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Printf("  %s input.txt output.txt\n", os.Args[0])
	fmt.Printf("  %s -format=json input.txt output.json\n", os.Args[0])
	fmt.Println()
	fmt.Println("For more information, visit: https://github.com/g-laliotis/go-reloaded")
}

// showUsage displays brief usage information
func showUsage() {
	fmt.Printf("Usage: %s [OPTIONS] <input_file> <output_file>\n", os.Args[0])
	fmt.Printf("       %s --help\n", os.Args[0])
	fmt.Printf("       %s --version\n", os.Args[0])
}
