// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageCanny */


/* debug [class_header]: Header for MPSImageCanny */
// The class instance for the [ImageCanny] class.
var (
	ImageCannyClass     _ImageCannyClass
	ImageCannyClassOnce sync.Once
)

func getImageCannyClass() _ImageCannyClass {
	ImageCannyClassOnce.Do(func() {
		ImageCannyClass = _ImageCannyClass{objc.GetClass("MPSImageCanny")}
	})
	return ImageCannyClass
}

type _ImageCannyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageCanny */
// An interface definition for the [ImageCanny] class.
type IImageCanny interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageCanny */
	// properties:
	ColorTransform() objectivec.IObject
	SetColorTransform(value objectivec.IObject)
	HighThreshold() objectivec.IObject
	SetHighThreshold(value objectivec.IObject)
	LowThreshold() objectivec.IObject
	SetLowThreshold(value objectivec.IObject)
	Sigma() objectivec.IObject
	SetSigma(value objectivec.IObject)
	UseFastMode() objectivec.IObject
	SetUseFastMode(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageCanny */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageCanny */
// Alloc allocates a new instance without initialization.
func (ic _ImageCannyClass) Alloc() ImageCanny {
	rv := objc.Send[ImageCanny](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageCannyClass) New() ImageCanny {
	rv := objc.Send[ImageCanny](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageCanny) Init() ImageCanny {
	rv := objc.Send[ImageCanny](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageCanny) Autorelease() ImageCanny {
	rv := objc.Send[ImageCanny](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageCanny creates a new ImageCanny instance.
func NewImageCanny() ImageCanny {
	return getImageCannyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageCanny */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageCanny
type ImageCanny struct {
	UnaryImageKernel
}

// ImageCannyFrom constructs a [ImageCanny] from an unsafe.Pointer.
func ImageCannyFrom(ptr unsafe.Pointer) ImageCanny {
	return ImageCanny{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageCanny */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547971-initwithcoder
func NewImageCannyWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageCannyWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547972-initwithdevice
func NewImageCannyWithDevice(device unsafe.Pointer) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageCannyWithDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547973-initwithdevice
func NewImageCannyWithDeviceLinearToGrayScaleTransformSigma(device unsafe.Pointer, transform objectivec.IObject, sigma float32) ImageCanny {
	instance := getImageCannyClass().Alloc()
	rv := objc.Send[ImageCanny](instance.ID, objc.Sel("initWithDevice:linearToGrayScaleTransform:sigma:"), device, transform, sigma)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageCannyWithDeviceLinearToGrayScaleTransformSigma */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageCanny */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageCanny */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageCanny */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageCanny */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547969-colortransform
func (i_ ImageCanny) ColorTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("colorTransform"))
	return rv
}/* debug [instance_properties/getter]: colorTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547969-colortransform
func (i_ ImageCanny) SetColorTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorTransform:"), value)
}/* debug [instance_properties/setter]: colorTransform */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold
func (i_ ImageCanny) HighThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("highThreshold"))
	return rv
}/* debug [instance_properties/getter]: highThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547970-highthreshold
func (i_ ImageCanny) SetHighThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHighThreshold:"), value)
}/* debug [instance_properties/setter]: highThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold
func (i_ ImageCanny) LowThreshold() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("lowThreshold"))
	return rv
}/* debug [instance_properties/getter]: lowThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547974-lowthreshold
func (i_ ImageCanny) SetLowThreshold(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLowThreshold:"), value)
}/* debug [instance_properties/setter]: lowThreshold */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547975-sigma
func (i_ ImageCanny) Sigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("sigma"))
	return rv
}/* debug [instance_properties/getter]: sigma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547975-sigma
func (i_ ImageCanny) SetSigma(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSigma:"), value)
}/* debug [instance_properties/setter]: sigma */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode
func (i_ ImageCanny) UseFastMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("useFastMode"))
	return rv
}/* debug [instance_properties/getter]: useFastMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecanny/3547976-usefastmode
func (i_ ImageCanny) SetUseFastMode(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setUseFastMode:"), value)
}/* debug [instance_properties/setter]: useFastMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageCanny */


