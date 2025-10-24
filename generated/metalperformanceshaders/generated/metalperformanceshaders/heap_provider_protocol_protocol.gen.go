// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
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
	NewHeap()/* debug [protocol_interface/required_method]: NewHeap */
	NewHeapWithDescriptor(descriptor HeapDescriptor /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewHeapWithDescriptor */
	RetireHeapCacheDelay(heap unsafe.Pointer, seconds float64)/* debug [protocol_interface/required_method]: RetireHeapCacheDelay */
	// Optional methods
	Retire()
	HasRetire() bool
}
