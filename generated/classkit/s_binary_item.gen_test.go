// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSBinaryItem

// ExampleNewSBinaryItemWithIdentifierTitleType demonstrates how to create a SBinaryItem instance using NewSBinaryItemWithIdentifierTitleType.
// Initializes a new binary activity item of the given type.
func ExampleNewSBinaryItemWithIdentifierTitleType() {
	_ = classkit.NewSBinaryItemWithIdentifierTitleType(
		"identifier", // identifier string
		"title", // title string
		classkit.CLSBinaryValueType{}, // valueType CLSBinaryValueType
	)
	// Output:
}
