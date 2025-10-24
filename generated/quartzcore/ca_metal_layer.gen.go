// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CAMetalLayer */


/* debug [class_header]: Header for CAMetalLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetalLayer */
// An interface definition for the [MetalLayer] class.
type IMetalLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for MetalLayer */
	// properties:
	AllowsNextDrawableTimeout() bool
	SetAllowsNextDrawableTimeout(value bool)
	Colorspace() ColorSpaceRef /* not a class type */
	SetColorspace(value ColorSpaceRef /* not a class type */)
	DeveloperHUDProperties() objc.IObject /* cross-framework: NSDictionary */
	SetDeveloperHUDProperties(value objc.IObject /* cross-framework: NSDictionary */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetalLayer */
	// methods:
	NextDrawable() unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetalLayer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetalLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetalLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetalLayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetalLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetalLayer */

// Waits until a Metal drawable is available, and then returns it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/nextDrawable()
func (m_ MetalLayer) NextDrawable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nextDrawable"))
	return rv
}/* debug [instance_methods/method]: NextDrawable */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetalLayer */

// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}/* debug [instance_properties/getter]: allowsNextDrawableTimeout */


// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m_ MetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsNextDrawableTimeout:"), value)
}/* debug [instance_properties/setter]: allowsNextDrawableTimeout */


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) Colorspace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](m_.ID, objc.Sel("colorspace"))
	return rv
}/* debug [instance_properties/getter]: colorspace */


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m_ MetalLayer) SetColorspace(value ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColorspace:"), value)
}/* debug [instance_properties/setter]: colorspace */


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) DeveloperHUDProperties() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](m_.ID, objc.Sel("developerHUDProperties"))
	return rv
}/* debug [instance_properties/getter]: developerHUDProperties */


// The properties of the Metal performance heads-up display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m_ MetalLayer) SetDeveloperHUDProperties(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeveloperHUDProperties:"), value)
}/* debug [instance_properties/setter]: developerHUDProperties */


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The Metal device responsible for the layer’s drawable resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m_ MetalLayer) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("displaySyncEnabled"))
	return rv
}/* debug [instance_properties/getter]: displaySyncEnabled */


// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m_ MetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplaySyncEnabled:"), value)
}/* debug [instance_properties/setter]: displaySyncEnabled */


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) DrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("drawableSize"))
	return rv
}/* debug [instance_properties/getter]: drawableSize */


// The size, in pixels, of textures for rendering layer content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m_ MetalLayer) SetDrawableSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableSize:"), value)
}/* debug [instance_properties/setter]: drawableSize */


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) EDRMetadata() IEDRMetadata {
	rv := objc.Send[EDRMetadata](m_.ID, objc.Sel("EDRMetadata"))
	return rv
}/* debug [instance_properties/getter]: EDRMetadata */


// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m_ MetalLayer) SetEDRMetadata(value IEDRMetadata) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEDRMetadata:"), value)
}/* debug [instance_properties/setter]: EDRMetadata */


// A Boolean value that determines whether the layer’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) FramebufferOnly() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("framebufferOnly"))
	return rv
}/* debug [instance_properties/getter]: framebufferOnly */


// A Boolean value that determines whether the layer’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m_ MetalLayer) SetFramebufferOnly(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFramebufferOnly:"), value)
}/* debug [instance_properties/setter]: framebufferOnly */


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) MaximumDrawableCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maximumDrawableCount"))
	return rv
}/* debug [instance_properties/getter]: maximumDrawableCount */


// The number of Metal drawables in the resource pool managed by Core Animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m_ MetalLayer) SetMaximumDrawableCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumDrawableCount:"), value)
}/* debug [instance_properties/setter]: maximumDrawableCount */


// The pixel format of the layer’s textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) PixelFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](m_.ID, objc.Sel("pixelFormat"))
	return rv
}/* debug [instance_properties/getter]: pixelFormat */


// The pixel format of the layer’s textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m_ MetalLayer) SetPixelFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPixelFormat:"), value)
}/* debug [instance_properties/setter]: pixelFormat */


// The device object that the system recommends using for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/preferredDevice
func (m_ MetalLayer) PreferredDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("preferredDevice"))
	return rv
}/* debug [instance_properties/getter]: preferredDevice */


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}/* debug [instance_properties/getter]: presentsWithTransaction */


// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m_ MetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}/* debug [instance_properties/setter]: presentsWithTransaction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/residencySet
func (m_ MetalLayer) ResidencySet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("residencySet"))
	return rv
}/* debug [instance_properties/getter]: residencySet */


// Enables extended dynamic range values onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}/* debug [instance_properties/getter]: wantsExtendedDynamicRangeContent */


// Enables extended dynamic range values onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/wantsExtendedDynamicRangeContent
func (m_ MetalLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}/* debug [instance_properties/setter]: wantsExtendedDynamicRangeContent */


// A positive integer that identifies the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) DrawableID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("drawableID"))
	return rv
}/* debug [instance_properties/getter]: drawableID */


// A positive integer that identifies the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m_ MetalLayer) SetDrawableID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDrawableID:"), value)
}/* debug [instance_properties/setter]: drawableID */


// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) PresentedTime() float64 {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("presentedTime"))
	return rv
}/* debug [instance_properties/getter]: presentedTime */


// The host time, in seconds, when the drawable was displayed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m_ MetalLayer) SetPresentedTime(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresentedTime:"), value)
}/* debug [instance_properties/setter]: presentedTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAMetalLayer */



