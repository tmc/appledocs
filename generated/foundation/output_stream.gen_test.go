// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewOutputStream

// ExampleNewOutputStreamToMemory demonstrates how to create a OutputStream instance using NewOutputStreamToMemory.
// Returns an initialized output stream that will write to memory.
func ExampleNewOutputStreamToMemory() {
	_ = foundation.NewOutputStreamToMemory()
	// Output:
}
// ExampleNewOutputStreamWithURLAppend demonstrates how to create a OutputStream instance using NewOutputStreamWithURLAppend.
// Returns an initialized output stream for writing to a specified URL.
func ExampleNewOutputStreamWithURLAppend() {
	_ = foundation.NewOutputStreamWithURLAppend(
		foundation.URL{}, // url URL
		false, // shouldAppend bool
	)
	// Output:
}
