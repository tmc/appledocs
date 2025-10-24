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
	Image()/* debug [protocol_interface/required_method]: Image */
	ImageForCommandBufferImageDescriptorKernel(cmdBuf unsafe.Pointer, descriptor IImageDescriptor, kernel IKernel) Image/* debug [protocol_interface/required_method]: ImageForCommandBufferImageDescriptorKernel */
	ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf unsafe.Pointer, descriptor IImageDescriptor, kernel IKernel, count uint) ImageBatch/* debug [protocol_interface/required_method]: ImageBatchForCommandBufferImageDescriptorKernelCount */
	// Optional methods
	ImageBatch()
	HasImageBatch() bool
}
