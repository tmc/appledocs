// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
)

// PNDArrayAllocator is the MPSNDArrayAllocator protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSNDArrayAllocator
type PNDArrayAllocator interface {
	// Required methods
	Array()
	ArrayForCommandBufferArrayDescriptorKernel(cmdBuf unsafe.Pointer, descriptor INDArrayDescriptor, kernel IKernel) INDArray
}
