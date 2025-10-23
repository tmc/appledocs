// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXCallAction

// ExampleNewCXCallActionWithCallUUID demonstrates how to create a CXCallAction instance using NewCXCallActionWithCallUUID.
// Initializes a new action for a call identified by a given UUID.
func ExampleNewCXCallActionWithCallUUID() {
	_ = callkit.NewCXCallActionWithCallUUID(
		callkit.UUID{}, // callUUID UUID
	)
	// Output:
}
// ExampleNewCXCallActionWithCoder demonstrates how to create a CXCallAction instance using NewCXCallActionWithCoder.
// Creates a new action for a call with data in an unarchiver.
func ExampleNewCXCallActionWithCoder() {
	_ = callkit.NewCXCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
