// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFXTemporalScalerDescriptor */


/* debug [class_header]: Header for MTLFXTemporalScalerDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FXTemporalScalerDescriptor */
// An interface definition for the [FXTemporalScalerDescriptor] class.
type IFXTemporalScalerDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FXTemporalScalerDescriptor */
	// properties:
	ColorTextureFormat() PixelFormat /* not a class type */
	SetColorTextureFormat(value PixelFormat /* not a class type */)
	DepthTextureFormat() PixelFormat /* not a class type */
	SetDepthTextureFormat(value PixelFormat /* not a class type */)
	InputContentMaxScale() float32
	SetInputContentMaxScale(value float32)
	InputContentMinScale() float32
	SetInputContentMinScale(value float32)
	InputHeight() uint
	SetInputHeight(value uint)
	InputWidth() uint
	SetInputWidth(value uint)
	AutoExposureEnabled() bool
	SetAutoExposureEnabled(value bool)
	InputContentPropertiesEnabled() bool
	SetInputContentPropertiesEnabled(value bool)
	ReactiveMaskTextureEnabled() bool
	SetReactiveMaskTextureEnabled(value bool)
	MotionTextureFormat() PixelFormat /* not a class type */
	SetMotionTextureFormat(value PixelFormat /* not a class type */)
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
	IsAutoExposureEnabled() bool
	SetIsAutoExposureEnabled(value bool)
	IsInputContentPropertiesEnabled() bool
	SetIsInputContentPropertiesEnabled(value bool)
	IsReactiveMaskTextureEnabled() bool
	SetIsReactiveMaskTextureEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FXTemporalScalerDescriptor */
	// methods:
	NewTemporalScalerWithDevice(device unsafe.Pointer) unsafe.Pointer
	NewTemporalScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FXTemporalScalerDescriptor */
// Alloc allocates a new instance without initialization.
func (fc _FXTemporalScalerDescriptorClass) Alloc() FXTemporalScalerDescriptor {
	rv := objc.Send[FXTemporalScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FXTemporalScalerDescriptor */
// A set of properties that configure a temporal scaling effect, and a factory method that creates the effect.


// A set of properties that configure a temporal scaling effect, and a factory method that creates the effect.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FXTemporalScalerDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FXTemporalScalerDescriptor */

// Returns the largest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportedInputContentMaxScale(device:)
func (fc _FXTemporalScalerDescriptorClass) SupportedInputContentMaxScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMaxScaleForDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedInputContentMaxScaleForDevice) */


// Returns the smallest temporal scaling factor the device supports as a floating-point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportedInputContentMinScale(device:)
func (fc _FXTemporalScalerDescriptorClass) SupportedInputContentMinScaleForDevice(device unsafe.Pointer) float32 {
	rv := objc.Send[float32](objc.ID(fc.class), objc.Sel("supportedInputContentMinScaleForDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedInputContentMinScaleForDevice) */


// Returns a Boolean value that indicates whether the temporal scaler works with a GPU.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportsDevice(_:)
func (fc _FXTemporalScalerDescriptorClass) SupportsDevice(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsDevice) */


// Queries whether a Metal device supports temporal scaling compatible with Metal 4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXTemporalScalerDescriptorClass) SupportsMetal4FX(device unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsMetal4FX) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FXTemporalScalerDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FXTemporalScalerDescriptor */

// Creates a temporal scaler instance from this descriptor’s current property values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/makeTemporalScaler(device:)
func (f_ FXTemporalScalerDescriptor) NewTemporalScalerWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalScalerWithDevice:"), device)
	return rv
}/* debug [instance_methods/method]: NewTemporalScalerWithDevice */


// Creates a temporal scaler instance for a Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/makeTemporalScaler(device:compiler:)
func (f_ FXTemporalScalerDescriptor) NewTemporalScalerWithDeviceCompiler(device unsafe.Pointer, compiler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("newTemporalScalerWithDevice:compiler:"), device, compiler)
	return rv
}/* debug [instance_methods/method]: NewTemporalScalerWithDeviceCompiler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FXTemporalScalerDescriptor */

// The pixel format of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/colorTextureFormat
func (f_ FXTemporalScalerDescriptor) ColorTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: colorTextureFormat */


// The pixel format of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/colorTextureFormat
func (f_ FXTemporalScalerDescriptor) SetColorTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}/* debug [instance_properties/setter]: colorTextureFormat */


// The pixel format of the input depth texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/depthTextureFormat
func (f_ FXTemporalScalerDescriptor) DepthTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("depthTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: depthTextureFormat */


// The pixel format of the input depth texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/depthTextureFormat
func (f_ FXTemporalScalerDescriptor) SetDepthTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDepthTextureFormat:"), value)
}/* debug [instance_properties/setter]: depthTextureFormat */


// The largest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMaxScale
func (f_ FXTemporalScalerDescriptor) InputContentMaxScale() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("inputContentMaxScale"))
	return rv
}/* debug [instance_properties/getter]: inputContentMaxScale */


// The largest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMaxScale
func (f_ FXTemporalScalerDescriptor) SetInputContentMaxScale(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentMaxScale:"), value)
}/* debug [instance_properties/setter]: inputContentMaxScale */


// The smallest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMinScale
func (f_ FXTemporalScalerDescriptor) InputContentMinScale() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("inputContentMinScale"))
	return rv
}/* debug [instance_properties/getter]: inputContentMinScale */


// The smallest scale factor the temporal scaler you create with this descriptor can use to generate output textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputContentMinScale
func (f_ FXTemporalScalerDescriptor) SetInputContentMinScale(value float32) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentMinScale:"), value)
}/* debug [instance_properties/setter]: inputContentMinScale */


// The height of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputHeight
func (f_ FXTemporalScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}/* debug [instance_properties/getter]: inputHeight */


