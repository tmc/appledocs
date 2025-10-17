// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats demonstrates how to create a Timer instance using NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats.
// Initializes a timer using the specified object and selector.
func ExampleNewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats() {
	_ = foundation.NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(
		nil, // date unsafe.Pointer
		foundation.TimeInterval(0), // ti TimeInterval
		0, // t objc.ID
		0, // s objc.SEL
		0, // ui objc.ID
		false, // rep bool
	)
	// Output:
}

// ExampleNewTimerWithTimeIntervalInvocationRepeats demonstrates how to create a Timer instance using NewTimerWithTimeIntervalInvocationRepeats.
// Initializes a timer object with the specified invocation object.
func ExampleNewTimerWithTimeIntervalInvocationRepeats() {
	_ = foundation.NewTimerWithTimeIntervalInvocationRepeats(
		foundation.TimeInterval(0), // ti TimeInterval
		nil, // invocation unsafe.Pointer
		false, // yesOrNo bool
	)
	// Output:
}

// ExampleNewTimerWithTimeIntervalRepeatsBlock demonstrates how to create a Timer instance using NewTimerWithTimeIntervalRepeatsBlock.
// Initializes a timer object with the specified time interval and block.
func ExampleNewTimerWithTimeIntervalRepeatsBlock() {
	_ = foundation.NewTimerWithTimeIntervalRepeatsBlock(
		foundation.TimeInterval(0), // interval TimeInterval
		false, // repeats bool
		nil, // block unsafe.Pointer
	)
	// Output:
}

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

// ExampleNewTimerWithFireDateIntervalRepeatsBlock demonstrates how to create a Timer instance using NewTimerWithFireDateIntervalRepeatsBlock.
// Initializes a timer for the specified date and time interval with the specified block.
func ExampleNewTimerWithFireDateIntervalRepeatsBlock() {
	_ = foundation.NewTimerWithFireDateIntervalRepeatsBlock(
		nil, // date unsafe.Pointer
		foundation.TimeInterval(0), // interval TimeInterval
		false, // repeats bool
		nil, // block unsafe.Pointer
	)
	// Output:
}


