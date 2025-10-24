// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSImageConvolution */


/* debug [class_header]: Header for MPSImageConvolution */
// The class instance for the [ImageConvolution] class.
var (
	ImageConvolutionClass     _ImageConvolutionClass
	ImageConvolutionClassOnce sync.Once
)

func getImageConvolutionClass() _ImageConvolutionClass {
	ImageConvolutionClassOnce.Do(func() {
		ImageConvolutionClass = _ImageConvolutionClass{objc.GetClass("MPSImageConvolution")}
	})
	return ImageConvolutionClass
}

type _ImageConvolutionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageConvolution */
// An interface definition for the [ImageConvolution] class.
type IImageConvolution interface {
	IUnaryImageKernel
	
/* debug [class_interface_properties]: Properties for ImageConvolution */
	// properties:
	Bias() objectivec.IObject
	SetBias(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageConvolution */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageConvolution */
// Alloc allocates a new instance without initialization.
func (ic _ImageConvolutionClass) Alloc() ImageConvolution {
	rv := objc.Send[ImageConvolution](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageConvolutionClass) New() ImageConvolution {
	rv := objc.Send[ImageConvolution](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageConvolution) Init() ImageConvolution {
	rv := objc.Send[ImageConvolution](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageConvolution) Autorelease() ImageConvolution {
	rv := objc.Send[ImageConvolution](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageConvolution creates a new ImageConvolution instance.
func NewImageConvolution() ImageConvolution {
	return getImageConvolutionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageConvolution */
// A filter that convolves an image with a given kernel of odd width and height.
//
// Filter width and height can be either 3, 5, 7 or 9. If there are multiple channels in the source image, each channel is processed independently. A convolution filter may perform better when done in two passes. . A convolution filter is separable if the ratio of filter values between all rows is constant over the whole row. For example, this edge detection filter: Can instead be separated into the product of two vectors, like so: And consequently can be done as two, one-dimensional convolution passes back to back on the same image. In this way, the number of multiplies (ignoring the fact that we could skip zeros here) is reduced from to . There are similar savings for addition. For large filters, the savings can be profound.


// A filter that convolves an image with a given kernel of odd width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution
type ImageConvolution struct {
	UnaryImageKernel
}

// ImageConvolutionFrom constructs a [ImageConvolution] from an unsafe.Pointer.
//
// A filter that convolves an image with a given kernel of odd width and height.
func ImageConvolutionFrom(ptr unsafe.Pointer) ImageConvolution {
	return ImageConvolution{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageConvolution */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/2866148-initwithcoder
func NewImageConvolutionWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageConvolution {
	instance := getImageConvolutionClass().Alloc()
	rv := objc.Send[ImageConvolution](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageConvolutionWithCoderDevice */


// Initializes a convolution filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618902-initwithdevice
func NewImageConvolutionWithDeviceKernelWidthKernelHeightWeights(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, kernelWeights objectivec.IObject) ImageConvolution {
	instance := getImageConvolutionClass().Alloc()
	rv := objc.Send[ImageConvolution](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewImageConvolutionWithDeviceKernelWidthKernelHeightWeights */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageConvolution */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageConvolution */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageConvolution */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageConvolution */

// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618841-bias
func (i_ ImageConvolution) Bias() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("bias"))
	return rv
}/* debug [instance_properties/getter]: bias */


// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618841-bias
func (i_ ImageConvolution) SetBias(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBias:"), value)
}/* debug [instance_properties/setter]: bias */


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618842-kernelheight
func (i_ ImageConvolution) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}/* debug [instance_properties/getter]: kernelHeight */


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618842-kernelheight
func (i_ ImageConvolution) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}/* debug [instance_properties/setter]: kernelHeight */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618868-kernelwidth
func (i_ ImageConvolution) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}/* debug [instance_properties/getter]: kernelWidth */


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618868-kernelwidth
func (i_ ImageConvolution) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}/* debug [instance_properties/setter]: kernelWidth */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSImageConvolution */


