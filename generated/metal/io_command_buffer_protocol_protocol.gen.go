// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PIOCommandBuffer is the MTLIOCommandBuffer protocol interface.
//
// A command buffer that contains input/output commands that work with files in the file systems and Metal resources.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLIOCommandBuffer
type PIOCommandBuffer interface {
	// Required methods
	AddBarrier()
	AddCompletedHandler(block CommandBufferHandler)
	Commit()
	CopyStatusToBufferOffset(buffer unsafe.Pointer, offset uint)
	Enqueue()
	LoadBufferOffsetSizeSourceHandleSourceHandleOffset(buffer unsafe.Pointer, offset uint, size uint, sourceHandle unsafe.Pointer, sourceHandleOffset uint)
	LoadTextureSliceLevelSizeSourceBytesPerRowSourceBytesPerImageDestinationOriginSourceHandleSourceHandleOffset(texture unsafe.Pointer, slice uint, level uint, size Size, sourceBytesPerRow uint, sourceBytesPerImage uint, destinationOrigin Origin, sourceHandle unsafe.Pointer, sourceHandleOffset uint)
	LoadBytesSizeSourceHandleSourceHandleOffset(pointer objectivec.IObject, size uint, sourceHandle unsafe.Pointer, sourceHandleOffset uint)
	PopDebugGroup()
	PushDebugGroup(string_ foundation.foundation.INSString)
	SignalEventValue(event unsafe.Pointer, value uint64)
	TryCancel()
	WaitForEventValue(event unsafe.Pointer, value uint64)
	WaitUntilCompleted()
}
