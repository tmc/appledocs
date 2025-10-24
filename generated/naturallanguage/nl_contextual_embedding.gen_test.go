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
		naturallanguage.Language /* typedef */{}, // language Language /* typedef */
	)
	// Output:
}
// ExampleNewContextualEmbeddingWithScript demonstrates how to create a ContextualEmbedding instance using NewContextualEmbeddingWithScript.
// Creates a contextual embedding from a script.
func ExampleNewContextualEmbeddingWithScript() {
	_ = naturallanguage.NewContextualEmbeddingWithScript(
		naturallanguage.Script /* typedef */{}, // script Script /* typedef */
	)
	// Output:
}
// ExampleContextualEmbedding_Unload demonstrates using Unload on a ContextualEmbedding instance.
// Unloads the embedding model.
func ExampleContextualEmbedding_Unload() {
	obj := naturallanguage.NewContextualEmbedding()
	obj.Unload()
	// Output:
	}

