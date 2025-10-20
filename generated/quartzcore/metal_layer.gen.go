// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [MetalLayer] class.
var (
	metalLayerClass     _MetalLayerClass
	metalLayerClassOnce sync.Once
)

func getMetalLayerClass() _MetalLayerClass {
	metalLayerClassOnce.Do(func() {
		metalLayerClass = _MetalLayerClass{objc.GetClass("CAMetalLayer")}
	})
	return metalLayerClass
}

type _MetalLayerClass struct {
	class objc.Class
}

// An interface definition for the [MetalLayer] class.
type IMetalLayer interface {
	ILayer
}

// A Core Animation layer that Metal can render into, typically displayed onscreen.
//
// Use a when you want to use Metal to render a layer’s contents; for example, to render into a view. Consider using instead, because this class automatically wraps a object and provides a higher-level abstraction. If you’re using UIKit, to create a view that uses a , create a subclass of and override its class method to return a : If you’re using AppKit, configure an object to use a backing layer and assign a object to the view: Adjust the layer’s properties to configure its underlying pixel format and other display behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
type MetalLayer struct {
	Layer
}

// MetalLayerFrom constructs a [MetalLayer] from an unsafe.Pointer.
//
// A Core Animation layer that Metal can render into, typically displayed onscreen.
func MetalLayerFrom(ptr unsafe.Pointer) MetalLayer {
	return MetalLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MetalLayerClass) Alloc() MetalLayer {
	rv := objc.Send[MetalLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetalLayerClass) New() MetalLayer {
	rv := objc.Send[MetalLayer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetalLayer) Init() MetalLayer {
	rv := objc.Send[MetalLayer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetalLayer) Autorelease() MetalLayer {
	rv := objc.Send[MetalLayer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalLayer creates a new MetalLayer instance.
func NewMetalLayer() MetalLayer {
	return getMetalLayerClass().New()
}


// The color space of the rendered content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) Colorspace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](m_.ID, objc.Sel("colorspace"))
	return rv
}


// SetColorspace sets the value of the colorspace property.
// The color space of the rendered content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) SetColorspace(value coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorspace:"), value)
}
// A Boolean value that determines whether the layer’s textures are used only for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) FramebufferOnly() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("framebufferOnly"))
	return rv
}


// SetFramebufferOnly sets the value of the framebufferOnly property.
// A Boolean value that determines whether the layer’s textures are used only for rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) SetFramebufferOnly(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFramebufferOnly:"), value)
}
// The pixel format of the layer’s textures.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) PixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("pixelFormat"))
	return rv
}


// SetPixelFormat sets the value of the pixelFormat property.
// The pixel format of the layer’s textures.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) SetPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}
// Enables extended dynamic range values onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}


// SetWantsExtendedDynamicRangeContent sets the value of the wantsExtendedDynamicRangeContent property.
// Enables extended dynamic range values onscreen.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}


