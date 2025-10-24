// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage_test

import (
	"github.com/tmc/appledocs/generated/naturallanguage"
)

// Suppress unused import errors
var _ = naturallanguage.NewContextualEmbedding

// ExampleContextualEmbedding_Unload demonstrates using Unload on a ContextualEmbedding instance.
// Unloads the embedding model.
func ExampleContextualEmbedding_Unload() {
	obj := naturallanguage.NewContextualEmbedding()
	obj.Unload()
	// Output:
	}

