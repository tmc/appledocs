// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXAction

// ExampleNewCXAction demonstrates how to create a CXAction instance.
// Initializes a new telephony action.
func ExampleNewCXAction() {
	_ = callkit.NewCXAction()
	// Output:
}
// ExampleNewCXActionWithCoder demonstrates how to create a CXAction instance using NewCXActionWithCoder.
// Creates a new telephony action with data in an unarchiver.
func ExampleNewCXActionWithCoder() {
	_ = callkit.NewCXActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
