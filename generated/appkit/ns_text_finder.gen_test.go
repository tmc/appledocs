// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewTextFinder

// ExampleNewTextFinder demonstrates how to create a TextFinder instance.
// Initializes and returns a new   instance.
func ExampleNewTextFinder() {
	_ = appkit.NewTextFinder()
	// Output:
}
// ExampleTextFinder_CancelFindIndicator demonstrates using CancelFindIndicator on a TextFinder instance.
// Cancels the find indicator immediately.
func ExampleTextFinder_CancelFindIndicator() {
	obj := appkit.NewTextFinder()
	obj.CancelFindIndicator()
	// Output:
	}

// ExampleTextFinder_NoteClientStringWillChange demonstrates using NoteClientStringWillChange on a TextFinder instance.
// Invoke this method when the searched content will change.
func ExampleTextFinder_NoteClientStringWillChange() {
	obj := appkit.NewTextFinder()
	obj.NoteClientStringWillChange()
	// Output:
	}

