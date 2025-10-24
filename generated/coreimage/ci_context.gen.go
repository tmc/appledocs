// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [Context] class.
var (
	ContextClass     _ContextClass
	ContextClassOnce sync.Once
)

func getContextClass() _ContextClass {
	ContextClassOnce.Do(func() {
		ContextClass = _ContextClass{objc.GetClass("CIContext")}
	})
	return ContextClass
}

type _ContextClass struct {
	class objc.Class
}





// An interface definition for the [Context] class.
type IContext interface {
	objectivec.IObject
	

	// properties:
	WorkingColorSpace() ColorSpaceRef /* not a class type */
	WorkingFormat() Format /* typedef */
	AlphaMode() RenderDestinationAlphaMode
	SetAlphaMode(value RenderDestinationAlphaMode)


	

	// methods:
	CalculateHDRStatsForCGImage(cgimage ImageRef /* not a class type */) ImageRef /* not a class type */
	CalculateHDRStatsForIOSurface(surface SurfaceRef /* not a class type */)
	CalculateHDRStatsForCVPixelBuffer(buffer PixelBufferRef /* not a class type */)
	CalculateHDRStatsForImage(image ICIImage) IImage
	ClearCaches()
	CreateCGImageFromRect(image ICIImage, fromRect corefoundation.CGRect) ImageRef /* not a class type */
	CreateCGImageFromRectFormatColorSpace(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */) ImageRef /* not a class type */
	CreateCGImageFromRectFormatColorSpaceDeferred(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, deferred bool) ImageRef /* not a class type */
	CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, deferred bool, calculateHDRStats bool) ImageRef /* not a class type */
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, glassesMatte ICIImage, gainMap ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter
	DepthBlurEffectFilterForImageDataOptions(data objc.IObject /* cross-framework: NSData */, options objc.IObject /* cross-framework: NSDictionary */) IFilter
	DepthBlurEffectFilterForImageURLOptions(url objc.IObject /* cross-framework: NSURL */, options objc.IObject /* cross-framework: NSDictionary */) IFilter
	DrawImageInRectFromRect(image ICIImage, inRect corefoundation.CGRect, fromRect corefoundation.CGRect)
	HEIF10RepresentationOfImageColorSpaceOptionsError(image ICIImage, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) foundation.Data
	HEIFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data
	JPEGRepresentationOfImageColorSpaceOptions(image ICIImage, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data
	OpenEXRRepresentationOfImageOptionsError(image ICIImage, options foundation.IDictionary, errorPtr objectivec.IObject) foundation.Data
	PNGRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data
	PrepareRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint, error_ objectivec.IObject) bool
	ReclaimResources()
	RenderToCVPixelBuffer(image ICIImage, buffer PixelBufferRef /* not a class type */)
	RenderToCVPixelBufferBoundsColorSpace(image ICIImage, buffer PixelBufferRef /* not a class type */, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */)
	RenderToIOSurfaceBoundsColorSpace(image ICIImage, surface SurfaceRef /* not a class type */, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */)
	RenderToMTLTextureCommandBufferBoundsColorSpace(image ICIImage, texture unsafe.Pointer, commandBuffer unsafe.Pointer, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */)
	RenderToBitmapRowBytesBoundsFormatColorSpace(image ICIImage, data objectivec.IObject, rowBytes objectivec.IObject, bounds corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */)
	StartTaskToClearError(destination ICIRenderDestination, error_ objectivec.IObject) IRenderTask
	StartTaskToRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint, error_ objectivec.IObject) IRenderTask
	StartTaskToRenderToDestinationError(image ICIImage, destination ICIRenderDestination, error_ objectivec.IObject) IRenderTask
	TIFFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data
	WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool
	WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool
	WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool
	WriteOpenEXRRepresentationOfImageToURLOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, errorPtr objectivec.IObject) bool
	WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool
	WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool


}





