// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewController

// ExampleNewController demonstrates how to create a Controller instance.
func ExampleNewController() {
	_ = appkit.NewController()
	// Output:
}
// ExampleController_CommitEditing demonstrates using CommitEditing on a Controller instance.
// Attempts to commit any pending edits.
//
// Note: This example is not executed because CommitEditing crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleController_CommitEditing() {
	obj := appkit.NewController()
	_ = obj.CommitEditing()
	}

// ExampleController_DiscardEditing demonstrates using DiscardEditing on a Controller instance.
// Discards any pending changes by registered editors.
//
// Note: This example is not executed because DiscardEditing crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleController_DiscardEditing() {
	obj := appkit.NewController()
	obj.DiscardEditing()
	}

