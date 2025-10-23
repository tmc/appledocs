// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPressureConfiguration

// ExampleNewPressureConfigurationWithPressureBehavior demonstrates how to create a PressureConfiguration instance using NewPressureConfigurationWithPressureBehavior.
// Initializes a pressure configuration object with a specified pressure behavior.
func ExampleNewPressureConfigurationWithPressureBehavior() {
	_ = appkit.NewPressureConfigurationWithPressureBehavior(
		appkit.PressureBehavior{}, // pressureBehavior PressureBehavior
	)
	// Output:
}
