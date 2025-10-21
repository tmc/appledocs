// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewAccelerationStructure

// ExampleNewAccelerationStructureWithGroup demonstrates how to create a AccelerationStructure instance using NewAccelerationStructureWithGroup.
func ExampleNewAccelerationStructureWithGroup() {
	_ = metalperformanceshaders.NewAccelerationStructureWithGroup(
		metalperformanceshaders.MPSAccelerationStructureGroup{}, // group MPSAccelerationStructureGroup
	)
	// Output:
}
