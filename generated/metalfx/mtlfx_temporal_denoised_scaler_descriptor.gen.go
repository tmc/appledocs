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
	

	// properties:
	ColorTextureFormat() PixelFormat /* not a class type */
	SetColorTextureFormat(value PixelFormat /* not a class type */)
	DenoiseStrengthMaskTextureFormat() PixelFormat /* not a class type */
	SetDenoiseStrengthMaskTextureFormat(value PixelFormat /* not a class type */)
	DepthTextureFormat() PixelFormat /* not a class type */
	SetDepthTextureFormat(value PixelFormat /* not a class type */)
	DiffuseAlbedoTextureFormat() PixelFormat /* not a class type */
	SetDiffuseAlbedoTextureFormat(value PixelFormat /* not a class type */)
	InputHeight() uint
	SetInputHeight(value uint)
	InputWidth() uint
	SetInputWidth(value uint)
	AutoExposureEnabled() bool
	SetAutoExposureEnabled(value bool)
	DenoiseStrengthMaskTextureEnabled() bool
	SetDenoiseStrengthMaskTextureEnabled(value bool)
	ReactiveMaskTextureEnabled() bool
	SetReactiveMaskTextureEnabled(value bool)
	SpecularHitDistanceTextureEnabled() bool
	SetSpecularHitDistanceTextureEnabled(value bool)
	TransparencyOverlayTextureEnabled() bool
	SetTransparencyOverlayTextureEnabled(value bool)
	MotionTextureFormat() PixelFormat /* not a class type */
	SetMotionTextureFormat(value PixelFormat /* not a class type */)
	NormalTextureFormat() PixelFormat /* not a class type */
	SetNormalTextureFormat(value PixelFormat /* not a class type */)
	OutputHeight() uint
	SetOutputHeight(value uint)
	OutputTextureFormat() PixelFormat /* not a class type */
	SetOutputTextureFormat(value PixelFormat /* not a class type */)
	OutputWidth() uint
	SetOutputWidth(value uint)
	ReactiveMaskTextureFormat() PixelFormat /* not a class type */
	SetReactiveMaskTextureFormat(value PixelFormat /* not a class type */)
	RequiresSynchronousInitialization() bool
	SetRequiresSynchronousInitialization(value bool)
	RoughnessTextureFormat() PixelFormat /* not a class type */
	SetRoughnessTextureFormat(value PixelFormat /* not a class type */)
	SpecularAlbedoTextureFormat() PixelFormat /* not a class type */
	SetSpecularAlbedoTextureFormat(value PixelFormat /* not a class type */)
	SpecularHitDistanceTextureFormat() PixelFormat /* not a class type */
	SetSpecularHitDistanceTextureFormat(value PixelFormat /* not a class type */)
	TransparencyOverlayTextureFormat() PixelFormat /* not a class type */
	SetTransparencyOverlayTextureFormat(value PixelFormat /* not a class type */)
	IsAutoExposureEnabled() bool
	SetIsAutoExposureEnabled(value bool)
	IsDenoiseStrengthMaskTextureEnabled() bool
	SetIsDenoiseStrengthMaskTextureEnabled(value bool)
	IsReactiveMaskTextureEnabled() bool
	SetIsReactiveMaskTextureEnabled(value bool)
	IsSpecularHitDistanceTextureEnabled() bool
	SetIsSpecularHitDistanceTextureEnabled(value bool)
	IsTransparencyOverlayTextureEnabled() bool
	SetIsTransparencyOverlayTextureEnabled(value bool)


	

	// methods:
	NewTemporalDenoisedScalerWithDevice(device unsafe.Pointer) unsafe.Pointer
	NewTemporalDenoisedScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (fc _FXTemporalDenoisedScalerDescriptorClass) Alloc() FXTemporalDenoisedScalerDescriptor {
	rv := objc.Send[FXTemporalDenoisedScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor
type FXTemporalDenoisedScalerDescriptor struct {
	objectivec.Object
}

// FXTemporalDenoisedScalerDescriptorFrom constructs a [FXTemporalDenoisedScalerDescriptor] from an unsafe.Pointer.
func FXTemporalDenoisedScalerDescriptorFrom(ptr unsafe.Pointer) FXTemporalDenoisedScalerDescriptor {
	return FXTemporalDenoisedScalerDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// Returns the largest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMaxScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMaxScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMaxScaleForDevice:"), device)
	return rv
}


// Returns the smallest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMinScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMinScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMinScaleForDevice:"), device)
	return rv
}


// Queries whether a Metal device supports denoising scaling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsDevice(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsDevice(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}


// Queries whether a Metal device supports denosing scaling compatible on Metal 4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsMetal4FX(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}












// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:"), device)
	return rv
}


// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:compiler:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:compiler:"), device, compiler)
	return rv
}







// The pixel format of the input color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// The pixel format of the input color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}


// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("denoiseStrengthMaskTextureFormat"))
	return rv
}


// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureFormat:"), value)
}


// The pixel format of the input depth texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DepthTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}


// The pixel format of the input depth texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDepthTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}


// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DiffuseAlbedoTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("diffuseAlbedoTextureFormat"))
	return rv
}


// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDiffuseAlbedoTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDiffuseAlbedoTextureFormat:"), value)
}


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}


// The width, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// The width, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) AutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("autoExposureEnabled"))
	return rv
}


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAutoExposureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("denoiseStrengthMaskTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureEnabled:"), value)
}


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reactiveMaskTextureEnabled"))
	return rv
}


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("specularHitDistanceTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("transparencyOverlayTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureEnabled:"), value)
}


// The pixel format of the input motion texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) MotionTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}


// The pixel format of the input motion texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetMotionTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}


// The pixel format of the input normal texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) NormalTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("normalTextureFormat"))
	return rv
}


// The pixel format of the input normal texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetNormalTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNormalTextureFormat:"), value)
}


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}


// The pixel format of the output color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// The pixel format of the output color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}


// The width, in pixels, of the output color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// The width, in pixels, of the output color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}


// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("reactiveMaskTextureFormat"))
	return rv
}


// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureFormat:"), value)
}


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) RequiresSynchronousInitialization() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresSynchronousInitialization"))
	return rv
}


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) SetRequiresSynchronousInitialization(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresSynchronousInitialization:"), value)
}


// The pixel format of the input roughness texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) RoughnessTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("roughnessTextureFormat"))
	return rv
}


// The pixel format of the input roughness texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetRoughnessTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRoughnessTextureFormat:"), value)
}


// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularAlbedoTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("specularAlbedoTextureFormat"))
	return rv
}


// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularAlbedoTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularAlbedoTextureFormat:"), value)
}


// The pixel format of the input specular hit texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("specularHitDistanceTextureFormat"))
	return rv
}


// The pixel format of the input specular hit texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureFormat:"), value)
}


// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("transparencyOverlayTextureFormat"))
	return rv
}


// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureFormat:"), value)
}


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsAutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAutoExposureEnabled"))
	return rv
}


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsAutoExposureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsDenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDenoiseStrengthMaskTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDenoiseStrengthMaskTextureEnabled:"), value)
}


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReactiveMaskTextureEnabled"))
	return rv
}


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsReactiveMaskTextureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsSpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSpecularHitDistanceTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSpecularHitDistanceTextureEnabled:"), value)
}


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsTransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isTransparencyOverlayTextureEnabled"))
	return rv
}


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsTransparencyOverlayTextureEnabled:"), value)
}








