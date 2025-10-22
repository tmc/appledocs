// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage_test

import (
	"github.com/tmc/appledocs/generated/naturallanguage"
)

// Suppress unused import errors
var _ = naturallanguage.NewContextualEmbedding

// ExampleNewContextualEmbeddingWithLanguage demonstrates how to create a ContextualEmbedding instance using NewContextualEmbeddingWithLanguage.
// Creates a contextual embedding from a language.
func ExampleNewContextualEmbeddingWithLanguage() {
	_ = naturallanguage.NewContextualEmbeddingWithLanguage(
		naturallanguage.Language{}, // language Language
	)
	// Output:
}
// ExampleNewContextualEmbeddingWithModelIdentifier demonstrates how to create a ContextualEmbedding instance using NewContextualEmbeddingWithModelIdentifier.
// Creates a contextual embedding from a model identifier.
func ExampleNewContextualEmbeddingWithModelIdentifier() {
	_ = naturallanguage.NewContextualEmbeddingWithModelIdentifier(
		"modelIdentifier", // modelIdentifier string
	)
	// Output:
}
// ExampleNewContextualEmbeddingWithScript demonstrates how to create a ContextualEmbedding instance using NewContextualEmbeddingWithScript.
// Creates a contextual embedding from a script.
func ExampleNewContextualEmbeddingWithScript() {
	_ = naturallanguage.NewContextualEmbeddingWithScript(
		naturallanguage.Script{}, // script Script
	)
	// Output:
}
