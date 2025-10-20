// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewData




// ExampleNewDataWithContentsOfFile demonstrates how to create a Data instance using NewDataWithContentsOfFile.
// Initializes a data object with the content of the file at a given path.
func ExampleNewDataWithContentsOfFile() {
	_ = foundation.NewDataWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}


// ExampleNewDataWithContentsOfMappedFile demonstrates how to create a Data instance using NewDataWithContentsOfMappedFile.
// Initializes a data object with the contents of the mapped file specified by a given path.
func ExampleNewDataWithContentsOfMappedFile() {
	_ = foundation.NewDataWithContentsOfMappedFile(
		"/tmp/test", // path string
	)
	// Output:
}



// ExampleNewDataWithBase64Encoding demonstrates how to create a Data instance using NewDataWithBase64Encoding.
// Initializes a data object initialized with the given Base64 encoded string.
func ExampleNewDataWithBase64Encoding() {
	_ = foundation.NewDataWithBase64Encoding(
		"base64String", // base64String string
	)
	// Output:
}







