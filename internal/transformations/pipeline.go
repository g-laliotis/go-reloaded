// pipeline.go defines the "pipeline" — a list of agents (stages) that
// process text one after another. Each agent implements the Agent interface.

package transformations

import "strings"

// Agent is any stage that can transform text.
type Agent interface {
	Name() string
	Process(input string) string
}

// RunPipeline executes all agents in a fixed order.
// The order matters (as per the assignment rules).
func RunPipeline(text string) string {
	agents := []Agent{
		HexBinAgent{},      // 1) convert N (hex)/(bin) → decimal
		CaseConvAgent{},    // 2) apply (up)/(low)/(cap[, N])
		PunctuationAgent{}, // 3) fix punctuation spacing & quotes
		ArticleAgent{},     // 4) "a" → "an" before vowel or 'h'
	}

	// Pass the text through each agent.
	for _, a := range agents {
		text = a.Process(text)
		// (Optional) debugging:
		// fmt.Printf("---- %s ----\n%s\n\n", a.Name(), text)
	}

	// Keep original line breaks; just trim trailing space/newlines.
	return strings.TrimRight(text, "\n ")
}
