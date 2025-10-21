// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewDrawing

// ExampleNewDrawing demonstrates how to create a Drawing instance.
// Creates an empty drawing object.
func ExampleNewDrawing() {
	_ = pencilkit.NewDrawing()
	// Output:
}
// ExampleNewDrawingWithStrokes demonstrates how to create a Drawing instance using NewDrawingWithStrokes.
// Creates a drawing object with the strokes you supply.
func ExampleNewDrawingWithStrokes() {
	_ = pencilkit.NewDrawingWithStrokes(
		[]pencilkit.Stroke{}, // strokes []Stroke
	)
	// Output:
}
