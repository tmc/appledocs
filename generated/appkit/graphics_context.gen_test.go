// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGraphicsContext

// ExampleNewGraphicsContextWithBitmapImageRep demonstrates how to create a GraphicsContext instance using NewGraphicsContextWithBitmapImageRep.
// Creates a new graphics context using the specified bitmap image representation object as the context destination.
func ExampleNewGraphicsContextWithBitmapImageRep() {
	_ = appkit.NewGraphicsContextWithBitmapImageRep(
		appkit.NSBitmapImageRep{}, // bitmapRep NSBitmapImageRep
	)
	// Output:
}
// ExampleNewGraphicsContextWithWindow demonstrates how to create a GraphicsContext instance using NewGraphicsContextWithWindow.
// Creates a new graphics context for drawing into a window.
func ExampleNewGraphicsContextWithWindow() {
	_ = appkit.NewGraphicsContextWithWindow(
		appkit.NSWindow{}, // window NSWindow
	)
	// Output:
}
