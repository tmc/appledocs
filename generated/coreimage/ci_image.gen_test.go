// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage_test

import (
	"github.com/tmc/appledocs/generated/coreimage"
)

// Suppress unused import errors
var _ = coreimage.NewImage

// ExampleNewImageWithCGImage demonstrates how to create a Image instance using NewImageWithCGImage.
// Initializes an image object with a Quartz 2D image.
func ExampleNewImageWithCGImage() {
	_ = coreimage.NewImageWithCGImage(
		coreimage.ImageRef /* not a class type */{}, // image ImageRef /* not a class type */
	)
	// Output:
}
// ExampleNewImageWithCGLayer demonstrates how to create a Image instance using NewImageWithCGLayer.
// Initializes an image object  from the contents supplied by a CGLayer object.
func ExampleNewImageWithCGLayer() {
	_ = coreimage.NewImageWithCGLayer(
		coreimage.LayerRef /* not a class type */{}, // layer LayerRef /* not a class type */
	)
	// Output:
}
// ExampleNewImageWithCVImageBuffer demonstrates how to create a Image instance using NewImageWithCVImageBuffer.
// Initializes an image object from the contents of a Core Video image buffer.
func ExampleNewImageWithCVImageBuffer() {
	_ = coreimage.NewImageWithCVImageBuffer(
		coreimage.ImageBufferRef /* not a class type */{}, // imageBuffer ImageBufferRef /* not a class type */
	)
	// Output:
}
// ExampleNewImageWithCVPixelBuffer demonstrates how to create a Image instance using NewImageWithCVPixelBuffer.
// Initializes an image object from the contents of a Core Video pixel buffer.
func ExampleNewImageWithCVPixelBuffer() {
	_ = coreimage.NewImageWithCVPixelBuffer(
		coreimage.PixelBufferRef /* not a class type */{}, // pixelBuffer PixelBufferRef /* not a class type */
	)
	// Output:
}
// ExampleNewImageWithIOSurface demonstrates how to create a Image instance using NewImageWithIOSurface.
// Initializes an image with the contents of an IOSurface.
func ExampleNewImageWithIOSurface() {
	_ = coreimage.NewImageWithIOSurface(
		coreimage.SurfaceRef /* not a class type */{}, // surface SurfaceRef /* not a class type */
	)
	// Output:
}
// ExampleImage_AutoAdjustmentFilters demonstrates using AutoAdjustmentFilters on a Image instance.
// Returns all possible automatically selected and configured filters for adjusting the image.
func ExampleImage_AutoAdjustmentFilters() {
	obj := coreimage.NewImage()
	_ = obj.AutoAdjustmentFilters()
	// Output:
	}

// ExampleImage_ImageByClampingToExtent demonstrates using ImageByClampingToExtent on a Image instance.
// Returns a new image created by making the pixel colors along its edges extend infinitely in all directions.
//
// Note: This example is not executed because ImageByClampingToExtent crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByClampingToExtent() {
	obj := coreimage.NewImage()
	_ = obj.ImageByClampingToExtent()
	}

// ExampleImage_ImageByConvertingLabToWorkingSpace demonstrates using ImageByConvertingLabToWorkingSpace on a Image instance.
//
// Note: This example is not executed because ImageByConvertingLabToWorkingSpace crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByConvertingLabToWorkingSpace() {
	obj := coreimage.NewImage()
	_ = obj.ImageByConvertingLabToWorkingSpace()
	}

// ExampleImage_ImageByConvertingWorkingSpaceToLab demonstrates using ImageByConvertingWorkingSpaceToLab on a Image instance.
//
// Note: This example is not executed because ImageByConvertingWorkingSpaceToLab crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByConvertingWorkingSpaceToLab() {
	obj := coreimage.NewImage()
	_ = obj.ImageByConvertingWorkingSpaceToLab()
	}

// ExampleImage_ImageByInsertingIntermediate demonstrates using ImageByInsertingIntermediate on a Image instance.
// Create an image that inserts a intermediate that is cacheable
//
// Note: This example is not executed because ImageByInsertingIntermediate crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByInsertingIntermediate() {
	obj := coreimage.NewImage()
	_ = obj.ImageByInsertingIntermediate()
	}

// ExampleImage_ImageByInsertingTiledIntermediate demonstrates using ImageByInsertingTiledIntermediate on a Image instance.
// Create an image that inserts a intermediate that is cached in tiles
//
// Note: This example is not executed because ImageByInsertingTiledIntermediate crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByInsertingTiledIntermediate() {
	obj := coreimage.NewImage()
	_ = obj.ImageByInsertingTiledIntermediate()
	}

// ExampleImage_ImageByPremultiplyingAlpha demonstrates using ImageByPremultiplyingAlpha on a Image instance.
// Returns a new image created by multiplying the image’s RGB values by its alpha values.
//
// Note: This example is not executed because ImageByPremultiplyingAlpha crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByPremultiplyingAlpha() {
	obj := coreimage.NewImage()
	_ = obj.ImageByPremultiplyingAlpha()
	}

// ExampleImage_ImageBySamplingLinear demonstrates using ImageBySamplingLinear on a Image instance.
// Create an image by changing the receiver’s sample mode to bilinear interpolation.
//
// Note: This example is not executed because ImageBySamplingLinear crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageBySamplingLinear() {
	obj := coreimage.NewImage()
	_ = obj.ImageBySamplingLinear()
	}

// ExampleImage_ImageBySamplingNearest demonstrates using ImageBySamplingNearest on a Image instance.
// Create an image by changing the receiver’s sample mode to nearest neighbor.
//
// Note: This example is not executed because ImageBySamplingNearest crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageBySamplingNearest() {
	obj := coreimage.NewImage()
	_ = obj.ImageBySamplingNearest()
	}

// ExampleImage_ImageByUnpremultiplyingAlpha demonstrates using ImageByUnpremultiplyingAlpha on a Image instance.
// Returns a new image created by dividing the image’s RGB values by its alpha values.
//
// Note: This example is not executed because ImageByUnpremultiplyingAlpha crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleImage_ImageByUnpremultiplyingAlpha() {
	obj := coreimage.NewImage()
	_ = obj.ImageByUnpremultiplyingAlpha()
	}

