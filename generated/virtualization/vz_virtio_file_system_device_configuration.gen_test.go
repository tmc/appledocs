// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZVirtioFileSystemDeviceConfiguration

// ExampleNewVZVirtioFileSystemDeviceConfigurationWithTag demonstrates how to create a VZVirtioFileSystemDeviceConfiguration instance using NewVZVirtioFileSystemDeviceConfigurationWithTag.
// Creates a configuration for a VIRTIO file system device.
func ExampleNewVZVirtioFileSystemDeviceConfigurationWithTag() {
	_ = virtualization.NewVZVirtioFileSystemDeviceConfigurationWithTag(
		"tag", // tag string
	)
	// Output:
}
