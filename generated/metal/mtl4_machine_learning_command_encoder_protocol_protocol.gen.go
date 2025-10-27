// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PMTL4MachineLearningCommandEncoder is the MTL4MachineLearningCommandEncoder protocol interface.
//
// Encodes dispatch commands that run machine-learning model inference on Apple silicon.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4MachineLearningCommandEncoder
type PMTL4MachineLearningCommandEncoder interface {
	// Required methods
	DispatchNetworkWithIntermediatesHeap(heap unsafe.Pointer)
	SetArgumentTable(argumentTable unsafe.Pointer)
	SetPipelineState(pipelineState unsafe.Pointer)
}
