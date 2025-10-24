// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSContext

// ExampleNewSContextWithTypeIdentifierTitle demonstrates how to create a SContext instance using NewSContextWithTypeIdentifierTitle.
// Initializes a new context.
func ExampleNewSContextWithTypeIdentifierTitle() {
	_ = classkit.NewSContextWithTypeIdentifierTitle(
		classkit.SContextType{}, // type SContextType
		"identifier",            // identifier string
		"title",                 // title string
	)
	// Output:
}
