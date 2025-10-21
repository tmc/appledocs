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
	NewFrameInterpolatorWithDevice(device objectivec.IObject) objc.ID
	NewFrameInterpolatorWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID
}

// A set of properties that configure a frame interpolator, and a factory method that creates the effect.
//
// A frame interpolator inspects two frames your game or app renders and, based on their properties, generates an extra frame at a fraction of the cost, helping you to increase your frame rate. When you configure this descriptor, set the properties that determine the pixel format for each texture to the respective format of the texture you later assign to the scaler. For example, make sure that the format to which you set the property matches the format of the texture you later assign to the interpolator’s property.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsDevice(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}

// Queries whether a Metal device supports frame interpolation compatible with a Metal 4 command buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/supportsMetal4FX(_:)
func (fc _FXFrameInterpolatorDescriptorClass) SupportsMetal4FX(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}

// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDevice(device objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:"), device)
	return rv
}

// Creates a frame interpolator instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/makeFrameInterpolator(device:compiler:)
func (f_ FXFrameInterpolatorDescriptor) NewFrameInterpolatorWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newFrameInterpolatorWithDevice:compiler:"), device, compiler)
	return rv
}

// The pixel format of the input color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) ColorTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// SetColorTextureFormat sets the value of the colorTextureFormat property.
// The pixel format of the input color texture for the frame interpolator you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/colorTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetColorTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}

// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) DepthTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}


// SetDepthTextureFormat sets the value of the depthTextureFormat property.
// The pixel format of the input depth texture for the frame interpolator you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/depthTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetDepthTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}

// The height, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// SetInputHeight sets the value of the inputHeight property.
// The height, in pixels, of the input motion and depth texture for the frame interpolator.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputHeight
func (f_ FXFrameInterpolatorDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}

// The width, in pixels, of the input motion and depth texture for the frame interpolator.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// SetInputWidth sets the value of the inputWidth property.
// The width, in pixels, of the input motion and depth texture for the frame interpolator.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/inputWidth
func (f_ FXFrameInterpolatorDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}

// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) MotionTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}


// SetMotionTextureFormat sets the value of the motionTextureFormat property.
// The pixel format of the input motion texture for the frame interpolator you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/motionTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetMotionTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}

// The height, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// SetOutputHeight sets the value of the outputHeight property.
// The height, in pixels, of the output color texture for the frame interpolator.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputHeight
func (f_ FXFrameInterpolatorDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}

// The pixel format of the output color texture for the frame interpolator you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) OutputTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// SetOutputTextureFormat sets the value of the outputTextureFormat property.
// The pixel format of the output color texture for the frame interpolator you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetOutputTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}

// The width, in pixels, of the output color texture for the frame interpolator.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// SetOutputWidth sets the value of the outputWidth property.
// The width, in pixels, of the output color texture for the frame interpolator.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/outputWidth
func (f_ FXFrameInterpolatorDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) Scaler() objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("scaler"))
	return rv
}


// SetScaler sets the value of the scaler property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/scaler
func (f_ FXFrameInterpolatorDescriptor) SetScaler(value objc.ID) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setScaler:"), value)
}

// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) UiTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("uiTextureFormat"))
	return rv
}


// SetUiTextureFormat sets the value of the uiTextureFormat property.
// The pixel format for the frame interpolator of an input texture containing your game’s custom UI.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXFrameInterpolatorDescriptor/uiTextureFormat
func (f_ FXFrameInterpolatorDescriptor) SetUiTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUiTextureFormat:"), value)
}



