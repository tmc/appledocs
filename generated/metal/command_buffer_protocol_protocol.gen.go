// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PCommandBuffer is the MTLCommandBuffer protocol interface.
//
// A container that stores a sequence of GPU commands that you encode into it.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLCommandBuffer
type PCommandBuffer interface {
	// Required methods
	AddCompletedHandler(block CommandBufferHandler /* not a class type */)
	AddScheduledHandler(block CommandBufferHandler /* not a class type */)
	Commit()
	EncodeSignalEventValue(event unsafe.Pointer, value uint64)
	EncodeWaitForEventValue(event unsafe.Pointer, value uint64)
	Enqueue()
	AccelerationStructureCommandEncoder() unsafe.Pointer
	AccelerationStructureCommandEncoderWithDescriptor(descriptor IMTLAccelerationStructurePassDescriptor) unsafe.Pointer
	BlitCommandEncoder() unsafe.Pointer
	BlitCommandEncoderWithDescriptor(blitPassDescriptor IMTLBlitPassDescriptor) unsafe.Pointer
	ComputeCommandEncoder() unsafe.Pointer
	ComputeCommandEncoderWithDescriptor(computePassDescriptor IMTLComputePassDescriptor) unsafe.Pointer
	ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) unsafe.Pointer
	ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) unsafe.Pointer
	RenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) unsafe.Pointer
	ResourceStateCommandEncoder() unsafe.Pointer
	PopDebugGroup()
	PresentDrawable(drawable unsafe.Pointer)
	PresentDrawableAfterMinimumDuration(drawable unsafe.Pointer, duration float64)
	PresentDrawableAtTime(drawable unsafe.Pointer, presentationTime float64)
	PushDebugGroup(string_ foundation.foundation.INSString)
	ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor IMTLResourceStatePassDescriptor) unsafe.Pointer
	UseResidencySet(residencySet unsafe.Pointer)
	UseResidencySetsCount(residencySets []objc.ID, count uint)
	WaitUntilCompleted()
	WaitUntilScheduled()
}
