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
// ExampleNewStringWithCoder demonstrates how to create a String instance using NewStringWithCoder.
func ExampleNewStringWithCoder() {
	_ = foundation.NewStringWithCoder(
		foundation.Coder{}, // coder Coder
	)
	// Output:
}
// ExampleNewStringWithContentsOfFile demonstrates how to create a String instance using NewStringWithContentsOfFile.
// Initializes the receiver, a newly allocated   object, by reading data from the file named by  .
func ExampleNewStringWithContentsOfFile() {
	_ = foundation.NewStringWithContentsOfFile(
		foundation.NSString{}, // path NSString
	)
	// Output:
}
// ExampleNewStringWithContentsOfURL demonstrates how to create a String instance using NewStringWithContentsOfURL.
// Initializes the receiver, a newly allocated   object, by reading data from the location named by a given URL.
func ExampleNewStringWithContentsOfURL() {
	_ = foundation.NewStringWithContentsOfURL(
		foundation.NSURL{}, // url NSURL
	)
	// Output:
}
// ExampleNewStringWithDataEncoding demonstrates how to create a String instance using NewStringWithDataEncoding.
// Returns an   object initialized by converting given data into UTF-16 code units using a given encoding.
func ExampleNewStringWithDataEncoding() {
	_ = foundation.NewStringWithDataEncoding(
		foundation.NSData{}, // data NSData
		foundation.StringEncoding{}, // encoding StringEncoding
	)
	// Output:
}
// ExampleNewStringWithFormat demonstrates how to create a String instance using NewStringWithFormat.
// Returns an   object initialized by using a given format string as a template into which the remaining argument values are substituted.
func ExampleNewStringWithFormat() {
	_ = foundation.NewStringWithFormat(
		foundation.NSString{}, // format NSString
	)
	// Output:
}
// ExampleNewStringWithString demonstrates how to create a String instance using NewStringWithString.
// Returns an   object initialized by copying the characters from another given string.
func ExampleNewStringWithString() {
	_ = foundation.NewStringWithString(
		foundation.NSString{}, // aString NSString
	)
	// Output:
}

