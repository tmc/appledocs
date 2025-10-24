// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class CIImage */


/* debug [class_header]: Header for CIImage */
// The class instance for the [Image] class.
var (
	ImageClass     _ImageClass
	ImageClassOnce sync.Once
)

func getImageClass() _ImageClass {
	ImageClassOnce.Do(func() {
		ImageClass = _ImageClass{objc.GetClass("CIImage")}
	})
	return ImageClass
}

type _ImageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Image */
// An interface definition for the [Image] class.
type IImage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Image */
	// properties:
	CGImage() ImageRef /* not a class type */
	ColorSpace() ColorSpaceRef /* not a class type */
	ContentAverageLightLevel() float32
	ContentHeadroom() float32
	Definition() ICIFilterShape
	DepthData() objectivec.IObject
	Extent() corefoundation.CGRect
	Opaque() bool
	MetalTexture() unsafe.Pointer
	PixelBuffer() PixelBufferRef /* not a class type */
	PortraitEffectsMatte() objectivec.IObject
	Properties() foundation.IDictionary
	SemanticSegmentationMatte() objectivec.IObject
	Url() objc.IObject /* cross-framework: NSURL */
	IsOpaque() bool
	SetIsOpaque(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Image */
	// methods:
	ImageByApplyingFilter(filterName objc.IObject /* cross-framework: NSString */) IImage
	ImageByApplyingFilterWithInputParameters(filterName objc.IObject /* cross-framework: NSString */, params foundation.IDictionary) IImage
	ImageByApplyingGainMap(gainmap ICIImage) IImage
	ImageByApplyingGainMapHeadroom(gainmap ICIImage, headroom float32) IImage
	ImageByApplyingGaussianBlurWithSigma(sigma float64) IImage
	AutoAdjustmentFilters() []Filter
	AutoAdjustmentFiltersWithOptions(options foundation.IDictionary) []Filter
	ImageByClampingToRect(rect corefoundation.CGRect) IImage
	ImageByClampingToExtent() IImage
	ImageByCompositingOverImage(dest ICIImage) IImage
	ImageByConvertingLabToWorkingSpace() IImage
	ImageByConvertingWorkingSpaceToLab() IImage
	ImageByCroppingToRect(rect corefoundation.CGRect) IImage
	DrawAtPointFromRectOperationFraction(point vision.Point, fromRect Rect /* not a class type */, op CompositingOperation /* not a class type */, delta float64)
	DrawInRectFromRectOperationFraction(rect Rect /* not a class type */, fromRect Rect /* not a class type */, op CompositingOperation /* not a class type */, delta float64)
	ImageByInsertingIntermediate() IImage
	ImageByInsertingIntermediateWithCache(cache bool) IImage
	ImageByInsertingTiledIntermediate() IImage
	ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace ColorSpaceRef /* not a class type */) IImage
	ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace ColorSpaceRef /* not a class type */) IImage
	ImageTransformForCGOrientation(orientation ImagePropertyOrientation /* not a class type */) corefoundation.CGAffineTransform
	ImageTransformForOrientation(orientation int) corefoundation.CGAffineTransform
	ImageByApplyingCGOrientation(orientation ImagePropertyOrientation /* not a class type */) IImage
	ImageByApplyingOrientation(orientation int) IImage
	ImageByPremultiplyingAlpha() IImage
	RegionOfInterestForImageInRect(image ICIImage, rect corefoundation.CGRect) corefoundation.CGRect
	ImageBySamplingLinear() IImage
	ImageBySamplingNearest() IImage
	ImageBySettingAlphaOneInExtent(extent corefoundation.CGRect) IImage
	ImageBySettingContentAverageLightLevel(average float32) IImage
	ImageBySettingContentHeadroom(headroom float32) IImage
	ImageBySettingProperties(properties objc.IObject /* cross-framework: NSDictionary */) IImage
	ImageByApplyingTransform(matrix corefoundation.CGAffineTransform) IImage
	ImageByApplyingTransformHighQualityDownsample(matrix corefoundation.CGAffineTransform, highQualityDownsample bool) IImage
	ImageByUnpremultiplyingAlpha() IImage
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Image */
// Alloc allocates a new instance without initialization.
func (ic _ImageClass) Alloc() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getImageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Image */
// A representation of an image to be processed or produced by Core Image filters.
//
// You use objects in conjunction with other Core Image classes—such as , , , and —to take advantage of the built-in Core Image filters when processing images. You can create objects with data supplied from a variety of sources, including Quartz 2D images, Core Video image buffers ( ), URL-based objects, and objects. Although a object has image data associated with it, it is not an image. You can think of a object as an image “recipe.” A object has all the information necessary to produce an image, but Core Image doesn’t actually render an image until it is told to do so. This lazy evaluation allows Core Image to operate as efficiently as possible. To show a object as an on-screen image, you can display it as a in : and objects are immutable, which means each can be shared safely among threads. Multiple threads can use the same GPU or CPU object to render objects. However, this is not the case for objects, which are mutable. A object cannot be shared safely among threads. If you app is multithreaded, each thread must create its own objects. Otherwise, your app could behave unexpectedly. Core Image also provides auto-adjustment methods. These methods analyze an image for common deficiencies and return a set of filters to correct those deficiencies. The filters are preset with values for improving image quality by altering values for skin tones, saturation, contrast, and shadows and for removing red-eye or other artifacts caused by flash. (See Getting Autoadjustment Filters.) For a discussion of all the methods you can use to create objects on iOS and macOS, see .


// A representation of an image to be processed or produced by Core Image filters.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Image */

// Initializes an image object with bitmap data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapData:bytesPerRow:size:format:colorSpace:)
func NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data objc.IObject /* cross-framework: NSData */, bytesPerRow uintptr /* not a class type */, size corefoundation.CGSize, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace */


// Initializes an image object with the specified bitmap image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(bitmapImageRep:)
func NewImageWithBitmapImageRep(bitmapImageRep objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithBitmapImageRep:"), bitmapImageRep)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithBitmapImageRep */


// Initializes an image object with a Quartz 2D image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:)
func NewImageWithCGImage(image ImageRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGImage */


// Initializes an image object with a Quartz 2D image, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImage:options:)
func NewImageWithCGImageOptions(image ImageRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImage:options:"), image, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGImageOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgImageSource:index:options:)
func NewImageWithCGImageSourceIndexOptions(source ImageSourceRef /* not a class type */, index uintptr /* not a class type */, dict foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGImageSource:index:options:"), source, index, dict)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGImageSourceIndexOptions */


// Initializes an image object from the contents supplied by a CGLayer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgLayer:)
func NewImageWithCGLayer(layer LayerRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGLayer:"), layer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGLayer */


// Initializes an image object from the contents supplied by a CGLayer object, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cgLayer:options:)
func NewImageWithCGLayerOptions(layer LayerRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCGLayer:options:"), layer, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCGLayerOptions */


// Initializes an image object from the contents of a Core Video image buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:)
func NewImageWithCVImageBuffer(imageBuffer ImageBufferRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVImageBuffer:"), imageBuffer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCVImageBuffer */


// Initializes an image object from the contents of a Core Video image buffer, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvImageBuffer:options:)
func NewImageWithCVImageBufferOptions(imageBuffer ImageBufferRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVImageBuffer:options:"), imageBuffer, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCVImageBufferOptions */


// Initializes an image object from the contents of a Core Video pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:)
func NewImageWithCVPixelBuffer(pixelBuffer PixelBufferRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVPixelBuffer:"), pixelBuffer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCVPixelBuffer */


// Initializes an image object from the contents of a Core Video pixel buffer using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(cvPixelBuffer:options:)
func NewImageWithCVPixelBufferOptions(pixelBuffer PixelBufferRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithCVPixelBuffer:options:"), pixelBuffer, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithCVPixelBufferOptions */


// Initializes an image of infinite extent whose entire content is the specified color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(color:)
func NewImageWithColor(color ICIColor) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithColor:"), color)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithColor */


// Initializes an image object by reading an image from a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:)
func NewImageWithContentsOfURL(url objc.IObject /* cross-framework: NSURL */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithContentsOfURL */


// Initializes an image object by reading an image from a URL, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(contentsOf:options:)
func NewImageWithContentsOfURLOptions(url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithContentsOfURL:options:"), url, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithContentsOfURLOptions */


// Initializes an image object with the supplied image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:)
func NewImageWithData(data objc.IObject /* cross-framework: NSData */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithData */


// Initializes an image object with the supplied image data, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(data:options:)
func NewImageWithDataOptions(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithData:options:"), data, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithDataOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:)
func NewImageWithDepthData(data objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDepthData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithDepthData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(depthData:options:)
func NewImageWithDepthDataOptions(data objectivec.IObject, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithDepthData:options:"), data, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithDepthDataOptions */


// Initializes an image with the contents of an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:)
func NewImageWithIOSurface(surface SurfaceRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithIOSurface */


// Initializes, using the specified options, an image with the contents of an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:options:)
func NewImageWithIOSurfaceOptions(surface SurfaceRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:options:"), surface, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithIOSurfaceOptions */


// Initializes, using the specified format and options, an image with the contents of a specific data plane in an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(ioSurface:plane:format:options:)
func NewImageWithIOSurfacePlaneFormatOptions(surface SurfaceRef /* not a class type */, plane uintptr /* not a class type */, format Format /* typedef */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithIOSurface:plane:format:options:"), surface, plane, format, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithIOSurfacePlaneFormatOptions */


// Initializes an image object with the specified UIKit image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:)
func NewImageWithImage(image IImage) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithImage */


// Initializes an image object with the specified UIKit image object, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(image:options:)
func NewImageWithImageOptions(image IImage, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImage:options:"), image, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithImageOptions */


// Initializes an image object based on pixels from an image provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(imageProvider:size:_:format:colorSpace:options:)
func NewImageWithImageProviderSizeFormatColorSpaceOptions(provider objc.IObject, width uintptr /* not a class type */, height uintptr /* not a class type */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithImageProviderSizeFormatColorSpaceOptions */


// Initializes an image object with data supplied by a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(mtlTexture:options:)
func NewImageWithMTLTextureOptions(texture unsafe.Pointer, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithMTLTexture:options:"), texture, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithMTLTextureOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:)
func NewImageWithPortaitEffectsMatte(matte objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPortaitEffectsMatte:"), matte)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithPortaitEffectsMatte */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(portaitEffectsMatte:options:)
func NewImageWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithPortaitEffectsMatte:options:"), matte, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithPortaitEffectsMatteOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:)
func NewImageWithSemanticSegmentationMatte(matte objectivec.IObject) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:"), matte)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSemanticSegmentationMatte */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(semanticSegmentationMatte:options:)
func NewImageWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithSemanticSegmentationMatte:options:"), matte, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithSemanticSegmentationMatteOptions */


// Initializes an image object with data supplied by an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(texture:size:flipped:colorSpace:)
func NewImageWithTextureSizeFlippedColorSpace(name objectivec.IObject, size corefoundation.CGSize, flipped bool, colorSpace ColorSpaceRef /* not a class type */) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithTexture:size:flipped:colorSpace:"), name, size, flipped, colorSpace)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithTextureSizeFlippedColorSpace */


// Initializes an image object with data supplied by an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/init(texture:size:flipped:options:)
func NewImageWithTextureSizeFlippedOptions(name objectivec.IObject, size corefoundation.CGSize, flipped bool, options foundation.IDictionary) Image {
	instance := getImageClass().Alloc()
	rv := objc.Send[Image](instance.ID, objc.Sel("initWithTexture:size:flipped:options:"), name, size, flipped, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageWithTextureSizeFlippedOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Image */

// Creates and returns an empty image object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/empty()
func (ic _ImageClass) EmptyImage() IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("emptyImage"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EmptyImage) */


// Creates and returns an image object from bitmap data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithBitmapData:bytesPerRow:size:format:colorSpace:
func (ic _ImageClass) ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data objc.IObject /* cross-framework: NSData */, bytesPerRow uintptr /* not a class type */, size corefoundation.CGSize, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithBitmapDataBytesPerRowSizeFormatColorSpace) */


// Creates and returns an image object from a Quartz 2D image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:
func (ic _ImageClass) ImageWithCGImage(image ImageRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCGImage:"), image)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCGImage) */


// Creates and returns an image object from a Quartz 2D image using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImage:options:
func (ic _ImageClass) ImageWithCGImageOptions(image ImageRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCGImage:options:"), image, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCGImageOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGImageSource:index:options:
func (ic _ImageClass) ImageWithCGImageSourceIndexOptions(source ImageSourceRef /* not a class type */, index uintptr /* not a class type */, dict foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCGImageSource:index:options:"), source, index, dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCGImageSourceIndexOptions) */


// Creates and returns an image object from the contents supplied by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGLayer:
func (ic _ImageClass) ImageWithCGLayer(layer LayerRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCGLayer:"), layer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCGLayer) */


// Creates and returns an image object from the contents supplied by a object, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCGLayer:options:
func (ic _ImageClass) ImageWithCGLayerOptions(layer LayerRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCGLayer:options:"), layer, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCGLayerOptions) */


// Creates and returns an image of infinite extent whose entire content is the specified color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithColor:
func (ic _ImageClass) ImageWithColor(color ICIColor) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithColor:"), color)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithColor) */


// Creates and returns an image object from the contents of a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:
func (ic _ImageClass) ImageWithContentsOfURL(url objc.IObject /* cross-framework: NSURL */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithContentsOfURL:"), url)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithContentsOfURL) */


// Creates and returns an image object from the contents of a file, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithContentsOfURL:options:
func (ic _ImageClass) ImageWithContentsOfURLOptions(url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithContentsOfURL:options:"), url, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithContentsOfURLOptions) */


// Creates and returns an image object from the contents of object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:
func (ic _ImageClass) ImageWithCVImageBuffer(imageBuffer ImageBufferRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCVImageBuffer:"), imageBuffer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCVImageBuffer) */


// Creates and returns an image object from the contents of object, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVImageBuffer:options:
func (ic _ImageClass) ImageWithCVImageBufferOptions(imageBuffer ImageBufferRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCVImageBuffer:options:"), imageBuffer, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCVImageBufferOptions) */


// Creates and returns an image object from the contents of object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:
func (ic _ImageClass) ImageWithCVPixelBuffer(pixelBuffer PixelBufferRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCVPixelBuffer:"), pixelBuffer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCVPixelBuffer) */


// Creates and returns an image object from the contents of object, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithCVPixelBuffer:options:
func (ic _ImageClass) ImageWithCVPixelBufferOptions(pixelBuffer PixelBufferRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithCVPixelBufferOptions) */


// Creates and returns an image object initialized with the supplied image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:
func (ic _ImageClass) ImageWithData(data objc.IObject /* cross-framework: NSData */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithData) */


// Creates and returns an image object initialized with the supplied image data, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithData:options:
func (ic _ImageClass) ImageWithDataOptions(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithData:options:"), data, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithDataOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:
func (ic _ImageClass) ImageWithDepthData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithDepthData:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithDepthData) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithDepthData:options:
func (ic _ImageClass) ImageWithDepthDataOptions(data objectivec.IObject, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithDepthData:options:"), data, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithDepthDataOptions) */


// Create an image object based on pixels from an image provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithImageProvider:size::format:colorSpace:options:
func (ic _ImageClass) ImageWithImageProviderSizeFormatColorSpaceOptions(provider objc.IObject, width uintptr /* not a class type */, height uintptr /* not a class type */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithImageProvider:size::format:colorSpace:options:"), provider, width, height, format, colorSpace, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithImageProviderSizeFormatColorSpaceOptions) */


// Creates and returns an image from the contents of an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:
func (ic _ImageClass) ImageWithIOSurface(surface SurfaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithIOSurface:"), surface)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithIOSurface) */


// Creates, using the specified options, and returns an image from the contents of an IOSurface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithIOSurface:options:
func (ic _ImageClass) ImageWithIOSurfaceOptions(surface SurfaceRef /* not a class type */, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithIOSurface:options:"), surface, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithIOSurfaceOptions) */


// Creates and returns an image object with data supplied by a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithMTLTexture:options:
func (ic _ImageClass) ImageWithMTLTextureOptions(texture unsafe.Pointer, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithMTLTexture:options:"), texture, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithMTLTextureOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:
func (ic _ImageClass) ImageWithPortaitEffectsMatte(matte objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithPortaitEffectsMatte:"), matte)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithPortaitEffectsMatte) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithPortaitEffectsMatte:options:
func (ic _ImageClass) ImageWithPortaitEffectsMatteOptions(matte objectivec.IObject, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithPortaitEffectsMatte:options:"), matte, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithPortaitEffectsMatteOptions) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:
func (ic _ImageClass) ImageWithSemanticSegmentationMatte(matte objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSemanticSegmentationMatte:"), matte)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSemanticSegmentationMatte) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithSemanticSegmentationMatte:options:
func (ic _ImageClass) ImageWithSemanticSegmentationMatteOptions(matte objectivec.IObject, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("imageWithSemanticSegmentationMatte:options:"), matte, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithSemanticSegmentationMatteOptions) */


// Creates and returns an image object initialized with data supplied by an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithTexture:size:flipped:colorSpace:
func (ic _ImageClass) ImageWithTextureSizeFlippedColorSpace(name objectivec.IObject, size corefoundation.CGSize, flipped bool, colorSpace ColorSpaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithTexture:size:flipped:colorSpace:"), name, size, flipped, colorSpace)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithTextureSizeFlippedColorSpace) */


// Creates and returns an image object initialized with data supplied by an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/imageWithTexture:size:flipped:options:
func (ic _ImageClass) ImageWithTextureSizeFlippedOptions(name objectivec.IObject, size corefoundation.CGSize, flipped bool, options foundation.IDictionary) IImage {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("imageWithTexture:size:flipped:options:"), name, size, flipped, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageWithTextureSizeFlippedOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Image */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/black
func (ic _ImageClass) BlackImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("blackImage"))
	return rv
}/* debug [class_properties_class/property]: blackImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/blue
func (ic _ImageClass) BlueImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("blueImage"))
	return rv
}/* debug [class_properties_class/property]: blueImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clear
func (ic _ImageClass) ClearImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("clearImage"))
	return rv
}/* debug [class_properties_class/property]: clearImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/cyan
func (ic _ImageClass) CyanImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("cyanImage"))
	return rv
}/* debug [class_properties_class/property]: cyanImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/gray
func (ic _ImageClass) GrayImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("grayImage"))
	return rv
}/* debug [class_properties_class/property]: grayImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/green
func (ic _ImageClass) GreenImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("greenImage"))
	return rv
}/* debug [class_properties_class/property]: greenImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/magenta
func (ic _ImageClass) MagentaImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("magentaImage"))
	return rv
}/* debug [class_properties_class/property]: magentaImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/red
func (ic _ImageClass) RedImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("redImage"))
	return rv
}/* debug [class_properties_class/property]: redImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/white
func (ic _ImageClass) WhiteImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("whiteImage"))
	return rv
}/* debug [class_properties_class/property]: whiteImage */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/yellow
func (ic _ImageClass) YellowImage() Image {
	rv := objc.Send[Image](objc.ID(ic.class), objc.Sel("yellowImage"))
	return rv
}/* debug [class_properties_class/property]: yellowImage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Image */

// Applies the filter to an image and returns the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:)
func (i_ Image) ImageByApplyingFilter(filterName objc.IObject /* cross-framework: NSString */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingFilter:"), filterName)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingFilter */


// Returns a new image created by applying a filter to the original image with the specified name and parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingFilter(_:parameters:)
func (i_ Image) ImageByApplyingFilterWithInputParameters(filterName objc.IObject /* cross-framework: NSString */, params foundation.IDictionary) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingFilter:withInputParameters:"), filterName, params)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingFilterWithInputParameters */


// Create an image that applies a gain map Core Image image to the received Core Image image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:)
func (i_ Image) ImageByApplyingGainMap(gainmap ICIImage) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingGainMap:"), gainmap)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingGainMap */


// Create an image that applies a gain map Core Image image with a specified headroom to the received Core Image image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGainMap(_:headroom:)
func (i_ Image) ImageByApplyingGainMapHeadroom(gainmap ICIImage, headroom float32) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingGainMap:headroom:"), gainmap, headroom)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingGainMapHeadroom */


// Create an image by applying a gaussian blur to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/applyingGaussianBlur(sigma:)
func (i_ Image) ImageByApplyingGaussianBlurWithSigma(sigma float64) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingGaussianBlurWithSigma:"), sigma)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingGaussianBlurWithSigma */


// Returns all possible automatically selected and configured filters for adjusting the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters()
func (i_ Image) AutoAdjustmentFilters() []Filter {
	rv := objc.Send[[]Filter](i_.ID, objc.Sel("autoAdjustmentFilters"))
	return rv
}/* debug [instance_methods/method]: AutoAdjustmentFilters */


// Returns a subset of automatically selected and configured filters for adjusting the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/autoAdjustmentFilters(options:)
func (i_ Image) AutoAdjustmentFiltersWithOptions(options foundation.IDictionary) []Filter {
	rv := objc.Send[[]Filter](i_.ID, objc.Sel("autoAdjustmentFiltersWithOptions:"), options)
	return rv
}/* debug [instance_methods/method]: AutoAdjustmentFiltersWithOptions */


// Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clamped(to:)
func (i_ Image) ImageByClampingToRect(rect corefoundation.CGRect) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByClampingToRect:"), rect)
	return rv
}/* debug [instance_methods/method]: ImageByClampingToRect */


// Returns a new image created by making the pixel colors along its edges extend infinitely in all directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clampedToExtent()
func (i_ Image) ImageByClampingToExtent() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByClampingToExtent"))
	return rv
}/* debug [instance_methods/method]: ImageByClampingToExtent */


// Returns a new image created by compositing the original image over the specified destination image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/composited(over:)
func (i_ Image) ImageByCompositingOverImage(dest ICIImage) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByCompositingOverImage:"), dest)
	return rv
}/* debug [instance_methods/method]: ImageByCompositingOverImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/convertingLabToWorkingSpace()
func (i_ Image) ImageByConvertingLabToWorkingSpace() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByConvertingLabToWorkingSpace"))
	return rv
}/* debug [instance_methods/method]: ImageByConvertingLabToWorkingSpace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/convertingWorkingSpaceToLab()
func (i_ Image) ImageByConvertingWorkingSpaceToLab() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByConvertingWorkingSpaceToLab"))
	return rv
}/* debug [instance_methods/method]: ImageByConvertingWorkingSpaceToLab */


// Returns a new image with a cropped portion of the original image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/cropped(to:)
func (i_ Image) ImageByCroppingToRect(rect corefoundation.CGRect) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByCroppingToRect:"), rect)
	return rv
}/* debug [instance_methods/method]: ImageByCroppingToRect */


// Draws all or part of the image at the specified point in the current coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/draw(at:from:operation:fraction:)
func (i_ Image) DrawAtPointFromRectOperationFraction(point vision.Point, fromRect Rect /* not a class type */, op CompositingOperation /* not a class type */, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}/* debug [instance_methods/method]: DrawAtPointFromRectOperationFraction */


// Draws all or part of the image in the specified rectangle in the current coordinate system
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/draw(in:from:operation:fraction:)
func (i_ Image) DrawInRectFromRectOperationFraction(rect Rect /* not a class type */, fromRect Rect /* not a class type */, op CompositingOperation /* not a class type */, delta float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}/* debug [instance_methods/method]: DrawInRectFromRectOperationFraction */


// Create an image that inserts a intermediate that is cacheable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate()
func (i_ Image) ImageByInsertingIntermediate() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByInsertingIntermediate"))
	return rv
}/* debug [instance_methods/method]: ImageByInsertingIntermediate */


// Create an image that inserts a intermediate that is cacheable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingIntermediate(cache:)
func (i_ Image) ImageByInsertingIntermediateWithCache(cache bool) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByInsertingIntermediate:"), cache)
	return rv
}/* debug [instance_methods/method]: ImageByInsertingIntermediateWithCache */


// Create an image that inserts a intermediate that is cached in tiles
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/insertingTiledIntermediate()
func (i_ Image) ImageByInsertingTiledIntermediate() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByInsertingTiledIntermediate"))
	return rv
}/* debug [instance_methods/method]: ImageByInsertingTiledIntermediate */


// Returns a new image created by color matching from the context’s working color space to the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/matchedFromWorkingSpace(to:)
func (i_ Image) ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace ColorSpaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByColorMatchingWorkingSpaceToColorSpace:"), colorSpace)
	return rv
}/* debug [instance_methods/method]: ImageByColorMatchingWorkingSpaceToColorSpace */


// Returns a new image created by color matching from the specified color space to the context’s working color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/matchedToWorkingSpace(from:)
func (i_ Image) ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace ColorSpaceRef /* not a class type */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByColorMatchingColorSpaceToWorkingSpace:"), colorSpace)
	return rv
}/* debug [instance_methods/method]: ImageByColorMatchingColorSpaceToWorkingSpace */


// The affine transform for changing the image to the given orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(for:)
func (i_ Image) ImageTransformForCGOrientation(orientation ImagePropertyOrientation /* not a class type */) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i_.ID, objc.Sel("imageTransformForCGOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ImageTransformForCGOrientation */


// Returns the transformation needed to reorient the image to the specified orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/orientationTransform(forExifOrientation:)
func (i_ Image) ImageTransformForOrientation(orientation int) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](i_.ID, objc.Sel("imageTransformForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ImageTransformForOrientation */


// Transforms the original image by a given orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(_:)
func (i_ Image) ImageByApplyingCGOrientation(orientation ImagePropertyOrientation /* not a class type */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingCGOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingCGOrientation */


// Returns a new image created by transforming the original image to the specified EXIF orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/oriented(forExifOrientation:)
func (i_ Image) ImageByApplyingOrientation(orientation int) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingOrientation */


// Returns a new image created by multiplying the image’s RGB values by its alpha values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/premultiplyingAlpha()
func (i_ Image) ImageByPremultiplyingAlpha() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByPremultiplyingAlpha"))
	return rv
}/* debug [instance_methods/method]: ImageByPremultiplyingAlpha */


// Returns the region of interest for the filter chain that generates the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/regionOfInterest(for:in:)
func (i_ Image) RegionOfInterestForImageInRect(image ICIImage, rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("regionOfInterestForImage:inRect:"), image, rect)
	return rv
}/* debug [instance_methods/method]: RegionOfInterestForImageInRect */


// Create an image by changing the receiver’s sample mode to bilinear interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/samplingLinear()
func (i_ Image) ImageBySamplingLinear() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySamplingLinear"))
	return rv
}/* debug [instance_methods/method]: ImageBySamplingLinear */


// Create an image by changing the receiver’s sample mode to nearest neighbor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/samplingNearest()
func (i_ Image) ImageBySamplingNearest() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySamplingNearest"))
	return rv
}/* debug [instance_methods/method]: ImageBySamplingNearest */


// Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingAlphaOne(in:)
func (i_ Image) ImageBySettingAlphaOneInExtent(extent corefoundation.CGRect) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySettingAlphaOneInExtent:"), extent)
	return rv
}/* debug [instance_methods/method]: ImageBySettingAlphaOneInExtent */


// Create an image by changing the receiver’s contentAverageLightLevel property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentAverageLightLevel(_:)
func (i_ Image) ImageBySettingContentAverageLightLevel(average float32) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySettingContentAverageLightLevel:"), average)
	return rv
}/* debug [instance_methods/method]: ImageBySettingContentAverageLightLevel */


// Create an image by changing the receiver’s contentHeadroom property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingContentHeadroom(_:)
func (i_ Image) ImageBySettingContentHeadroom(headroom float32) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySettingContentHeadroom:"), headroom)
	return rv
}/* debug [instance_methods/method]: ImageBySettingContentHeadroom */


// Return a new image by changing the receiver’s metadata properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/settingProperties(_:)
func (i_ Image) ImageBySettingProperties(properties objc.IObject /* cross-framework: NSDictionary */) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageBySettingProperties:"), properties)
	return rv
}/* debug [instance_methods/method]: ImageBySettingProperties */


// Returns a new image that represents the original image after applying an affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:)
func (i_ Image) ImageByApplyingTransform(matrix corefoundation.CGAffineTransform) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingTransform:"), matrix)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/transformed(by:highQualityDownsample:)
func (i_ Image) ImageByApplyingTransformHighQualityDownsample(matrix corefoundation.CGAffineTransform, highQualityDownsample bool) IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByApplyingTransform:highQualityDownsample:"), matrix, highQualityDownsample)
	return rv
}/* debug [instance_methods/method]: ImageByApplyingTransformHighQualityDownsample */


// Returns a new image created by dividing the image’s RGB values by its alpha values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/unpremultiplyingAlpha()
func (i_ Image) ImageByUnpremultiplyingAlpha() IImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("imageByUnpremultiplyingAlpha"))
	return rv
}/* debug [instance_methods/method]: ImageByUnpremultiplyingAlpha */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Image */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/black
func (i_ Image) BlackImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("blackImage"))
	return rv
}/* debug [instance_properties/getter]: blackImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/blue
func (i_ Image) BlueImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("blueImage"))
	return rv
}/* debug [instance_properties/getter]: blueImage */


// The CoreGraphics image object this image was created from, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/cgImage
func (i_ Image) CGImage() ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](i_.ID, objc.Sel("CGImage"))
	return rv
}/* debug [instance_properties/getter]: CGImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/clear
func (i_ Image) ClearImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("clearImage"))
	return rv
}/* debug [instance_properties/getter]: clearImage */


// The color space of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/colorSpace
func (i_ Image) ColorSpace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](i_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// Returns the content average light level of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/contentAverageLightLevel
func (i_ Image) ContentAverageLightLevel() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("contentAverageLightLevel"))
	return rv
}/* debug [instance_properties/getter]: contentAverageLightLevel */


// Returns the content headroom of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/contentHeadroom
func (i_ Image) ContentHeadroom() float32 {
	rv := objc.Send[float32](i_.ID, objc.Sel("contentHeadroom"))
	return rv
}/* debug [instance_properties/getter]: contentHeadroom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/cyan
func (i_ Image) CyanImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("cyanImage"))
	return rv
}/* debug [instance_properties/getter]: cyanImage */


// Returns a filter shape object that represents the domain of definition of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/definition
func (i_ Image) Definition() ICIFilterShape {
	rv := objc.Send[FilterShape](i_.ID, objc.Sel("definition"))
	return rv
}/* debug [instance_properties/getter]: definition */


// Depth data associated with the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/depthData
func (i_ Image) DepthData() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("depthData"))
	return rv
}/* debug [instance_properties/getter]: depthData */


// A rectangle that specifies the extent of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/extent
func (i_ Image) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i_.ID, objc.Sel("extent"))
	return rv
}/* debug [instance_properties/getter]: extent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/gray
func (i_ Image) GrayImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("grayImage"))
	return rv
}/* debug [instance_properties/getter]: grayImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/green
func (i_ Image) GreenImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("greenImage"))
	return rv
}/* debug [instance_properties/getter]: greenImage */


// Returns YES if the image is known to have and alpha value of over the entire image extent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/isOpaque
func (i_ Image) Opaque() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/magenta
func (i_ Image) MagentaImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("magentaImage"))
	return rv
}/* debug [instance_properties/getter]: magentaImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/metalTexture
func (i_ Image) MetalTexture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("metalTexture"))
	return rv
}/* debug [instance_properties/getter]: metalTexture */


// The CoreVideo pixel buffer this image was created from, if applicable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/pixelBuffer
func (i_ Image) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](i_.ID, objc.Sel("pixelBuffer"))
	return rv
}/* debug [instance_properties/getter]: pixelBuffer */


// The portrait effects matte associated with the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/portraitEffectsMatte
func (i_ Image) PortraitEffectsMatte() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("portraitEffectsMatte"))
	return rv
}/* debug [instance_properties/getter]: portraitEffectsMatte */


// Returns the metadata properties dictionary of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/properties
func (i_ Image) Properties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](i_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/red
func (i_ Image) RedImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("redImage"))
	return rv
}/* debug [instance_properties/getter]: redImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/semanticSegmentationMatte
func (i_ Image) SemanticSegmentationMatte() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("semanticSegmentationMatte"))
	return rv
}/* debug [instance_properties/getter]: semanticSegmentationMatte */


// The URL from which the image was loaded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/url
func (i_ Image) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](i_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/white
func (i_ Image) WhiteImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("whiteImage"))
	return rv
}/* debug [instance_properties/getter]: whiteImage */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIImage/yellow
func (i_ Image) YellowImage() ICIImage {
	rv := objc.Send[Image](i_.ID, objc.Sel("yellowImage"))
	return rv
}/* debug [instance_properties/getter]: yellowImage */


// Returns YES if the image is known to have and alpha value of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/isopaque
func (i_ Image) IsOpaque() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// Returns YES if the image is known to have and alpha value of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/isopaque
func (i_ Image) SetIsOpaque(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIImage */


