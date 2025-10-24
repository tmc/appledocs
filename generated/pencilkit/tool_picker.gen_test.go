// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewToolPicker

// ExampleNewToolPicker demonstrates how to create a ToolPicker instance.
// Creates a new tool picker with a default set of tools.
func ExampleNewToolPicker() {
	_ = pencilkit.NewToolPicker()
	// Output:
}

// ExampleNewToolPickerWithToolItems demonstrates how to create a ToolPicker instance using NewToolPickerWithToolItems.
// Creates a new tool picker with the tools you specify.
func ExampleNewToolPickerWithToolItems() {
	_ = pencilkit.NewToolPickerWithToolItems(
		[]pencilkit.ToolPickerItem{}, // items []ToolPickerItem
	)
	// Output:
}
