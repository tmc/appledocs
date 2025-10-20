// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZMacGraphicsDisplayConfiguration



// ExampleNewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch demonstrates how to create a VZMacGraphicsDisplayConfiguration instance using NewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch.
// Create a display configuration with the specified pixel dimensions and pixel density.
func ExampleNewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch() {
	_ = virtualization.NewVZMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(
		0, // widthInPixels int
		0, // heightInPixels int
		0, // pixelsPerInch int
	)
	// Output:
}


