// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewEraserTool

// ExampleNewEraserToolWithEraserType demonstrates how to create a EraserTool instance using NewEraserToolWithEraserType.
// Creates an eraser tool object that removes objects wholly or partially from a canvas view.
func ExampleNewEraserToolWithEraserType() {
	_ = pencilkit.NewEraserToolWithEraserType(
		pencilkit.EraserType{}, // eraserType EraserType
	)
	// Output:
}

// ExampleNewEraserToolWithEraserTypeWidth demonstrates how to create a EraserTool instance using NewEraserToolWithEraserTypeWidth.
func ExampleNewEraserToolWithEraserTypeWidth() {
	_ = pencilkit.NewEraserToolWithEraserTypeWidth(
		pencilkit.EraserType{}, // eraserType EraserType
		0.0,                    // width float64
	)
	// Output:
}
