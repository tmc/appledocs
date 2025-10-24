// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageDilate] class.
var (
	ImageDilateClass     _ImageDilateClass
	ImageDilateClassOnce sync.Once
)

func getImageDilateClass() _ImageDilateClass {
	ImageDilateClassOnce.Do(func() {
		ImageDilateClass = _ImageDilateClass{objc.GetClass("MPSImageDilate")}
	})
	return ImageDilateClass
}

type _ImageDilateClass struct {
	class objc.Class
}

// An interface definition for the [ImageDilate] class.
type IImageDilate interface {
	IUnaryImageKernel
	// properties:
	KernelHeight() uint /* primitive/slice/pointer. */
	KernelWidth() uint /* primitive/slice/pointer. */
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
	// methods:
}

// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
//
// An filter behaves like the filter, except Metal calculates the intensity at each position relative to a different value before determining which is the maximum pixel value, allowing for shaped, nonrectangular morphological probes. The code example below shows pseudocode for the calculation that returns each pixel value: A filter that contains all zeros is identical to an filter. Metal handles the center filter element as to avoid causing a general darkening of the image, and it handles the property as for this filter.


// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate
type ImageDilate struct {
	UnaryImageKernel
}

// ImageDilateFrom constructs a [ImageDilate] from an unsafe.Pointer.
//
// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function.
func ImageDilateFrom(ptr unsafe.Pointer) ImageDilate {
	return ImageDilate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageDilateClass) Alloc() ImageDilate {
	rv := objc.Send[ImageDilate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageDilateClass) New() ImageDilate {
	rv := objc.Send[ImageDilate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageDilate) Init() ImageDilate {
	rv := objc.Send[ImageDilate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageDilate) Autorelease() ImageDilate {
	rv := objc.Send[ImageDilate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageDilate creates a new ImageDilate instance.
func NewImageDilate() ImageDilate {
	return getImageDilateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(coder:device:)
func NewImageDilateWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImageDilate {
	instance := getImageDilateClass().Alloc()
	rv := objc.Send[ImageDilate](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes the kernel with a specified width, height, and weight values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/init(device:kernelWidth:kernelHeight:values:)
func NewImageDilateWithDeviceKernelWidthKernelHeightValues(device objectivec.IObject, kernelWidth uint /* primitive/slice/pointer. */, kernelHeight uint /* primitive/slice/pointer. */, values unsafe.Pointer) ImageDilate {
	instance := getImageDilateClass().Alloc()
	rv := objc.Send[ImageDilate](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), device, kernelWidth, kernelHeight, values)
	rv.Autorelease()
	return rv
}



// The height of the filter window. which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/kernelHeight
func (i_ ImageDilate) KernelHeight() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The width of the filter window which must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageDilate/kernelWidth
func (i_ ImageDilate) KernelWidth() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelWidth"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageDilate) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageDilate) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}


