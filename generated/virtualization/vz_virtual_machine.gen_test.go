// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization_test

import (
	"github.com/tmc/appledocs/generated/virtualization"
)

// Suppress unused import errors
var _ = virtualization.NewVZVirtualMachine

// ExampleNewVZVirtualMachineWithConfiguration demonstrates how to create a VZVirtualMachine instance using NewVZVirtualMachineWithConfiguration.
// Creates the VM and configures it with the specified data.
func ExampleNewVZVirtualMachineWithConfiguration() {
	_ = virtualization.NewVZVirtualMachineWithConfiguration(
		virtualization.VZVirtualMachineConfiguration{}, // configuration VZVirtualMachineConfiguration
	)
	// Output:
}
