// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewFileManager

// ExampleNewFileManagerWithAuthorization demonstrates how to create a FileManager instance using NewFileManagerWithAuthorization.
// Initializes a file manager object that is authorized to perform privileged file system operations.
func ExampleNewFileManagerWithAuthorization() {
	_ = foundation.NewFileManagerWithAuthorization(
		foundation.WorkspaceAuthorization /* not a class type */{}, // authorization WorkspaceAuthorization /* not a class type */
	)
	// Output:
}