// Alloc allocates a new instance without initialization.
func (cc _ContextClass) Alloc() Context {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContextClass) New() Context {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Context) Init() Context {
	rv := objc.Send[Context](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Context) Autorelease() Context {
	rv := objc.Send[Context](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContext creates a new Context instance.
func NewContext() Context {
	return getContextClass().New()
}





// The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL.
//
// You use a instance to render a instance which represents a graph of image processing operations which are built using other Core Image classes, such as , , and . You can also use a with the class to analyze images — for example, to detect faces or barcodes. Contexts support automatic color management by performing all processing operations in a working color space. This means that unless told otherwise: All input images are color matched from the input’s color space to the working space. All renders are color matched from the working space to the destination space. (For more information on see ) and instances are immutable, so multiple threads can use the same instance to render instances. However, instances are mutable and thus cannot be shared safely among threads. Each thread must take case not to access or modify a instance while it is being used by another thread. The manages various internal state such as and caches for compiled kernels and intermediate buffers. For this reason it is not recommended to create many instances. As a rule, it recommended that you create one instance for each view that renders or each background task.


// The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext
type Context struct {
	objectivec.Object
}

// ContextFrom constructs a [Context] from an unsafe.Pointer.
//
// The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL.
func ContextFrom(ptr unsafe.Pointer) Context {
	return Context{objectivec.Object{objc.ID(ptr)}}
}






// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:)
func NewContextForOfflineGPUAtIndex(index objectivec.IObject) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextForOfflineGPUAtIndex:"), index)
	return rv
}


// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display, with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)
func NewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext(index objectivec.IObject, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, sharedContext LContextObj /* not a class type */) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextForOfflineGPUAtIndex:colorSpace:options:sharedContext:"), index, colorSpace, options, sharedContext)
	return rv
}


// Creates a Core Image context from a Quartz context, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cgContext:options:)
func NewContextWithCGContextOptions(cgctx ContextRef /* not a class type */, options foundation.IDictionary) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return rv
}


// Creates a Core Image context from a CGL context, using the specified options, color space, and pixel format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cglContext:pixelFormat:colorSpace:options:)
func NewContextWithCGLContextPixelFormatColorSpaceOptions(cglctx LContextObj /* not a class type */, pixelFormat LPixelFormatObj /* not a class type */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithCGLContext:pixelFormat:colorSpace:options:"), cglctx, pixelFormat, colorSpace, options)
	return rv
}


// Creates a Core Image context from an EAGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:)
func NewContextWithEAGLContext(eaglContext objectivec.IObject) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithEAGLContext:"), eaglContext)
	return rv
}


// Creates a Core Image context from an EAGL context using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:options:)
func NewContextWithEAGLContextOptions(eaglContext objectivec.IObject, options foundation.IDictionary) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithEAGLContext:options:"), eaglContext, options)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:)
func NewContextWithMTLCommandQueue(commandQueue unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLCommandQueue:"), commandQueue)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:options:)
func NewContextWithMTLCommandQueueOptions(commandQueue unsafe.Pointer, options foundation.IDictionary) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLCommandQueue:options:"), commandQueue, options)
	return rv
}


// Creates a Core Image context using the specified Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:)
func NewContextWithMTLDevice(device unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLDevice:"), device)
	return rv
}


// Creates a Core Image context using the specified Metal device and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:options:)
func NewContextWithMTLDeviceOptions(device unsafe.Pointer, options foundation.IDictionary) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLDevice:options:"), device, options)
	return rv
}


// Initializes a context without a specific rendering destination, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(options:)
func NewContextWithOptions(options foundation.IDictionary) Context {
	instance := getContextClass().Alloc()
	rv := objc.Send[Context](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}







// Creates a context without a specific rendering destination, using default options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/context
func (cc _ContextClass) Context() IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("context"))
	return rv
}


// Creates a Core Image context from a CGL context, using the specified options and pixel format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/contextWithCGLContext:pixelFormat:options:
func (cc _ContextClass) ContextWithCGLContextPixelFormatOptions(cglctx LContextObj /* not a class type */, pixelFormat LPixelFormatObj /* not a class type */, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithCGLContext:pixelFormat:options:"), cglctx, pixelFormat, options)
	return rv
}


// Initializes a context without a specific rendering destination, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/contextWithOptions:
func (cc _ContextClass) ContextWithOptions(options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithOptions:"), options)
	return rv
}


// Creates a Core Image context from a Quartz context, using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cgContext:options:)
func (cc _ContextClass) ContextWithCGContextOptions(cgctx ContextRef /* not a class type */, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return rv
}


