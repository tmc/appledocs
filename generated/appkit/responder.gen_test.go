// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewResponder

// ExampleNewResponder demonstrates how to create a Responder instance.
// Creates a new responder object.
func ExampleNewResponder() {
	_ = appkit.NewResponder()
	// Output:
}
// ExampleNewResponderWithCoder demonstrates how to create a Responder instance using NewResponderWithCoder.
// Creates a new responder object with data in an unarchiver.
func ExampleNewResponderWithCoder() {
	_ = appkit.NewResponderWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
