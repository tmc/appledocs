// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
)

// PTextureViewPool is the MTLTextureViewPool protocol interface.
//
// A pool of lightweight texture views.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLTextureViewPool
type PTextureViewPool interface {
	// Required methods
	SetTextureViewFromBufferDescriptorOffsetBytesPerRowAtIndex(buffer unsafe.Pointer, descriptor IMTLTextureDescriptor, offset uint, bytesPerRow uint, index uint) MTLResourceID/* debug [protocol_interface/required_method]: SetTextureViewFromBufferDescriptorOffsetBytesPerRowAtIndex */
	SetTextureViewDescriptorAtIndex(texture unsafe.Pointer, descriptor IMTLTextureViewDescriptor, index uint) MTLResourceID/* debug [protocol_interface/required_method]: SetTextureViewDescriptorAtIndex */
	SetTextureViewAtIndex(texture unsafe.Pointer, index uint) MTLResourceID/* debug [protocol_interface/required_method]: SetTextureViewAtIndex */
}
