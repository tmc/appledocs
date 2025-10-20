// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewString







// ExampleNewStringWithFormatLocale demonstrates how to create a String instance using NewStringWithFormatLocale.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted according to given locale.
func ExampleNewStringWithFormatLocale() {
	_ = foundation.NewStringWithFormatLocale(
		"format", // format string
		0, // locale objc.ID
	)
	// Output:
}


// ExampleNewStringWithContentsOfFile demonstrates how to create a String instance using NewStringWithContentsOfFile.
// Initializes the receiver, a newly allocated   object, by reading data from the file named by  .
func ExampleNewStringWithContentsOfFile() {
	_ = foundation.NewStringWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}

// ExampleNewString demonstrates how to create a String instance.
// Returns an initialized   object that contains no characters.
func ExampleNewString() {
	_ = foundation.NewString()
	// Output:
}




// ExampleNewStringWithString demonstrates how to create a String instance using NewStringWithString.
// Returns an   object initialized by copying the characters from another given string.
func ExampleNewStringWithString() {
	_ = foundation.NewStringWithString(
		"aString", // aString string
	)
	// Output:
}






// ExampleNewStringWithFormat demonstrates how to create a String instance using NewStringWithFormat.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted.
func ExampleNewStringWithFormat() {
	_ = foundation.NewStringWithFormat(
		"format", // format string
	)
	// Output:
}












