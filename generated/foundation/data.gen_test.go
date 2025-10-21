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
