// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewProgressWithParentUserInfo demonstrates how to create a Progress instance using NewProgressWithParentUserInfo.
// Creates a new progress instance.
func ExampleNewProgressWithParentUserInfo() {
	_ = foundation.NewProgressWithParentUserInfo(
		nil, // parentProgressOrNil unsafe.Pointer
		nil, // userInfoOrNil unsafe.Pointer
	)
	// Output:
}

// ExampleNewProgressWithTotalUnitCount demonstrates how to create a Progress instance using NewProgressWithTotalUnitCount.
// Creates and returns a progress instance.
func ExampleNewProgressWithTotalUnitCount() {
	_ = foundation.NewProgressWithTotalUnitCount(
		nil, // unitCount unsafe.Pointer
	)
	// Output:
}


