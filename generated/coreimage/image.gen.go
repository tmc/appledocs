// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Image] class.
var imageClass = _ImageClass{objc.GetClass("CIImage")}

type _ImageClass struct {
	class objc.Class
}

// A representation of an image to be processed or produced by Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage

type Image struct {
	objectivec.Object
}

// ImageFrom constructs a [Image] from an unsafe.Pointer.
//
// A representation of an image to be processed or produced by Core Image filters.
func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _ImageClass) New() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ Image) Init() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ Image) Autorelease() Image {
	rv := objc.Send[Image](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImage creates a new Image instance.
func NewImage() Image {
	return imageClass.New()
}
// Initializes an image object with a Quartz 2D image, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:options:)
func NewImageWithCGImageOptions(image unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:options:"), image, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents supplied by a CGLayer object, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgLayer:options:)
func NewImageWithCGLayerOptions(layer unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGLayer:options:"), layer, options)
	rv.Autorelease()
	return rv
}
// Initializes an image of infinite extent whose entire content is the specified color. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(color:)
func NewImageWithColor(color unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithColor:"), color)
	rv.Autorelease()
	return rv
}
// Initializes an image with the contents of an IOSurface. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:)
func NewImageWithIOSurface(surface unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:options:)
func NewImageWithSemanticSegmentationMatteOptions(matte unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:options:"), matte, options)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImageSource:index:options:)
func NewImageWithCGImageSourceIndexOptions(source unsafe.Pointer, index uintptr, dict unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImageSource:index:options:"), source, index, dict)
	rv.Autorelease()
	return rv
}
// Initializes an image object by reading an image from a URL, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:options:)
func NewImageWithContentsOfURLOptions(url unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:options:"), url, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents of a Core Video image buffer, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:options:)
func NewImageWithCVImageBufferOptions(imageBuffer unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVImageBuffer:options:"), imageBuffer, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents of a Core Video pixel buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:)
func NewImageWithCVPixelBuffer(pixelBuffer unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVPixelBuffer:"), pixelBuffer)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents of a Core Video pixel buffer using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:options:)
func NewImageWithCVPixelBufferOptions(pixelBuffer unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object with the specified UIKit image object, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:options:)
func NewImageWithImageOptions(image unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object with the supplied image data, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:options:)
func NewImageWithDataOptions(data unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:options:"), data, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object with data supplied by a Metal texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(mtlTexture:options:)
func NewImageWithMTLTextureOptions(texture unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithMTLTexture:options:"), texture, options)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:)
func NewImageWithSemanticSegmentationMatte(matte unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:"), matte)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents supplied by a CGLayer object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgLayer:)
func NewImageWithCGLayer(layer unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGLayer:"), layer)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:)
func NewImageWithDepthData(data unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDepthData:"), data)
	rv.Autorelease()
	return rv
}
// Initializes an image object with data supplied by an OpenGL texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(texture:size:flipped:colorSpace:)
func NewImageWithTextureSizeFlippedColorSpace(name unsafe.Pointer, size unsafe.Pointer, flipped bool, colorSpace unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithTexture:size:flipped:colorSpace:"), name, size, flipped, colorSpace)
	rv.Autorelease()
	return rv
}
// Initializes an image object with a Quartz 2D image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:)
func NewImageWithCGImage(image unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:"), image)
	rv.Autorelease()
	return rv
}
// Initializes an image object by reading an image from a URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:)
func NewImageWithContentsOfURL(url unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}
// Initializes an image object from the contents of a Core Video image buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:)
func NewImageWithCVImageBuffer(imageBuffer unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVImageBuffer:"), imageBuffer)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:options:)
func NewImageWithDepthDataOptions(data unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDepthData:options:"), data, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object based on pixels from an image provider object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(imageProvider:size:_:format:colorSpace:options:)
func NewImageWithImageProviderSizeFormatColorSpaceOptions(provider objc.ID, width uintptr, height uintptr, format unsafe.Pointer, colorSpace unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object with data supplied by an OpenGL texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(texture:size:flipped:options:)
func NewImageWithTextureSizeFlippedOptions(name unsafe.Pointer, size unsafe.Pointer, flipped bool, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithTexture:size:flipped:options:"), name, size, flipped, options)
	rv.Autorelease()
	return rv
}
// Initializes, using the specified format and options, an image with the contents of a specific data plane in an IOSurface. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:plane:format:options:)
func NewImageWithIOSurfacePlaneFormatOptions(surface unsafe.Pointer, plane uintptr, format unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:plane:format:options:"), surface, plane, format, options)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:)
func NewImageWithPortaitEffectsMatte(matte unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPortaitEffectsMatte:"), matte)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:options:)
func NewImageWithPortaitEffectsMatteOptions(matte unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPortaitEffectsMatte:options:"), matte, options)
	rv.Autorelease()
	return rv
}
// Initializes, using the specified options, an image with the contents of an IOSurface. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:options:)
func NewImageWithIOSurfaceOptions(surface unsafe.Pointer, options unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:options:"), surface, options)
	rv.Autorelease()
	return rv
}
// Initializes an image object with bitmap data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapData:bytesPerRow:size:format:colorSpace:)
func NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data unsafe.Pointer, bytesPerRow uintptr, size unsafe.Pointer, format unsafe.Pointer, colorSpace unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	rv.Autorelease()
	return rv
}
// Initializes an image object with the specified bitmap image representation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapImageRep:)
func NewImageWithBitmapImageRep(bitmapImageRep unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithBitmapImageRep:"), bitmapImageRep)
	rv.Autorelease()
	return rv
}
// Initializes an image object with the supplied image data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:)
func NewImageWithData(data unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}
// Initializes an image object with the specified UIKit image object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:)
func NewImageWithImage(image unsafe.Pointer) Image {
	instance := imageClass.Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}


// Creates and returns an empty image object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/empty()
func (ic _ImageClass) EmptyImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("emptyImage"))
	return rv
}
// Creates and returns an image object from bitmap data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithBitmapData:bytesPerRow:size:format:colorSpace:
func (ic _ImageClass) ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data unsafe.Pointer, bytesPerRow uintptr, size unsafe.Pointer, format unsafe.Pointer, colorSpace unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return rv
}
// Creates and returns an image object from a Quartz 2D image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:
func (ic _ImageClass) ImageWithCGImage(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGImage:"), image)
	return rv
}
// Creates and returns an image object from a Quartz 2D image using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:options:
func (ic _ImageClass) ImageWithCGImageOptions(image unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGImage:options:"), image, options)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImageSource:index:options:
func (ic _ImageClass) ImageWithCGImageSourceIndexOptions(source unsafe.Pointer, index uintptr, dict unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGImageSource:index:options:"), source, index, dict)
	return rv
}
// Creates and returns an image object from the contents supplied by a object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGLayer:
func (ic _ImageClass) ImageWithCGLayer(layer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGLayer:"), layer)
	return rv
}
// Creates and returns an image object from the contents supplied by a object, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGLayer:options:
func (ic _ImageClass) ImageWithCGLayerOptions(layer unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCGLayer:options:"), layer, options)
	return rv
}
// Creates and returns an image object from the contents of object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:
func (ic _ImageClass) ImageWithCVImageBuffer(imageBuffer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCVImageBuffer:"), imageBuffer)
	return rv
}
// Creates and returns an image object from the contents of object, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:options:
func (ic _ImageClass) ImageWithCVImageBufferOptions(imageBuffer unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCVImageBuffer:options:"), imageBuffer, options)
	return rv
}
// Creates and returns an image object from the contents of object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:
func (ic _ImageClass) ImageWithCVPixelBuffer(pixelBuffer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCVPixelBuffer:"), pixelBuffer)
	return rv
}
// Creates and returns an image object from the contents of object, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:options:
func (ic _ImageClass) ImageWithCVPixelBufferOptions(pixelBuffer unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}
// Creates and returns an image of infinite extent whose entire content is the specified color. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithColor:
func (ic _ImageClass) ImageWithColor(color unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithColor:"), color)
	return rv
}
// Creates and returns an image object from the contents of a file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:
func (ic _ImageClass) ImageWithContentsOfURL(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithContentsOfURL:"), url)
	return rv
}
// Creates and returns an image object from the contents of a file, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:options:
func (ic _ImageClass) ImageWithContentsOfURLOptions(url unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithContentsOfURL:options:"), url, options)
	return rv
}
// Creates and returns an image object initialized with the supplied image data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:
func (ic _ImageClass) ImageWithData(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithData:"), data)
	return rv
}
// Creates and returns an image object initialized with the supplied image data, using the specified options. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:options:
func (ic _ImageClass) ImageWithDataOptions(data unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithData:options:"), data, options)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:
func (ic _ImageClass) ImageWithDepthData(data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithDepthData:"), data)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:options:
func (ic _ImageClass) ImageWithDepthDataOptions(data unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithDepthData:options:"), data, options)
	return rv
}
// Creates and returns an image from the contents of an IOSurface. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:
func (ic _ImageClass) ImageWithIOSurface(surface unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithIOSurface:"), surface)
	return rv
}
// Creates, using the specified options, and returns an image from the contents of an IOSurface. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:options:
func (ic _ImageClass) ImageWithIOSurfaceOptions(surface unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithIOSurface:options:"), surface, options)
	return rv
}
// Create an image object based on pixels from an image provider object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithImageProvider:size::format:colorSpace:options:
func (ic _ImageClass) ImageWithImageProviderSizeFormatColorSpaceOptions(provider objc.ID, width uintptr, height uintptr, format unsafe.Pointer, colorSpace unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	return rv
}
// Creates and returns an image object with data supplied by a Metal texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithMTLTexture:options:
func (ic _ImageClass) ImageWithMTLTextureOptions(texture unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithMTLTexture:options:"), texture, options)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:
func (ic _ImageClass) ImageWithPortaitEffectsMatte(matte unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithPortaitEffectsMatte:"), matte)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:options:
func (ic _ImageClass) ImageWithPortaitEffectsMatteOptions(matte unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithPortaitEffectsMatte:options:"), matte, options)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:
func (ic _ImageClass) ImageWithSemanticSegmentationMatte(matte unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSemanticSegmentationMatte:"), matte)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:options:
func (ic _ImageClass) ImageWithSemanticSegmentationMatteOptions(matte unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithSemanticSegmentationMatte:options:"), matte, options)
	return rv
}
// Creates and returns an image object initialized with data supplied by an OpenGL texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithTexture:size:flipped:colorSpace:
func (ic _ImageClass) ImageWithTextureSizeFlippedColorSpace(name unsafe.Pointer, size unsafe.Pointer, flipped bool, colorSpace unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithTexture:size:flipped:colorSpace:"), name, size, flipped, colorSpace)
	return rv
}
// Creates and returns an image object initialized with data supplied by an OpenGL texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithTexture:size:flipped:options:
func (ic _ImageClass) ImageWithTextureSizeFlippedOptions(name unsafe.Pointer, size unsafe.Pointer, flipped bool, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageWithTexture:size:flipped:options:"), name, size, flipped, options)
	return rv
}
// Applies the filter to an image and returns the output. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:)
func (i_ Image) ImageByApplyingFilter(filterName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingFilter:"), filterName)
	return rv
}
// Returns a new image created by applying a filter to the original image with the specified name and parameters. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:parameters:)
func (i_ Image) ImageByApplyingFilterWithInputParameters(filterName string, params unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingFilter:withInputParameters:"), filterName, params)
	return rv
}
// Create an image that applies a gain map Core Image image to the received Core Image image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:)
func (i_ Image) ImageByApplyingGainMap(gainmap unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingGainMap:"), gainmap)
	return rv
}
// Create an image that applies a gain map Core Image image with a specified headroom to the received Core Image image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:headroom:)
func (i_ Image) ImageByApplyingGainMapHeadroom(gainmap unsafe.Pointer, headroom float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingGainMap:headroom:"), gainmap, headroom)
	return rv
}
// Create an image by applying a gaussian blur to the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGaussianBlur(sigma:)
func (i_ Image) ImageByApplyingGaussianBlurWithSigma(sigma float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingGaussianBlurWithSigma:"), sigma)
	return rv
}
// Returns all possible automatically selected and configured filters for adjusting the image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters()
func (i_ Image) AutoAdjustmentFilters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("autoAdjustmentFilters"))
	return rv
}
// Returns a subset of automatically selected and configured filters for adjusting the image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters(options:)
func (i_ Image) AutoAdjustmentFiltersWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("autoAdjustmentFiltersWithOptions:"), options)
	return rv
}
// Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clamped(to:)
func (i_ Image) ImageByClampingToRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByClampingToRect:"), rect)
	return rv
}
// Returns a new image created by making the pixel colors along its edges extend infinitely in all directions. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clampedToExtent()
func (i_ Image) ImageByClampingToExtent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByClampingToExtent"))
	return rv
}
// Returns a new image created by compositing the original image over the specified destination image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/composited(over:)
func (i_ Image) ImageByCompositingOverImage(dest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByCompositingOverImage:"), dest)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/convertingLabToWorkingSpace()
func (i_ Image) ImageByConvertingLabToWorkingSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByConvertingLabToWorkingSpace"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/convertingWorkingSpaceToLab()
func (i_ Image) ImageByConvertingWorkingSpaceToLab() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByConvertingWorkingSpaceToLab"))
	return rv
}
// Returns a new image with a cropped portion of the original image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/cropped(to:)
func (i_ Image) ImageByCroppingToRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByCroppingToRect:"), rect)
	return rv
}
// Draws all or part of the image at the specified point in the current coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point unsafe.Pointer, fromRect unsafe.Pointer, op unsafe.Pointer, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}
// Draws all or part of the image in the specified rectangle in the current coordinate system [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect unsafe.Pointer, fromRect unsafe.Pointer, op unsafe.Pointer, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}
// Create an image that inserts a intermediate that is cacheable [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate()
func (i_ Image) ImageByInsertingIntermediate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByInsertingIntermediate"))
	return rv
}
// Create an image that inserts a intermediate that is cacheable. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate(cache:)
func (i_ Image) ImageByInsertingIntermediate(cache bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByInsertingIntermediate:"), cache)
	return rv
}
// Create an image that inserts a intermediate that is cached in tiles [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingTiledIntermediate()
func (i_ Image) ImageByInsertingTiledIntermediate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByInsertingTiledIntermediate"))
	return rv
}
// Returns a new image created by color matching from the context’s working color space to the specified color space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/matchedFromWorkingSpace(to:)
func (i_ Image) ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByColorMatchingWorkingSpaceToColorSpace:"), colorSpace)
	return rv
}
// Returns a new image created by color matching from the specified color space to the context’s working color space. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/matchedToWorkingSpace(from:)
func (i_ Image) ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByColorMatchingColorSpaceToWorkingSpace:"), colorSpace)
	return rv
}
// The affine transform for changing the image to the given orientation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(for:)
func (i_ Image) ImageTransformForCGOrientation(orientation unsafe.Pointer) coregraphics.AffineTransform {
	rv := objc.Send[coregraphics.AffineTransform](i_.ID, objc.Sel("imageTransformForCGOrientation:"), orientation)
	return rv
}
// Returns the transformation needed to reorient the image to the specified orientation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(forExifOrientation:)
func (i_ Image) ImageTransformForOrientation(orientation int) coregraphics.AffineTransform {
	rv := objc.Send[coregraphics.AffineTransform](i_.ID, objc.Sel("imageTransformForOrientation:"), orientation)
	return rv
}
// Transforms the original image by a given orientation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(_:)
func (i_ Image) ImageByApplyingCGOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingCGOrientation:"), orientation)
	return rv
}
// Returns a new image created by transforming the original image to the specified EXIF orientation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(forExifOrientation:)
func (i_ Image) ImageByApplyingOrientation(orientation int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingOrientation:"), orientation)
	return rv
}
// Returns a new image created by multiplying the image’s RGB values by its alpha values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/premultiplyingAlpha()
func (i_ Image) ImageByPremultiplyingAlpha() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByPremultiplyingAlpha"))
	return rv
}
// Returns the region of interest for the filter chain that generates the image. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/regionOfInterest(for:in:)
func (i_ Image) RegionOfInterestForImageInRect(image unsafe.Pointer, rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("regionOfInterestForImage:inRect:"), image, rect)
	return rv
}
// Create an image by changing the receiver’s sample mode to bilinear interpolation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/samplingLinear()
func (i_ Image) ImageBySamplingLinear() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySamplingLinear"))
	return rv
}
// Create an image by changing the receiver’s sample mode to nearest neighbor. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/samplingNearest()
func (i_ Image) ImageBySamplingNearest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySamplingNearest"))
	return rv
}
// Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingAlphaOne(in:)
func (i_ Image) ImageBySettingAlphaOneInExtent(extent unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySettingAlphaOneInExtent:"), extent)
	return rv
}
// Create an image by changing the receiver’s contentAverageLightLevel property. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentAverageLightLevel(_:)
func (i_ Image) ImageBySettingContentAverageLightLevel(average float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySettingContentAverageLightLevel:"), average)
	return rv
}
// Create an image by changing the receiver’s contentHeadroom property. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentHeadroom(_:)
func (i_ Image) ImageBySettingContentHeadroom(headroom float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySettingContentHeadroom:"), headroom)
	return rv
}
// Return a new image by changing the receiver’s metadata properties. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingProperties(_:)
func (i_ Image) ImageBySettingProperties(properties unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageBySettingProperties:"), properties)
	return rv
}
// Returns a new image that represents the original image after applying an affine transform. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:)
func (i_ Image) ImageByApplyingTransform(matrix coregraphics.AffineTransform) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingTransform:"), matrix)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:highQualityDownsample:)
func (i_ Image) ImageByApplyingTransformHighQualityDownsample(matrix coregraphics.AffineTransform, highQualityDownsample bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByApplyingTransform:highQualityDownsample:"), matrix, highQualityDownsample)
	return rv
}
// Returns a new image created by dividing the image’s RGB values by its alpha values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/unpremultiplyingAlpha()
func (i_ Image) ImageByUnpremultiplyingAlpha() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("imageByUnpremultiplyingAlpha"))
	return rv
}

