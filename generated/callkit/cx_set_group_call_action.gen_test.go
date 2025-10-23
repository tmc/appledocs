// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXSetGroupCallAction

// ExampleNewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith demonstrates how to create a CXSetGroupCallAction instance using NewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith.
// Initializes a new action for a call identified by a given UUID, as well as a call to group with identified by another UUID.
func ExampleNewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith() {
	_ = callkit.NewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith(
		callkit.UUID{}, // callUUID UUID
		callkit.UUID{}, // callUUIDToGroupWith UUID
	)
	// Output:
}
// ExampleNewCXSetGroupCallActionWithCoder demonstrates how to create a CXSetGroupCallAction instance using NewCXSetGroupCallActionWithCoder.
// Creates a new action to group calls with data in an unarchiver.
func ExampleNewCXSetGroupCallActionWithCoder() {
	_ = callkit.NewCXSetGroupCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
