// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PComputeCommandEncoder is the MTLComputeCommandEncoder protocol interface.
//
// An interface for dispatching commands to encode in a compute pass.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLComputeCommandEncoder
type PComputeCommandEncoder interface {
	// Required methods
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadgroupsThreadsPerThreadgroup */
	DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer unsafe.Pointer, indirectBufferOffset uint, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup */
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: DispatchThreadsThreadsPerThreadgroup */
	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(indirectCommandbuffer unsafe.Pointer, indirectRangeBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset */
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange corefoundation.Range)/* debug [protocol_interface/required_method]: ExecuteCommandsInBufferWithRange */
	MemoryBarrierWithScope(scope BarrierScope)/* debug [protocol_interface/required_method]: MemoryBarrierWithScope */
	MemoryBarrierWithResourcesCount(resources []objc.ID, count uint)/* debug [protocol_interface/required_method]: MemoryBarrierWithResourcesCount */
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer unsafe.Pointer, sampleIndex uint, barrier bool)/* debug [protocol_interface/required_method]: SampleCountersInBufferAtSampleIndexWithBarrier */
	SetAccelerationStructureAtBufferIndex(accelerationStructure unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetAccelerationStructureAtBufferIndex */
	SetBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetBufferOffsetAttributeStrideAtIndex */
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetBufferOffsetAtIndex */
	SetBuffersOffsetsAttributeStridesWithRange(buffers []objc.ID, offsets uint, strides uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetBuffersOffsetsAttributeStridesWithRange */
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetBuffersOffsetsWithRange */
	SetBytesLengthAttributeStrideAtIndex(bytes objectivec.IObject, length uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetBytesLengthAttributeStrideAtIndex */
	SetBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)/* debug [protocol_interface/required_method]: SetBytesLengthAtIndex */
	SetComputePipelineState(state unsafe.Pointer)/* debug [protocol_interface/required_method]: SetComputePipelineState */
	SetImageblockWidthHeight(width uint, height uint)/* debug [protocol_interface/required_method]: SetImageblockWidthHeight */
	SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetIntersectionFunctionTableAtBufferIndex */
	SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetIntersectionFunctionTablesWithBufferRange */
	SetSamplerStateAtIndex(sampler unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetSamplerStateAtIndex */
	SetSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)/* debug [protocol_interface/required_method]: SetSamplerStateLodMinClampLodMaxClampAtIndex */
	SetSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetSamplerStatesLodMinClampsLodMaxClampsWithRange */
	SetSamplerStatesWithRange(samplers []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetSamplerStatesWithRange */
	SetStageInRegion(region objc.IObject /* cross-framework: MTLRegion */)/* debug [protocol_interface/required_method]: SetStageInRegion */
	SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer unsafe.Pointer, indirectBufferOffset uint)/* debug [protocol_interface/required_method]: SetStageInRegionWithIndirectBufferIndirectBufferOffset */
	SetTextureAtIndex(texture unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: SetTextureAtIndex */
	SetTexturesWithRange(textures []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetTexturesWithRange */
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetThreadgroupMemoryLengthAtIndex */
	SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable unsafe.Pointer, bufferIndex uint)/* debug [protocol_interface/required_method]: SetVisibleFunctionTableAtBufferIndex */
	SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables []objc.ID, range_ corefoundation.Range)/* debug [protocol_interface/required_method]: SetVisibleFunctionTablesWithBufferRange */
	UpdateFence(fence unsafe.Pointer)/* debug [protocol_interface/required_method]: UpdateFence */
	UseHeap(heap unsafe.Pointer)/* debug [protocol_interface/required_method]: UseHeap */
	UseHeapsCount(heaps []objc.ID, count uint)/* debug [protocol_interface/required_method]: UseHeapsCount */
	UseResourceUsage(resource unsafe.Pointer, usage ResourceUsage)/* debug [protocol_interface/required_method]: UseResourceUsage */
	UseResourcesCountUsage(resources []objc.ID, count uint, usage ResourceUsage)/* debug [protocol_interface/required_method]: UseResourcesCountUsage */
	WaitForFence(fence unsafe.Pointer)/* debug [protocol_interface/required_method]: WaitForFence */
}
