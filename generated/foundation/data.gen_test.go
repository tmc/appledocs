// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewDataWithBase64EncodedDataOptions demonstrates how to create a Data instance using NewDataWithBase64EncodedDataOptions.
// Initializes a data object with the given Base64 encoded data.
func ExampleNewDataWithBase64EncodedDataOptions() {
	_ = foundation.NewDataWithBase64EncodedDataOptions(
		nil, // base64Data unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithBase64EncodedStringOptions demonstrates how to create a Data instance using NewDataWithBase64EncodedStringOptions.
// Initializes a data object with the given Base64 encoded string.
func ExampleNewDataWithBase64EncodedStringOptions() {
	_ = foundation.NewDataWithBase64EncodedStringOptions(
		"base64String", // base64String string
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithBytesLength demonstrates how to create a Data instance using NewDataWithBytesLength.
// Initializes a data object filled with a given number of bytes copied from a given buffer.
func ExampleNewDataWithBytesLength() {
	_ = foundation.NewDataWithBytesLength(
		nil, // bytes unsafe.Pointer
		0, // length uint
	)
	// Output:
}

// ExampleNewDataWithBytesNoCopyLength demonstrates how to create a Data instance using NewDataWithBytesNoCopyLength.
// Initializes a data object filled with a given number of bytes of data from a given buffer.
func ExampleNewDataWithBytesNoCopyLength() {
	_ = foundation.NewDataWithBytesNoCopyLength(
		nil, // bytes unsafe.Pointer
		0, // length uint
	)
	// Output:
}

// ExampleNewDataWithBytesNoCopyLengthDeallocator demonstrates how to create a Data instance using NewDataWithBytesNoCopyLengthDeallocator.
// Initializes a data object filled with a given number of bytes of data from a given buffer, with a custom deallocator block.
func ExampleNewDataWithBytesNoCopyLengthDeallocator() {
	_ = foundation.NewDataWithBytesNoCopyLengthDeallocator(
		nil, // bytes unsafe.Pointer
		0, // length uint
		nil, // deallocator unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithContentsOfFileOptionsError demonstrates how to create a Data instance using NewDataWithContentsOfFileOptionsError.
// Initializes a data object with the content of the file at a given path.
func ExampleNewDataWithContentsOfFileOptionsError() {
	_ = foundation.NewDataWithContentsOfFileOptionsError(
		"path", // path string
		nil, // readOptionsMask unsafe.Pointer
		nil, // errorPtr unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithContentsOfMappedFile demonstrates how to create a Data instance using NewDataWithContentsOfMappedFile.
// Initializes a data object with the contents of the mapped file specified by a given path.
func ExampleNewDataWithContentsOfMappedFile() {
	_ = foundation.NewDataWithContentsOfMappedFile(
		"path", // path string
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

// ExampleNewDataWithBytesNoCopyLengthFreeWhenDone demonstrates how to create a Data instance using NewDataWithBytesNoCopyLengthFreeWhenDone.
// Initializes a newly allocated data object by adding the given number of bytes from the given buffer.
func ExampleNewDataWithBytesNoCopyLengthFreeWhenDone() {
	_ = foundation.NewDataWithBytesNoCopyLengthFreeWhenDone(
		nil, // bytes unsafe.Pointer
		0, // length uint
		false, // b bool
	)
	// Output:
}

// ExampleNewDataWithContentsOfFile demonstrates how to create a Data instance using NewDataWithContentsOfFile.
// Initializes a data object with the content of the file at a given path.
func ExampleNewDataWithContentsOfFile() {
	_ = foundation.NewDataWithContentsOfFile(
		"path", // path string
	)
	// Output:
}

// ExampleNewDataWithContentsOfURL demonstrates how to create a Data instance using NewDataWithContentsOfURL.
// Creates a data object from the data at the specified file URL, or returns   if the system can’t create one.
func ExampleNewDataWithContentsOfURL() {
	_ = foundation.NewDataWithContentsOfURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithContentsOfURLOptionsError demonstrates how to create a Data instance using NewDataWithContentsOfURLOptionsError.
// Creates a data object from the data at the provided file URL using specific reading options.
func ExampleNewDataWithContentsOfURLOptionsError() {
	_ = foundation.NewDataWithContentsOfURLOptionsError(
		nil, // url unsafe.Pointer
		nil, // readOptionsMask unsafe.Pointer
		nil, // errorPtr unsafe.Pointer
	)
	// Output:
}

// ExampleNewDataWithData demonstrates how to create a Data instance using NewDataWithData.
// Initializes a data object with the contents of another data object.
func ExampleNewDataWithData() {
	_ = foundation.NewDataWithData(
		nil, // data unsafe.Pointer
	)
	// Output:
}


