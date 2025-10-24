// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FXFrameInterpolatorDescriptor] class.
type IFXFrameInterpolatorDescriptor interface {
	objectivec.IObject
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
	Scaler() objc.ID
	SetScaler(value objc.ID)
	UiTextureFormat() PixelFormat /* not a class type */
	SetUiTextureFormat(value PixelFormat /* not a class type */)
	// methods:
	NewFrameInterpolatorWithDevice(device objectivec.IObject) objc.ID
	NewFrameInterpolatorWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID
}

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

// Alloc allocates a new instance without initialization.
func (fc _FXFrameInterpolatorDescriptorClass) Alloc() FXFrameInterpolatorDescriptor {
	rv := objc.Send[FXFrameInterpolatorDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Queries whether a Metal device supports frame interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsDevice(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}


// Queries whether a Metal device supports frame interpolation compatible with a Metal 4 command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsMetal4FX(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsMetal4FX(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}


// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDevice(device objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:"), device)
	return rv
}


// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:compiler:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:compiler:"), device, compiler)
	return rv
}


// The pixel format of the input color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// The pixel format of the input color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}


// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) DepthTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}


// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetDepthTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}


// The height, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// The height, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}


// The width, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// The width, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}


// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) MotionTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}


// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetMotionTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}


// The height, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// The height, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}


// The pixel format of the output color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// The pixel format of the output color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}


// The width, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// The width, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) Scaler() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("scaler"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) SetScaler(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setScaler:"), value)
}


// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) UiTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("uiTextureFormat"))
	return rv
}


// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetUiTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUiTextureFormat:"), value)
}



