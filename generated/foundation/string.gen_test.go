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
		foundation.unichar{}, // characters unichar
		0, // length uint
	)
	// Output:
}
// ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone demonstrates how to create a String instance using NewStringWithCharactersNoCopyLengthFreeWhenDone.
// Returns an initialized   object that contains a given number of characters from a given C array of UTF-16 code units.
func ExampleNewStringWithCharactersNoCopyLengthFreeWhenDone() {
	_ = foundation.NewStringWithCharactersNoCopyLengthFreeWhenDone(
		foundation.unichar{}, // characters unichar
		0, // length uint
		false, // freeBuffer bool
	)
	// Output:
}
// ExampleNewStringWithCoder demonstrates how to create a String instance using NewStringWithCoder.
func ExampleNewStringWithCoder() {
	_ = foundation.NewStringWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewStringWithContentsOfURL demonstrates how to create a String instance using NewStringWithContentsOfURL.
// Initializes the receiver, a newly allocated   object, by reading data from the location named by a given URL.
func ExampleNewStringWithContentsOfURL() {
	_ = foundation.NewStringWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
