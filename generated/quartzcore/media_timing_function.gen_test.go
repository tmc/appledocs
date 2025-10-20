// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewMediaTimingFunction


// ExampleNewMediaTimingFunctionWithControlPoints demonstrates how to create a MediaTimingFunction instance using NewMediaTimingFunctionWithControlPoints.
// Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
func ExampleNewMediaTimingFunctionWithControlPoints() {
	_ = quartzcore.NewMediaTimingFunctionWithControlPoints(
		0.0, // c1x float32
		0.0, // c1y float32
		0.0, // c2x float32
		0.0, // c2y float32
	)
	// Output:
}



