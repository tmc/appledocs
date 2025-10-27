// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

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
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size)
	DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer unsafe.Pointer, indirectBufferOffset uint, threadsPerThreadgroup Size)
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size)
	ExecuteCommandsInBufferIndirectBufferIndirectBufferOffset(indirectCommandbuffer unsafe.Pointer, indirectRangeBuffer unsafe.Pointer, indirectBufferOffset uint)
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer unsafe.Pointer, executionRange foundation.Range)
	MemoryBarrierWithScope(scope BarrierScope)
	MemoryBarrierWithResourcesCount(resources []objc.ID, count uint)
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer unsafe.Pointer, sampleIndex uint, barrier bool)
	SetAccelerationStructureAtBufferIndex(accelerationStructure unsafe.Pointer, bufferIndex uint)
	SetBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)
	SetBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetBufferOffsetAtIndex(offset uint, index uint)
	SetBufferOffsetAttributeStrideAtIndex(offset uint, stride uint, index uint)
	SetBuffersOffsetsAttributeStridesWithRange(buffers []objc.ID, offsets uint, strides uint, range_ foundation.Range)
	SetBuffersOffsetsWithRange(buffers []objc.ID, offsets uint, range_ foundation.Range)
	SetBytesLengthAttributeStrideAtIndex(bytes objectivec.IObject, length uint, stride uint, index uint)
	SetBytesLengthAtIndex(bytes objectivec.IObject, length uint, index uint)
	SetComputePipelineState(state unsafe.Pointer)
	SetImageblockWidthHeight(width uint, height uint)
	SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable unsafe.Pointer, bufferIndex uint)
	SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables []objc.ID, range_ foundation.Range)
	SetSamplerStateAtIndex(sampler unsafe.Pointer, index uint)
	SetSamplerStateLodMinClampLodMaxClampAtIndex(sampler unsafe.Pointer, lodMinClamp float32, lodMaxClamp float32, index uint)
	SetSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers []objc.ID, lodMinClamps []float32, lodMaxClamps []float32, range_ foundation.Range)
	SetSamplerStatesWithRange(samplers []objc.ID, range_ foundation.Range)
	SetStageInRegion(region Region)
	SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer unsafe.Pointer, indirectBufferOffset uint)
	SetTextureAtIndex(texture unsafe.Pointer, index uint)
	SetTexturesWithRange(textures []objc.ID, range_ foundation.Range)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)
	SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable unsafe.Pointer, bufferIndex uint)
	SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables []objc.ID, range_ foundation.Range)
	UpdateFence(fence unsafe.Pointer)
	UseHeap(heap unsafe.Pointer)
	UseHeapsCount(heaps []objc.ID, count uint)
	UseResourceUsage(resource unsafe.Pointer, usage ResourceUsage)
	UseResourcesCountUsage(resources []objc.ID, count uint, usage ResourceUsage)
	WaitForFence(fence unsafe.Pointer)
}
