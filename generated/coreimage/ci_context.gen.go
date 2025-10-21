// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
	CalculateHDRStatsForCGImage(cgimage coregraphics.CGImageRef) coregraphics.CGImageRef
	CalculateHDRStatsForIOSurface(surface unsafe.Pointer)
	CalculateHDRStatsForCVPixelBuffer(buffer unsafe.Pointer)
	CalculateHDRStatsForImage(image unsafe.Pointer) unsafe.Pointer
	ClearCaches()
	CreateCGImageFromRect(image unsafe.Pointer, fromRect coregraphics.CGRect) coregraphics.CGImageRef
	CreateCGImageFromRectFormatColorSpace(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) coregraphics.CGImageRef
	CreateCGImageFromRectFormatColorSpaceDeferred(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, deferred bool) coregraphics.CGImageRef
	CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, deferred bool, calculateHDRStats bool) coregraphics.CGImageRef
	CreateCGLayerWithSizeInfo(size coregraphics.CGSize, info unsafe.Pointer) coregraphics.CGLayerRef
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, hairSemanticSegmentation unsafe.Pointer, glassesMatte unsafe.Pointer, gainMap unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, hairSemanticSegmentation unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer
	DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer
	DepthBlurEffectFilterForImageDataOptions(data unsafe.Pointer, options objc.ID) unsafe.Pointer
	DepthBlurEffectFilterForImageURLOptions(url foundation.URL, options objc.ID) unsafe.Pointer
	DrawImageAtPointFromRect(image unsafe.Pointer, atPoint coregraphics.CGPoint, fromRect coregraphics.CGRect)
	DrawImageInRectFromRect(image unsafe.Pointer, inRect coregraphics.CGRect, fromRect coregraphics.CGRect)
	HEIF10RepresentationOfImageColorSpaceOptionsError(image unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer
	HEIFRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer
	InputImageMaximumSize() coregraphics.CGSize
	JPEGRepresentationOfImageColorSpaceOptions(image unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer
	OpenEXRRepresentationOfImageOptionsError(image unsafe.Pointer, options unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer
	OutputImageMaximumSize() coregraphics.CGSize
	PNGRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer
	PrepareRenderFromRectToDestinationAtPointError(image unsafe.Pointer, fromRect coregraphics.CGRect, destination unsafe.Pointer, atPoint coregraphics.CGPoint, error_ unsafe.Pointer) bool
	ReclaimResources()
	RenderToCVPixelBuffer(image unsafe.Pointer, buffer unsafe.Pointer)
	RenderToCVPixelBufferBoundsColorSpace(image unsafe.Pointer, buffer unsafe.Pointer, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef)
	RenderToIOSurfaceBoundsColorSpace(image unsafe.Pointer, surface unsafe.Pointer, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef)
	RenderToMTLTextureCommandBufferBoundsColorSpace(image unsafe.Pointer, texture objc.ID, commandBuffer objc.ID, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef)
	RenderToBitmapRowBytesBoundsFormatColorSpace(image unsafe.Pointer, data unsafe.Pointer, rowBytes unsafe.Pointer, bounds coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef)
	StartTaskToClearError(destination unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	StartTaskToRenderFromRectToDestinationAtPointError(image unsafe.Pointer, fromRect coregraphics.CGRect, destination unsafe.Pointer, atPoint coregraphics.CGPoint, error_ unsafe.Pointer) unsafe.Pointer
	StartTaskToRenderToDestinationError(image unsafe.Pointer, destination unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	TIFFRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer
	WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
	WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
	WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
	WriteOpenEXRRepresentationOfImageToURLOptionsError(image unsafe.Pointer, url foundation.URL, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
	WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
	WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool
}

// The Core Image context class provides an evaluation context for Core Image processing with Metal, OpenGL, or OpenCL.
//
// You use a instance to render a instance which represents a graph of image processing operations which are built using other Core Image classes, such as , , and . You can also use a with the class to analyze images — for example, to detect faces or barcodes. Contexts support automatic color management by performing all processing operations in a working color space. This means that unless told otherwise: All input images are color matched from the input’s color space to the working space. All renders are color matched from the working space to the destination space. (For more information on see ) and instances are immutable, so multiple threads can use the same instance to render instances. However, instances are mutable and thus cannot be shared safely among threads. Each thread must take case not to access or modify a instance while it is being used by another thread. The manages various internal state such as and caches for compiled kernels and intermediate buffers. For this reason it is not recommended to create many instances. As a rule, it recommended that you create one instance for each view that renders or each background task.
//
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

// Alloc allocates a new instance without initialization.
func (cc _ContextClass) Alloc() Context {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:)
func NewContextForOfflineGPUAtIndex(index unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextForOfflineGPUAtIndex:"), index)
	return rv
}



// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display, with the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)
func NewContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext(index unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, sharedContext unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextForOfflineGPUAtIndex:colorSpace:options:sharedContext:"), index, colorSpace, options, sharedContext)
	return rv
}



// Creates a Core Image context from a Quartz context, using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cgContext:options:)
func NewContextWithCGContextOptions(cgctx coregraphics.CGContextRef, options unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return rv
}



// Creates a Core Image context from a CGL context, using the specified options, color space, and pixel format object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cglContext:pixelFormat:colorSpace:options:)
func NewContextWithCGLContextPixelFormatColorSpaceOptions(cglctx unsafe.Pointer, pixelFormat unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithCGLContext:pixelFormat:colorSpace:options:"), cglctx, pixelFormat, colorSpace, options)
	return rv
}



// Creates a Core Image context from an EAGL context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:)
func NewContextWithEAGLContext(eaglContext unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithEAGLContext:"), eaglContext)
	return rv
}



// Creates a Core Image context from an EAGL context using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:options:)
func NewContextWithEAGLContextOptions(eaglContext unsafe.Pointer, options unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithEAGLContext:options:"), eaglContext, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:)
func NewContextWithMTLCommandQueue(commandQueue objc.ID) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLCommandQueue:"), commandQueue)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:options:)
func NewContextWithMTLCommandQueueOptions(commandQueue objc.ID, options unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLCommandQueue:options:"), commandQueue, options)
	return rv
}



// Creates a Core Image context using the specified Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:)
func NewContextWithMTLDevice(device objc.ID) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLDevice:"), device)
	return rv
}



// Creates a Core Image context using the specified Metal device and options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:options:)
func NewContextWithMTLDeviceOptions(device objc.ID, options unsafe.Pointer) Context {
	rv := objc.Send[Context](objc.ID(getContextClass().class), objc.Sel("contextWithMTLDevice:options:"), device, options)
	return rv
}



// Initializes a context without a specific rendering destination, using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(options:)
func NewContextWithOptions(options unsafe.Pointer) Context {
	instance := getContextClass().Alloc()
	rv := objc.Send[Context](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}


// Creates a context without a specific rendering destination, using default options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/context
func (cc _ContextClass) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("context"))
	return rv
}

// Creates a Core Image context from a CGL context, using the specified options and pixel format object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/contextWithCGLContext:pixelFormat:options:
func (cc _ContextClass) ContextWithCGLContextPixelFormatOptions(cglctx unsafe.Pointer, pixelFormat unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithCGLContext:pixelFormat:options:"), cglctx, pixelFormat, options)
	return rv
}

// Initializes a context without a specific rendering destination, using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/contextWithOptions:
func (cc _ContextClass) ContextWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithOptions:"), options)
	return rv
}

// Creates a Core Image context from a Quartz context, using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cgContext:options:)
func (cc _ContextClass) ContextWithCGContextOptions(cgctx coregraphics.CGContextRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithCGContext:options:"), cgctx, options)
	return rv
}

// Creates a Core Image context from a CGL context, using the specified options, color space, and pixel format object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(cglContext:pixelFormat:colorSpace:options:)
func (cc _ContextClass) ContextWithCGLContextPixelFormatColorSpaceOptions(cglctx unsafe.Pointer, pixelFormat unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithCGLContext:pixelFormat:colorSpace:options:"), cglctx, pixelFormat, colorSpace, options)
	return rv
}

// Creates a Core Image context from an EAGL context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:)
func (cc _ContextClass) ContextWithEAGLContext(eaglContext unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithEAGLContext:"), eaglContext)
	return rv
}

// Creates a Core Image context from an EAGL context using the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(eaglContext:options:)
func (cc _ContextClass) ContextWithEAGLContextOptions(eaglContext unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithEAGLContext:options:"), eaglContext, options)
	return rv
}

// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:)
func (cc _ContextClass) ContextForOfflineGPUAtIndex(index unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextForOfflineGPUAtIndex:"), index)
	return rv
}

// Creates an OpenGL-based Core Image context using a GPU that is not currently driving a display, with the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(forOfflineGPUAtIndex:colorSpace:options:sharedContext:)
func (cc _ContextClass) ContextForOfflineGPUAtIndexColorSpaceOptionsSharedContext(index unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, sharedContext unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextForOfflineGPUAtIndex:colorSpace:options:sharedContext:"), index, colorSpace, options, sharedContext)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:)
func (cc _ContextClass) ContextWithMTLCommandQueue(commandQueue objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithMTLCommandQueue:"), commandQueue)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlCommandQueue:options:)
func (cc _ContextClass) ContextWithMTLCommandQueueOptions(commandQueue objc.ID, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithMTLCommandQueue:options:"), commandQueue, options)
	return rv
}

// Creates a Core Image context using the specified Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:)
func (cc _ContextClass) ContextWithMTLDevice(device objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithMTLDevice:"), device)
	return rv
}

// Creates a Core Image context using the specified Metal device and options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/init(mtlDevice:options:)
func (cc _ContextClass) ContextWithMTLDeviceOptions(device objc.ID, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("contextWithMTLDevice:options:"), device, options)
	return rv
}

// Returns the number of GPUs not currently driving a display.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/offlineGPUCount()
func (cc _ContextClass) OfflineGPUCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("offlineGPUCount"))
	return rv
}

// Given a Core Graphics image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Graphics image that has the calculated values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-3ia7r
func (c_ Context) CalculateHDRStatsForCGImage(cgimage coregraphics.CGImageRef) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("calculateHDRStatsForCGImage:"), cgimage)
	return rv
}

