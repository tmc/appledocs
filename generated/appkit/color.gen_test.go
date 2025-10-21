// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewColor

// ExampleNewColorForControlTint demonstrates how to create a Color instance using NewColorForControlTint.
// Returns the color object specified by the given control tint.
func ExampleNewColorForControlTint() {
	_ = appkit.NewColorForControlTint(
		appkit.ControlTint{}, // controlTint ControlTint
	)
	// Output:
}
