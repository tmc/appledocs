// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit_test

import (
	"github.com/tmc/appledocs/generated/callkit"
)

// Suppress unused import errors
var _ = callkit.NewCXPlayDTMFCallAction

// ExampleNewCXPlayDTMFCallActionWithCallUUIDDigitsType demonstrates how to create a CXPlayDTMFCallAction instance using NewCXPlayDTMFCallActionWithCallUUIDDigitsType.
// Initializes a new action for a call identified by a given UUID, as well as a specified type and sequence of digits.
func ExampleNewCXPlayDTMFCallActionWithCallUUIDDigitsType() {
	_ = callkit.NewCXPlayDTMFCallActionWithCallUUIDDigitsType(
		callkit.UUID{}, // callUUID UUID
		"digits", // digits string
		callkit.CXPlayDTMFCallActionType{}, // type CXPlayDTMFCallActionType
	)
	// Output:
}
// ExampleNewCXPlayDTMFCallActionWithCoder demonstrates how to create a CXPlayDTMFCallAction instance using NewCXPlayDTMFCallActionWithCoder.
// Creates a new action to play dual-tone multifrequency (DTMF) tones with data in an unarchiver.
func ExampleNewCXPlayDTMFCallActionWithCoder() {
	_ = callkit.NewCXPlayDTMFCallActionWithCoder(
		callkit.Coder{}, // aDecoder Coder
	)
	// Output:
}
