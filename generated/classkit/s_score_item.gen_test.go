// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSScoreItem

// ExampleNewSScoreItemWithIdentifierTitleScoreMaxScore demonstrates how to create a SScoreItem instance using NewSScoreItemWithIdentifierTitleScoreMaxScore.
// Initializes an activity item that holds a score value.
func ExampleNewSScoreItemWithIdentifierTitleScoreMaxScore() {
	_ = classkit.NewSScoreItemWithIdentifierTitleScoreMaxScore(
		"identifier", // identifier string
		"title",      // title string
		0.0,          // score float64
		0.0,          // maxScore float64
	)
	// Output:
}
