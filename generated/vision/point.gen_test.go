// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewPoint

// ExampleNewPointWithXY demonstrates how to create a Point instance using NewPointWithXY.
// Creates a point object with the specified coordinates.
func ExampleNewPointWithXY() {
	_ = vision.NewPointWithXY(
		0.0, // x float64
		0.0, // y float64
	)
	// Output:
}
