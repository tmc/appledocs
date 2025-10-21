// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewToolPickerEraserItem

// ExampleNewToolPickerEraserItemWithEraserType demonstrates how to create a ToolPickerEraserItem instance using NewToolPickerEraserItemWithEraserType.
// Creates a new eraser item.
func ExampleNewToolPickerEraserItemWithEraserType() {
	_ = pencilkit.NewToolPickerEraserItemWithEraserType(
		pencilkit.EraserType{}, // eraserType EraserType
	)
	// Output:
}
// ExampleNewToolPickerEraserItemWithEraserTypeWidth demonstrates how to create a ToolPickerEraserItem instance using NewToolPickerEraserItemWithEraserTypeWidth.
// Creates a new eraser item with the specified width.
func ExampleNewToolPickerEraserItemWithEraserTypeWidth() {
	_ = pencilkit.NewToolPickerEraserItemWithEraserTypeWidth(
		pencilkit.EraserType{}, // eraserType EraserType
		0.0, // width float64
	)
	// Output:
}
