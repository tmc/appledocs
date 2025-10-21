// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight_test

import (
	"github.com/tmc/appledocs/generated/corespotlight"
)

// Suppress unused import errors
var _ = corespotlight.NewCSUserQueryContext

// ExampleNewCSUserQueryContextWithCurrentSuggestion demonstrates how to create a CSUserQueryContext instance using NewCSUserQueryContextWithCurrentSuggestion.
// Creates a new query context object with an optional suggested search string.
func ExampleNewCSUserQueryContextWithCurrentSuggestion() {
	_ = corespotlight.NewCSUserQueryContextWithCurrentSuggestion(
		corespotlight.CSSuggestion{}, // currentSuggestion CSSuggestion
	)
	// Output:
}

