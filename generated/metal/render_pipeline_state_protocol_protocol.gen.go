// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PRenderPipelineState is the MTLRenderPipelineState protocol interface.
//
// An interface that represents a graphics pipeline configuration for a render pass, which the pass applies to the draw commands you encode.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLRenderPipelineState
type PRenderPipelineState interface {
	// Required methods
	FunctionHandleWithBinaryFunctionStage(function unsafe.Pointer, stage RenderStages) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithBinaryFunctionStage */
	FunctionHandleWithFunctionStage(function unsafe.Pointer, stage RenderStages) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithFunctionStage */
	FunctionHandleWithNameStage(name objc.IObject /* cross-framework: NSString */, stage RenderStages) unsafe.Pointer/* debug [protocol_interface/required_method]: FunctionHandleWithNameStage */
	ImageblockMemoryLengthForDimensions(imageblockDimensions objc.IObject /* cross-framework: MTLSize */) uint/* debug [protocol_interface/required_method]: ImageblockMemoryLengthForDimensions */
	NewIntersectionFunctionTableWithDescriptorStage(descriptor IMTLIntersectionFunctionTableDescriptor, stage RenderStages) unsafe.Pointer/* debug [protocol_interface/required_method]: NewIntersectionFunctionTableWithDescriptorStage */
	NewRenderPipelineDescriptorForSpecialization() MTL4PipelineDescriptor/* debug [protocol_interface/required_method]: NewRenderPipelineDescriptorForSpecialization */
	NewRenderPipelineStateWithBinaryFunctionsError(binaryFunctionsDescriptor IMTL4RenderPipelineBinaryFunctionsDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithBinaryFunctionsError */
	NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions IMTLRenderPipelineFunctionsDescriptor, error_ objectivec.IObject) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRenderPipelineStateWithAdditionalBinaryFunctionsError */
	NewVisibleFunctionTableWithDescriptorStage(descriptor IMTLVisibleFunctionTableDescriptor, stage RenderStages) unsafe.Pointer/* debug [protocol_interface/required_method]: NewVisibleFunctionTableWithDescriptorStage */
}
