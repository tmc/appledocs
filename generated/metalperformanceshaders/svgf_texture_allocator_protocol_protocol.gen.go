// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"unsafe"
)

// PSVGFTextureAllocator is the MPSSVGFTextureAllocator protocol interface.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metalperformanceshaders/documentation/MetalPerformanceShaders/MPSSVGFTextureAllocator
type PSVGFTextureAllocator interface {
	// Required methods
	`return`()
	ReturnTexture(texture unsafe.Pointer)
	Texture()
	TextureWithPixelFormatWidthHeight(pixelFormat PixelFormat /* not a class type */, width uint, height uint) unsafe.Pointer
}
