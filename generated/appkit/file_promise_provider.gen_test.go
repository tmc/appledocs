// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewFilePromiseProvider

// ExampleNewFilePromiseProviderWithFileTypeDelegate demonstrates how to create a FilePromiseProvider instance using NewFilePromiseProviderWithFileTypeDelegate.
// Initializes a file promise provider for a certain file type.
func ExampleNewFilePromiseProviderWithFileTypeDelegate() {
	_ = appkit.NewFilePromiseProviderWithFileTypeDelegate(
		"fileType", // fileType string
		0,          // delegate objc.ID
	)
	// Output:
}
