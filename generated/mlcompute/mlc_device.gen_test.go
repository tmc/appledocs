// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCDevice

// ExampleNewCDeviceWithType demonstrates how to create a CDevice instance using NewCDeviceWithType.
// Creates a device of the type you specify.
func ExampleNewCDeviceWithType() {
	_ = mlcompute.NewCDeviceWithType(
		mlcompute.CDeviceType{}, // type CDeviceType
	)
	// Output:
}
// ExampleNewCDeviceWithTypeSelectsMultipleComputeDevices demonstrates how to create a CDevice instance using NewCDeviceWithTypeSelectsMultipleComputeDevices.
// Creates a device that you can configure to use multiple compute devices.
func ExampleNewCDeviceWithTypeSelectsMultipleComputeDevices() {
	_ = mlcompute.NewCDeviceWithTypeSelectsMultipleComputeDevices(
		mlcompute.CDeviceType{}, // type CDeviceType
		false, // selectsMultipleComputeDevices bool
	)
	// Output:
}
