// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewAffineTransform

// ExampleNewAffineTransform demonstrates how to create a AffineTransform instance.
// Initializes an affine transform matrix to the identity matrix.
func ExampleNewAffineTransform() {
	_ = foundation.NewAffineTransform()
	// Output:
}
// ExampleNewAffineTransformWithTransform demonstrates how to create a AffineTransform instance using NewAffineTransformWithTransform.
// Initializes the receiver’s matrix using another transform object.
func ExampleNewAffineTransformWithTransform() {
	_ = foundation.NewAffineTransformWithTransform(
		foundation.NSAffineTransform{}, // transform NSAffineTransform
	)
	// Output:
}
