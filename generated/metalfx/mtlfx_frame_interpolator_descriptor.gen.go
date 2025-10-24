// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFXFrameInterpolatorDescriptor */


/* debug [class_header]: Header for MTLFXFrameInterpolatorDescriptor */
// The class instance for the [FXFrameInterpolatorDescriptor] class.
var (
	FXFrameInterpolatorDescriptorClass     _FXFrameInterpolatorDescriptorClass
	FXFrameInterpolatorDescriptorClassOnce sync.Once
)

func getFXFrameInterpolatorDescriptorClass() _FXFrameInterpolatorDescriptorClass {
	FXFrameInterpolatorDescriptorClassOnce.Do(func() {
		FXFrameInterpolatorDescriptorClass = _FXFrameInterpolatorDescriptorClass{objc.GetClass("MTLFXFrameInterpolatorDescriptor")}
	})
	return FXFrameInterpolatorDescriptorClass
}

type _FXFrameInterpolatorDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FXFrameInterpolatorDescriptor */
// An interface definition for the [FXFrameInterpolatorDescriptor] class.
type IFXFrameInterpolatorDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FXFrameInterpolatorDescriptor */
	// properties:
	ColorTextureFormat() PixelFormat /* not a class type */
	SetColorTextureFormat(value PixelFormat /* not a class type */)
	DepthTextureFormat() PixelFormat /* not a class type */
	SetDepthTextureFormat(value PixelFormat /* not a class type */)
	InputHeight() uint
	SetInputHeight(value uint)
	InputWidth() uint
	SetInputWidth(value uint)
	MotionTextureFormat() PixelFormat /* not a class type */
	SetMotionTextureFormat(value PixelFormat /* not a class type */)
	OutputHeight() uint
	SetOutputHeight(value uint)
	OutputTextureFormat() PixelFormat /* not a class type */
	SetOutputTextureFormat(value PixelFormat /* not a class type */)
	OutputWidth() uint
	SetOutputWidth(value uint)
	Scaler() unsafe.Pointer
	SetScaler(value unsafe.Pointer)
	UiTextureFormat() PixelFormat /* not a class type */
	SetUiTextureFormat(value PixelFormat /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FXFrameInterpolatorDescriptor */
	// methods:
	NewFrameInterpolatorWithDevice(device unsafe.Pointer) unsafe.Pointer
	NewFrameInterpolatorWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FXFrameInterpolatorDescriptor */
// Alloc allocates a new instance without initialization.
func (fc _FXFrameInterpolatorDescriptorClass) Alloc() FXFrameInterpolatorDescriptor {
	rv := objc.Send[FXFrameInterpolatorDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FXFrameInterpolatorDescriptorClass) New() FXFrameInterpolatorDescriptor {
	rv := objc.Send[FXFrameInterpolatorDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FXFrameInterpolatorDescriptor) Init() FXFrameInterpolatorDescriptor {
	rv := objc.Send[FXFrameInterpolatorDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FXFrameInterpolatorDescriptor) Autorelease() FXFrameInterpolatorDescriptor {
	rv := objc.Send[FXFrameInterpolatorDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFXFrameInterpolatorDescriptor creates a new FXFrameInterpolatorDescriptor instance.
func NewFXFrameInterpolatorDescriptor() FXFrameInterpolatorDescriptor {
	return getFXFrameInterpolatorDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FXFrameInterpolatorDescriptor */
// A set of properties that configure a frame interpolator, and a factory method that creates the effect.
//
// A frame interpolator inspects two frames your game or app renders and, based on their properties, generates an extra frame at a fraction of the cost, helping you to increase your frame rate. When you configure this descriptor, set the properties that determine the pixel format for each texture to the respective format of the texture you later assign to the scaler. For example, make sure that the format to which you set the property matches the format of the texture you later assign to the interpolator’s property.


// A set of properties that configure a frame interpolator, and a factory method that creates the effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor
type FXFrameInterpolatorDescriptor struct {
	objectivec.Object
}

// FXFrameInterpolatorDescriptorFrom constructs a [FXFrameInterpolatorDescriptor] from an unsafe.Pointer.
//
// A set of properties that configure a frame interpolator, and a factory method that creates the effect.
func FXFrameInterpolatorDescriptorFrom(ptr unsafe.Pointer) FXFrameInterpolatorDescriptor {
	return FXFrameInterpolatorDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FXFrameInterpolatorDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FXFrameInterpolatorDescriptor */

// Queries whether a Metal device supports frame interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsDevice(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsDevice(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsDevice) */


// Queries whether a Metal device supports frame interpolation compatible with a Metal 4 command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsMetal4FX(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsMetal4FX(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsMetal4FX) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FXFrameInterpolatorDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FXFrameInterpolatorDescriptor */

// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:"), device)
	return rv
}/* debug [instance_methods/method]: NewFrameInterpolatorWithDevice */


// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:compiler:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:compiler:"), device, compiler)
	return rv
}/* debug [instance_methods/method]: NewFrameInterpolatorWithDeviceCompiler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FXFrameInterpolatorDescriptor */

// The pixel format of the input color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: colorTextureFormat */


// The pixel format of the input color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}/* debug [instance_properties/setter]: colorTextureFormat */


// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) DepthTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: depthTextureFormat */


// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetDepthTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}/* debug [instance_properties/setter]: depthTextureFormat */


// The height, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}/* debug [instance_properties/getter]: inputHeight */


// The height, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}/* debug [instance_properties/setter]: inputHeight */


// The width, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}/* debug [instance_properties/getter]: inputWidth */


// The width, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}/* debug [instance_properties/setter]: inputWidth */


// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) MotionTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: motionTextureFormat */


// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetMotionTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}/* debug [instance_properties/setter]: motionTextureFormat */


// The height, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}/* debug [instance_properties/getter]: outputHeight */


// The height, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}/* debug [instance_properties/setter]: outputHeight */


// The pixel format of the output color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: outputTextureFormat */


// The pixel format of the output color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}/* debug [instance_properties/setter]: outputTextureFormat */


// The width, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}/* debug [instance_properties/getter]: outputWidth */


// The width, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}/* debug [instance_properties/setter]: outputWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) Scaler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("scaler"))
	return rv
}/* debug [instance_properties/getter]: scaler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) SetScaler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setScaler:"), value)
}/* debug [instance_properties/setter]: scaler */


// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) UiTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("uiTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: uiTextureFormat */


// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetUiTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUiTextureFormat:"), value)
}/* debug [instance_properties/setter]: uiTextureFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFXFrameInterpolatorDescriptor */



