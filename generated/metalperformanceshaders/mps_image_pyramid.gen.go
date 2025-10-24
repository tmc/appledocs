// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImagePyramid] class.
var (
	ImagePyramidClass     _ImagePyramidClass
	ImagePyramidClassOnce sync.Once
)

func getImagePyramidClass() _ImagePyramidClass {
	ImagePyramidClassOnce.Do(func() {
		ImagePyramidClass = _ImagePyramidClass{objc.GetClass("MPSImagePyramid")}
	})
	return ImagePyramidClass
}

type _ImagePyramidClass struct {
	class objc.Class
}

// An interface definition for the [ImagePyramid] class.
type IImagePyramid interface {
	IUnaryImageKernel
	// properties:
	KernelHeight() uint /* primitive/slice/pointer. */
	KernelWidth() uint /* primitive/slice/pointer. */
	// methods:
}

// A base class for creating different kinds of pyramid images.


// A base class for creating different kinds of pyramid images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid
type ImagePyramid struct {
	UnaryImageKernel
}

// ImagePyramidFrom constructs a [ImagePyramid] from an unsafe.Pointer.
//
// A base class for creating different kinds of pyramid images.
func ImagePyramidFrom(ptr unsafe.Pointer) ImagePyramid {
	return ImagePyramid{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImagePyramidClass) Alloc() ImagePyramid {
	rv := objc.Send[ImagePyramid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImagePyramidClass) New() ImagePyramid {
	rv := objc.Send[ImagePyramid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImagePyramid) Init() ImagePyramid {
	rv := objc.Send[ImagePyramid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImagePyramid) Autorelease() ImagePyramid {
	rv := objc.Send[ImagePyramid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImagePyramid creates a new ImagePyramid instance.
func NewImagePyramid() ImagePyramid {
	return getImagePyramidClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(coder:device:)
func NewImagePyramidWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a downwards 5-tap image pyramid with the default filter kernel and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:)
func NewImagePyramidWithDevice(device objectivec.IObject) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Initialize a downwards 5-tap image pyramid with a central weight parameter and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:centerWeight:)
func NewImagePyramidWithDeviceCenterWeight(device objectivec.IObject, centerWeight float32 /* primitive/slice/pointer. */) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	rv.Autorelease()
	return rv
}


// Initialize a downwards n-tap image pyramid with a custom filter kernel and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/init(device:kernelWidth:kernelHeight:weights:)
func NewImagePyramidWithDeviceKernelWidthKernelHeightWeights(device objectivec.IObject, kernelWidth uint /* primitive/slice/pointer. */, kernelHeight uint /* primitive/slice/pointer. */, kernelWeights unsafe.Pointer) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	rv.Autorelease()
	return rv
}



// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/kernelHeight
func (i_ ImagePyramid) KernelHeight() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImagePyramid/kernelWidth
func (i_ ImagePyramid) KernelWidth() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](i_.ID, objc.Sel("kernelWidth"))
	return rv
}


