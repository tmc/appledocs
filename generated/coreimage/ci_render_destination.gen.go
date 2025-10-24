// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RenderDestination] class.
var (
	RenderDestinationClass     _RenderDestinationClass
	RenderDestinationClassOnce sync.Once
)

func getRenderDestinationClass() _RenderDestinationClass {
	RenderDestinationClassOnce.Do(func() {
		RenderDestinationClass = _RenderDestinationClass{objc.GetClass("CIRenderDestination")}
	})
	return RenderDestinationClass
}

type _RenderDestinationClass struct {
	class objc.Class
}





// An interface definition for the [RenderDestination] class.
type IRenderDestination interface {
	objectivec.IObject
	

	// properties:
	AlphaMode() RenderDestinationAlphaMode
	SetAlphaMode(value RenderDestinationAlphaMode)
	BlendKernel() ICIBlendKernel
	SetBlendKernel(value ICIBlendKernel)
	BlendsInDestinationColorSpace() bool
	SetBlendsInDestinationColorSpace(value bool)
	CaptureTraceURL() objc.IObject /* cross-framework: NSURL */
	SetCaptureTraceURL(value objc.IObject /* cross-framework: NSURL */)
	ColorSpace() ColorSpaceRef /* not a class type */
	SetColorSpace(value ColorSpaceRef /* not a class type */)
	Height() uint
	Clamped() bool
	SetClamped(value bool)
	Dithered() bool
	SetDithered(value bool)
	Flipped() bool
	SetFlipped(value bool)
	Width() uint
	IsClamped() bool
	SetIsClamped(value bool)
	IsDithered() bool
	SetIsDithered(value bool)
	IsFlipped() bool
	SetIsFlipped(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RenderDestinationClass) Alloc() RenderDestination {
	rv := objc.Send[RenderDestination](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderDestinationClass) New() RenderDestination {
	rv := objc.Send[RenderDestination](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderDestination) Init() RenderDestination {
	rv := objc.Send[RenderDestination](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderDestination) Autorelease() RenderDestination {
	rv := objc.Send[RenderDestination](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderDestination creates a new RenderDestination instance.
func NewRenderDestination() RenderDestination {
	return getRenderDestinationClass().New()
}





// A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks.
//
// The class provides an API for specifying a render task destination’s properties, such as buffer format, alpha mode, clamping behavior, blending, and color space, properties formerly tied to . You can create a object for each surface or buffer to which you must render. You can also render multiple times to a single destination with different settings such as colorspace and blend mode by mutating a single object between renders. Renders issued to a return to the caller as soon as the CPU has issued the task, rather than after the GPU has performed the task, so you can start render tasks on subsequent frames without waiting for previous renders to finish. If the render fails, a will return immediately.


// A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination
type RenderDestination struct {
	objectivec.Object
}

// RenderDestinationFrom constructs a [RenderDestination] from an unsafe.Pointer.
//
// A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks.
func RenderDestinationFrom(ptr unsafe.Pointer) RenderDestination {
	return RenderDestination{objectivec.Object{objc.ID(ptr)}}
}






// Creates a render destination based on a client-managed buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(bitmapData:width:height:bytesPerRow:format:)
func NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat(data objectivec.IObject, width uint, height uint, bytesPerRow uint, format Format /* typedef */) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithBitmapData:width:height:bytesPerRow:format:"), data, width, height, bytesPerRow, format)
	rv.Autorelease()
	return rv
}


// Creates a render destination based on an OpenGL texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(glTexture:target:width:height:)
func NewRenderDestinationWithGLTextureTargetWidthHeight(texture objectivec.IObject, target objectivec.IObject, width uint, height uint) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithGLTexture:target:width:height:"), texture, target, width, height)
	rv.Autorelease()
	return rv
}


// Creates a render destination based on an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(ioSurface:)
func NewRenderDestinationWithIOSurface(surface iosurface.Surface) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	rv.Autorelease()
	return rv
}


// Creates a render destination based on a Metal texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(mtlTexture:commandBuffer:)
func NewRenderDestinationWithMTLTextureCommandBuffer(texture unsafe.Pointer, commandBuffer unsafe.Pointer) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithMTLTexture:commandBuffer:"), texture, commandBuffer)
	rv.Autorelease()
	return rv
}


// Creates a render destination based on a Core Video pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(pixelBuffer:)
func NewRenderDestinationWithPixelBuffer(pixelBuffer PixelBufferRef /* not a class type */) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithPixelBuffer:"), pixelBuffer)
	rv.Autorelease()
	return rv
}


// Creates a render destination based on a Metal texture with specified pixel format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(width:height:pixelFormat:commandBuffer:mtlTextureProvider:)
func NewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat PixelFormat /* not a class type */, commandBuffer unsafe.Pointer, block unsafe.Pointer) RenderDestination {
	instance := getRenderDestinationClass().Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithWidth:height:pixelFormat:commandBuffer:mtlTextureProvider:"), width, height, pixelFormat, commandBuffer, block)
	rv.Autorelease()
	return rv
}






















