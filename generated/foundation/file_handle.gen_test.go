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


// ExampleNewFileHandleForUpdatingAtPath demonstrates how to create a FileHandle instance using NewFileHandleForUpdatingAtPath.
// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
func ExampleNewFileHandleForUpdatingAtPath() {
	_ = foundation.NewFileHandleForUpdatingAtPath(
		"/tmp/test", // path string
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






