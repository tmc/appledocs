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
// ExampleAffineTransform_Concat demonstrates using Concat on a AffineTransform instance.
// Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result.
func ExampleAffineTransform_Concat() {
	obj := foundation.NewAffineTransform()
	obj.Concat()
	// Output:
	}

// ExampleAffineTransform_Invert demonstrates using Invert on a AffineTransform instance.
// Replaces the receiver’s matrix with its inverse matrix.
func ExampleAffineTransform_Invert() {
	obj := foundation.NewAffineTransform()
	obj.Invert()
	// Output:
	}

// ExampleAffineTransform_Set demonstrates using Set on a AffineTransform instance.
// Sets the current transformation matrix to the receiver’s transformation matrix.
func ExampleAffineTransform_Set() {
	obj := foundation.NewAffineTransform()
	obj.Set()
	// Output:
	}

