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
	NextDrawable() objc.ID
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


// Waits until a Metal drawable is available, and then returns it.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/nextDrawable()
func (m_ MetalLayer) NextDrawable() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("nextDrawable"))
	return rv
}

// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}


// SetAllowsNextDrawableTimeout sets the value of the allowsNextDrawableTimeout property.
// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsNextDrawableTimeout:"), value)
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

// The properties of the Metal performance heads-up display.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) DeveloperHUDProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("developerHUDProperties"))
	return rv
}


// SetDeveloperHUDProperties sets the value of the developerHUDProperties property.
// The properties of the Metal performance heads-up display.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) SetDeveloperHUDProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeveloperHUDProperties:"), value)
}

// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) Device() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The Metal device responsible for the layer’s drawable resources.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) SetDevice(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}

// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("displaySyncEnabled"))
	return rv
}


// SetDisplaySyncEnabled sets the value of the displaySyncEnabled property.
// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplaySyncEnabled:"), value)
}

// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) DrawableSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("drawableSize"))
	return rv
}


// SetDrawableSize sets the value of the drawableSize property.
// The size, in pixels, of textures for rendering layer content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) SetDrawableSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableSize:"), value)
}

// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) EDRMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("EDRMetadata"))
	return rv
}


// SetEDRMetadata sets the value of the EDRMetadata property.
// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) SetEDRMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEDRMetadata:"), value)
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

// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) MaximumDrawableCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maximumDrawableCount"))
	return rv
}


// SetMaximumDrawableCount sets the value of the maximumDrawableCount property.
// The number of Metal drawables in the resource pool managed by Core Animation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) SetMaximumDrawableCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumDrawableCount:"), value)
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

// The device object that the system recommends using for this layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/preferredDevice
func (m_ MetalLayer) PreferredDevice() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("preferredDevice"))
	return rv
}

// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}


// SetPresentsWithTransaction sets the value of the presentsWithTransaction property.
// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/residencySet
func (m_ MetalLayer) ResidencySet() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("residencySet"))
	return rv
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

// A positive integer that identifies the drawable.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) DrawableID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("drawableID"))
	return rv
}


// SetDrawableID sets the value of the drawableID property.
// A positive integer that identifies the drawable.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) SetDrawableID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableID:"), value)
}

// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) PresentedTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presentedTime"))
	return rv
}


// SetPresentedTime sets the value of the presentedTime property.
// The host time, in seconds, when the drawable was displayed onscreen.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) SetPresentedTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentedTime:"), value)
}



