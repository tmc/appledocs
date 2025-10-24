// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFXTemporalDenoisedScalerDescriptor */


/* debug [class_header]: Header for MTLFXTemporalDenoisedScalerDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FXTemporalDenoisedScalerDescriptor */
// An interface definition for the [FXTemporalDenoisedScalerDescriptor] class.
type IFXTemporalDenoisedScalerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FXTemporalDenoisedScalerDescriptor */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FXTemporalDenoisedScalerDescriptor */
	// methods:
	NewTemporalDenoisedScalerWithDevice(device unsafe.Pointer) unsafe.Pointer
	NewTemporalDenoisedScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FXTemporalDenoisedScalerDescriptor */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FXTemporalDenoisedScalerDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor
type FXTemporalDenoisedScalerDescriptor struct {
	objectivec.Object
}

// FXTemporalDenoisedScalerDescriptorFrom constructs a [FXTemporalDenoisedScalerDescriptor] from an unsafe.Pointer.
func FXTemporalDenoisedScalerDescriptorFrom(ptr unsafe.Pointer) FXTemporalDenoisedScalerDescriptor {
	return FXTemporalDenoisedScalerDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FXTemporalDenoisedScalerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FXTemporalDenoisedScalerDescriptor */

// Returns the largest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMaxScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMaxScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMaxScaleForDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedInputContentMaxScaleForDevice) */


// Returns the smallest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportedInputContentMinScale(device:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportedInputContentMinScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMinScaleForDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedInputContentMinScaleForDevice) */


// Queries whether a Metal device supports denoising scaling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsDevice(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsDevice(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsDevice) */


// Queries whether a Metal device supports denosing scaling compatible on Metal 4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXTemporalDenoisedScalerDescriptorClass) SupportsMetal4FX(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsMetal4FX) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FXTemporalDenoisedScalerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FXTemporalDenoisedScalerDescriptor */

// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:"), device)
	return rv
}/* debug [instance_methods/method]: NewTemporalDenoisedScalerWithDevice */


// Creates a denoiser scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/makeTemporalDenoisedScaler(device:compiler:)
func (f_ FXTemporalDenoisedScalerDescriptor) NewTemporalDenoisedScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalDenoisedScalerWithDevice:compiler:"), device, compiler)
	return rv
}/* debug [instance_methods/method]: NewTemporalDenoisedScalerWithDeviceCompiler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FXTemporalDenoisedScalerDescriptor */

// The pixel format of the input color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: colorTextureFormat */


// The pixel format of the input color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/colorTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}/* debug [instance_properties/setter]: colorTextureFormat */


// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("denoiseStrengthMaskTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: denoiseStrengthMaskTextureFormat */


// The pixel format of the input denoise strength mask texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/denoiseStrengthMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureFormat:"), value)
}/* debug [instance_properties/setter]: denoiseStrengthMaskTextureFormat */


// The pixel format of the input depth texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DepthTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: depthTextureFormat */


// The pixel format of the input depth texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/depthTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDepthTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}/* debug [instance_properties/setter]: depthTextureFormat */


// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) DiffuseAlbedoTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("diffuseAlbedoTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: diffuseAlbedoTextureFormat */


// The pixel format of the input diffuse albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/diffuseAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetDiffuseAlbedoTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDiffuseAlbedoTextureFormat:"), value)
}/* debug [instance_properties/setter]: diffuseAlbedoTextureFormat */


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}/* debug [instance_properties/getter]: inputHeight */


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}/* debug [instance_properties/setter]: inputHeight */


// The width, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}/* debug [instance_properties/getter]: inputWidth */


// The width, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/inputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}/* debug [instance_properties/setter]: inputWidth */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) AutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("autoExposureEnabled"))
	return rv
}/* debug [instance_properties/getter]: autoExposureEnabled */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAutoExposureEnabled:"), value)
}/* debug [instance_properties/setter]: autoExposureEnabled */


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) DenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("denoiseStrengthMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: denoiseStrengthMaskTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isDenoiseStrengthMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDenoiseStrengthMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: denoiseStrengthMaskTextureEnabled */


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reactiveMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: reactiveMaskTextureEnabled */


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: reactiveMaskTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("specularHitDistanceTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: specularHitDistanceTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isSpecularHitDistanceTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureEnabled:"), value)
}/* debug [instance_properties/setter]: specularHitDistanceTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("transparencyOverlayTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: transparencyOverlayTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/isTransparencyOverlayTextureEnabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureEnabled:"), value)
}/* debug [instance_properties/setter]: transparencyOverlayTextureEnabled */


