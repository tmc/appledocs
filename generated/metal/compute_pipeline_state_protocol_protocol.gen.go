// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PComputePipelineState is the MTLComputePipelineState protocol interface.
//
// An interface that represents a GPU pipeline configuration for running kernels in a compute pass.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLComputePipelineState
type PComputePipelineState interface {
	// Required methods
	FunctionHandleWithFunction(function unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithFunction */
	FunctionHandleWithBinaryFunction(function unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithBinaryFunction */
	FunctionHandleWithName(name objc.IObject /* cross-framework: NSString */) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithName */
	ImageblockMemoryLengthForDimensions(imageblockDimensions objc.IObject /* cross-framework: MTLSize */) uint/* debug [protocol_interface/required_method]: ImageblockMemoryLengthForDimensions */
	NewComputePipelineStateWithBinaryFunctionsError(additionalBinaryFunctions []objc.ID, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithBinaryFunctionsError */
	NewComputePipelineStateWithAdditionalBinaryFunctionsError(functions []objc.ID, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewComputePipelineStateWithAdditionalBinaryFunctionsError */
	NewIntersectionFunctionTableWithDescriptor(descriptor IMTLIntersectionFunctionTableDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIntersectionFunctionTableWithDescriptor */
	NewVisibleFunctionTableWithDescriptor(descriptor IMTLVisibleFunctionTableDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewVisibleFunctionTableWithDescriptor */
}