// Given an IOSurface, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the surface’s attachments to store the values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-6lwmz
func (c_ Context) CalculateHDRStatsForIOSurface(surface unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calculateHDRStatsForIOSurface:"), surface)
}

// Given a CVPixelBuffer, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then update the buffers’s attachments to store the values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-7bcki
func (c_ Context) CalculateHDRStatsForCVPixelBuffer(buffer unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("calculateHDRStatsForCVPixelBuffer:"), buffer)
}

// Given a Core Image image, use the receiving Core Image context to calculate its HDR statistics (content headroom and content average light level) and then return a new Core Image image that has the calculated values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/calculateHDRStats(for:)-l1rj
func (c_ Context) CalculateHDRStatsForImage(image unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("calculateHDRStatsForImage:"), image)
	return rv
}

// Frees any cached data, such as temporary images, associated with the context and runs the garbage collector.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/clearCaches()
func (c_ Context) ClearCaches() {
	objc.Send[objc.ID](c_.ID, objc.Sel("clearCaches"))
}

// Creates a Core Graphics image from a region of a Core Image image instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:)
func (c_ Context) CreateCGImageFromRect(image unsafe.Pointer, fromRect coregraphics.CGRect) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("createCGImage:fromRect:"), image, fromRect)
	return rv
}

// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling the pixel format and color space of the .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:)
func (c_ Context) CreateCGImageFromRectFormatColorSpace(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:"), image, fromRect, format, colorSpace)
	return rv
}

// Creates a Core Graphics image from a region of a Core Image image instance with an option for controlling when the image is rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:)
func (c_ Context) CreateCGImageFromRectFormatColorSpaceDeferred(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, deferred bool) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:"), image, fromRect, format, colorSpace, deferred)
	return rv
}

// Creates a Core Graphics image from a region of a Core Image image instance with an option for calculating HDR statistics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGImage(_:from:format:colorSpace:deferred:calculateHDRStats:)
func (c_ Context) CreateCGImageFromRectFormatColorSpaceDeferredCalculateHDRStats(image unsafe.Pointer, fromRect coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, deferred bool, calculateHDRStats bool) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c_.ID, objc.Sel("createCGImage:fromRect:format:colorSpace:deferred:calculateHDRStats:"), image, fromRect, format, colorSpace, deferred, calculateHDRStats)
	return rv
}

// Creates a CGLayer object from the provided parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/createCGLayer(with:info:)
func (c_ Context) CreateCGLayerWithSizeInfo(size coregraphics.CGSize, info unsafe.Pointer) coregraphics.CGLayerRef {
	rv := objc.Send[coregraphics.CGLayerRef](c_.ID, objc.Sel("createCGLayerWithSize:info:"), size, info)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationGlassesMatteGainMapOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, hairSemanticSegmentation unsafe.Pointer, glassesMatte unsafe.Pointer, gainMap unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:glassesMatte:gainMap:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, glassesMatte, gainMap, orientation, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteHairSemanticSegmentationOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, hairSemanticSegmentation unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:hairSemanticSegmentation:orientation:options:"), image, disparityImage, portraitEffectsMatte, hairSemanticSegmentation, orientation, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(for:disparityImage:portraitEffectsMatte:orientation:options:)
func (c_ Context) DepthBlurEffectFilterForImageDisparityImagePortraitEffectsMatteOrientationOptions(image unsafe.Pointer, disparityImage unsafe.Pointer, portraitEffectsMatte unsafe.Pointer, orientation unsafe.Pointer, options objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("depthBlurEffectFilterForImage:disparityImage:portraitEffectsMatte:orientation:options:"), image, disparityImage, portraitEffectsMatte, orientation, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageData:options:)
func (c_ Context) DepthBlurEffectFilterForImageDataOptions(data unsafe.Pointer, options objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("depthBlurEffectFilterForImageData:options:"), data, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/depthBlurEffectFilter(forImageURL:options:)
func (c_ Context) DepthBlurEffectFilterForImageURLOptions(url foundation.URL, options objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("depthBlurEffectFilterForImageURL:options:"), url, options)
	return rv
}

// Renders a region of an image to a point in the context destination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/draw(_:at:from:)
func (c_ Context) DrawImageAtPointFromRect(image unsafe.Pointer, atPoint coregraphics.CGPoint, fromRect coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawImage:atPoint:fromRect:"), image, atPoint, fromRect)
}

// Renders a region of an image to a rectangle in the context destination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/draw(_:in:from:)
func (c_ Context) DrawImageInRectFromRect(image unsafe.Pointer, inRect coregraphics.CGRect, fromRect coregraphics.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("drawImage:inRect:fromRect:"), image, inRect, fromRect)
}

// Renders the image and exports the resulting image data in HEIF10 format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/heif10Representation(of:colorSpace:options:)
func (c_ Context) HEIF10RepresentationOfImageColorSpaceOptionsError(image unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("HEIF10RepresentationOfImage:colorSpace:options:error:"), image, colorSpace, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data in HEIF format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/heifRepresentation(of:format:colorSpace:options:)
func (c_ Context) HEIFRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("HEIFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}

// Returns the maximum size allowed for any image rendered into the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/inputImageMaximumSize()
func (c_ Context) InputImageMaximumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("inputImageMaximumSize"))
	return rv
}

// Renders the image and exports the resulting image data in JPEG format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/jpegRepresentation(of:colorSpace:options:)
func (c_ Context) JPEGRepresentationOfImageColorSpaceOptions(image unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("JPEGRepresentationOfImage:colorSpace:options:"), image, colorSpace, options)
	return rv
}

// Renders the image and exports the resulting image data in open EXR format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/openEXRRepresentation(of:options:)
func (c_ Context) OpenEXRRepresentationOfImageOptionsError(image unsafe.Pointer, options unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("OpenEXRRepresentationOfImage:options:error:"), image, options, errorPtr)
	return rv
}

// Returns the maximum size allowed for any image created by the context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/outputImageMaximumSize()
func (c_ Context) OutputImageMaximumSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("outputImageMaximumSize"))
	return rv
}

// Renders the image and exports the resulting image data in PNG format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/pngRepresentation(of:format:colorSpace:options:)
func (c_ Context) PNGRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("PNGRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}

// An optional call to warm up a so that subsequent calls to render with the same arguments run more efficiently.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/prepareRender(_:from:to:at:)
func (c_ Context) PrepareRenderFromRectToDestinationAtPointError(image unsafe.Pointer, fromRect coregraphics.CGRect, destination unsafe.Pointer, atPoint coregraphics.CGPoint, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("prepareRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, error_)
	return rv
}

// Runs the garbage collector to reclaim any resources that the context no longer requires.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/reclaimResources()
func (c_ Context) ReclaimResources() {
	objc.Send[objc.ID](c_.ID, objc.Sel("reclaimResources"))
}

// Renders an image into a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:)
func (c_ Context) RenderToCVPixelBuffer(image unsafe.Pointer, buffer unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toCVPixelBuffer:"), image, buffer)
}

// Renders a region of an image into a pixel buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-2k8l2
func (c_ Context) RenderToCVPixelBufferBoundsColorSpace(image unsafe.Pointer, buffer unsafe.Pointer, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toCVPixelBuffer:bounds:colorSpace:"), image, buffer, bounds, colorSpace)
}

// Renders a region of an image into an IOSurface object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:bounds:colorSpace:)-54b9l
func (c_ Context) RenderToIOSurfaceBoundsColorSpace(image unsafe.Pointer, surface unsafe.Pointer, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toIOSurface:bounds:colorSpace:"), image, surface, bounds, colorSpace)
}

// Renders a region of an image to a Metal texture.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:to:commandBuffer:bounds:colorSpace:)
func (c_ Context) RenderToMTLTextureCommandBufferBoundsColorSpace(image unsafe.Pointer, texture objc.ID, commandBuffer objc.ID, bounds coregraphics.CGRect, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toMTLTexture:commandBuffer:bounds:colorSpace:"), image, texture, commandBuffer, bounds, colorSpace)
}

// Renders to the given bitmap.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/render(_:toBitmap:rowBytes:bounds:format:colorSpace:)
func (c_ Context) RenderToBitmapRowBytesBoundsFormatColorSpace(image unsafe.Pointer, data unsafe.Pointer, rowBytes unsafe.Pointer, bounds coregraphics.CGRect, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](c_.ID, objc.Sel("render:toBitmap:rowBytes:bounds:format:colorSpace:"), image, data, rowBytes, bounds, format, colorSpace)
}

// Fills the entire destination with black or clear depending on its .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toClear:)
func (c_ Context) StartTaskToClearError(destination unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("startTaskToClear:error:"), destination, error_)
	return rv
}

