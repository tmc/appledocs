// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage_test

import (
	"github.com/tmc/appledocs/generated/naturallanguage"
)

// Suppress unused import errors
var _ = naturallanguage.NewContextualEmbedding

// ExampleNewContextualEmbeddingWithModelIdentifier demonstrates how to create a ContextualEmbedding instance using NewContextualEmbeddingWithModelIdentifier.
// Creates a contextual embedding from a model identifier.
func ExampleNewContextualEmbeddingWithModelIdentifier() {
	_ = naturallanguage.NewContextualEmbeddingWithModelIdentifier(
		"modelIdentifier", // modelIdentifier string
	)
	// Output:
}
