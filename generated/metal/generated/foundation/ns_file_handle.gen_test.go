// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileHandle

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
// ExampleFileHandle_AcceptConnectionInBackgroundAndNotify demonstrates using AcceptConnectionInBackgroundAndNotify on a FileHandle instance.
// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
func ExampleFileHandle_AcceptConnectionInBackgroundAndNotify() {
	obj := foundation.NewFileHandle()
	obj.AcceptConnectionInBackgroundAndNotify()
	// Output:
	}

// ExampleFileHandle_ReadInBackgroundAndNotify demonstrates using ReadInBackgroundAndNotify on a FileHandle instance.
// Reads from the file or communications channel in the background and posts a notification when finished.
func ExampleFileHandle_ReadInBackgroundAndNotify() {
	obj := foundation.NewFileHandle()
	obj.ReadInBackgroundAndNotify()
	// Output:
	}

// ExampleFileHandle_ReadToEndOfFileInBackgroundAndNotify demonstrates using ReadToEndOfFileInBackgroundAndNotify on a FileHandle instance.
// Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
func ExampleFileHandle_ReadToEndOfFileInBackgroundAndNotify() {
	obj := foundation.NewFileHandle()
	obj.ReadToEndOfFileInBackgroundAndNotify()
	// Output:
	}

// ExampleFileHandle_WaitForDataInBackgroundAndNotify demonstrates using WaitForDataInBackgroundAndNotify on a FileHandle instance.
// Asynchronously checks to see if data is available.
func ExampleFileHandle_WaitForDataInBackgroundAndNotify() {
	obj := foundation.NewFileHandle()
	obj.WaitForDataInBackgroundAndNotify()
	// Output:
	}

