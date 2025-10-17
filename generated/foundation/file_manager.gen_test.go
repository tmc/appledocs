// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)


// ExampleNewFileManagerWithAuthorization demonstrates how to create a FileManager instance using NewFileManagerWithAuthorization.
// Initializes a file manager object that is authorized to perform privileged file system operations.
func ExampleNewFileManagerWithAuthorization() {
	_ = foundation.NewFileManagerWithAuthorization(
		nil, // authorization unsafe.Pointer
	)
	// Output:
}


