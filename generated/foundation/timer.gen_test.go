// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewTimer

// ExampleNewTimerWithTimeIntervalInvocationRepeats demonstrates how to create a Timer instance using NewTimerWithTimeIntervalInvocationRepeats.
// Initializes a timer object with the specified invocation object.
func ExampleNewTimerWithTimeIntervalInvocationRepeats() {
	_ = foundation.NewTimerWithTimeIntervalInvocationRepeats(
		foundation.TimeInterval(0.0), // ti TimeInterval
		foundation.NSInvocation{}, // invocation NSInvocation
		false, // yesOrNo bool
	)
	// Output:
}
