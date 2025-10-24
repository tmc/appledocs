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
	Colorspace() ColorSpaceRef /* not a class type */
	SetColorspace(value ColorSpaceRef /* not a class type */)
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	PixelFormat() PixelFormat /* not a class type */
	SetPixelFormat(value PixelFormat /* not a class type */)
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)
	DrawableID() int
	SetDrawableID(value int)
	PresentedTime() float64
	SetPresentedTime(value float64)
	AllowsNextDrawableTimeout() bool
	SetAllowsNextDrawableTimeout(value bool)
	DeveloperHUDProperties() unsafe.Pointer
	SetDeveloperHUDProperties(value unsafe.Pointer)
	Device() Device /* not a class type */
	SetDevice(value Device /* not a class type */)
	DisplaySyncEnabled() bool
	SetDisplaySyncEnabled(value bool)
	DrawableSize() objc.IObject /* cross-framework: Size */
	SetDrawableSize(value objc.IObject /* cross-framework: Size */)
	EdrMetadata() IEDRMetadata
	SetEdrMetadata(value IEDRMetadata)
	MaximumDrawableCount() int
	SetMaximumDrawableCount(value int)
	PreferredDevice() Device /* not a class type */
	SetPreferredDevice(value Device /* not a class type */)
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)
	ResidencySet() ResidencySet /* not a class type */
	SetResidencySet(value ResidencySet /* not a class type */)
	// methods:
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
	rv := objc.Send[TimeInterval](m_.ID, objc.Sel("presentedTime"))
	return rv
}


// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) SetPresentedTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentedTime:"), value)
}


// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/allowsnextdrawabletimeout
func (m_ MetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}


// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/allowsnextdrawabletimeout
func (m_ MetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsNextDrawableTimeout:"), value)
}


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/developerhudproperties
func (m_ MetalLayer) DeveloperHUDProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("developerHUDProperties"))
	return rv
}


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/developerhudproperties
func (m_ MetalLayer) SetDeveloperHUDProperties(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeveloperHUDProperties:"), value)
}


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/device
func (m_ MetalLayer) Device() Device /* not a class type */ {
	rv := objc.Send[Device](m_.ID, objc.Sel("device"))
	return rv
}


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/device
func (m_ MetalLayer) SetDevice(value Device /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/displaysyncenabled
func (m_ MetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("displaySyncEnabled"))
	return rv
}


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/displaysyncenabled
func (m_ MetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplaySyncEnabled:"), value)
}


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/drawablesize
func (m_ MetalLayer) DrawableSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](m_.ID, objc.Sel("drawableSize"))
	return rv
}


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/drawablesize
func (m_ MetalLayer) SetDrawableSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableSize:"), value)
}


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/edrmetadata
func (m_ MetalLayer) EdrMetadata() IEDRMetadata {
	rv := objc.Send[EDRMetadata](m_.ID, objc.Sel("edrMetadata"))
	return rv
}


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/edrmetadata
func (m_ MetalLayer) SetEdrMetadata(value IEDRMetadata) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEdrMetadata:"), value)
}


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/maximumdrawablecount
func (m_ MetalLayer) MaximumDrawableCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maximumDrawableCount"))
	return rv
}


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/maximumdrawablecount
func (m_ MetalLayer) SetMaximumDrawableCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumDrawableCount:"), value)
}


// The device object that the system recommends using for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/preferreddevice
func (m_ MetalLayer) PreferredDevice() Device /* not a class type */ {
	rv := objc.Send[Device](m_.ID, objc.Sel("preferredDevice"))
	return rv
}


// The device object that the system recommends using for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/preferreddevice
func (m_ MetalLayer) SetPreferredDevice(value Device /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreferredDevice:"), value)
}


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/presentswithtransaction
func (m_ MetalLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/presentswithtransaction
func (m_ MetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/residencyset
func (m_ MetalLayer) ResidencySet() ResidencySet /* not a class type */ {
	rv := objc.Send[ResidencySet](m_.ID, objc.Sel("residencySet"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/residencyset
func (m_ MetalLayer) SetResidencySet(value ResidencySet /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setResidencySet:"), value)
}



