// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewObjectController

// ExampleObjectController_DefaultFetchRequest demonstrates using DefaultFetchRequest on a ObjectController instance.
// Returns the default fetch request used by the receiver.
func ExampleObjectController_DefaultFetchRequest() {
	obj := appkit.NewObjectController()
	_ = obj.DefaultFetchRequest()
	// Output:
	}

// ExampleObjectController_NewObject demonstrates using NewObject on a ObjectController instance.
// Creates and returns a new object of the appropriate class.
func ExampleObjectController_NewObject() {
	obj := appkit.NewObjectController()
	_ = obj.NewObject()
	// Output:
	}

// ExampleObjectController_PrepareContent demonstrates using PrepareContent on a ObjectController instance.
// Typically overridden by subclasses that require additional control over the creation of new objects.
//
// Note: This example is not executed because PrepareContent crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleObjectController_PrepareContent() {
	obj := appkit.NewObjectController()
	obj.PrepareContent()
	}

