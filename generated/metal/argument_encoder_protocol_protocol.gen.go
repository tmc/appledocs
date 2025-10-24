// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
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
	ConstantDataAtIndex(index uint)/* debug [protocol_interface/required_method]: ConstantDataAtIndex */
	NewArgumentEncoderForBufferAtIndex(index uint) unsafe.Pointer/* debug [protocol_interface/required_method]: NewArgumentEncoderForBufferAtIndex */
	SetAccelerationStructureAtIndex(accelerationStructure unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetAccelerationStructureAtIndex */
	SetArgumentBufferOffset(argumentBuffer unsafe.Pointer, offset uint)/* debug [protocol_interface/required_method]: SetArgumentBufferOffset */
	SetArgumentBufferStartOffsetArrayElement(argumentBuffer unsafe.Pointer, startOffset uint, arrayElement uint)/* debug [protocol_interface/required_method]: SetArgumentBufferStartOffsetArrayElement */
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetBufferOffsetAtIndex */
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetBuffersOffsetsWithRange */
	SetComputePipelineStateAtIndex(pipeline unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetComputePipelineStateAtIndex */
	SetComputePipelineStatesWithRange(pipelines []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetComputePipelineStatesWithRange */
	SetDepthStencilStateAtIndex(depthStencilState unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetDepthStencilStateAtIndex */
	SetDepthStencilStatesWithRange(depthStencilStates []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetDepthStencilStatesWithRange */
	SetIndirectCommandBufferAtIndex(indirectCommandBuffer unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetIndirectCommandBufferAtIndex */
	SetIndirectCommandBuffersWithRange(buffers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetIndirectCommandBuffersWithRange */
	SetIntersectionFunctionTableAtIndex(intersectionFunctionTable unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetIntersectionFunctionTableAtIndex */
	SetIntersectionFunctionTablesWithRange(intersectionFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetIntersectionFunctionTablesWithRange */
	SetRenderPipelineStateAtIndex(pipeline unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetRenderPipelineStateAtIndex */
	SetRenderPipelineStatesWithRange(pipelines []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetRenderPipelineStatesWithRange */
	SetSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetSamplerStateAtIndex */
	SetSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetSamplerStatesWithRange */
	SetTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetTextureAtIndex */
	SetTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTexturesWithRange */
	SetVisibleFunctionTableAtIndex(visibleFunctionTable unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetVisibleFunctionTableAtIndex */
	SetVisibleFunctionTablesWithRange(visibleFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVisibleFunctionTablesWithRange */
}
