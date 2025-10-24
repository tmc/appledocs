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
	AddCompletedHandler(block CommandBufferHandler /* not a class type */)/* debug [protocol_interface/required_method]: AddCompletedHandler */
	AddScheduledHandler(block CommandBufferHandler /* not a class type */)/* debug [protocol_interface/required_method]: AddScheduledHandler */
	Commit()/* debug [protocol_interface/required_method]: Commit */
	EncodeSignalEventValue(event unsafe.Pointer, value uint64)/* debug [protocol_interface/required_method]: EncodeSignalEventValue */
	EncodeWaitForEventValue(event unsafe.Pointer, value uint64)/* debug [protocol_interface/required_method]: EncodeWaitForEventValue */
	Enqueue()/* debug [protocol_interface/required_method]: Enqueue */
	AccelerationStructureCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: AccelerationStructureCommandEncoder */
	AccelerationStructureCommandEncoderWithDescriptor(descriptor IMTLAccelerationStructurePassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: AccelerationStructureCommandEncoderWithDescriptor */
	BlitCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: BlitCommandEncoder */
	BlitCommandEncoderWithDescriptor(blitPassDescriptor IMTLBlitPassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: BlitCommandEncoderWithDescriptor */
	ComputeCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: ComputeCommandEncoder */
	ComputeCommandEncoderWithDescriptor(computePassDescriptor IMTLComputePassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: ComputeCommandEncoderWithDescriptor */
	ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) unsafe.Pointer/* debug [protocol_interface/required_method]: ComputeCommandEncoderWithDispatchType */
	ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: ParallelRenderCommandEncoderWithDescriptor */
	RenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: RenderCommandEncoderWithDescriptor */
	ResourceStateCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: ResourceStateCommandEncoder */
	PopDebugGroup()/* debug [protocol_interface/required_method]: PopDebugGroup */
	PresentDrawable(drawable unsafe.Pointer)/* debug [protocol_interface/required_method]: PresentDrawable */
	PresentDrawableAfterMinimumDuration(drawable unsafe.Pointer, duration float64)/* debug [protocol_interface/required_method]: PresentDrawableAfterMinimumDuration */
	PresentDrawableAtTime(drawable unsafe.Pointer, presentationTime float64)/* debug [protocol_interface/required_method]: PresentDrawableAtTime */
	PushDebugGroup(string_ objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: PushDebugGroup */
	ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor IMTLResourceStatePassDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: ResourceStateCommandEncoderWithDescriptor */
	UseResidencySet(residencySet unsafe.Pointer)/* debug [protocol_interface/required_method]: UseResidencySet */
	UseResidencySetsCount(residencySets []objc.ID, count uint)/* debug [protocol_interface/required_method]: UseResidencySetsCount */
	WaitUntilCompleted()/* debug [protocol_interface/required_method]: WaitUntilCompleted */
	WaitUntilScheduled()/* debug [protocol_interface/required_method]: WaitUntilScheduled */
}
