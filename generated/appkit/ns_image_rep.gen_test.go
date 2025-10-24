// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewImageRep

// ExampleNewImageRep demonstrates how to create a ImageRep instance.
// Creates and returns an image representation object.
func ExampleNewImageRep() {
	_ = appkit.NewImageRep()
	// Output:
}
// ExampleImageRep_Draw demonstrates using Draw on a ImageRep instance.
// Implemented by subclasses to draw the image in the current coordinate system.
func ExampleImageRep_Draw() {
	obj := appkit.NewImageRep()
	_ = obj.Draw()
	// Output:
	}

