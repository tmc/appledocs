// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewImage

// ExampleNewImageWithBitmapImageRep demonstrates how to create a Image instance using NewImageWithBitmapImageRep.
// Initializes an image object with the specified bitmap image representation.
func ExampleNewImageWithBitmapImageRep() {
	_ = coreimage.NewImageWithBitmapImageRep(
		coreimage.BitmapImageRep{}, // bitmapImageRep BitmapImageRep
	)
	// Output:
}
// ExampleNewImageWithCVImageBuffer demonstrates how to create a Image instance using NewImageWithCVImageBuffer.
// Initializes an image object from the contents of a Core Video image buffer.
func ExampleNewImageWithCVImageBuffer() {
	_ = coreimage.NewImageWithCVImageBuffer(
		coreimage.ImageBufferRef{}, // imageBuffer ImageBufferRef
	)
	// Output:
}
// ExampleNewImageWithCVPixelBuffer demonstrates how to create a Image instance using NewImageWithCVPixelBuffer.
// Initializes an image object from the contents of a Core Video pixel buffer.
func ExampleNewImageWithCVPixelBuffer() {
	_ = coreimage.NewImageWithCVPixelBuffer(
		coreimage.PixelBufferRef{}, // pixelBuffer PixelBufferRef
	)
	// Output:
}
// ExampleNewImageWithColor demonstrates how to create a Image instance using NewImageWithColor.
// Initializes an image of infinite extent whose entire content is the specified color.
func ExampleNewImageWithColor() {
	_ = coreimage.NewImageWithColor(
		coreimage.CIColor{}, // color CIColor
	)
	// Output:
}
// ExampleNewImageWithDepthData demonstrates how to create a Image instance using NewImageWithDepthData.
func ExampleNewImageWithDepthData() {
	_ = coreimage.NewImageWithDepthData(
		coreimage.DepthData{}, // data DepthData
	)
	// Output:
}
// ExampleNewImageWithIOSurface demonstrates how to create a Image instance using NewImageWithIOSurface.
// Initializes an image with the contents of an IOSurface.
func ExampleNewImageWithIOSurface() {
	_ = coreimage.NewImageWithIOSurface(
		coreimage.SurfaceRef{}, // surface SurfaceRef
	)
	// Output:
}
// ExampleNewImageWithImage demonstrates how to create a Image instance using NewImageWithImage.
// Initializes an image object with the specified UIKit image object.
func ExampleNewImageWithImage() {
	_ = coreimage.NewImageWithImage(
		coreimage.Image{}, // image Image
	)
	// Output:
}
// ExampleNewImageWithPortaitEffectsMatte demonstrates how to create a Image instance using NewImageWithPortaitEffectsMatte.
func ExampleNewImageWithPortaitEffectsMatte() {
	_ = coreimage.NewImageWithPortaitEffectsMatte(
		coreimage.PortraitEffectsMatte{}, // matte PortraitEffectsMatte
	)
	// Output:
}
// ExampleNewImageWithSemanticSegmentationMatte demonstrates how to create a Image instance using NewImageWithSemanticSegmentationMatte.
func ExampleNewImageWithSemanticSegmentationMatte() {
	_ = coreimage.NewImageWithSemanticSegmentationMatte(
		coreimage.SemanticSegmentationMatte{}, // matte SemanticSegmentationMatte
	)
	// Output:
}
