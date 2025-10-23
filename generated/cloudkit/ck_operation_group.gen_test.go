// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit_test

import (
	"github.com/tmc/appledocs/generated/cloudkit"
)

// Suppress unused import errors
var _ = cloudkit.NewCKOperationGroup

// ExampleNewCKOperationGroup demonstrates how to create a CKOperationGroup instance.
// Creates an operation group.
func ExampleNewCKOperationGroup() {
	_ = cloudkit.NewCKOperationGroup()
	// Output:
}
// ExampleNewCKOperationGroupWithCoder demonstrates how to create a CKOperationGroup instance using NewCKOperationGroupWithCoder.
// Creates an operation group from a serialized instance.
func ExampleNewCKOperationGroupWithCoder() {
	_ = cloudkit.NewCKOperationGroupWithCoder(
		cloudkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
