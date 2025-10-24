// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageGaussianBlur */


/* debug [class_header]: Header for MPSImageGaussianBlur */
// The class instance for the [ImageGaussianBlur] class.
var (
	ImageGaussianBlurClass     _ImageGaussianBlurClass
	ImageGaussianBlurClassOnce sync.Once
)

func getImageGaussianBlurClass() _ImageGaussianBlurClass {
	ImageGaussianBlurClassOnce.Do(func() {
		ImageGaussianBlurClass = _ImageGaussianBlurClass{objc.GetClass("MPSImageGaussianBlur")}
	})
	return ImageGaussianBlurClass
}

type _ImageGaussianBlurClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageGaussianBlur */
// An interface definition for the [ImageGaussianBlur] class.
type IImageGaussianBlur interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageGaussianBlur */
	// properties:
	Sigma() objectivec.IObject
	SetSigma(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageGaussianBlur */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageGaussianBlur */
// Alloc allocates a new instance without initialization.
func (ic _ImageGaussianBlurClass) Alloc() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageGaussianBlurClass) New() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageGaussianBlur) Init() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageGaussianBlur) Autorelease() ImageGaussianBlur {
	rv := objc.Send[ImageGaussianBlur](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageGaussianBlur creates a new ImageGaussianBlur instance.
func NewImageGaussianBlur() ImageGaussianBlur {
	return getImageGaussianBlurClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageGaussianBlur */
// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.


// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageGaussianBlur
type ImageGaussianBlur struct {
	UnaryImageKernel
}

// ImageGaussianBlurFrom constructs a [ImageGaussianBlur] from an unsafe.Pointer.
//
// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions.
func ImageGaussianBlurFrom(ptr unsafe.Pointer) ImageGaussianBlur {
	return ImageGaussianBlur{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageGaussianBlur */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/2866150-initwithcoder
func NewImageGaussianBlurWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) ImageGaussianBlur {
	instance := getImageGaussianBlurClass().Alloc()
	rv := objc.Send[ImageGaussianBlur](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageGaussianBlurWithCoderDevice */


// Initializes a Gaussian blur filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/1618813-initwithdevice
func NewImageGaussianBlurWithDeviceSigma(device unsafe.Pointer, sigma float32) ImageGaussianBlur {
	instance := getImageGaussianBlurClass().Alloc()
	rv := objc.Send[ImageGaussianBlur](instance.ID, objc.Sel("initWithDevice:sigma:"), device, sigma)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageGaussianBlurWithDeviceSigma */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageGaussianBlur */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageGaussianBlur */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageGaussianBlur */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageGaussianBlur */

// The sigma value with which the filter was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/1618850-sigma
func (i_ ImageGaussianBlur) Sigma() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("sigma"))
	return rv
}/* debug [instance_properties/getter]: sigma */


// The sigma value with which the filter was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/1618850-sigma
func (i_ ImageGaussianBlur) SetSigma(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSigma:"), value)
}/* debug [instance_properties/setter]: sigma */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageGaussianBlur */


