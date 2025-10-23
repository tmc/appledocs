// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXStartCallAction

// ExampleNewCXStartCallActionWithCallUUIDHandle demonstrates how to create a CXStartCallAction instance using NewCXStartCallActionWithCallUUIDHandle.
// Initializes a new action to start a call with the specified UUID to a recipient with the specified handle.
func ExampleNewCXStartCallActionWithCallUUIDHandle() {
	_ = callkit.NewCXStartCallActionWithCallUUIDHandle(
		callkit.UUID{}, // callUUID UUID
		callkit.CXHandle{}, // handle CXHandle
	)
	// Output:
}
// ExampleNewCXStartCallActionWithCoder demonstrates how to create a CXStartCallAction instance using NewCXStartCallActionWithCoder.
// Creates a new action to start a call with data in an unarchiver.
func ExampleNewCXStartCallActionWithCoder() {
	_ = callkit.NewCXStartCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
