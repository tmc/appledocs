// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFXSpatialScalerDescriptor */


/* debug [class_header]: Header for MTLFXSpatialScalerDescriptor */
// The class instance for the [FXSpatialScalerDescriptor] class.
var (
	FXSpatialScalerDescriptorClass     _FXSpatialScalerDescriptorClass
	FXSpatialScalerDescriptorClassOnce sync.Once
)

func getFXSpatialScalerDescriptorClass() _FXSpatialScalerDescriptorClass {
	FXSpatialScalerDescriptorClassOnce.Do(func() {
		FXSpatialScalerDescriptorClass = _FXSpatialScalerDescriptorClass{objc.GetClass("MTLFXSpatialScalerDescriptor")}
	})
	return FXSpatialScalerDescriptorClass
}

type _FXSpatialScalerDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FXSpatialScalerDescriptor */
// An interface definition for the [FXSpatialScalerDescriptor] class.
type IFXSpatialScalerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FXSpatialScalerDescriptor */
	// properties:
	ColorProcessingMode() FXSpatialScalerColorProcessingMode
	SetColorProcessingMode(value FXSpatialScalerColorProcessingMode)
	ColorTextureFormat() PixelFormat /* not a class type */
	SetColorTextureFormat(value PixelFormat /* not a class type */)
	InputHeight() uint
	SetInputHeight(value uint)
	InputWidth() uint
	SetInputWidth(value uint)
	OutputHeight() uint
	SetOutputHeight(value uint)
	OutputTextureFormat() PixelFormat /* not a class type */
	SetOutputTextureFormat(value PixelFormat /* not a class type */)
	OutputWidth() uint
	SetOutputWidth(value uint)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FXSpatialScalerDescriptor */
	// methods:
	NewSpatialScalerWithDevice(device unsafe.Pointer) unsafe.Pointer
	NewSpatialScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FXSpatialScalerDescriptor */
// Alloc allocates a new instance without initialization.
func (fc _FXSpatialScalerDescriptorClass) Alloc() FXSpatialScalerDescriptor {
	rv := objc.Send[FXSpatialScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FXSpatialScalerDescriptorClass) New() FXSpatialScalerDescriptor {
	rv := objc.Send[FXSpatialScalerDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FXSpatialScalerDescriptor) Init() FXSpatialScalerDescriptor {
	rv := objc.Send[FXSpatialScalerDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FXSpatialScalerDescriptor) Autorelease() FXSpatialScalerDescriptor {
	rv := objc.Send[FXSpatialScalerDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFXSpatialScalerDescriptor creates a new FXSpatialScalerDescriptor instance.
func NewFXSpatialScalerDescriptor() FXSpatialScalerDescriptor {
	return getFXSpatialScalerDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FXSpatialScalerDescriptor */
// A set of properties that configure a spatial scaling effect, and a factory method that creates the effect.


// A set of properties that configure a spatial scaling effect, and a factory method that creates the effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor
type FXSpatialScalerDescriptor struct {
	objectivec.Object
}

// FXSpatialScalerDescriptorFrom constructs a [FXSpatialScalerDescriptor] from an unsafe.Pointer.
//
// A set of properties that configure a spatial scaling effect, and a factory method that creates the effect.
func FXSpatialScalerDescriptorFrom(ptr unsafe.Pointer) FXSpatialScalerDescriptor {
	return FXSpatialScalerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FXSpatialScalerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FXSpatialScalerDescriptor */

// Returns a Boolean value that indicates whether the spatial scaler works with a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/supportsDevice(_:)
func (fc _FXSpatialScalerDescriptorClass) SupportsDevice(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsDevice) */


// Queries whether a Metal device supports spatial scaling compatible with Metal 4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXSpatialScalerDescriptorClass) SupportsMetal4FX(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsMetal4FX) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FXSpatialScalerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FXSpatialScalerDescriptor */

// Creates a spatial scaler instance from this descriptor’s current property values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/makeSpatialScaler(device:)
func (f_ FXSpatialScalerDescriptor) NewSpatialScalerWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newSpatialScalerWithDevice:"), device)
	return rv
}/* debug [instance_methods/method]: NewSpatialScalerWithDevice */


// Creates a spatial scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/makeSpatialScaler(device:compiler:)
func (f_ FXSpatialScalerDescriptor) NewSpatialScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newSpatialScalerWithDevice:compiler:"), device, compiler)
	return rv
}/* debug [instance_methods/method]: NewSpatialScalerWithDeviceCompiler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FXSpatialScalerDescriptor */

// The color space of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorProcessingMode
func (f_ FXSpatialScalerDescriptor) ColorProcessingMode() FXSpatialScalerColorProcessingMode {
	rv := objc.Send[FXSpatialScalerColorProcessingMode](f_.ID, objc.Sel("colorProcessingMode"))
	return rv
}/* debug [instance_properties/getter]: colorProcessingMode */


// The color space of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorProcessingMode
func (f_ FXSpatialScalerDescriptor) SetColorProcessingMode(value FXSpatialScalerColorProcessingMode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorProcessingMode:"), value)
}/* debug [instance_properties/setter]: colorProcessingMode */


// The pixel format of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorTextureFormat
func (f_ FXSpatialScalerDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: colorTextureFormat */


// The pixel format of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorTextureFormat
func (f_ FXSpatialScalerDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}/* debug [instance_properties/setter]: colorTextureFormat */


// The height of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputHeight
func (f_ FXSpatialScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}/* debug [instance_properties/getter]: inputHeight */


// The height of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputHeight
func (f_ FXSpatialScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}/* debug [instance_properties/setter]: inputHeight */


// The width of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputWidth
func (f_ FXSpatialScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}/* debug [instance_properties/getter]: inputWidth */


// The width of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputWidth
func (f_ FXSpatialScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}/* debug [instance_properties/setter]: inputWidth */


// The height of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputHeight
func (f_ FXSpatialScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}/* debug [instance_properties/getter]: outputHeight */


// The height of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputHeight
func (f_ FXSpatialScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}/* debug [instance_properties/setter]: outputHeight */


// The pixel format of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputTextureFormat
func (f_ FXSpatialScalerDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: outputTextureFormat */


// The pixel format of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputTextureFormat
func (f_ FXSpatialScalerDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}/* debug [instance_properties/setter]: outputTextureFormat */


// The width of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputWidth
func (f_ FXSpatialScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}/* debug [instance_properties/getter]: outputWidth */


// The width of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputWidth
func (f_ FXSpatialScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}/* debug [instance_properties/setter]: outputWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFXSpatialScalerDescriptor */



