// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [MetalLayer] class.
var (
	MetalLayerClass     _MetalLayerClass
	MetalLayerClassOnce sync.Once
)

func getMetalLayerClass() _MetalLayerClass {
	MetalLayerClassOnce.Do(func() {
		MetalLayerClass = _MetalLayerClass{objc.GetClass("CAMetalLayer")}
	})
	return MetalLayerClass
}

type _MetalLayerClass struct {
	class objc.Class
}





// An interface definition for the [MetalLayer] class.
type IMetalLayer interface {
	ILayer
	

	// properties:
	AllowsNextDrawableTimeout() bool
	SetAllowsNextDrawableTimeout(value bool)
	Colorspace() ColorSpaceRef /* not a class type */
	SetColorspace(value ColorSpaceRef /* not a class type */)
	DeveloperHUDProperties() foundation.foundation.INSDictionary
	SetDeveloperHUDProperties(value foundation.foundation.INSDictionary)
	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	DisplaySyncEnabled() bool
	SetDisplaySyncEnabled(value bool)
	DrawableSize() corefoundation.CGSize
	SetDrawableSize(value corefoundation.CGSize)
	EDRMetadata() IEDRMetadata
	SetEDRMetadata(value IEDRMetadata)
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	MaximumDrawableCount() uint
	SetMaximumDrawableCount(value uint)
	PixelFormat() PixelFormat /* not a class type */
	SetPixelFormat(value PixelFormat /* not a class type */)
	PreferredDevice() unsafe.Pointer
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)
	ResidencySet() unsafe.Pointer
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)
	DrawableID() int
	SetDrawableID(value int)
	PresentedTime() float64
	SetPresentedTime(value float64)


	

	// methods:
	NextDrawable() unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (mc _MetalLayerClass) Alloc() MetalLayer {
	rv := objc.Send[MetalLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A Core Animation layer that Metal can render into, typically displayed onscreen.
//
// Use a when you want to use Metal to render a layer’s contents; for example, to render into a view. Consider using instead, because this class automatically wraps a object and provides a higher-level abstraction. If you’re using UIKit, to create a view that uses a , create a subclass of and override its class method to return a : If you’re using AppKit, configure an object to use a backing layer and assign a object to the view: Adjust the layer’s properties to configure its underlying pixel format and other display behaviors.


// A Core Animation layer that Metal can render into, typically displayed onscreen.
//
// [Full Topic]
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




















// Waits until a Metal drawable is available, and then returns it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/nextDrawable()
func (m_ MetalLayer) NextDrawable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nextDrawable"))
	return rv
}







// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}


// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsNextDrawableTimeout:"), value)
}


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) Colorspace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](m_.ID, objc.Sel("colorspace"))
	return rv
}


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) SetColorspace(value ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorspace:"), value)
}


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) DeveloperHUDProperties() foundation.foundation.INSDictionary {
	rv := objc.Send[foundation.NSDictionary](m_.ID, objc.Sel("developerHUDProperties"))
	return rv
}


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) SetDeveloperHUDProperties(value foundation.foundation.INSDictionary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeveloperHUDProperties:"), value)
}


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("device"))
	return rv
}


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("displaySyncEnabled"))
	return rv
}


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplaySyncEnabled:"), value)
}


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) DrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("drawableSize"))
	return rv
}


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) SetDrawableSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableSize:"), value)
}


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) EDRMetadata() IEDRMetadata {
	rv := objc.Send[EDRMetadata](m_.ID, objc.Sel("EDRMetadata"))
	return rv
}


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) SetEDRMetadata(value IEDRMetadata) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEDRMetadata:"), value)
}


// A Boolean value that determines whether the layer’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) FramebufferOnly() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("framebufferOnly"))
	return rv
}


// A Boolean value that determines whether the layer’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) SetFramebufferOnly(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFramebufferOnly:"), value)
}


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) MaximumDrawableCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maximumDrawableCount"))
	return rv
}


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) SetMaximumDrawableCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumDrawableCount:"), value)
}


// The pixel format of the layer’s textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) PixelFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("pixelFormat"))
	return rv
}


// The pixel format of the layer’s textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) SetPixelFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}


// The device object that the system recommends using for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/preferredDevice
func (m_ MetalLayer) PreferredDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredDevice"))
	return rv
}


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/residencySet
func (m_ MetalLayer) ResidencySet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("residencySet"))
	return rv
}


// Enables extended dynamic range values onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}


// Enables extended dynamic range values onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}


// A positive integer that identifies the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) DrawableID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("drawableID"))
	return rv
}


// A positive integer that identifies the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) SetDrawableID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableID:"), value)
}


// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) PresentedTime() float64 {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("presentedTime"))
	return rv
}


// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) SetPresentedTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentedTime:"), value)
}








