// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewProgress

// ExampleNewProgressWithParentUserInfo demonstrates how to create a Progress instance using NewProgressWithParentUserInfo.
// Creates a new progress instance.
func ExampleNewProgressWithParentUserInfo() {
	_ = foundation.NewProgressWithParentUserInfo(
		foundation.NSProgress{}, // parentProgressOrNil NSProgress
		foundation.IDictionary{}, // userInfoOrNil IDictionary
	)
	// Output:
}
