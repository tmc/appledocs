// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewData

// ExampleNewDataWithBase64EncodedDataOptions demonstrates how to create a Data instance using NewDataWithBase64EncodedDataOptions.
// Initializes a data object with the given Base64 encoded data.
func ExampleNewDataWithBase64EncodedDataOptions() {
	_ = foundation.NewDataWithBase64EncodedDataOptions(
		foundation.NSData{}, // base64Data NSData
		foundation.DataBase64DecodingOptions{}, // options DataBase64DecodingOptions
	)
	// Output:
}
// ExampleNewDataWithBase64EncodedStringOptions demonstrates how to create a Data instance using NewDataWithBase64EncodedStringOptions.
// Initializes a data object with the given Base64 encoded string.
func ExampleNewDataWithBase64EncodedStringOptions() {
	_ = foundation.NewDataWithBase64EncodedStringOptions(
		"base64String", // base64String string
		foundation.DataBase64DecodingOptions{}, // options DataBase64DecodingOptions
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
// ExampleNewDataWithContentsOfFile demonstrates how to create a Data instance using NewDataWithContentsOfFile.
// Initializes a data object with the content of the file at a given path.
func ExampleNewDataWithContentsOfFile() {
	_ = foundation.NewDataWithContentsOfFile(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewDataWithContentsOfFileOptionsError demonstrates how to create a Data instance using NewDataWithContentsOfFileOptionsError.
// Initializes a data object with the content of the file at a given path.
func ExampleNewDataWithContentsOfFileOptionsError() {
	_ = foundation.NewDataWithContentsOfFileOptionsError(
		"/tmp/test", // path string
		foundation.DataReadingOptions{}, // readOptionsMask DataReadingOptions
		foundation.NSError{}, // errorPtr NSError
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
// ExampleNewDataWithContentsOfURL demonstrates how to create a Data instance using NewDataWithContentsOfURL.
// Creates a data object from the data at the specified file URL, or returns   if the system can’t create one.
func ExampleNewDataWithContentsOfURL() {
	_ = foundation.NewDataWithContentsOfURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
// ExampleNewDataWithContentsOfURLOptionsError demonstrates how to create a Data instance using NewDataWithContentsOfURLOptionsError.
// Creates a data object from the data at the provided file URL using specific reading options.
func ExampleNewDataWithContentsOfURLOptionsError() {
	_ = foundation.NewDataWithContentsOfURLOptionsError(
		foundation.URL{}, // url URL
		foundation.DataReadingOptions{}, // readOptionsMask DataReadingOptions
		foundation.NSError{}, // errorPtr NSError
	)
	// Output:
}
// ExampleNewDataWithData demonstrates how to create a Data instance using NewDataWithData.
// Initializes a data object with the contents of another data object.
func ExampleNewDataWithData() {
	_ = foundation.NewDataWithData(
		foundation.NSData{}, // data NSData
	)
	// Output:
}