// Creates a Core Image context from a CGL context, using the specified options, color space, and pixel format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cglContext:pixelFormat:colorSpace:options:)
func (cc _ContextClass) ContextWithCGLContextPixelFormatColorSpaceOptions(cglctx LContextObj /* not a class type */, pixelFormat LPixelFormatObj /* not a class type */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithCGLContext:pixelFormat:colorSpace:options:"), cglctx, pixelFormat, colorSpace, options)
	return rv
}


// Creates a Core Image context from an EAGL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:)
func (cc _ContextClass) ContextWithEAGLContext(eaglContext objectivec.IObject) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithEAGLContext:"), eaglContext)
	return rv
}


// Creates a Core Image context from an EAGL context using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:options:)
func (cc _ContextClass) ContextWithEAGLContextOptions(eaglContext objectivec.IObject, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithEAGLContext:options:"), eaglContext, options)
	return rv
}


// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:)
func (cc _ContextClass) ContextForOfflineGPUAtIndex(index objectivec.IObject) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextForOfflineGPUAtIndex:"), index)
	return rv
}


// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display, with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)
func (cc _ContextClass) ContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext(index objectivec.IObject, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, sharedContext LContextObj /* not a class type */) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextForOfflineGPUAtIndex:colorSpace:options:sharedContext:"), index, colorSpace, options, sharedContext)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:)
func (cc _ContextClass) ContextWithMTLCommandQueue(commandQueue unsafe.Pointer) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithMTLCommandQueue:"), commandQueue)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:options:)
func (cc _ContextClass) ContextWithMTLCommandQueueOptions(commandQueue unsafe.Pointer, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithMTLCommandQueue:options:"), commandQueue, options)
	return rv
}


// Creates a Core Image context using the specified Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:)
func (cc _ContextClass) ContextWithMTLDevice(device unsafe.Pointer) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithMTLDevice:"), device)
	return rv
}


// Creates a Core Image context using the specified Metal device and options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:options:)
func (cc _ContextClass) ContextWithMTLDeviceOptions(device unsafe.Pointer, options foundation.IDictionary) IContext {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("contextWithMTLDevice:options:"), device, options)
	return rv
}


// Returns the number of GPUs not currently driving a display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/offlineGPUCount()
func (cc _ContextClass) OfflineGPUCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("offlineGPUCount"))
	return rv
}












// Given a Core Graphics image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Graphics image that has the calculated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-3ia7r
func (c_ Context) CalculateHDRStatsForCGImage(cgimage ImageRef /* not a class type */) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("calculateHDRStatsForCGImage:"), cgimage)
	return rv
}


// Given an IOSurface, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the surface’s attachments to store the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-6lwmz
func (c_ Context) CalculateHDRStatsForIOSurface(surface SurfaceRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calculateHDRStatsForIOSurface:"), surface)
}


// Given a CVPixelBuffer, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the buffers’s attachments to store the values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-7bcki
func (c_ Context) CalculateHDRStatsForCVPixelBuffer(buffer PixelBufferRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calculateHDRStatsForCVPixelBuffer:"), buffer)
}


// Given a Core Image image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Image image that has the calculated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-l1rj
func (c_ Context) CalculateHDRStatsForImage(image ICIImage) IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("calculateHDRStatsForImage:"), image)
	return rv
}


// Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/clearCaches()
func (c_ Context) ClearCaches() {
	objc.Send[objc.ID](c_.ID, objc.Sel("clearCaches"))
}


// Creates a Core Graphics image from a region of a Core Image image instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:)
func (c_ Context) CreateCGImageFromRect(image ICIImage, fromRect corefoundation.CGRect) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("createCGImage:fromRect:"), image, fromRect)
	return rv
}


// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling the pixel format and color space of the .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:)
func (c_ Context) CreateCGImageFromRectFormatColorSpace(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:"), image, fromRect, format, colorSpace)
	return rv
}


// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling when the image is rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:)
func (c_ Context) CreateCGImageFromRectFormatColorSpaceDeferred(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, deferred bool) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:"), image, fromRect, format, colorSpace, deferred)
	return rv
}


