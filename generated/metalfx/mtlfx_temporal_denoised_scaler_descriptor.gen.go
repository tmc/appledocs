// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FXTemporalDenoisedScalerDescriptor] class.
var (
	FXTemporalDenoisedScalerDescriptorClass     _FXTemporalDenoisedScalerDescriptorClass
	FXTemporalDenoisedScalerDescriptorClassOnce sync.Once
)

func getFXTemporalDenoisedScalerDescriptorClass() _FXTemporalDenoisedScalerDescriptorClass {
	FXTemporalDenoisedScalerDescriptorClassOnce.Do(func() {
		FXTemporalDenoisedScalerDescriptorClass = _FXTemporalDenoisedScalerDescriptorClass{objc.GetClass("MTLFXTemporalDenoisedScalerDescriptor")}
	})
	return FXTemporalDenoisedScalerDescriptorClass
}

type _FXTemporalDenoisedScalerDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [FXTemporalDenoisedScalerDescriptor] class.
type IFXTemporalDenoisedScalerDescriptor interface {
	objectivec.IObject
	NewTemporalDenoisedScalerWithDevice(device objectivec.IObject) objc.ID
	NewTemporalDenoisedScalerWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor
type FXTemporalDenoisedScalerDescriptor struct {
	objectivec.Object
}

// FXTemporalDenoisedScalerDescriptorFrom constructs a [FXTemporalDenoisedScalerDescriptor] from an unsafe.Pointer.
func FXTemporalDenoisedScalerDescriptorFrom(ptr unsafe.Pointer) FXTemporalDenoisedScalerDescriptor {
	return FXTemporalDenoisedScalerDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FXTemporalDenoisedScalerDescriptorClass) Alloc() FXTemporalDenoisedScalerDescriptor {
	rv := objc.Send[FXTemporalDenoisedScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FXTemporalDenoisedScalerDescriptorClass) New() FXTemporalDenoisedScalerDescriptor {
	rv := objc.Send[FXTemporalDenoisedScalerDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FXTemporalDenoisedScalerDescriptor) Init() FXTemporalDenoisedScalerDescriptor {
	rv := objc.Send[FXTemporalDenoisedScalerDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FXTemporalDenoisedScalerDescriptor) Autorelease() FXTemporalDenoisedScalerDescriptor {
	rv := objc.Send[FXTemporalDenoisedScalerDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFXTemporalDenoisedScalerDescriptor creates a new FXTemporalDenoisedScalerDescriptor instance.
func NewFXTemporalDenoisedScalerDescriptor() FXTemporalDenoisedScalerDescriptor {
	return getFXTemporalDenoisedScalerDescriptorClass().New()
}


// Returns the largest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMaxScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMaxScaleForDevice(device objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("supportedInputContentMaxScaleForDevice:"), device)
	return rv
}

// Returns the smallest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMinScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMinScaleForDevice(device objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("supportedInputContentMinScaleForDevice:"), device)
	return rv
}

// Queries whether a Metal device supports denoising scaling.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsDevice(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}

// Queries whether a Metal device supports denosing scaling compatible on Metal 4.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsMetal4FX(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}

// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDevice(device objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:"), device)
	return rv
}

// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:compiler:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:compiler:"), device, compiler)
	return rv
}

// The pixel format of the input color texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ColorTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// SetColorTextureFormat sets the value of the colorTextureFormat property.
// The pixel format of the input color texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetColorTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}

// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("denoiseStrengthMaskTextureFormat"))
	return rv
}


// SetDenoiseStrengthMaskTextureFormat sets the value of the denoiseStrengthMaskTextureFormat property.
// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureFormat:"), value)
}

// The pixel format of the input depth texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DepthTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}


// SetDepthTextureFormat sets the value of the depthTextureFormat property.
// The pixel format of the input depth texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDepthTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}

// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DiffuseAlbedoTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("diffuseAlbedoTextureFormat"))
	return rv
}


// SetDiffuseAlbedoTextureFormat sets the value of the diffuseAlbedoTextureFormat property.
// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDiffuseAlbedoTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDiffuseAlbedoTextureFormat:"), value)
}

// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// SetInputHeight sets the value of the inputHeight property.
// The height, in pixels, of the input color texture for the denoiser scaler.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}

// The width, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// SetInputWidth sets the value of the inputWidth property.
// The width, in pixels, of the input color texture for the denoiser scaler.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}

// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) AutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("autoExposureEnabled"))
	return rv
}


