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
	RenderCommandEncoder() unsafe.Pointer/* debug [protocol_interface/required_method]: RenderCommandEncoder */
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)/* debug [protocol_interface/required_method]: SetColorStoreActionAtIndex */
	SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint)/* debug [protocol_interface/required_method]: SetColorStoreActionOptionsAtIndex */
	SetDepthStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetDepthStoreAction */
	SetDepthStoreActionOptions(storeActionOptions StoreActionOptions)/* debug [protocol_interface/required_method]: SetDepthStoreActionOptions */
	SetStencilStoreAction(storeAction StoreAction)/* debug [protocol_interface/required_method]: SetStencilStoreAction */
	SetStencilStoreActionOptions(storeActionOptions StoreActionOptions)/* debug [protocol_interface/required_method]: SetStencilStoreActionOptions */
}
