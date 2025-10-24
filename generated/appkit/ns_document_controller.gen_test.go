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
// ExampleDocumentController_StandardShareMenuItem demonstrates using StandardShareMenuItem on a DocumentController instance.
// Returns a menu item that your app uses for sharing the current document.
func ExampleDocumentController_StandardShareMenuItem() {
	obj := appkit.NewDocumentController()
	_ = obj.StandardShareMenuItem()
	// Output:
	}

// ExampleDocumentController_URLsFromRunningOpenPanel demonstrates using URLsFromRunningOpenPanel on a DocumentController instance.
// An array of URLs that correspond to the selected files in a running Open dialog.
func ExampleDocumentController_URLsFromRunningOpenPanel() {
	obj := appkit.NewDocumentController()
	_ = obj.URLsFromRunningOpenPanel()
	// Output:
	}

