package transformations

import "testing"

func TestHexBinAgent(t *testing.T) {
	agent := HexBinAgent{}
	tests := []struct {
		input, expected string
	}{
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"FF (hex) and 1010 (bin)", "255 and 10"},
		{"invalid G1 (hex)", "invalid G1 (hex)"}, // should remain unchanged
	}

	for _, test := range tests {
		result := agent.Process(test.input)
		if result != test.expected {
			t.Errorf("HexBin: input %q, got %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestCaseConvAgent(t *testing.T) {
	agent := CaseConvAgent{}
	tests := []struct {
		input, expected string
	}{
		{"hello (up)", "HELLO"},
		{"WORLD (low)", "world"},
		{"test (cap)", "Test"},
		{"one two (up, 2)", "ONE TWO"},
	}

	for _, test := range tests {
		result := agent.Process(test.input)
		if result != test.expected {
			t.Errorf("CaseConv: input %q, got %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestPunctuationAgent(t *testing.T) {
	agent := PunctuationAgent{}
	tests := []struct {
		input, expected string
	}{
		{"hello , world", "hello, world"},
		{"test ... ok", "test... ok"},
		{"wow !! really ?", "wow!! really?"},
		{"' spaced '", "'spaced'"},
	}

	for _, test := range tests {
		result := agent.Process(test.input)
		if result != test.expected {
			t.Errorf("Punctuation: input %q, got %q, want %q", test.input, result, test.expected)
		}
	}
}

func TestArticleAgent(t *testing.T) {
	agent := ArticleAgent{}
	tests := []struct {
		input, expected string
	}{
		{"a apple", "an apple"},
		{"A honest", "An honest"},
		{"a car", "a car"},       // should not change
		{"a 8-hour", "a 8-hour"}, // should not change (starts with number)
	}

	for _, test := range tests {
		result := agent.Process(test.input)
		if result != test.expected {
			t.Errorf("Article: input %q, got %q, want %q", test.input, result, test.expected)
		}
	}
}
