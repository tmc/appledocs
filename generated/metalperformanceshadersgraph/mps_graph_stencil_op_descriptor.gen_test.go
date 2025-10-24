// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshadersgraph"
)

// Suppress unused import errors
var _ = metalperformanceshadersgraph.NewGraphStencilOpDescriptor

// ExampleNewGraphStencilOpDescriptorWithExplicitPadding demonstrates how to create a GraphStencilOpDescriptor instance using NewGraphStencilOpDescriptorWithExplicitPadding.
// Creates a stencil operation descriptor with default values.
func ExampleNewGraphStencilOpDescriptorWithExplicitPadding() {
	_ = metalperformanceshadersgraph.NewGraphStencilOpDescriptorWithExplicitPadding(
		metalperformanceshadersgraph.Shape /* not a class type */{}, // explicitPadding Shape /* not a class type */
	)
	// Output:
}
// ExampleNewGraphStencilOpDescriptorWithOffsetsExplicitPadding demonstrates how to create a GraphStencilOpDescriptor instance using NewGraphStencilOpDescriptorWithOffsetsExplicitPadding.
// Creates a stencil operation descriptor with default values.
func ExampleNewGraphStencilOpDescriptorWithOffsetsExplicitPadding() {
	_ = metalperformanceshadersgraph.NewGraphStencilOpDescriptorWithOffsetsExplicitPadding(
		metalperformanceshadersgraph.Shape /* not a class type */{}, // offsets Shape /* not a class type */
		metalperformanceshadersgraph.Shape /* not a class type */{}, // explicitPadding Shape /* not a class type */
	)
	// Output:
}
// ExampleNewGraphStencilOpDescriptorWithPaddingStyle demonstrates how to create a GraphStencilOpDescriptor instance using NewGraphStencilOpDescriptorWithPaddingStyle.
// Creates a stencil operation descriptor with default values.
func ExampleNewGraphStencilOpDescriptorWithPaddingStyle() {
	_ = metalperformanceshadersgraph.NewGraphStencilOpDescriptorWithPaddingStyle(
		metalperformanceshadersgraph.GraphPaddingStyle{}, // paddingStyle GraphPaddingStyle
	)
	// Output:
}
// ExampleNewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant demonstrates how to create a GraphStencilOpDescriptor instance using NewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant.
// Creates a stencil operation descriptor with given values.
func ExampleNewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant() {
	_ = metalperformanceshadersgraph.NewGraphStencilOpDescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(
		metalperformanceshadersgraph.GraphReductionMode{}, // reductionMode GraphReductionMode
		metalperformanceshadersgraph.Shape /* not a class type */{}, // offsets Shape /* not a class type */
		metalperformanceshadersgraph.Shape /* not a class type */{}, // strides Shape /* not a class type */
		metalperformanceshadersgraph.Shape /* not a class type */{}, // dilationRates Shape /* not a class type */
		metalperformanceshadersgraph.Shape /* not a class type */{}, // explicitPadding Shape /* not a class type */
		metalperformanceshadersgraph.GraphPaddingMode{}, // boundaryMode GraphPaddingMode
		metalperformanceshadersgraph.GraphPaddingStyle{}, // paddingStyle GraphPaddingStyle
		0.0, // paddingConstant float32
	)
	// Output:
}