// The pixel format of the input motion texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) MotionTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: motionTextureFormat */


// The pixel format of the input motion texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/motionTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetMotionTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}/* debug [instance_properties/setter]: motionTextureFormat */


// The pixel format of the input normal texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) NormalTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("normalTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: normalTextureFormat */


// The pixel format of the input normal texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/normalTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetNormalTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNormalTextureFormat:"), value)
}/* debug [instance_properties/setter]: normalTextureFormat */


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}/* debug [instance_properties/getter]: outputHeight */


// The height, in pixels, of the input color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputHeight
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}/* debug [instance_properties/setter]: outputHeight */


// The pixel format of the output color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: outputTextureFormat */


// The pixel format of the output color texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}/* debug [instance_properties/setter]: outputTextureFormat */


// The width, in pixels, of the output color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}/* debug [instance_properties/getter]: outputWidth */


// The width, in pixels, of the output color texture for the denoiser scaler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/outputWidth
func (f_ FXTemporalDenoisedScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}/* debug [instance_properties/setter]: outputWidth */


// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) ReactiveMaskTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("reactiveMaskTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: reactiveMaskTextureFormat */


// The pixel format of the reactive mask input texture for a scaler you create from this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetReactiveMaskTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureFormat:"), value)
}/* debug [instance_properties/setter]: reactiveMaskTextureFormat */


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) RequiresSynchronousInitialization() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresSynchronousInitialization"))
	return rv
}/* debug [instance_properties/getter]: requiresSynchronousInitialization */


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalDenoisedScalerDescriptor) SetRequiresSynchronousInitialization(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresSynchronousInitialization:"), value)
}/* debug [instance_properties/setter]: requiresSynchronousInitialization */


// The pixel format of the input roughness texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) RoughnessTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("roughnessTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: roughnessTextureFormat */


// The pixel format of the input roughness texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/roughnessTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetRoughnessTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRoughnessTextureFormat:"), value)
}/* debug [instance_properties/setter]: roughnessTextureFormat */


// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularAlbedoTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("specularAlbedoTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: specularAlbedoTextureFormat */


// The pixel format of the input specular albedo texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularAlbedoTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularAlbedoTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularAlbedoTextureFormat:"), value)
}/* debug [instance_properties/setter]: specularAlbedoTextureFormat */


// The pixel format of the input specular hit texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SpecularHitDistanceTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("specularHitDistanceTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: specularHitDistanceTextureFormat */


// The pixel format of the input specular hit texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/specularHitDistanceTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetSpecularHitDistanceTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecularHitDistanceTextureFormat:"), value)
}/* debug [instance_properties/setter]: specularHitDistanceTextureFormat */


// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) TransparencyOverlayTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("transparencyOverlayTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: transparencyOverlayTextureFormat */


// The pixel format of the input transparency overlay texture for the scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalDenoisedScalerDescriptor/transparencyOverlayTextureFormat
func (f_ FXTemporalDenoisedScalerDescriptor) SetTransparencyOverlayTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTransparencyOverlayTextureFormat:"), value)
}/* debug [instance_properties/setter]: transparencyOverlayTextureFormat */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsAutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAutoExposureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoExposureEnabled */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsAutoExposureEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoExposureEnabled */


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsDenoiseStrengthMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isDenoiseStrengthMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDenoiseStrengthMaskTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a denoise strength mask texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isdenoisestrengthmasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsDenoiseStrengthMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsDenoiseStrengthMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: isDenoiseStrengthMaskTextureEnabled */


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReactiveMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isReactiveMaskTextureEnabled */


// A Boolean value that indicates whether a scaler you create from this descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsReactiveMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: isReactiveMaskTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsSpecularHitDistanceTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isSpecularHitDistanceTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isSpecularHitDistanceTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a specular hit distance texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/isspecularhitdistancetextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsSpecularHitDistanceTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsSpecularHitDistanceTextureEnabled:"), value)
}/* debug [instance_properties/setter]: isSpecularHitDistanceTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) IsTransparencyOverlayTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isTransparencyOverlayTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isTransparencyOverlayTextureEnabled */


// A Boolean value indicating whether the scaler evaluates a transparency overlay texture as part of its operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporaldenoisedscalerdescriptor/istransparencyoverlaytextureenabled
func (f_ FXTemporalDenoisedScalerDescriptor) SetIsTransparencyOverlayTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsTransparencyOverlayTextureEnabled:"), value)
}/* debug [instance_properties/setter]: isTransparencyOverlayTextureEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFXTemporalDenoisedScalerDescriptor */



