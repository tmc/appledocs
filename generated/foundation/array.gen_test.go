// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewArray

// ExampleNewArray demonstrates how to create a Array instance.
// Initializes a newly allocated array.
func ExampleNewArray() {
	_ = foundation.NewArray()
	// Output:
}
// ExampleNewArrayWithCoder demonstrates how to create a Array instance using NewArrayWithCoder.
func ExampleNewArrayWithCoder() {
	_ = foundation.NewArrayWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
// ExampleNewArrayWithContentsOfFile demonstrates how to create a Array instance using NewArrayWithContentsOfFile.
// Initializes a newly allocated array with the contents of the file specified by a given path.
func ExampleNewArrayWithContentsOfFile() {
	_ = foundation.NewArrayWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewArrayWithContentsOfURL demonstrates how to create a Array instance using NewArrayWithContentsOfURL.
// Initializes a newly allocated array with the contents of the location specified by a given URL.
func ExampleNewArrayWithContentsOfURL() {
	_ = foundation.NewArrayWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
// ExampleNewArrayWithContentsOfURLError demonstrates how to create a Array instance using NewArrayWithContentsOfURLError.
func ExampleNewArrayWithContentsOfURLError() {
	_ = foundation.NewArrayWithContentsOfURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
