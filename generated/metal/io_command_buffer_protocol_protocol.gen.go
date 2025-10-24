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
	AddBarrier()/* debug [protocol_interface/required_method]: AddBarrier */
	AddCompletedHandler(block CommandBufferHandler)/* debug [protocol_interface/required_method]: AddCompletedHandler */
	Commit()/* debug [protocol_interface/required_method]: Commit */
	CopyStatusToBufferOffset(buffer unsafe.Pointer, offset uint)/* debug [protocol_interface/required_method]: CopyStatusToBufferOffset */
	Enqueue()/* debug [protocol_interface/required_method]: Enqueue */
	LoadBufferOffsetSizeSourceHandleSourceHandleOffset(buffer unsafe.Pointer, offset uint, size uint, sourceHandle unsafe.Pointer, sourceHandleOffset uint)/* debug [protocol_interface/required_method]: LoadBufferOffsetSizeSourceHandleSourceHandleOffset */
	LoadTextureSliceLevelSizeSourceBytesPerRowSourceBytesPerImageDestinationOriginSourceHandleSourceHandleOffset(texture unsafe.Pointer, slice uint, level uint, size objc.IObject /* cross-framework: MTLSize */, sourceBytesPerRow uint, sourceBytesPerImage uint, destinationOrigin objc.IObject /* cross-framework: MTLOrigin */, sourceHandle unsafe.Pointer, sourceHandleOffset uint)/* debug [protocol_interface/required_method]: LoadTextureSliceLevelSizeSourceBytesPerRowSourceBytesPerImageDestinationOriginSourceHandleSourceHandleOffset */
	LoadBytesSizeSourceHandleSourceHandleOffset(pointer objectivec.IObject, size uint, sourceHandle unsafe.Pointer, sourceHandleOffset uint)/* debug [protocol_interface/required_method]: LoadBytesSizeSourceHandleSourceHandleOffset */
	PopDebugGroup()/* debug [protocol_interface/required_method]: PopDebugGroup */
	PushDebugGroup(string_ objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: PushDebugGroup */
	SignalEventValue(event unsafe.Pointer, value uint64)/* debug [protocol_interface/required_method]: SignalEventValue */
	TryCancel()/* debug [protocol_interface/required_method]: TryCancel */
	WaitForEventValue(event unsafe.Pointer, value uint64)/* debug [protocol_interface/required_method]: WaitForEventValue */
	WaitUntilCompleted()/* debug [protocol_interface/required_method]: WaitUntilCompleted */
}
