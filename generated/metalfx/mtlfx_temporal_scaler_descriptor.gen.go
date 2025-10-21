// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FXTemporalScalerDescriptor] class.
var (
	FXTemporalScalerDescriptorClass     _FXTemporalScalerDescriptorClass
	FXTemporalScalerDescriptorClassOnce sync.Once
)

func getFXTemporalScalerDescriptorClass() _FXTemporalScalerDescriptorClass {
	FXTemporalScalerDescriptorClassOnce.Do(func() {
		FXTemporalScalerDescriptorClass = _FXTemporalScalerDescriptorClass{objc.GetClass("MTLFXTemporalScalerDescriptor")}
	})
	return FXTemporalScalerDescriptorClass
}

type _FXTemporalScalerDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [FXTemporalScalerDescriptor] class.
type IFXTemporalScalerDescriptor interface {
	objectivec.IObject
	NewTemporalScalerWithDevice(device objc.ID) objc.ID
	NewTemporalScalerWithDeviceCompiler(device objc.ID, compiler objc.ID) objc.ID
}

// A set of properties that configure a temporal scaling effect, and a factory method that creates the effect.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor
type FXTemporalScalerDescriptor struct {
	objectivec.Object
}

// FXTemporalScalerDescriptorFrom constructs a [FXTemporalScalerDescriptor] from an unsafe.Pointer.
//
// A set of properties that configure a temporal scaling effect, and a factory method that creates the effect.
func FXTemporalScalerDescriptorFrom(ptr unsafe.Pointer) FXTemporalScalerDescriptor {
	return FXTemporalScalerDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FXTemporalScalerDescriptorClass) Alloc() FXTemporalScalerDescriptor {
	rv := objc.Send[FXTemporalScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FXTemporalScalerDescriptorClass) New() FXTemporalScalerDescriptor {
	rv := objc.Send[FXTemporalScalerDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FXTemporalScalerDescriptor) Init() FXTemporalScalerDescriptor {
	rv := objc.Send[FXTemporalScalerDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FXTemporalScalerDescriptor) Autorelease() FXTemporalScalerDescriptor {
	rv := objc.Send[FXTemporalScalerDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFXTemporalScalerDescriptor creates a new FXTemporalScalerDescriptor instance.
func NewFXTemporalScalerDescriptor() FXTemporalScalerDescriptor {
	return getFXTemporalScalerDescriptorClass().New()
}


// Returns the largest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportedInputContentMaxScale(device:)
func (fc _FXTemporalScalerDescriptorClass) SupportedInputContentMaxScaleForDevice(device objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("supportedInputContentMaxScaleForDevice:"), device)
	return rv
}

// Returns the smallest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportedInputContentMinScale(device:)
func (fc _FXTemporalScalerDescriptorClass) SupportedInputContentMinScaleForDevice(device objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("supportedInputContentMinScaleForDevice:"), device)
	return rv
}

// Returns a Boolean value that indicates whether the temporal scaler works with a GPU.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportsDevice(_:)
func (fc _FXTemporalScalerDescriptorClass) SupportsDevice(device objc.ID) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}

// Queries whether a Metal device supports temporal scaling compatible with Metal 4.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXTemporalScalerDescriptorClass) SupportsMetal4FX(device objc.ID) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}

// Creates a temporal scaler instance from this descriptor’s current property values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/makeTemporalScaler(device:)
func (f_ FXTemporalScalerDescriptor) NewTemporalScalerWithDevice(device objc.ID) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newTemporalScalerWithDevice:"), device)
	return rv
}

// Creates a temporal scaler instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/makeTemporalScaler(device:compiler:)
func (f_ FXTemporalScalerDescriptor) NewTemporalScalerWithDeviceCompiler(device objc.ID, compiler objc.ID) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newTemporalScalerWithDevice:compiler:"), device, compiler)
	return rv
}

// The pixel format of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/colorTextureFormat
func (f_ FXTemporalScalerDescriptor) ColorTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// SetColorTextureFormat sets the value of the colorTextureFormat property.
// The pixel format of the input color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/colorTextureFormat
func (f_ FXTemporalScalerDescriptor) SetColorTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}

// The pixel format of the input depth texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/depthTextureFormat
func (f_ FXTemporalScalerDescriptor) DepthTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}


