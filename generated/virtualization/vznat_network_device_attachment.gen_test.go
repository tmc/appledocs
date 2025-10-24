// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZNATNetworkDeviceAttachment

// ExampleNewVZNATNetworkDeviceAttachment demonstrates how to create a VZNATNetworkDeviceAttachment instance.
// Creates an attachment that performs network address translation on the guest system’s network packets.
func ExampleNewVZNATNetworkDeviceAttachment() {
	_ = virtualization.NewVZNATNetworkDeviceAttachment()
	// Output:
}
