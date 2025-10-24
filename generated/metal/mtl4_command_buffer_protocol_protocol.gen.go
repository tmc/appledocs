// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

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
	BeginCommandBufferWithAllocator(allocator unsafe.Pointer)/* debug [protocol_interface/required_method]: BeginCommandBufferWithAllocator */
	BeginCommandBufferWithAllocatorOptions(allocator unsafe.Pointer, options IMTL4CommandBufferOptions)/* debug [protocol_interface/required_method]: BeginCommandBufferWithAllocatorOptions */
	EndCommandBuffer()/* debug [protocol_interface/required_method]: EndCommandBuffer */
	ComputeCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: ComputeCommandEncoder */
	MachineLearningCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: MachineLearningCommandEncoder */
	RenderCommandEncoderWithDescriptorOptions(descriptor IMTL4RenderPassDescriptor, options MTL4RenderEncoderOptions) unsafe.Pointer/* debug [protocol_interface/required_method]: RenderCommandEncoderWithDescriptorOptions */
	PopDebugGroup()/* debug [protocol_interface/required_method]: PopDebugGroup */
	PushDebugGroup(string_ objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: PushDebugGroup */
	RenderCommandEncoderWithDescriptor(descriptor IMTL4RenderPassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: RenderCommandEncoderWithDescriptor */
	ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence(counterHeap unsafe.Pointer, range_ corefoundation.Range, bufferRange objc.IObject /* cross-framework: MTL4BufferRange */, fenceToWait unsafe.Pointer, fenceToUpdate unsafe.Pointer)/* debug [protocol_interface/required_method]: ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence */
	UseResidencySet(residencySet unsafe.Pointer)/* debug [protocol_interface/required_method]: UseResidencySet */
	UseResidencySetsCount(residencySets []objc.ID, count uint)/* debug [protocol_interface/required_method]: UseResidencySetsCount */
	WriteTimestampIntoHeapAtIndex(counterHeap unsafe.Pointer, index uint)/* debug [protocol_interface/required_method]: WriteTimestampIntoHeapAtIndex */
}
