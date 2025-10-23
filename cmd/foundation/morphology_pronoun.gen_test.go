// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMorphologyPronoun

// ExampleNewMorphologyPronounWithPronounMorphologyDependentMorphology demonstrates how to create a MorphologyPronoun instance using NewMorphologyPronounWithPronounMorphologyDependentMorphology.
func ExampleNewMorphologyPronounWithPronounMorphologyDependentMorphology() {
	_ = foundation.NewMorphologyPronounWithPronounMorphologyDependentMorphology(
		"pronoun", // pronoun string
		foundation.NSMorphology{}, // morphology NSMorphology
		foundation.NSMorphology{}, // dependentMorphology NSMorphology
	)
	// Output:
}
