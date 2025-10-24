// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColorWell

// ExampleNewColorWellWithStyle demonstrates how to create a ColorWell instance using NewColorWellWithStyle.
// Creates a color well that adopts the specified appearance style.
func ExampleNewColorWellWithStyle() {
	_ = appkit.NewColorWellWithStyle(
		appkit.ColorWellStyle{}, // style ColorWellStyle
	)
	// Output:
}
// ExampleColorWell_Deactivate demonstrates using Deactivate on a ColorWell instance.
// Deactivates the color well.
func ExampleColorWell_Deactivate() {
	obj := appkit.NewColorWell()
	obj.Deactivate()
	// Output:
	}

