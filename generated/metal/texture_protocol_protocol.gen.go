// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"

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
	GetBytesBytesPerRowBytesPerImageFromRegionMipmapLevelSlice(pixelBytes objectivec.IObject, bytesPerRow uint, bytesPerImage uint, region Region, level uint, slice uint)
	GetBytesBytesPerRowFromRegionMipmapLevel(pixelBytes objectivec.IObject, bytesPerRow uint, region Region, level uint)
	NewRemoteTextureViewForDevice(device unsafe.Pointer) unsafe.Pointer
	NewSharedTextureHandle() ISharedTextureHandle
	NewTextureViewWithPixelFormat(pixelFormat PixelFormat) unsafe.Pointer
	NewTextureViewWithDescriptor(descriptor IMTLTextureViewDescriptor) unsafe.Pointer
	NewTextureViewWithPixelFormatTextureTypeLevelsSlices(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range) unsafe.Pointer
	NewTextureViewWithPixelFormatTextureTypeLevelsSlicesSwizzle(pixelFormat PixelFormat, textureType TextureType, levelRange foundation.Range, sliceRange foundation.Range, swizzle TextureSwizzleChannels) unsafe.Pointer
	ReplaceRegionMipmapLevelSliceWithBytesBytesPerRowBytesPerImage(region Region, level uint, slice uint, pixelBytes objectivec.IObject, bytesPerRow uint, bytesPerImage uint)
	ReplaceRegionMipmapLevelWithBytesBytesPerRow(region Region, level uint, pixelBytes objectivec.IObject, bytesPerRow uint)
}
