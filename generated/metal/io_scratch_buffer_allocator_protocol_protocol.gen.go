// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PIOScratchBufferAllocator is the MTLIOScratchBufferAllocator protocol interface.
//
// A protocol your app implements to provide scratch memory to an input/output command queue.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIOScratchBufferAllocator
type PIOScratchBufferAllocator interface {
	// Required methods
	NewScratchBufferWithMinimumSize(minimumSize uint) unsafe.Pointer
}