// Renders a portion of an image to a point in the destination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:from:to:at:)
func (c_ Context) StartTaskToRenderFromRectToDestinationAtPointError(image unsafe.Pointer, fromRect coregraphics.CGRect, destination unsafe.Pointer, atPoint coregraphics.CGPoint, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("startTaskToRender:fromRect:toDestination:atPoint:error:"), image, fromRect, destination, atPoint, error_)
	return rv
}

// Renders an image to a destination so that point (0, 0) of the image maps to point (0, 0) of the destination.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/startTask(toRender:to:)
func (c_ Context) StartTaskToRenderToDestinationError(image unsafe.Pointer, destination unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("startTaskToRender:toDestination:error:"), image, destination, error_)
	return rv
}

// Renders the image and exports the resulting image data in TIFF format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/tiffRepresentation(of:format:colorSpace:options:)
func (c_ Context) TIFFRepresentationOfImageFormatColorSpaceOptions(image unsafe.Pointer, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("TIFFRepresentationOfImage:format:colorSpace:options:"), image, format, colorSpace, options)
	return rv
}

// Renders the image and exports the resulting image data as a file in HEIF10 format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIF10Representation(of:to:colorSpace:options:)
func (c_ Context) WriteHEIF10RepresentationOfImageToURLColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeHEIF10RepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data as a file in HEIF format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeHEIFRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WriteHEIFRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeHEIFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data as a file in JPEG format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeJPEGRepresentation(of:to:colorSpace:options:)
func (c_ Context) WriteJPEGRepresentationOfImageToURLColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeJPEGRepresentationOfImage:toURL:colorSpace:options:error:"), image, url, colorSpace, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data as a file in open EXR format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeOpenEXRRepresentation(of:to:options:)
func (c_ Context) WriteOpenEXRRepresentationOfImageToURLOptionsError(image unsafe.Pointer, url foundation.URL, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeOpenEXRRepresentationOfImage:toURL:options:error:"), image, url, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data as a file in PNG format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writePNGRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WritePNGRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writePNGRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}

// Renders the image and exports the resulting image data as a file in TIFF format.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/writeTIFFRepresentation(of:to:format:colorSpace:options:)
func (c_ Context) WriteTIFFRepresentationOfImageToURLFormatColorSpaceOptionsError(image unsafe.Pointer, url foundation.URL, format unsafe.Pointer, colorSpace coregraphics.CGColorSpaceRef, options unsafe.Pointer, errorPtr unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("writeTIFFRepresentationOfImage:toURL:format:colorSpace:options:error:"), image, url, format, colorSpace, options, errorPtr)
	return rv
}

// The render destination’s representation of alpha (transparency) values.
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/alphamode
func (c_ Context) AlphaMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("alphaMode"))
	return rv
}


// SetAlphaMode sets the value of the alphaMode property.
// The render destination’s representation of alpha (transparency) values.

//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/alphamode
func (c_ Context) SetAlphaMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlphaMode:"), value)
}

// The working color space of the Core Image context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/workingColorSpace
func (c_ Context) WorkingColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](c_.ID, objc.Sel("workingColorSpace"))
	return rv
}

// The working pixel format of the Core Image context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/workingFormat
func (c_ Context) WorkingFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("workingFormat"))
	return rv
}


