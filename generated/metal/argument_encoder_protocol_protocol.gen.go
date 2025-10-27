// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PArgumentEncoder is the MTLArgumentEncoder protocol interface.
//
// An interface you can use to encode argument data into an argument buffer.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLArgumentEncoder
type PArgumentEncoder interface {
	// Required methods
	ConstantDataAtIndex(index uint)
	NewArgumentEncoderForBufferAtIndex(index uint) unsafe.Pointer
	SetAccelerationStructureAtIndex(accelerationStructure unsafe.Pointer, index uint)
	SetArgumentBufferOffset(argumentBuffer unsafe.Pointer, offset uint)
	SetArgumentBufferStartOffsetArrayElement(argumentBuffer unsafe.Pointer, startOffset uint, arrayElement uint)
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ foundation.Range)
	SetComputePipelineStateAtIndex(pipeline unsafe.Pointer, index uint)
	SetComputePipelineStatesWithRange(pipelines []objc.ID, range_ foundation.Range)
	SetDepthStencilStateAtIndex(depthStencilState unsafe.Pointer, index uint)
	SetDepthStencilStatesWithRange(depthStencilStates []objc.ID, range_ foundation.Range)
	SetIndirectCommandBufferAtIndex(indirectCommandBuffer unsafe.Pointer, index uint)
	SetIndirectCommandBuffersWithRange(buffers []objc.ID, range_ foundation.Range)
	SetIntersectionFunctionTableAtIndex(intersectionFunctionTable unsafe.Pointer, index uint)
	SetIntersectionFunctionTablesWithRange(intersectionFunctionTables []objc.ID, range_ foundation.Range)
	SetRenderPipelineStateAtIndex(pipeline unsafe.Pointer, index uint)
	SetRenderPipelineStatesWithRange(pipelines []objc.ID, range_ foundation.Range)
	SetSamplerStateAtIndex(sampler unsafe.Pointer, index uint)
	SetSamplerStatesWithRange(samplers []objc.ID, range_ foundation.Range)
	SetTextureAtIndex(texture unsafe.Pointer, index uint)
	SetTexturesWithRange(textures []objc.ID, range_ foundation.Range)
	SetVisibleFunctionTableAtIndex(visibleFunctionTable unsafe.Pointer, index uint)
	SetVisibleFunctionTablesWithRange(visibleFunctionTables []objc.ID, range_ foundation.Range)
}
