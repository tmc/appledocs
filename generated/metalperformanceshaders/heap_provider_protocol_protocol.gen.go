// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/metal"
)

// PHeapProvider is the MPSHeapProvider protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSHeapProvider
type PHeapProvider interface {
	// Required methods
	NewHeap()
	NewHeapWithDescriptor(descriptor metal.HeapDescriptor) unsafe.Pointer
	RetireHeapCacheDelay(heap unsafe.Pointer, seconds float64)
	// Optional methods
	Retire()
	HasRetire() bool
}
