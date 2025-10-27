// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	FunctionHandleWithBinaryFunctionStage(function unsafe.Pointer, stage RenderStages) unsafe.Pointer
	FunctionHandleWithFunctionStage(function unsafe.Pointer, stage RenderStages) unsafe.Pointer
	FunctionHandleWithNameStage(name foundation.foundation.INSString, stage RenderStages) unsafe.Pointer
	ImageblockMemoryLengthForDimensions(imageblockDimensions Size) uint
	NewIntersectionFunctionTableWithDescriptorStage(descriptor IMTLIntersectionFunctionTableDescriptor, stage RenderStages) unsafe.Pointer
	NewRenderPipelineDescriptorForSpecialization() IMTL4PipelineDescriptor
	NewRenderPipelineStateWithBinaryFunctionsError(binaryFunctionsDescriptor IMTL4RenderPipelineBinaryFunctionsDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewRenderPipelineStateWithAdditionalBinaryFunctionsError(additionalBinaryFunctions IMTLRenderPipelineFunctionsDescriptor, error_ foundation.foundation.INSError) unsafe.Pointer
	NewVisibleFunctionTableWithDescriptorStage(descriptor IMTLVisibleFunctionTableDescriptor, stage RenderStages) unsafe.Pointer
}
