// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewGraphicsContext

// ExampleNewGraphicsContextWithCGContextFlipped demonstrates how to create a GraphicsContext instance using NewGraphicsContextWithCGContextFlipped.
// Creates a new graphics context from the specified Core Graphics context and the initial flipped state.
func ExampleNewGraphicsContextWithCGContextFlipped() {
	_ = appkit.NewGraphicsContextWithCGContextFlipped(
		appkit.ContextRef /* not a class type */{}, // graphicsPort ContextRef /* not a class type */
		false, // initialFlippedState bool
	)
	// Output:
}
// ExampleGraphicsContext_FlushGraphics demonstrates using FlushGraphics on a GraphicsContext instance.
// Forces any buffered operations or data to be sent to the graphics context’s destination.
func ExampleGraphicsContext_FlushGraphics() {
	obj := appkit.NewGraphicsContext()
	obj.FlushGraphics()
	// Output:
	}

// ExampleGraphicsContext_RestoreGraphicsState demonstrates using RestoreGraphicsState on a GraphicsContext instance.
// Removes the context’s graphics state from the top of the graphics state stack and makes the next graphics state the current graphics state.
func ExampleGraphicsContext_RestoreGraphicsState() {
	obj := appkit.NewGraphicsContext()
	obj.RestoreGraphicsState()
	// Output:
	}

// ExampleGraphicsContext_SaveGraphicsState demonstrates using SaveGraphicsState on a GraphicsContext instance.
// Saves the current graphics state and creates a new graphics state on the top of the stack.
func ExampleGraphicsContext_SaveGraphicsState() {
	obj := appkit.NewGraphicsContext()
	obj.SaveGraphicsState()
	// Output:
	}

