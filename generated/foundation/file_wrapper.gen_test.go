// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileWrapper

// ExampleNewFileWrapperDirectoryWithFileWrappers demonstrates how to create a FileWrapper instance using NewFileWrapperDirectoryWithFileWrappers.
// Initializes the receiver as a directory file wrapper, with a given file-wrapper list.
func ExampleNewFileWrapperDirectoryWithFileWrappers() {
	_ = foundation.NewFileWrapperDirectoryWithFileWrappers(
		foundation.IDictionary{}, // childrenByPreferredName IDictionary
	)
	// Output:
}
// ExampleNewFileWrapperRegularFileWithContents demonstrates how to create a FileWrapper instance using NewFileWrapperRegularFileWithContents.
// Initializes the receiver as a regular-file file wrapper.
func ExampleNewFileWrapperRegularFileWithContents() {
	_ = foundation.NewFileWrapperRegularFileWithContents(
		foundation.NSData{}, // contents NSData
	)
	// Output:
}
// ExampleNewFileWrapperSymbolicLinkWithDestination demonstrates how to create a FileWrapper instance using NewFileWrapperSymbolicLinkWithDestination.
// Initializes the receiver as a symbolic-link file wrapper.
func ExampleNewFileWrapperSymbolicLinkWithDestination() {
	_ = foundation.NewFileWrapperSymbolicLinkWithDestination(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewFileWrapperSymbolicLinkWithDestinationURL demonstrates how to create a FileWrapper instance using NewFileWrapperSymbolicLinkWithDestinationURL.
// Initializes the receiver as a symbolic-link file wrapper that links to a specified file.
func ExampleNewFileWrapperSymbolicLinkWithDestinationURL() {
	_ = foundation.NewFileWrapperSymbolicLinkWithDestinationURL(
		foundation.URL{}, // url URL
	)
	// Output:
}
// ExampleNewFileWrapperWithCoder demonstrates how to create a FileWrapper instance using NewFileWrapperWithCoder.
func ExampleNewFileWrapperWithCoder() {
	_ = foundation.NewFileWrapperWithCoder(
		foundation.NSCoder{}, // inCoder NSCoder
	)
	// Output:
}
// ExampleNewFileWrapperWithPath demonstrates how to create a FileWrapper instance using NewFileWrapperWithPath.
// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the path.
func ExampleNewFileWrapperWithPath() {
	_ = foundation.NewFileWrapperWithPath(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewFileWrapperWithSerializedRepresentation demonstrates how to create a FileWrapper instance using NewFileWrapperWithSerializedRepresentation.
// Initializes the receiver as a regular-file file wrapper from given serialized data.
func ExampleNewFileWrapperWithSerializedRepresentation() {
	_ = foundation.NewFileWrapperWithSerializedRepresentation(
		foundation.NSData{}, // serializeRepresentation NSData
	)
	// Output:
}
// ExampleNewFileWrapperWithURLOptionsError demonstrates how to create a FileWrapper instance using NewFileWrapperWithURLOptionsError.
// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL.
func ExampleNewFileWrapperWithURLOptionsError() {
	_ = foundation.NewFileWrapperWithURLOptionsError(
		foundation.URL{}, // url URL
		foundation.FileWrapperReadingOptions{}, // options FileWrapperReadingOptions
		foundation.NSError{}, // outError NSError
	)
	// Output:
}
