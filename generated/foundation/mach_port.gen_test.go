// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewMachPort

// ExampleNewMachPortWithMachPort demonstrates how to create a MachPort instance using NewMachPortWithMachPort.
// Initializes a newly allocated   object with a given Mach port.
func ExampleNewMachPortWithMachPort() {
	_ = foundation.NewMachPortWithMachPort(
		0, // machPort uint32
	)
	// Output:
}
// ExampleNewMachPortWithMachPortOptions demonstrates how to create a MachPort instance using NewMachPortWithMachPortOptions.
// Initializes a newly allocated   object with a given Mach port and the specified options.
func ExampleNewMachPortWithMachPortOptions() {
	_ = foundation.NewMachPortWithMachPortOptions(
		0, // machPort uint32
		foundation.MachPortOptions{}, // f MachPortOptions
	)
	// Output:
}
