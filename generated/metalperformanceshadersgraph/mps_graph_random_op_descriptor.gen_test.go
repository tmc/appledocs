// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphRandomOpDescriptor

// ExampleNewGraphRandomOpDescriptorWithDistributionDataType demonstrates how to create a GraphRandomOpDescriptor instance using NewGraphRandomOpDescriptorWithDistributionDataType.
// Class method to initialize a distribution descriptor.
func ExampleNewGraphRandomOpDescriptorWithDistributionDataType() {
	_ = metalperformanceshadersgraph.NewGraphRandomOpDescriptorWithDistributionDataType(
		metalperformanceshadersgraph.GraphRandomDistribution{}, // distribution GraphRandomDistribution
		metalperformanceshadersgraph.DataType /* not a class type */{}, // dataType DataType /* not a class type */
	)
	// Output:
}