// SetDepthTextureFormat sets the value of the depthTextureFormat property.
// The pixel format of the input depth texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/depthTextureFormat
func (f_ FXTemporalScalerDescriptor) SetDepthTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}

// The largest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMaxScale
func (f_ FXTemporalScalerDescriptor) InputContentMaxScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("inputContentMaxScale"))
	return rv
}


// SetInputContentMaxScale sets the value of the inputContentMaxScale property.
// The largest scale factor the temporal scaler you create with this descriptor can use to generate output textures.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMaxScale
func (f_ FXTemporalScalerDescriptor) SetInputContentMaxScale(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentMaxScale:"), value)
}

// The smallest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMinScale
func (f_ FXTemporalScalerDescriptor) InputContentMinScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("inputContentMinScale"))
	return rv
}


// SetInputContentMinScale sets the value of the inputContentMinScale property.
// The smallest scale factor the temporal scaler you create with this descriptor can use to generate output textures.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMinScale
func (f_ FXTemporalScalerDescriptor) SetInputContentMinScale(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentMinScale:"), value)
}

// The height of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputHeight
func (f_ FXTemporalScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// SetInputHeight sets the value of the inputHeight property.
// The height of the input color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputHeight
func (f_ FXTemporalScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}

// The width of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputWidth
func (f_ FXTemporalScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// SetInputWidth sets the value of the inputWidth property.
// The width of the input color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputWidth
func (f_ FXTemporalScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}

// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalScalerDescriptor) AutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("autoExposureEnabled"))
	return rv
}


// SetAutoExposureEnabled sets the value of the autoExposureEnabled property.
// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalScalerDescriptor) SetAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAutoExposureEnabled:"), value)
}

// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isInputContentPropertiesEnabled
func (f_ FXTemporalScalerDescriptor) InputContentPropertiesEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("inputContentPropertiesEnabled"))
	return rv
}


// SetInputContentPropertiesEnabled sets the value of the inputContentPropertiesEnabled property.
// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isInputContentPropertiesEnabled
func (f_ FXTemporalScalerDescriptor) SetInputContentPropertiesEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentPropertiesEnabled:"), value)
}

// A Boolean value that indicates whether a temporal scaler you create with the descriptor applies a reactive mask.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalScalerDescriptor) ReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reactiveMaskTextureEnabled"))
	return rv
}


// SetReactiveMaskTextureEnabled sets the value of the reactiveMaskTextureEnabled property.
// A Boolean value that indicates whether a temporal scaler you create with the descriptor applies a reactive mask.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalScalerDescriptor) SetReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureEnabled:"), value)
}

// The pixel format of the input motion texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/motionTextureFormat
func (f_ FXTemporalScalerDescriptor) MotionTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}


// SetMotionTextureFormat sets the value of the motionTextureFormat property.
// The pixel format of the input motion texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/motionTextureFormat
func (f_ FXTemporalScalerDescriptor) SetMotionTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}

// The height of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputHeight
func (f_ FXTemporalScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// SetOutputHeight sets the value of the outputHeight property.
// The height of the output color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputHeight
func (f_ FXTemporalScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}

// The pixel format of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputTextureFormat
func (f_ FXTemporalScalerDescriptor) OutputTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// SetOutputTextureFormat sets the value of the outputTextureFormat property.
// The pixel format of the output color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputTextureFormat
func (f_ FXTemporalScalerDescriptor) SetOutputTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}

// The width of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputWidth
func (f_ FXTemporalScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// SetOutputWidth sets the value of the outputWidth property.
// The width of the output color texture for the temporal scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputWidth
func (f_ FXTemporalScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}

// The pixel format of the reactive mask input texture for a temporal scaler you create with the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalScalerDescriptor) ReactiveMaskTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("reactiveMaskTextureFormat"))
	return rv
}


// SetReactiveMaskTextureFormat sets the value of the reactiveMaskTextureFormat property.
// The pixel format of the reactive mask input texture for a temporal scaler you create with the descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalScalerDescriptor) SetReactiveMaskTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureFormat:"), value)
}

// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalScalerDescriptor) RequiresSynchronousInitialization() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresSynchronousInitialization"))
	return rv
}


// SetRequiresSynchronousInitialization sets the value of the requiresSynchronousInitialization property.
// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalScalerDescriptor) SetRequiresSynchronousInitialization(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresSynchronousInitialization:"), value)
}




