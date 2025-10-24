// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ImageConvolution] class.
type IImageConvolution interface {
	IUnaryImageKernel
	// properties:
	Bias() float32 /* primitive/slice/pointer. */
	SetBias(value float32 /* primitive/slice/pointer. */)
	KernelHeight() uint /* primitive/slice/pointer. */
	KernelWidth() uint /* primitive/slice/pointer. */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageConvolutionClass) Alloc() ImageConvolution {
	rv := objc.Send[ImageConvolution](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/bias
func (i_ ImageConvolution) Bias() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("bias"))
	return rv
}


// The value added to a convolved pixel before it is converted back to its intended storage format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/bias
func (i_ ImageConvolution) SetBias(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBias:"), value)
}


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/kernelHeight
func (i_ ImageConvolution) KernelHeight() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageConvolution/kernelWidth
func (i_ ImageConvolution) KernelWidth() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelWidth"))
	return rv
}



