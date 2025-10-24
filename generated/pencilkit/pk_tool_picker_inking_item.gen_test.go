// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit_test

import (
	"github.com/tmc/appledocs/generated/pencilkit"
)

// Suppress unused import errors
var _ = pencilkit.NewToolPickerInkingItem

// ExampleNewToolPickerInkingItemWithInkType demonstrates how to create a ToolPickerInkingItem instance using NewToolPickerInkingItemWithInkType.
// Create a new tool picker item with a  .
func ExampleNewToolPickerInkingItemWithInkType() {
	_ = pencilkit.NewToolPickerInkingItemWithInkType(
		pencilkit.InkType /* not a class type */{}, // inkType InkType /* not a class type */
	)
	// Output:
}
// ExampleNewToolPickerInkingItemWithInkTypeWidth demonstrates how to create a ToolPickerInkingItem instance using NewToolPickerInkingItemWithInkTypeWidth.
// Create a new tool picker item with a  .
func ExampleNewToolPickerInkingItemWithInkTypeWidth() {
	_ = pencilkit.NewToolPickerInkingItemWithInkTypeWidth(
		pencilkit.InkType /* not a class type */{}, // inkType InkType /* not a class type */
		100.0, // width float64
	)
	// Output:
}
