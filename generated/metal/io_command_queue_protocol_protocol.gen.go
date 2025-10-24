// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PIOCommandQueue is the MTLIOCommandQueue protocol interface.
//
// A command queue that schedules input/output commands for reading files in the file system, and writing to GPU resources and memory.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIOCommandQueue
type PIOCommandQueue interface {
	// Required methods
	EnqueueBarrier()/* debug [protocol_interface/required_method]: EnqueueBarrier */
	CommandBuffer() unsafe.Pointer/* debug [protocol_interface/required_method]: CommandBuffer */
	CommandBufferWithUnretainedReferences() unsafe.Pointer/* debug [protocol_interface/required_method]: CommandBufferWithUnretainedReferences */
}
