// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewStoryboard

// ExampleStoryboard_InstantiateInitialController demonstrates using InstantiateInitialController on a Storyboard instance.
// Creates the initial view controller or window controller from a storyboard.
func ExampleStoryboard_InstantiateInitialController() {
	obj := appkit.NewStoryboard()
	_ = obj.InstantiateInitialController()
	// Output:
	}

