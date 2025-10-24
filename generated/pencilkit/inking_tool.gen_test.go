// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewInkingTool

// ExampleNewInkingToolWithInkWidth demonstrates how to create a InkingTool instance using NewInkingToolWithInkWidth.
// Create an inking tool with the specified ink and width.
func ExampleNewInkingToolWithInkWidth() {
	_ = pencilkit.NewInkingToolWithInkWidth(
		pencilkit.PKInk{}, // ink PKInk
		0.0,               // width float64
	)
	// Output:
}