// Creates a Core Graphics image from a region of a Core Image image instance with an option for calculating HDR statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:calculateHDRStats:)
func (c_ Context) CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image ICIImage, fromRect corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, deferred bool, calculateHDRStats bool) ImageRef /* not a class type */ {
	rv := objc.Send[ImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:calculateHDRStats:"), image, fromRect, format, colorSpace, deferred, calculateHDRStats)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, glassesMatte ICIImage, gainMap ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter {
	rv := objc.Send[Filter](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, glassesMatte, gainMap, orientation, options)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, hairSemanticSegmentation ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter {
	rv := objc.Send[Filter](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, orientation, options)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image ICIImage, disparityImage ICIImage, portraitEffectsMatte ICIImage, orientation ImagePropertyOrientation /* not a class type */, options objc.IObject /* cross-framework: NSDictionary */) IFilter {
	rv := objc.Send[Filter](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:orientation:options:"), image, disparityImage, portraitEffectsMatte, orientation, options)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageData:options:)
func (c_ Context) DepthBlurEffectFilterForImageDataOptions(data objc.IObject /* cross-framework: NSData */, options objc.IObject /* cross-framework: NSDictionary */) IFilter {
	rv := objc.Send[Filter](c_.ID, objc.Sel("depthBlurEffectFilterForImageData:options:"), data, options)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageURL:options:)
func (c_ Context) DepthBlurEffectFilterForImageURLOptions(url objc.IObject /* cross-framework: NSURL */, options objc.IObject /* cross-framework: NSDictionary */) IFilter {
	rv := objc.Send[Filter](c_.ID, objc.Sel("depthBlurEffectFilterForImageURL:options:"), url, options)
	return rv
}


// Renders a region of an image to a rectangle in the context destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/draw(_:in:from:)
func (c_ Context) DrawImageInRectFromRect(image ICIImage, inRect corefoundation.CGRect, fromRect corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawImage:inRect:fromRect:"), image, inRect, fromRect)
}


// Renders the image and exports the resulting image data in HEIF10 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/heif10Representation(of:colorSpace:options:)
func (c_ Context) HEIF10RepresentationOfImageColorSpaceOptionsError(image ICIImage, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("HEIF10RepresentationOfImage:colorSpace:options:error:"), image, colorSpace, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data in HEIF format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/heifRepresentation(of:format:colorSpace:options:)
func (c_ Context) HEIFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("HEIFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}


// Renders the image and exports the resulting image data in JPEG format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/jpegRepresentation(of:colorSpace:options:)
func (c_ Context) JPEGRepresentationOfImageColorSpaceOptions(image ICIImage, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("JPEGRepresentationOfImage:colorSpace:options:"), image, colorSpace, options)
	return rv
}


// Renders the image and exports the resulting image data in open EXR format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/openEXRRepresentation(of:options:)
func (c_ Context) OpenEXRRepresentationOfImageOptionsError(image ICIImage, options foundation.IDictionary, errorPtr objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("OpenEXRRepresentationOfImage:options:error:"), image, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data in PNG format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/pngRepresentation(of:format:colorSpace:options:)
func (c_ Context) PNGRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("PNGRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}


// An optional call to warm up a so that subsequent calls to render with the same arguments run more efficiently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/prepareRender(_:from:to:at:)
func (c_ Context) PrepareRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint, error_ objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("prepareRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, error_)
	return rv
}


// Runs the garbage collector to reclaim any resources that the context no longer requires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/reclaimResources()
func (c_ Context) ReclaimResources() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reclaimResources"))
}


// Renders an image into a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:)
func (c_ Context) RenderToCVPixelBuffer(image ICIImage, buffer PixelBufferRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toCVPixelBuffer:"), image, buffer)
}


// Renders a region of an image into a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-2k8l2
func (c_ Context) RenderToCVPixelBufferBoundsColorSpace(image ICIImage, buffer PixelBufferRef /* not a class type */, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toCVPixelBuffer:bounds:colorSpace:"), image, buffer, bounds, colorSpace)
}


// Renders a region of an image into an IOSurface object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-54b9l
func (c_ Context) RenderToIOSurfaceBoundsColorSpace(image ICIImage, surface SurfaceRef /* not a class type */, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toIOSurface:bounds:colorSpace:"), image, surface, bounds, colorSpace)
}


// Renders a region of an image to a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:commandBuffer:bounds:colorSpace:)
func (c_ Context) RenderToMTLTextureCommandBufferBoundsColorSpace(image ICIImage, texture unsafe.Pointer, commandBuffer unsafe.Pointer, bounds corefoundation.CGRect, colorSpace ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toMTLTexture:commandBuffer:bounds:colorSpace:"), image, texture, commandBuffer, bounds, colorSpace)
}


// Renders to the given bitmap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:toBitmap:rowBytes:bounds:format:colorSpace:)
func (c_ Context) RenderToBitmapRowBytesBoundsFormatColorSpace(image ICIImage, data objectivec.IObject, rowBytes objectivec.IObject, bounds corefoundation.CGRect, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toBitmap:rowBytes:bounds:format:colorSpace:"), image, data, rowBytes, bounds, format, colorSpace)
}


// Fills the entire destination with black or clear depending on its .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toClear:)
func (c_ Context) StartTaskToClearError(destination ICIRenderDestination, error_ objectivec.IObject) IRenderTask {
	rv := objc.Send[RenderTask](c_.ID, objc.Sel("startTaskToClear:error:"), destination, error_)
	return rv
}


// Renders a portion of an image to a point in the destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:from:to:at:)
func (c_ Context) StartTaskToRenderFromRectToDestinationAtPointError(image ICIImage, fromRect corefoundation.CGRect, destination ICIRenderDestination, atPoint corefoundation.CGPoint, error_ objectivec.IObject) IRenderTask {
	rv := objc.Send[RenderTask](c_.ID, objc.Sel("startTaskToRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, error_)
	return rv
}


// Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:to:)
func (c_ Context) StartTaskToRenderToDestinationError(image ICIImage, destination ICIRenderDestination, error_ objectivec.IObject) IRenderTask {
	rv := objc.Send[RenderTask](c_.ID, objc.Sel("startTaskToRender:toDestination:error:"), image, destination, error_)
	return rv
}


// Renders the image and exports the resulting image data in TIFF format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/tiffRepresentation(of:format:colorSpace:options:)
func (c_ Context) TIFFRepresentationOfImageFormatColorSpaceOptions(image ICIImage, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary) foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("TIFFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}


// Renders the image and exports the resulting image data as a file in HEIF10 format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIF10Representation(of:to:colorSpace:options:)
func (c_ Context) WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeHEIF10RepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data as a file in HEIF format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIFRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeHEIFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data as a file in JPEG format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeJPEGRepresentation(of:to:colorSpace:options:)
func (c_ Context) WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeJPEGRepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data as a file in open EXR format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeOpenEXRRepresentation(of:to:options:)
func (c_ Context) WriteOpenEXRRepresentationOfImageToURLOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeOpenEXRRepresentationOfImage:toURL:options:error:"), image, url, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data as a file in PNG format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writePNGRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writePNGRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}


// Renders the image and exports the resulting image data as a file in TIFF format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeTIFFRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image ICIImage, url objc.IObject /* cross-framework: NSURL */, format Format /* typedef */, colorSpace ColorSpaceRef /* not a class type */, options foundation.IDictionary, errorPtr objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeTIFFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}







// The working color space of the Core Image context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/workingColorSpace
func (c_ Context) WorkingColorSpace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](c_.ID, objc.Sel("workingColorSpace"))
	return rv
}


// The working pixel format of the Core Image context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/workingFormat
func (c_ Context) WorkingFormat() Format /* typedef */ {
	rv := objc.Send[int32](c_.ID, objc.Sel("workingFormat"))
	return rv
}


// The render destination’s representation of alpha (transparency) values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/alphamode
func (c_ Context) AlphaMode() RenderDestinationAlphaMode {
	rv := objc.Send[RenderDestinationAlphaMode](c_.ID, objc.Sel("alphaMode"))
	return rv
}


// The render destination’s representation of alpha (transparency) values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/alphamode
func (c_ Context) SetAlphaMode(value RenderDestinationAlphaMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlphaMode:"), value)
}







