// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
)

// PImageAllocator is the MPSImageAllocator protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSImageAllocator
type PImageAllocator interface {
	// Required methods
	Image()
	ImageForCommandBufferImageDescriptorKernel(cmdBuf unsafe.Pointer, descriptor IImageDescriptor, kernel IKernel) IImage
	ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf unsafe.Pointer, descriptor IImageDescriptor, kernel IKernel, count uint) ImageBatch /* not a class type */
	// Optional methods
	ImageBatch()
	HasImageBatch() bool
}
