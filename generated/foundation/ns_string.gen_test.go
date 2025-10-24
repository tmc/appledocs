// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewString

// ExampleNewString demonstrates how to create a String instance.
// Returns an initialized   object that contains no characters.
func ExampleNewString() {
	_ = foundation.NewString()
	// Output:
}
// ExampleNewStringWithCharactersLength demonstrates how to create a String instance using NewStringWithCharactersLength.
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units.
func ExampleNewStringWithCharactersLength() {
	_ = foundation.NewStringWithCharactersLength(
		foundation.Unichar /* typedef */{}, // characters Unichar /* typedef */
		0, // length uint
	)
	// Output:
}
// ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone demonstrates how to create a String instance using NewStringWithCharactersNoCopyLengthFreeWhenDone.
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units.
func ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone() {
	_ = foundation.NewStringWithCharactersNoCopyLengthFreeWhenDone(
		foundation.Unichar /* typedef */{}, // characters Unichar /* typedef */
		0, // length uint
		false, // freeBuffer bool
	)
	// Output:
}
// ExampleString_PropertyList demonstrates using PropertyList on a String instance.
// Parses the receiver as a text representation of a property list, returning an  ,  ,  , or   object, according to the topmost element.
func ExampleString_PropertyList() {
	obj := foundation.NewString()
	_ = obj.PropertyList()
	// Output:
	}

// ExampleString_PropertyListFromStringsFileFormat demonstrates using PropertyListFromStringsFileFormat on a String instance.
// Returns a dictionary object initialized with the keys and values found in the receiver.
func ExampleString_PropertyListFromStringsFileFormat() {
	obj := foundation.NewString()
	_ = obj.PropertyListFromStringsFileFormat()
	// Output:
	}

