// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit_test

import (
	"github.com/tmc/appledocs/generated/screencapturekit"
)

// Suppress unused import errors
var _ = screencapturekit.NewContentFilter

// ExampleNewContentFilterWithDisplayExcludingApplicationsExceptingWindows demonstrates how to create a ContentFilter instance using NewContentFilterWithDisplayExcludingApplicationsExceptingWindows.
// Creates a filter that captures a display, excluding windows of the specified apps.
func ExampleNewContentFilterWithDisplayExcludingApplicationsExceptingWindows() {
	_ = screencapturekit.NewContentFilterWithDisplayExcludingApplicationsExceptingWindows(
		screencapturekit.SCDisplay{},            // display SCDisplay
		[]screencapturekit.RunningApplication{}, // applications []RunningApplication
		[]screencapturekit.Window{},             // exceptingWindows []Window
	)
	// Output:
}

// ExampleNewContentFilterWithDisplayExcludingWindows demonstrates how to create a ContentFilter instance using NewContentFilterWithDisplayExcludingWindows.
// Creates a filter that captures the contents of a display, excluding the specified windows.
func ExampleNewContentFilterWithDisplayExcludingWindows() {
	_ = screencapturekit.NewContentFilterWithDisplayExcludingWindows(
		screencapturekit.SCDisplay{}, // display SCDisplay
		[]screencapturekit.Window{},  // excluded []Window
	)
	// Output:
}
