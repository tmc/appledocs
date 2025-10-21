// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewOutputStream



// ExampleNewOutputStreamToFileAtPathAppend demonstrates how to create a OutputStream instance using NewOutputStreamToFileAtPathAppend.
// Returns an initialized output stream for writing to a specified file.
func ExampleNewOutputStreamToFileAtPathAppend() {
	_ = foundation.NewOutputStreamToFileAtPathAppend(
		"/tmp/test", // path string
		false, // shouldAppend bool
	)
	// Output:
}

// ExampleNewOutputStreamToMemory demonstrates how to create a OutputStream instance using NewOutputStreamToMemory.
// Returns an initialized output stream that will write to memory.
func ExampleNewOutputStreamToMemory() {
	_ = foundation.NewOutputStreamToMemory()
	// Output:
}



