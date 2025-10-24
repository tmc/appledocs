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
		foundation.uint32 /* not a class type */{}, // machPort uint32 /* not a class type */
	)
	// Output:
}
// ExampleNewMachPortWithMachPortOptions demonstrates how to create a MachPort instance using NewMachPortWithMachPortOptions.
// Initializes a newly allocated   object with a given Mach port and the specified options.
func ExampleNewMachPortWithMachPortOptions() {
	_ = foundation.NewMachPortWithMachPortOptions(
		foundation.uint32 /* not a class type */{}, // machPort uint32 /* not a class type */
		foundation.MachPortOptions{}, // f MachPortOptions
	)
	// Output:
}
// ExampleMachPort_Delegate demonstrates using Delegate on a MachPort instance.
// Returns the receiver’s delegate.
func ExampleMachPort_Delegate() {
	obj := foundation.NewMachPort()
	_ = obj.Delegate()
	// Output:
	}

