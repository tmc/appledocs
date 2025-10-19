// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RenderDestination] class.
var renderDestinationClass = _RenderDestinationClass{objc.GetClass("CIRenderDestination")}

type _RenderDestinationClass struct {
	class objc.Class
}

// A specification for configuring all attributes of a render task’s destination and issuing asynchronous render tasks. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (rc _RenderDestinationClass) Alloc() RenderDestination {
	rv := objc.Send[RenderDestination](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return renderDestinationClass.New()
}
// Creates a render destination based on a client-managed buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(bitmapData:width:height:bytesPerRow:format:)
func NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format unsafe.Pointer) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithBitmapData:width:height:bytesPerRow:format:"), data, width, height, bytesPerRow, format)
	rv.Autorelease()
	return rv
}
// Creates a render destination based on an OpenGL texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(glTexture:target:width:height:)
func NewRenderDestinationWithGLTextureTargetWidthHeight(texture unsafe.Pointer, target unsafe.Pointer, width uint, height uint) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithGLTexture:target:width:height:"), texture, target, width, height)
	rv.Autorelease()
	return rv
}
// Creates a render destination based on an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(ioSurface:)
func NewRenderDestinationWithIOSurface(surface unsafe.Pointer) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithIOSurface:"), surface)
	rv.Autorelease()
	return rv
}
// Creates a render destination based on a Metal texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(mtlTexture:commandBuffer:)
func NewRenderDestinationWithMTLTextureCommandBuffer(texture unsafe.Pointer, commandBuffer unsafe.Pointer) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithMTLTexture:commandBuffer:"), texture, commandBuffer)
	rv.Autorelease()
	return rv
}
// Creates a render destination based on a Core Video pixel buffer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(pixelBuffer:)
func NewRenderDestinationWithPixelBuffer(pixelBuffer unsafe.Pointer) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithPixelBuffer:"), pixelBuffer)
	rv.Autorelease()
	return rv
}
// Creates a render destination based on a Metal texture with specified pixel format. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRenderDestination/init(width:height:pixelFormat:commandBuffer:mtlTextureProvider:)
func NewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat unsafe.Pointer, commandBuffer unsafe.Pointer, block unsafe.Pointer) RenderDestination {
	instance := renderDestinationClass.Alloc()
	rv := objc.Send[RenderDestination](instance.ID, objc.Sel("initWithWidth:height:pixelFormat:commandBuffer:mtlTextureProvider:"), width, height, pixelFormat, commandBuffer, block)
	rv.Autorelease()
	return rv
}



