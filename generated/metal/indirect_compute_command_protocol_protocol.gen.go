// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PIndirectComputeCommand is the MTLIndirectComputeCommand protocol interface.
//
// A compute command in an indirect command buffer.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 11.0+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIndirectComputeCommand
type PIndirectComputeCommand interface {
	// Required methods
	ClearBarrier()/* debug [protocol_interface/required_method]: ClearBarrier */
	ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: ConcurrentDispatchThreadgroupsThreadsPerThreadgroup */
	ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid objc.IObject /* cross-framework: MTLSize */, threadsPerThreadgroup objc.IObject /* cross-framework: MTLSize */)/* debug [protocol_interface/required_method]: ConcurrentDispatchThreadsThreadsPerThreadgroup */
	Reset()/* debug [protocol_interface/required_method]: Reset */
	SetBarrier()/* debug [protocol_interface/required_method]: SetBarrier */
	SetComputePipelineState(pipelineState unsafe.Pointer)/* debug [protocol_interface/required_method]: SetComputePipelineState */
	SetImageblockWidthHeight(width uint, height uint)/* debug [protocol_interface/required_method]: SetImageblockWidthHeight */
	SetKernelBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)/* debug [protocol_interface/required_method]: SetKernelBufferOffsetAtIndex */
	SetKernelBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)/* debug [protocol_interface/required_method]: SetKernelBufferOffsetAttributeStrideAtIndex */
	SetStageInRegion(region objc.IObject /* cross-framework: MTLRegion */)/* debug [protocol_interface/required_method]: SetStageInRegion */
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)/* debug [protocol_interface/required_method]: SetThreadgroupMemoryLengthAtIndex */
}
