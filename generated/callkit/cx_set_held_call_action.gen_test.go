// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXSetHeldCallAction

// ExampleNewCXSetHeldCallActionWithCallUUIDOnHold demonstrates how to create a CXSetHeldCallAction instance using NewCXSetHeldCallActionWithCallUUIDOnHold.
// Initializes a new action for a call identified by a given UUID, as well as whether the call is on hold.
func ExampleNewCXSetHeldCallActionWithCallUUIDOnHold() {
	_ = callkit.NewCXSetHeldCallActionWithCallUUIDOnHold(
		callkit.UUID{}, // callUUID UUID
		false, // onHold bool
	)
	// Output:
}
// ExampleNewCXSetHeldCallActionWithCoder demonstrates how to create a CXSetHeldCallAction instance using NewCXSetHeldCallActionWithCoder.
// Creates a new action to place a call on hold with data in an unarchiver.
func ExampleNewCXSetHeldCallActionWithCoder() {
	_ = callkit.NewCXSetHeldCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
