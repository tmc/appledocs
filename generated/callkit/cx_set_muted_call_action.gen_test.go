// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXSetMutedCallAction

// ExampleNewCXSetMutedCallActionWithCallUUIDMuted demonstrates how to create a CXSetMutedCallAction instance using NewCXSetMutedCallActionWithCallUUIDMuted.
// Initializes a new action for a call identified by a given UUID, as well as whether the call is muted.
func ExampleNewCXSetMutedCallActionWithCallUUIDMuted() {
	_ = callkit.NewCXSetMutedCallActionWithCallUUIDMuted(
		callkit.UUID{}, // callUUID UUID
		false, // muted bool
	)
	// Output:
}
// ExampleNewCXSetMutedCallActionWithCoder demonstrates how to create a CXSetMutedCallAction instance using NewCXSetMutedCallActionWithCoder.
// Creates a new action for a call with data in an unarchiver.
func ExampleNewCXSetMutedCallActionWithCoder() {
	_ = callkit.NewCXSetMutedCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
