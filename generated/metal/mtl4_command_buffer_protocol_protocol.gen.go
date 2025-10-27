// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMTL4CommandBuffer is the MTL4CommandBuffer protocol interface.
//
// Records a sequence of GPU commands.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4CommandBuffer
type PMTL4CommandBuffer interface {
	// Required methods
	BeginCommandBufferWithAllocator(allocator unsafe.Pointer)
	BeginCommandBufferWithAllocatorOptions(allocator unsafe.Pointer, options IMTL4CommandBufferOptions)
	EndCommandBuffer()
	ComputeCommandEncoder() unsafe.Pointer
	MachineLearningCommandEncoder() unsafe.Pointer
	RenderCommandEncoderWithDescriptorOptions(descriptor IMTL4RenderPassDescriptor, options MTL4RenderEncoderOptions) unsafe.Pointer
	PopDebugGroup()
	PushDebugGroup(string_ foundation.foundation.INSString)
	RenderCommandEncoderWithDescriptor(descriptor IMTL4RenderPassDescriptor) unsafe.Pointer
	ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence(counterHeap unsafe.Pointer, range_ foundation.Range, bufferRange MTL4BufferRange, fenceToWait unsafe.Pointer, fenceToUpdate unsafe.Pointer)
	UseResidencySet(residencySet unsafe.Pointer)
	UseResidencySetsCount(residencySets []objc.ID, count uint)
	WriteTimestampIntoHeapAtIndex(counterHeap unsafe.Pointer, index uint)
}
