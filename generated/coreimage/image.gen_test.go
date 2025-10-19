// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)


// ExampleNewImageWithBitmapImageRep demonstrates how to create a Image instance using NewImageWithBitmapImageRep.
// Initializes an image object with the specified bitmap image representation.
func ExampleNewImageWithBitmapImageRep() {
	_ = coreimage.NewImageWithBitmapImageRep(
		nil, // bitmapImageRep unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCGLayer demonstrates how to create a Image instance using NewImageWithCGLayer.
// Initializes an image object  from the contents supplied by a CGLayer object.
func ExampleNewImageWithCGLayer() {
	_ = coreimage.NewImageWithCGLayer(
		nil, // layer unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithColor demonstrates how to create a Image instance using NewImageWithColor.
// Initializes an image of infinite extent whose entire content is the specified color.
func ExampleNewImageWithColor() {
	_ = coreimage.NewImageWithColor(
		nil, // color unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithData demonstrates how to create a Image instance using NewImageWithData.
// Initializes an image object with the supplied image data.
func ExampleNewImageWithData() {
	_ = coreimage.NewImageWithData(
		nil, // data unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithDataOptions demonstrates how to create a Image instance using NewImageWithDataOptions.
// Initializes an image object with the supplied image data, using the specified options.
func ExampleNewImageWithDataOptions() {
	_ = coreimage.NewImageWithDataOptions(
		nil, // data unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithImageOptions demonstrates how to create a Image instance using NewImageWithImageOptions.
// Initializes an image object with the specified UIKit image object, using the specified options.
func ExampleNewImageWithImageOptions() {
	_ = coreimage.NewImageWithImageOptions(
		nil, // image unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithIOSurfacePlaneFormatOptions demonstrates how to create a Image instance using NewImageWithIOSurfacePlaneFormatOptions.
// Initializes, using the specified format and options, an image with the contents of a specific data plane in an IOSurface.
func ExampleNewImageWithIOSurfacePlaneFormatOptions() {
	_ = coreimage.NewImageWithIOSurfacePlaneFormatOptions(
		nil, // surface unsafe.Pointer
		coreimage.uintptr(0), // plane uintptr
		nil, // format unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithTextureSizeFlippedColorSpace demonstrates how to create a Image instance using NewImageWithTextureSizeFlippedColorSpace.
// Initializes an image object with data supplied by an OpenGL texture.
func ExampleNewImageWithTextureSizeFlippedColorSpace() {
	_ = coreimage.NewImageWithTextureSizeFlippedColorSpace(
		nil, // name unsafe.Pointer
		nil, // size unsafe.Pointer
		false, // flipped bool
		nil, // colorSpace unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithContentsOfURL demonstrates how to create a Image instance using NewImageWithContentsOfURL.
// Initializes an image object by reading an image from a URL.
func ExampleNewImageWithContentsOfURL() {
	_ = coreimage.NewImageWithContentsOfURL(
		nil, // url unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithContentsOfURLOptions demonstrates how to create a Image instance using NewImageWithContentsOfURLOptions.
// Initializes an image object by reading an image from a URL, using the specified options.
func ExampleNewImageWithContentsOfURLOptions() {
	_ = coreimage.NewImageWithContentsOfURLOptions(
		nil, // url unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCVPixelBuffer demonstrates how to create a Image instance using NewImageWithCVPixelBuffer.
// Initializes an image object from the contents of a Core Video pixel buffer.
func ExampleNewImageWithCVPixelBuffer() {
	_ = coreimage.NewImageWithCVPixelBuffer(
		nil, // pixelBuffer unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithSemanticSegmentationMatte demonstrates how to create a Image instance using NewImageWithSemanticSegmentationMatte.
func ExampleNewImageWithSemanticSegmentationMatte() {
	_ = coreimage.NewImageWithSemanticSegmentationMatte(
		nil, // matte unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithTextureSizeFlippedOptions demonstrates how to create a Image instance using NewImageWithTextureSizeFlippedOptions.
// Initializes an image object with data supplied by an OpenGL texture.
func ExampleNewImageWithTextureSizeFlippedOptions() {
	_ = coreimage.NewImageWithTextureSizeFlippedOptions(
		nil, // name unsafe.Pointer
		nil, // size unsafe.Pointer
		false, // flipped bool
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCGImageOptions demonstrates how to create a Image instance using NewImageWithCGImageOptions.
// Initializes an image object with a Quartz 2D image, using the specified options.
func ExampleNewImageWithCGImageOptions() {
	_ = coreimage.NewImageWithCGImageOptions(
		nil, // image unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithPortaitEffectsMatteOptions demonstrates how to create a Image instance using NewImageWithPortaitEffectsMatteOptions.
func ExampleNewImageWithPortaitEffectsMatteOptions() {
	_ = coreimage.NewImageWithPortaitEffectsMatteOptions(
		nil, // matte unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCGLayerOptions demonstrates how to create a Image instance using NewImageWithCGLayerOptions.
// Initializes an image object  from the contents supplied by a CGLayer object, using the  specified options.
func ExampleNewImageWithCGLayerOptions() {
	_ = coreimage.NewImageWithCGLayerOptions(
		nil, // layer unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithIOSurface demonstrates how to create a Image instance using NewImageWithIOSurface.
// Initializes an image with the contents of an IOSurface.
func ExampleNewImageWithIOSurface() {
	_ = coreimage.NewImageWithIOSurface(
		nil, // surface unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithDepthDataOptions demonstrates how to create a Image instance using NewImageWithDepthDataOptions.
func ExampleNewImageWithDepthDataOptions() {
	_ = coreimage.NewImageWithDepthDataOptions(
		nil, // data unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithImageProviderSizeFormatColorSpaceOptions demonstrates how to create a Image instance using NewImageWithImageProviderSizeFormatColorSpaceOptions.
// Initializes an image object based on pixels from an image provider object.
func ExampleNewImageWithImageProviderSizeFormatColorSpaceOptions() {
	_ = coreimage.NewImageWithImageProviderSizeFormatColorSpaceOptions(
		0, // provider objc.ID
		coreimage.uintptr(0), // width uintptr
		coreimage.uintptr(0), // height uintptr
		nil, // format unsafe.Pointer
		nil, // colorSpace unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCGImage demonstrates how to create a Image instance using NewImageWithCGImage.
// Initializes an image object with a Quartz 2D image.
func ExampleNewImageWithCGImage() {
	_ = coreimage.NewImageWithCGImage(
		nil, // image unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithDepthData demonstrates how to create a Image instance using NewImageWithDepthData.
func ExampleNewImageWithDepthData() {
	_ = coreimage.NewImageWithDepthData(
		nil, // data unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithSemanticSegmentationMatteOptions demonstrates how to create a Image instance using NewImageWithSemanticSegmentationMatteOptions.
func ExampleNewImageWithSemanticSegmentationMatteOptions() {
	_ = coreimage.NewImageWithSemanticSegmentationMatteOptions(
		nil, // matte unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithBitmapDataBytesPerRowSizeFormatColorSpace demonstrates how to create a Image instance using NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace.
// Initializes an image object with bitmap data.
func ExampleNewImageWithBitmapDataBytesPerRowSizeFormatColorSpace() {
	_ = coreimage.NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace(
		nil, // data unsafe.Pointer
		coreimage.uintptr(0), // bytesPerRow uintptr
		nil, // size unsafe.Pointer
		nil, // format unsafe.Pointer
		nil, // colorSpace unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCVImageBuffer demonstrates how to create a Image instance using NewImageWithCVImageBuffer.
// Initializes an image object from the contents of a Core Video image buffer.
func ExampleNewImageWithCVImageBuffer() {
	_ = coreimage.NewImageWithCVImageBuffer(
		nil, // imageBuffer unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCVImageBufferOptions demonstrates how to create a Image instance using NewImageWithCVImageBufferOptions.
// Initializes an image object from the contents of a Core Video image buffer, using the specified options.
func ExampleNewImageWithCVImageBufferOptions() {
	_ = coreimage.NewImageWithCVImageBufferOptions(
		nil, // imageBuffer unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCVPixelBufferOptions demonstrates how to create a Image instance using NewImageWithCVPixelBufferOptions.
// Initializes an image object from the contents of a Core Video pixel buffer using the specified options.
func ExampleNewImageWithCVPixelBufferOptions() {
	_ = coreimage.NewImageWithCVPixelBufferOptions(
		nil, // pixelBuffer unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithImage demonstrates how to create a Image instance using NewImageWithImage.
// Initializes an image object with the specified UIKit image object.
func ExampleNewImageWithImage() {
	_ = coreimage.NewImageWithImage(
		nil, // image unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithIOSurfaceOptions demonstrates how to create a Image instance using NewImageWithIOSurfaceOptions.
// Initializes, using the specified options, an image with the contents of an IOSurface.
func ExampleNewImageWithIOSurfaceOptions() {
	_ = coreimage.NewImageWithIOSurfaceOptions(
		nil, // surface unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithMTLTextureOptions demonstrates how to create a Image instance using NewImageWithMTLTextureOptions.
// Initializes an image object with data supplied by a Metal texture.
func ExampleNewImageWithMTLTextureOptions() {
	_ = coreimage.NewImageWithMTLTextureOptions(
		nil, // texture unsafe.Pointer
		nil, // options unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithCGImageSourceIndexOptions demonstrates how to create a Image instance using NewImageWithCGImageSourceIndexOptions.
func ExampleNewImageWithCGImageSourceIndexOptions() {
	_ = coreimage.NewImageWithCGImageSourceIndexOptions(
		nil, // source unsafe.Pointer
		coreimage.uintptr(0), // index uintptr
		nil, // dict unsafe.Pointer
	)
	// Output:
}

// ExampleNewImageWithPortaitEffectsMatte demonstrates how to create a Image instance using NewImageWithPortaitEffectsMatte.
func ExampleNewImageWithPortaitEffectsMatte() {
	_ = coreimage.NewImageWithPortaitEffectsMatte(
		nil, // matte unsafe.Pointer
	)
	// Output:
}