// SetAutoExposureEnabled sets the value of the autoExposureEnabled property.
// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAutoExposureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("denoiseStrengthMaskTextureEnabled"))
	return rv
}


// SetDenoiseStrengthMaskTextureEnabled sets the value of the denoiseStrengthMaskTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureEnabled:"), value)
}

// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reactiveMaskTextureEnabled"))
	return rv
}


// SetReactiveMaskTextureEnabled sets the value of the reactiveMaskTextureEnabled property.
// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("specularHitDistanceTextureEnabled"))
	return rv
}


// SetSpecularHitDistanceTextureEnabled sets the value of the specularHitDistanceTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("transparencyOverlayTextureEnabled"))
	return rv
}


// SetTransparencyOverlayTextureEnabled sets the value of the transparencyOverlayTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureEnabled:"), value)
}

// The pixel format of the input motion texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) MotionTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}


// SetMotionTextureFormat sets the value of the motionTextureFormat property.
// The pixel format of the input motion texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetMotionTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}

// The pixel format of the input normal texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) NormalTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("normalTextureFormat"))
	return rv
}


// SetNormalTextureFormat sets the value of the normalTextureFormat property.
// The pixel format of the input normal texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetNormalTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNormalTextureFormat:"), value)
}

// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// SetOutputHeight sets the value of the outputHeight property.
// The height, in pixels, of the input color texture for the denoiser scaler.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}

// The pixel format of the output color texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) OutputTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// SetOutputTextureFormat sets the value of the outputTextureFormat property.
// The pixel format of the output color texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}

// The width, in pixels, of the output color texture for the denoiser scaler.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// SetOutputWidth sets the value of the outputWidth property.
// The width, in pixels, of the output color texture for the denoiser scaler.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}

// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("reactiveMaskTextureFormat"))
	return rv
}


// SetReactiveMaskTextureFormat sets the value of the reactiveMaskTextureFormat property.
// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureFormat:"), value)
}

// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) RequiresSynchronousInitialization() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresSynchronousInitialization"))
	return rv
}


// SetRequiresSynchronousInitialization sets the value of the requiresSynchronousInitialization property.
// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) SetRequiresSynchronousInitialization(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresSynchronousInitialization:"), value)
}

// The pixel format of the input roughness texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) RoughnessTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("roughnessTextureFormat"))
	return rv
}


// SetRoughnessTextureFormat sets the value of the roughnessTextureFormat property.
// The pixel format of the input roughness texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetRoughnessTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRoughnessTextureFormat:"), value)
}

// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularAlbedoTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("specularAlbedoTextureFormat"))
	return rv
}


// SetSpecularAlbedoTextureFormat sets the value of the specularAlbedoTextureFormat property.
// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularAlbedoTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularAlbedoTextureFormat:"), value)
}

// The pixel format of the input specular hit texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("specularHitDistanceTextureFormat"))
	return rv
}


// SetSpecularHitDistanceTextureFormat sets the value of the specularHitDistanceTextureFormat property.
// The pixel format of the input specular hit texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureFormat:"), value)
}

// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("transparencyOverlayTextureFormat"))
	return rv
}


// SetTransparencyOverlayTextureFormat sets the value of the transparencyOverlayTextureFormat property.
// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureFormat:"), value)
}

// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsAutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAutoExposureEnabled"))
	return rv
}


// SetIsAutoExposureEnabled sets the value of the isAutoExposureEnabled property.
// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.

//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsAutoExposureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsDenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDenoiseStrengthMaskTextureEnabled"))
	return rv
}


// SetIsDenoiseStrengthMaskTextureEnabled sets the value of the isDenoiseStrengthMaskTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDenoiseStrengthMaskTextureEnabled:"), value)
}

// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReactiveMaskTextureEnabled"))
	return rv
}


// SetIsReactiveMaskTextureEnabled sets the value of the isReactiveMaskTextureEnabled property.
// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.

//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsReactiveMaskTextureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsSpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSpecularHitDistanceTextureEnabled"))
	return rv
}


// SetIsSpecularHitDistanceTextureEnabled sets the value of the isSpecularHitDistanceTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSpecularHitDistanceTextureEnabled:"), value)
}

// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsTransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isTransparencyOverlayTextureEnabled"))
	return rv
}


// SetIsTransparencyOverlayTextureEnabled sets the value of the isTransparencyOverlayTextureEnabled property.
// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.

//
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsTransparencyOverlayTextureEnabled:"), value)
}