// The height of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputHeight
func (f_ FXTemporalScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}/* debug [instance_properties/setter]: inputHeight */


// The width of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputWidth
func (f_ FXTemporalScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}/* debug [instance_properties/getter]: inputWidth */


// The width of the input color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/inputWidth
func (f_ FXTemporalScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}/* debug [instance_properties/setter]: inputWidth */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalScalerDescriptor) AutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("autoExposureEnabled"))
	return rv
}/* debug [instance_properties/getter]: autoExposureEnabled */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isAutoExposureEnabled
func (f_ FXTemporalScalerDescriptor) SetAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAutoExposureEnabled:"), value)
}/* debug [instance_properties/setter]: autoExposureEnabled */


// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isInputContentPropertiesEnabled
func (f_ FXTemporalScalerDescriptor) InputContentPropertiesEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("inputContentPropertiesEnabled"))
	return rv
}/* debug [instance_properties/getter]: inputContentPropertiesEnabled */


// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isInputContentPropertiesEnabled
func (f_ FXTemporalScalerDescriptor) SetInputContentPropertiesEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputContentPropertiesEnabled:"), value)
}/* debug [instance_properties/setter]: inputContentPropertiesEnabled */


// A Boolean value that indicates whether a temporal scaler you create with the descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalScalerDescriptor) ReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("reactiveMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: reactiveMaskTextureEnabled */


// A Boolean value that indicates whether a temporal scaler you create with the descriptor applies a reactive mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/isReactiveMaskTextureEnabled
func (f_ FXTemporalScalerDescriptor) SetReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: reactiveMaskTextureEnabled */


// The pixel format of the input motion texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/motionTextureFormat
func (f_ FXTemporalScalerDescriptor) MotionTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("motionTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: motionTextureFormat */


// The pixel format of the input motion texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/motionTextureFormat
func (f_ FXTemporalScalerDescriptor) SetMotionTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMotionTextureFormat:"), value)
}/* debug [instance_properties/setter]: motionTextureFormat */


// The height of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputHeight
func (f_ FXTemporalScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}/* debug [instance_properties/getter]: outputHeight */


// The height of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputHeight
func (f_ FXTemporalScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}/* debug [instance_properties/setter]: outputHeight */


// The pixel format of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputTextureFormat
func (f_ FXTemporalScalerDescriptor) OutputTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: outputTextureFormat */


// The pixel format of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputTextureFormat
func (f_ FXTemporalScalerDescriptor) SetOutputTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}/* debug [instance_properties/setter]: outputTextureFormat */


// The width of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputWidth
func (f_ FXTemporalScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}/* debug [instance_properties/getter]: outputWidth */


// The width of the output color texture for the temporal scaler you create with this descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/outputWidth
func (f_ FXTemporalScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}/* debug [instance_properties/setter]: outputWidth */


// The pixel format of the reactive mask input texture for a temporal scaler you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalScalerDescriptor) ReactiveMaskTextureFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](f_.ID, objc.Sel("reactiveMaskTextureFormat"))
	return rv
}/* debug [instance_properties/getter]: reactiveMaskTextureFormat */


// The pixel format of the reactive mask input texture for a temporal scaler you create with the descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/reactiveMaskTextureFormat
func (f_ FXTemporalScalerDescriptor) SetReactiveMaskTextureFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReactiveMaskTextureFormat:"), value)
}/* debug [instance_properties/setter]: reactiveMaskTextureFormat */


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalScalerDescriptor) RequiresSynchronousInitialization() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresSynchronousInitialization"))
	return rv
}/* debug [instance_properties/getter]: requiresSynchronousInitialization */


// A Boolean value that indicates whether MetalFX compiles a temporal scaling effect’s underlying upscaler as it creates the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXTemporalScalerDescriptor/requiresSynchronousInitialization
func (f_ FXTemporalScalerDescriptor) SetRequiresSynchronousInitialization(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresSynchronousInitialization:"), value)
}/* debug [instance_properties/setter]: requiresSynchronousInitialization */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalScalerDescriptor) IsAutoExposureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAutoExposureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAutoExposureEnabled */


// A Boolean value that indicates whether MetalFX calculates the exposure for each frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isautoexposureenabled
func (f_ FXTemporalScalerDescriptor) SetIsAutoExposureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsAutoExposureEnabled:"), value)
}/* debug [instance_properties/setter]: isAutoExposureEnabled */


// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isinputcontentpropertiesenabled
func (f_ FXTemporalScalerDescriptor) IsInputContentPropertiesEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isInputContentPropertiesEnabled"))
	return rv
}/* debug [instance_properties/getter]: isInputContentPropertiesEnabled */


// A Boolean value that indicates whether the temporal scaler you create with this descriptor uses dynamic resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isinputcontentpropertiesenabled
func (f_ FXTemporalScalerDescriptor) SetIsInputContentPropertiesEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsInputContentPropertiesEnabled:"), value)
}/* debug [instance_properties/setter]: isInputContentPropertiesEnabled */


// A Boolean value that indicates whether a temporal scaler you create with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalScalerDescriptor) IsReactiveMaskTextureEnabled() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isReactiveMaskTextureEnabled"))
	return rv
}/* debug [instance_properties/getter]: isReactiveMaskTextureEnabled */


// A Boolean value that indicates whether a temporal scaler you create with the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalfx/mtlfxtemporalscalerdescriptor/isreactivemasktextureenabled
func (f_ FXTemporalScalerDescriptor) SetIsReactiveMaskTextureEnabled(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsReactiveMaskTextureEnabled:"), value)
}/* debug [instance_properties/setter]: isReactiveMaskTextureEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFXTemporalScalerDescriptor */





