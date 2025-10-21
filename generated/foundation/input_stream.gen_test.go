// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewInputStream

// ExampleNewInputStreamWithData demonstrates how to create a InputStream instance using NewInputStreamWithData.
// Initializes and returns an   object for reading from a given   object.
func ExampleNewInputStreamWithData() {
	_ = foundation.NewInputStreamWithData(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
// ExampleNewInputStreamWithURL demonstrates how to create a InputStream instance using NewInputStreamWithURL.
// Initializes and returns an   object that reads data from the file at a given URL.
func ExampleNewInputStreamWithURL() {
	_ = foundation.NewInputStreamWithURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
