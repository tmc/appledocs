// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewTimer


// ExampleNewTimerWithTimeIntervalTargetSelectorUserInfoRepeats demonstrates how to create a Timer instance using NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats.
// Initializes a timer object with the specified object and selector.
func ExampleNewTimerWithTimeIntervalTargetSelectorUserInfoRepeats() {
	_ = foundation.NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats(
		foundation.TimeInterval(0), // ti TimeInterval
		0, // aTarget objc.ID
		0, // aSelector objc.SEL
		0, // userInfo objc.ID
		false, // yesOrNo bool
	)
	// Output:
}






