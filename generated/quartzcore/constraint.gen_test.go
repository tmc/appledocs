// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore_test

import (
	"github.com/tmc/appledocs/generated/quartzcore"
)

// Suppress unused import errors
var _ = quartzcore.NewConstraint

// ExampleNewConstraintWithAttributeRelativeToAttribute demonstrates how to create a Constraint instance using NewConstraintWithAttributeRelativeToAttribute.
// Creates and returns an   object with the specified parameters.
func ExampleNewConstraintWithAttributeRelativeToAttribute() {
	_ = quartzcore.NewConstraintWithAttributeRelativeToAttribute(
		quartzcore.ConstraintAttribute{}, // attr ConstraintAttribute
		"srcId", // srcId string
		quartzcore.ConstraintAttribute{}, // srcAttr ConstraintAttribute
	)
	// Output:
}
// ExampleNewConstraintWithAttributeRelativeToAttributeOffset demonstrates how to create a Constraint instance using NewConstraintWithAttributeRelativeToAttributeOffset.
// Creates and returns an   object with the specified parameters.
func ExampleNewConstraintWithAttributeRelativeToAttributeOffset() {
	_ = quartzcore.NewConstraintWithAttributeRelativeToAttributeOffset(
		quartzcore.ConstraintAttribute{}, // attr ConstraintAttribute
		"srcId", // srcId string
		quartzcore.ConstraintAttribute{}, // srcAttr ConstraintAttribute
		0.0, // c float64
	)
	// Output:
}
// ExampleNewConstraintWithAttributeRelativeToAttributeScaleOffset demonstrates how to create a Constraint instance using NewConstraintWithAttributeRelativeToAttributeScaleOffset.
// Returns an   object with the specified parameters. Designated initializer.
func ExampleNewConstraintWithAttributeRelativeToAttributeScaleOffset() {
	_ = quartzcore.NewConstraintWithAttributeRelativeToAttributeScaleOffset(
		quartzcore.ConstraintAttribute{}, // attr ConstraintAttribute
		"srcId", // srcId string
		quartzcore.ConstraintAttribute{}, // srcAttr ConstraintAttribute
		0.0, // m float64
		0.0, // c float64
	)
	// Output:
}
