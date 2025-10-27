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
	ClearBarrier()
	ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size)
	ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size)
	Reset()
	SetBarrier()
	SetComputePipelineState(pipelineState unsafe.Pointer)
	SetImageblockWidthHeight(width uint, height uint)
	SetKernelBufferOffsetAtIndex(buffer unsafe.Pointer, offset uint, index uint)
	SetKernelBufferOffsetAttributeStrideAtIndex(buffer unsafe.Pointer, offset uint, stride uint, index uint)
	SetStageInRegion(region Region)
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)
}
