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
// ExampleNewStringWithContentsOfFile demonstrates how to create a String instance using NewStringWithContentsOfFile.
// Initializes the receiver, a newly allocated   object, by reading data from the file named by  .
func ExampleNewStringWithContentsOfFile() {
	_ = foundation.NewStringWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewStringWithContentsOfFileEncodingError demonstrates how to create a String instance using NewStringWithContentsOfFileEncodingError.
// Returns an   object initialized by reading data from the file at a given path using a given encoding.
func ExampleNewStringWithContentsOfFileEncodingError() {
	_ = foundation.NewStringWithContentsOfFileEncodingError(
		"/tmp/test", // path string
		foundation.StringEncoding{}, // enc StringEncoding
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewStringWithContentsOfFileUsedEncodingError demonstrates how to create a String instance using NewStringWithContentsOfFileUsedEncodingError.
// Returns an   object initialized by reading data from the file at a given path and returns by reference the encoding used to interpret the characters.
func ExampleNewStringWithContentsOfFileUsedEncodingError() {
	_ = foundation.NewStringWithContentsOfFileUsedEncodingError(
		"/tmp/test", // path string
		foundation.StringEncoding{}, // enc StringEncoding
		foundation.NSError{}, // error NSError
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
// ExampleNewStringWithContentsOfURLEncodingError demonstrates how to create a String instance using NewStringWithContentsOfURLEncodingError.
// Returns an   object initialized by reading data from a given URL interpreted using a given encoding.
func ExampleNewStringWithContentsOfURLEncodingError() {
	_ = foundation.NewStringWithContentsOfURLEncodingError(
		foundation.URL{}, // url URL
		foundation.StringEncoding{}, // enc StringEncoding
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewStringWithContentsOfURLUsedEncodingError demonstrates how to create a String instance using NewStringWithContentsOfURLUsedEncodingError.
// Returns an   object initialized by reading data from a given URL and returns by reference the encoding used to interpret the data.
func ExampleNewStringWithContentsOfURLUsedEncodingError() {
	_ = foundation.NewStringWithContentsOfURLUsedEncodingError(
		foundation.URL{}, // url URL
		foundation.StringEncoding{}, // enc StringEncoding
		foundation.NSError{}, // error NSError
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
		"format", // format string
	)
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
