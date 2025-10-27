// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIndirectCommandBuffer is the MTLIndirectCommandBuffer protocol interface.
//
// A command buffer containing reusable commands, encoded either on the CPU or GPU.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIndirectCommandBuffer
type PIndirectCommandBuffer interface {
	// Required methods
	IndirectComputeCommandAtIndex(commandIndex uint) unsafe.Pointer
	IndirectRenderCommandAtIndex(commandIndex uint) unsafe.Pointer
	ResetWithRange(range_ foundation.Range)
}
