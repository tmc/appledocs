// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
)

// PImageTransformProvider is the MPSImageTransformProvider protocol interface.
//
// A general interface for objects that provide image resampling.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSImageTransformProvider
type PImageTransformProvider interface {
	// Required methods
	Transform()
	TransformForSourceImageHandle(image IImage, handle unsafe.Pointer) MPSScaleTransform
}
