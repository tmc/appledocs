// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PMTL4CommandAllocator is the MTL4CommandAllocator protocol interface.
//
// Manages the memory backing the encoding of GPU commands into command buffers.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTL4CommandAllocator
type PMTL4CommandAllocator interface {
	// Required methods
	AllocatedSize() uint64/* debug [protocol_interface/required_method]: AllocatedSize */
	Reset()/* debug [protocol_interface/required_method]: Reset */
}
