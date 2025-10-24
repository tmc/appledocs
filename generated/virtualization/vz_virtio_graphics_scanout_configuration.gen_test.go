// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZVirtioGraphicsScanoutConfiguration

// ExampleNewVZVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels demonstrates how to create a VZVirtioGraphicsScanoutConfiguration instance using NewVZVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels.
// Creates a Virtio graphics device with the specified dimensions.
func ExampleNewVZVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels() {
	_ = virtualization.NewVZVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels(
		100, // widthInPixels int
		100, // heightInPixels int
	)
	// Output:
}
