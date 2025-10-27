// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PCommandQueue is the MTLCommandQueue protocol interface.
//
// An instance you use to create, submit, and schedule command buffers to a specific GPU device to run the commands within those buffers.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLCommandQueue
type PCommandQueue interface {
	// Required methods
	AddResidencySet(residencySet unsafe.Pointer)
	AddResidencySetsCount(residencySets []objc.ID, count uint)
	InsertDebugCaptureBoundary()
	CommandBuffer() unsafe.Pointer
	CommandBufferWithDescriptor(descriptor IMTLCommandBufferDescriptor) unsafe.Pointer
	CommandBufferWithUnretainedReferences() unsafe.Pointer
	RemoveResidencySet(residencySet unsafe.Pointer)
	RemoveResidencySetsCount(residencySets []objc.ID, count uint)
}
