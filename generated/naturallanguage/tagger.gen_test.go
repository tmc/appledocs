// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage_test

import (
	"github.com/tmc/appledocs/generated/naturallanguage"
)

// Suppress unused import errors
var _ = naturallanguage.NewTagger

// ExampleNewTaggerWithTagSchemes demonstrates how to create a Tagger instance using NewTaggerWithTagSchemes.
// Creates a linguistic tagger instance using the specified tag schemes and options.
func ExampleNewTaggerWithTagSchemes() {
	_ = naturallanguage.NewTaggerWithTagSchemes(
		[]naturallanguage.string{}, // tagSchemes []string
	)
	// Output:
}
