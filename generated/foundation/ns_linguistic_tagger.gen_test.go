// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewLinguisticTagger

// ExampleNewLinguisticTaggerWithTagSchemesOptions demonstrates how to create a LinguisticTagger instance using NewLinguisticTaggerWithTagSchemesOptions.
// Creates a linguistic tagger instance using the specified tag schemes and options.
func ExampleNewLinguisticTaggerWithTagSchemesOptions() {
	_ = foundation.NewLinguisticTaggerWithTagSchemesOptions(
		[]foundation.string{}, // tagSchemes []string
		0, // opts uint
	)
	// Output:
}
