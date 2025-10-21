// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMutableCharacterSet

// ExampleNewMutableCharacterSetWithContentsOfFile demonstrates how to create a MutableCharacterSet instance using NewMutableCharacterSetWithContentsOfFile.
// Returns a character set read from the bitmap representation stored in the file a given path.
func ExampleNewMutableCharacterSetWithContentsOfFile() {
	_ = foundation.NewMutableCharacterSetWithContentsOfFile(
		"fName", // fName string
	)
	// Output:
}

// ExampleNewMutableCharacterSetWithRange demonstrates how to create a MutableCharacterSet instance using NewMutableCharacterSetWithRange.
// Returns a character set containing characters with Unicode values in a given range.
func ExampleNewMutableCharacterSetWithRange() {
	_ = foundation.NewMutableCharacterSetWithRange(
		foundation.Range{}, // aRange Range
	)
	// Output:
}


