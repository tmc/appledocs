// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileHandle

// ExampleNewFileHandleForReadingAtPath demonstrates how to create a FileHandle instance using NewFileHandleForReadingAtPath.
// Returns a file handle initialized for reading the file, device, or named socket at the specified path.
func ExampleNewFileHandleForReadingAtPath() {
	_ = foundation.NewFileHandleForReadingAtPath(
		"/tmp/test", // path string
	)
	// Output:
}
// ExampleNewFileHandleForReadingFromURLError demonstrates how to create a FileHandle instance using NewFileHandleForReadingFromURLError.
// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
func ExampleNewFileHandleForReadingFromURLError() {
	_ = foundation.NewFileHandleForReadingFromURLError(
		foundation.URL{}, // url URL
		foundation.NSError{}, // error NSError
	)
	// Output:
}
// ExampleNewFileHandleForUpdatingAtPath demonstrates how to create a FileHandle instance using NewFileHandleForUpdatingAtPath.
// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
func ExampleNewFileHandleForUpdatingAtPath() {
	_ = foundation.NewFileHandleForUpdatingAtPath(
		"/tmp/test", // path string
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
// ExampleNewFileHandleForWritingAtPath demonstrates how to create a FileHandle instance using NewFileHandleForWritingAtPath.
// Returns a file handle initialized for writing to the file, device, or named socket at the specified path.
func ExampleNewFileHandleForWritingAtPath() {
	_ = foundation.NewFileHandleForWritingAtPath(
		"/tmp/test", // path string
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
// ExampleNewFileHandleWithFileDescriptor demonstrates how to create a FileHandle instance using NewFileHandleWithFileDescriptor.
// Creates and returns a file handle object associated with the specified file descriptor.
func ExampleNewFileHandleWithFileDescriptor() {
	_ = foundation.NewFileHandleWithFileDescriptor(
		0, // fd int
	)
	// Output:
}
// ExampleNewFileHandleWithFileDescriptorCloseOnDealloc demonstrates how to create a FileHandle instance using NewFileHandleWithFileDescriptorCloseOnDealloc.
// Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
func ExampleNewFileHandleWithFileDescriptorCloseOnDealloc() {
	_ = foundation.NewFileHandleWithFileDescriptorCloseOnDealloc(
		0, // fd int
		false, // closeopt bool
	)
	// Output:
}