// The render destination’s representation of alpha (transparency) values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/alphaMode
func (r_ RenderDestination) AlphaMode() RenderDestinationAlphaMode {
	rv := objc.Send[RenderDestinationAlphaMode](r_.ID, objc.Sel("alphaMode"))
	return rv
}


// The render destination’s representation of alpha (transparency) values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/alphaMode
func (r_ RenderDestination) SetAlphaMode(value RenderDestinationAlphaMode) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAlphaMode:"), value)
}


// The destination’s blend kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendKernel
func (r_ RenderDestination) BlendKernel() ICIBlendKernel {
	rv := objc.Send[BlendKernel](r_.ID, objc.Sel("blendKernel"))
	return rv
}


// The destination’s blend kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendKernel
func (r_ RenderDestination) SetBlendKernel(value ICIBlendKernel) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBlendKernel:"), value)
}


// Indicator of whether to blend in the destination’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendsInDestinationColorSpace
func (r_ RenderDestination) BlendsInDestinationColorSpace() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("blendsInDestinationColorSpace"))
	return rv
}


// Indicator of whether to blend in the destination’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/blendsInDestinationColorSpace
func (r_ RenderDestination) SetBlendsInDestinationColorSpace(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBlendsInDestinationColorSpace:"), value)
}


// Tell the next render using this destination to capture a Metal trace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/captureTraceURL
func (r_ RenderDestination) CaptureTraceURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](r_.ID, objc.Sel("captureTraceURL"))
	return rv
}


// Tell the next render using this destination to capture a Metal trace.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/captureTraceURL
func (r_ RenderDestination) SetCaptureTraceURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCaptureTraceURL:"), value)
}


// The destination’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/colorSpace
func (r_ RenderDestination) ColorSpace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](r_.ID, objc.Sel("colorSpace"))
	return rv
}


// The destination’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/colorSpace
func (r_ RenderDestination) SetColorSpace(value ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setColorSpace:"), value)
}


// The render destination’s buffer height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/height
func (r_ RenderDestination) Height() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("height"))
	return rv
}


// Indicator of whether or not the destination clamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isClamped
func (r_ RenderDestination) Clamped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("clamped"))
	return rv
}


// Indicator of whether or not the destination clamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isClamped
func (r_ RenderDestination) SetClamped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClamped:"), value)
}


// Indicator of whether or not the destination dithers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isDithered
func (r_ RenderDestination) Dithered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("dithered"))
	return rv
}


// Indicator of whether or not the destination dithers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isDithered
func (r_ RenderDestination) SetDithered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDithered:"), value)
}


// Indicator of whether the destination is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isFlipped
func (r_ RenderDestination) Flipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("flipped"))
	return rv
}


// Indicator of whether the destination is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/isFlipped
func (r_ RenderDestination) SetFlipped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFlipped:"), value)
}


// The render destination’s row width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/width
func (r_ RenderDestination) Width() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("width"))
	return rv
}


// Indicator of whether or not the destination clamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isclamped
func (r_ RenderDestination) IsClamped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isClamped"))
	return rv
}


// Indicator of whether or not the destination clamps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isclamped
func (r_ RenderDestination) SetIsClamped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsClamped:"), value)
}


// Indicator of whether or not the destination dithers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isdithered
func (r_ RenderDestination) IsDithered() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDithered"))
	return rv
}


// Indicator of whether or not the destination dithers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isdithered
func (r_ RenderDestination) SetIsDithered(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDithered:"), value)
}


// Indicator of whether the destination is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isflipped
func (r_ RenderDestination) IsFlipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isFlipped"))
	return rv
}


// Indicator of whether the destination is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/isflipped
func (r_ RenderDestination) SetIsFlipped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsFlipped:"), value)
}







