// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PParallelRenderCommandEncoder is the MTLParallelRenderCommandEncoder protocol interface.
//
// An instance that splits up a single render pass so that it can be simultaneously encoded from multiple threads.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLParallelRenderCommandEncoder
type PParallelRenderCommandEncoder interface {
	// Required methods
	RenderCommandEncoder() unsafe.Pointer
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)
	SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint)
	SetDepthStoreAction(storeAction StoreAction)
	SetDepthStoreActionOptions(storeActionOptions StoreActionOptions)
	SetStencilStoreAction(storeAction StoreAction)
	SetStencilStoreActionOptions(storeActionOptions StoreActionOptions)
}
