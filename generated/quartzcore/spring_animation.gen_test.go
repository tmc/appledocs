// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewSpringAnimation

// ExampleNewSpringAnimationWithPerceptualDurationBounce demonstrates how to create a SpringAnimation instance using NewSpringAnimationWithPerceptualDurationBounce.
func ExampleNewSpringAnimationWithPerceptualDurationBounce() {
	_ = quartzcore.NewSpringAnimationWithPerceptualDurationBounce(
		quartzcore.TimeInterval(0.0), // perceptualDuration TimeInterval
		0.0, // bounce float64
	)
	// Output:
}
