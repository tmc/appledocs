// Code generated from Apple documentation for MetalFX. DO NOT EDIT.

package metalfx

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FXSpatialScalerDescriptor] class.
type IFXSpatialScalerDescriptor interface {
	objectivec.IObject
	NewSpatialScalerWithDevice(device objectivec.IObject) objc.ID
	NewSpatialScalerWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID
}

// A set of properties that configure a spatial scaling effect, and a factory method that creates the effect.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FXSpatialScalerDescriptorClass) Alloc() FXSpatialScalerDescriptor {
	rv := objc.Send[FXSpatialScalerDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a Boolean value that indicates whether the spatial scaler works with a GPU.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/supportsDevice(_:)
func (fc _FXSpatialScalerDescriptorClass) SupportsDevice(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsDevice:"), device)
	return rv
}

// Queries whether a Metal device supports spatial scaling compatible with Metal 4.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/supportsMetal4FX(_:)
func (fc _FXSpatialScalerDescriptorClass) SupportsMetal4FX(device objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(fc.class), objc.Sel("supportsMetal4FX:"), device)
	return rv
}

// Creates a spatial scaler instance from this descriptor’s current property values.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/makeSpatialScaler(device:)
func (f_ FXSpatialScalerDescriptor) NewSpatialScalerWithDevice(device objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newSpatialScalerWithDevice:"), device)
	return rv
}

// Creates a spatial scaler instance for a Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/makeSpatialScaler(device:compiler:)
func (f_ FXSpatialScalerDescriptor) NewSpatialScalerWithDeviceCompiler(device objectivec.IObject, compiler objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("newSpatialScalerWithDevice:compiler:"), device, compiler)
	return rv
}

// The color space of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorProcessingMode
func (f_ FXSpatialScalerDescriptor) ColorProcessingMode() FXSpatialScalerColorProcessingMode {
	rv := objc.Send[FXSpatialScalerColorProcessingMode](f_.ID, objc.Sel("colorProcessingMode"))
	return rv
}


// SetColorProcessingMode sets the value of the colorProcessingMode property.
// The color space of the input color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorProcessingMode
func (f_ FXSpatialScalerDescriptor) SetColorProcessingMode(value FXSpatialScalerColorProcessingMode) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorProcessingMode:"), value)
}

// The pixel format of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorTextureFormat
func (f_ FXSpatialScalerDescriptor) ColorTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("colorTextureFormat"))
	return rv
}


// SetColorTextureFormat sets the value of the colorTextureFormat property.
// The pixel format of the input color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/colorTextureFormat
func (f_ FXSpatialScalerDescriptor) SetColorTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setColorTextureFormat:"), value)
}

// The height of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputHeight
func (f_ FXSpatialScalerDescriptor) InputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputHeight"))
	return rv
}


// SetInputHeight sets the value of the inputHeight property.
// The height of the input color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputHeight
func (f_ FXSpatialScalerDescriptor) SetInputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputHeight:"), value)
}

// The width of the input color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputWidth
func (f_ FXSpatialScalerDescriptor) InputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("inputWidth"))
	return rv
}


// SetInputWidth sets the value of the inputWidth property.
// The width of the input color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/inputWidth
func (f_ FXSpatialScalerDescriptor) SetInputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setInputWidth:"), value)
}

// The height of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputHeight
func (f_ FXSpatialScalerDescriptor) OutputHeight() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputHeight"))
	return rv
}


// SetOutputHeight sets the value of the outputHeight property.
// The height of the output color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputHeight
func (f_ FXSpatialScalerDescriptor) SetOutputHeight(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputHeight:"), value)
}

// The pixel format of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputTextureFormat
func (f_ FXSpatialScalerDescriptor) OutputTextureFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("outputTextureFormat"))
	return rv
}


// SetOutputTextureFormat sets the value of the outputTextureFormat property.
// The pixel format of the output color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputTextureFormat
func (f_ FXSpatialScalerDescriptor) SetOutputTextureFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputTextureFormat:"), value)
}

// The width of the output color texture for the spatial scaler you create with this descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputWidth
func (f_ FXSpatialScalerDescriptor) OutputWidth() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("outputWidth"))
	return rv
}


// SetOutputWidth sets the value of the outputWidth property.
// The width of the output color texture for the spatial scaler you create with this descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalFX/MTLFXSpatialScalerDescriptor/outputWidth
func (f_ FXSpatialScalerDescriptor) SetOutputWidth(value uint) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOutputWidth:"), value)
}



