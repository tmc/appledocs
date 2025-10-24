// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis_test

import (
	"github.com/tmc/appledocs/generated/soundanalysis"
)

// Suppress unused import errors
var _ = soundanalysis.NewSNTimeDurationConstraint

// ExampleNewSNTimeDurationConstraintWithDurationRange demonstrates how to create a SNTimeDurationConstraint instance using NewSNTimeDurationConstraintWithDurationRange.
// Creates a constraint with a time duration range.
func ExampleNewSNTimeDurationConstraintWithDurationRange() {
	_ = soundanalysis.NewSNTimeDurationConstraintWithDurationRange(
		soundanalysis.TimeRange /* not a class type */ {}, // durationRange TimeRange /* not a class type */
	)
	// Output:
}
