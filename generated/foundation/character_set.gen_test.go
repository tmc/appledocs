// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewCharacterSet

// ExampleNewCharacterSetWithBitmapRepresentation demonstrates how to create a CharacterSet instance using NewCharacterSetWithBitmapRepresentation.
// Returns a character set containing characters determined by a given bitmap representation.
func ExampleNewCharacterSetWithBitmapRepresentation() {
	_ = foundation.NewCharacterSetWithBitmapRepresentation(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
// ExampleNewCharacterSetWithCharactersInString demonstrates how to create a CharacterSet instance using NewCharacterSetWithCharactersInString.
// Returns a character set containing the characters in a given string.
func ExampleNewCharacterSetWithCharactersInString() {
	_ = foundation.NewCharacterSetWithCharactersInString(
		"aString", // aString string
	)
	// Output:
}
// ExampleNewCharacterSetWithCoder demonstrates how to create a CharacterSet instance using NewCharacterSetWithCoder.
func ExampleNewCharacterSetWithCoder() {
	_ = foundation.NewCharacterSetWithCoder(
		foundation.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewCharacterSetWithContentsOfFile demonstrates how to create a CharacterSet instance using NewCharacterSetWithContentsOfFile.
// Returns a character set read from the bitmap representation stored in the file a given path.
func ExampleNewCharacterSetWithContentsOfFile() {
	_ = foundation.NewCharacterSetWithContentsOfFile(
		"fName", // fName string
	)
	// Output:
}
// ExampleNewCharacterSetWithRange demonstrates how to create a CharacterSet instance using NewCharacterSetWithRange.
// Returns a character set containing characters with Unicode values in a given range.
func ExampleNewCharacterSetWithRange() {
	_ = foundation.NewCharacterSetWithRange(
		foundation.Range{}, // aRange Range
	)
	// Output:
}
