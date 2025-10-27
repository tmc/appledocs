// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
)

// PPadding is the MPSNNPadding protocol interface.
//
// The protocol that provides a description of how kernels should pad images.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSNNPadding
type PPadding interface {
	// Required methods
	PaddingMethod()
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, kernel IKernel, inDescriptor IImageDescriptor) IImageDescriptor
	// Optional methods
	DestinationImageDescriptor()
	HasDestinationImageDescriptor() bool
	Label()
	HasLabel() bool
	Inverse()
	HasInverse() bool
}
