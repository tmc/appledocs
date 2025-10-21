// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewDocumentController

// ExampleNewDocumentController demonstrates how to create a DocumentController instance.
// This method is the designated initializer for  .
func ExampleNewDocumentController() {
	_ = appkit.NewDocumentController()
	// Output:
}
// ExampleNewDocumentControllerWithCoder demonstrates how to create a DocumentController instance using NewDocumentControllerWithCoder.
// This method initializes a new NSDocumentController from the coder.
func ExampleNewDocumentControllerWithCoder() {
	_ = appkit.NewDocumentControllerWithCoder(
		appkit.Coder{}, // coder Coder
	)
	// Output:
}
