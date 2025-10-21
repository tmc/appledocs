// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileHandle

// ExampleNewFileHandleForReadingFromURLError demonstrates how to create a FileHandle instance using NewFileHandleForReadingFromURLError.
// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
func ExampleNewFileHandleForReadingFromURLError() {
	_ = foundation.NewFileHandleForReadingFromURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewFileHandleForUpdatingURLError demonstrates how to create a FileHandle instance using NewFileHandleForUpdatingURLError.
// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL.
func ExampleNewFileHandleForUpdatingURLError() {
	_ = foundation.NewFileHandleForUpdatingURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewFileHandleForWritingToURLError demonstrates how to create a FileHandle instance using NewFileHandleForWritingToURLError.
// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL.
func ExampleNewFileHandleForWritingToURLError() {
	_ = foundation.NewFileHandleForWritingToURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewFileHandleWithCoder demonstrates how to create a FileHandle instance using NewFileHandleWithCoder.
// Returns a file handle initialized from data in an unarchiver.
func ExampleNewFileHandleWithCoder() {
	_ = foundation.NewFileHandleWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
