package main

import (
	"strings"
	"testing"

	"go-reloaded/internal/transformations"
)

func TestHexBin(t *testing.T) {
	in := "Simply add 1E (hex) and 10 (bin) to see the result."
	out := transformations.RunPipeline(in)
	if !strings.Contains(out, "30") || !strings.Contains(out, "2") {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestCaseConv(t *testing.T) {
	in := "this is so exciting (up, 2)\nSTOP SHOUTING (low)\nwelcome to the brooklyn bridge (cap, 2)"
	// (up, 2) -> "SO EXCITING"
	// (low)   -> only the previous word -> "STOP shouting"
	// (cap,2) -> "Brooklyn Bridge"
	exp := "this is SO EXCITING\nSTOP shouting\nwelcome to the Brooklyn Bridge"
	out := transformations.RunPipeline(in)
	if out != exp {
		t.Fatalf("got:\n%q\nwant:\n%q", out, exp)
	}
}

func TestPunctuation(t *testing.T) {
	in := "Punct ... test ,ok !? yes...now"
	exp := "Punct... test, ok!? yes... now"
	out := transformations.RunPipeline(in)
	if out != exp {
		t.Fatalf("got:%q want:%q", out, exp)
	}
}

func TestArticle(t *testing.T) {
	in := "a apple\nA honest person\na 'idea'\na (hour)"
	exp := "an apple\nAn honest person\nan 'idea'\nan (hour)"
	out := transformations.RunPipeline(in)
	if out != exp {
		t.Fatalf("got:\n%q\nwant:\n%q", out, exp)
	}
}
