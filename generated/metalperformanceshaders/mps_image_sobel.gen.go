// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageSobel */


/* debug [class_header]: Header for MPSImageSobel */
// The class instance for the [ImageSobel] class.
var (
	ImageSobelClass     _ImageSobelClass
	ImageSobelClassOnce sync.Once
)

func getImageSobelClass() _ImageSobelClass {
	ImageSobelClassOnce.Do(func() {
		ImageSobelClass = _ImageSobelClass{objc.GetClass("MPSImageSobel")}
	})
	return ImageSobelClass
}

type _ImageSobelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageSobel */
// An interface definition for the [ImageSobel] class.
type IImageSobel interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageSobel */
	// properties:
	ColorTransform() objectivec.IObject
	SetColorTransform(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageSobel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageSobel */
// Alloc allocates a new instance without initialization.
func (ic _ImageSobelClass) Alloc() ImageSobel {
	rv := objc.Send[ImageSobel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSobelClass) New() ImageSobel {
	rv := objc.Send[ImageSobel](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSobel) Init() ImageSobel {
	rv := objc.Send[ImageSobel](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSobel) Autorelease() ImageSobel {
	rv := objc.Send[ImageSobel](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSobel creates a new ImageSobel instance.
func NewImageSobel() ImageSobel {
	return getImageSobelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageSobel */
// A filter that convolves an image with the Sobel operator.
//
// When the color model (e.g. RGB, two-channel, grayscale, etc.) of the source and destination textures match, the filter is applied to each color channel separately. If the destination is single-channel (i.e. monochrome) but the source is multi-channel, the pixel values are converted to grayscale before applying the Sobel operator by using the linear gray color transform vector shown in the code listing below.


// A filter that convolves an image with the Sobel operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageSobel
type ImageSobel struct {
	UnaryImageKernel
}

// ImageSobelFrom constructs a [ImageSobel] from an unsafe.Pointer.
//
// A filter that convolves an image with the Sobel operator.
func ImageSobelFrom(ptr unsafe.Pointer) ImageSobel {
	return ImageSobel{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageSobel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/2866152-initwithcoder
func NewImageSobelWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageSobelWithCoderDevice */


// Initializes a Sobel filter on a given device using the default color transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618843-initwithdevice
func NewImageSobelWithDevice(device unsafe.Pointer) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageSobelWithDevice */


// Initializes a Sobel filter on a given device using a specific color transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618899-initwithdevice
func NewImageSobelWithDeviceLinearGrayColorTransform(device unsafe.Pointer, transform objectivec.IObject) ImageSobel {
	instance := getImageSobelClass().Alloc()
	rv := objc.Send[ImageSobel](instance.ID, objc.Sel("initWithDevice:linearGrayColorTransform:"), device, transform)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageSobelWithDeviceLinearGrayColorTransform */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageSobel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageSobel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageSobel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageSobel */

// The color transform used to initialize the Sobel filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618777-colortransform
func (i_ ImageSobel) ColorTransform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("colorTransform"))
	return rv
}/* debug [instance_properties/getter]: colorTransform */


// The color transform used to initialize the Sobel filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618777-colortransform
func (i_ ImageSobel) SetColorTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setColorTransform:"), value)
}/* debug [instance_properties/setter]: colorTransform */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageSobel */


