// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTensor is the MTLTensor protocol interface.
//
// A resource representing a multi-dimensional array that you can use with machine learning workloads.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLTensor
type PTensor interface {
	// Required methods
	GetBytesStridesFromSliceOriginSliceDimensions(bytes objectivec.IObject, strides IMTLTensorExtents, sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents)
	ReplaceSliceOriginSliceDimensionsWithBytesStrides(sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents, bytes objectivec.IObject, strides IMTLTensorExtents)
}
