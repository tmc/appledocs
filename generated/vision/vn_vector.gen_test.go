// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision_test

import (
	"github.com/tmc/appledocs/generated/vision"
)

// Suppress unused import errors
var _ = vision.NewVector

// ExampleNewVectorWithRTheta demonstrates how to create a Vector instance using NewVectorWithRTheta.
// Creates a new vector in polar coordinate space.
func ExampleNewVectorWithRTheta() {
	_ = vision.NewVectorWithRTheta(
		0.0, // r float64
		0.0, // theta float64
	)
	// Output:
}
// ExampleNewVectorWithXComponentYComponent demonstrates how to create a Vector instance using NewVectorWithXComponentYComponent.
// Creates a new vector in Cartesian coordinate space, based on its x-axis and y-axis projections.
func ExampleNewVectorWithXComponentYComponent() {
	_ = vision.NewVectorWithXComponentYComponent(
		0.0, // x float64
		0.0, // y float64
	)
	// Output:
}
