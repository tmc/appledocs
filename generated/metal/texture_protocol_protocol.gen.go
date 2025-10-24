// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PTexture is the MTLTexture protocol interface.
//
// A resource that holds formatted image data.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - tvOS +
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLTexture
type PTexture interface {
	// Required methods
	GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes objectivec.IObject, bytesPerRow uint, bytesPerImage uint, region objc.IObject /* cross-framework: MTLRegion */, level uint, slice uint)/* debug [protocol_interface/required_method]: GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice */
	GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes objectivec.IObject, bytesPerRow uint, region objc.IObject /* cross-framework: MTLRegion */, level uint)/* debug [protocol_interface/required_method]: GetBytesBytesPerRowFromRegionMipmapLevel */
	NewRemoteTextureViewForDevice(device unsafe.Pointer) unsafe.Pointer/* debug [protocol_interface/required_method]: NewRemoteTextureViewForDevice */
	NewSharedTextureHandle() SharedTextureHandle/* debug [protocol_interface/required_method]: NewSharedTextureHandle */
	NewTextureViewWithPixelFormat(pixelFormat PixelFormat) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureViewWithPixelFormat */
	NewTextureViewWithDescriptor(descriptor IMTLTextureViewDescriptor) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureViewWithDescriptor */
	NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat PixelFormat, textureType TextureType, levelRange corefoundation.Range, sliceRange corefoundation.Range) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureViewWithPixelFormatTextureTypeLevelsSlices */
	NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange corefoundation.Range, sliceRange corefoundation.Range, swizzle objc.IObject /* cross-framework: MTLTextureSwizzleChannels */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle */
	ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region objc.IObject /* cross-framework: MTLRegion */, level uint, slice uint, pixelBytes objectivec.IObject, bytesPerRow uint, bytesPerImage uint)/* debug [protocol_interface/required_method]: ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage */
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region objc.IObject /* cross-framework: MTLRegion */, level uint, pixelBytes objectivec.IObject, bytesPerRow uint)/* debug [protocol_interface/required_method]: ReplaceRegionMipmapLevelWithBytesBytesPerRow */
}
