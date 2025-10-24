// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageEuclideanDistanceTransform] class.
var (
	ImageEuclideanDistanceTransformClass     _ImageEuclideanDistanceTransformClass
	ImageEuclideanDistanceTransformClassOnce sync.Once
)

func getImageEuclideanDistanceTransformClass() _ImageEuclideanDistanceTransformClass {
	ImageEuclideanDistanceTransformClassOnce.Do(func() {
		ImageEuclideanDistanceTransformClass = _ImageEuclideanDistanceTransformClass{objc.GetClass("MPSImageEuclideanDistanceTransform")}
	})
	return ImageEuclideanDistanceTransformClass
}

type _ImageEuclideanDistanceTransformClass struct {
	class objc.Class
}

// An interface definition for the [ImageEuclideanDistanceTransform] class.
type IImageEuclideanDistanceTransform interface {
	IUnaryImageKernel
	// properties:
	SearchLimitRadius() float32 /* primitive/slice/pointer. */
	SetSearchLimitRadius(value float32 /* primitive/slice/pointer. */)
	// methods:
}

// A filter that performs a Euclidean distance transform on an image.


// A filter that performs a Euclidean distance transform on an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform
type ImageEuclideanDistanceTransform struct {
	UnaryImageKernel
}

// ImageEuclideanDistanceTransformFrom constructs a [ImageEuclideanDistanceTransform] from an unsafe.Pointer.
//
// A filter that performs a Euclidean distance transform on an image.
func ImageEuclideanDistanceTransformFrom(ptr unsafe.Pointer) ImageEuclideanDistanceTransform {
	return ImageEuclideanDistanceTransform{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageEuclideanDistanceTransformClass) Alloc() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageEuclideanDistanceTransformClass) New() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageEuclideanDistanceTransform) Init() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageEuclideanDistanceTransform) Autorelease() ImageEuclideanDistanceTransform {
	rv := objc.Send[ImageEuclideanDistanceTransform](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageEuclideanDistanceTransform creates a new ImageEuclideanDistanceTransform instance.
func NewImageEuclideanDistanceTransform() ImageEuclideanDistanceTransform {
	return getImageEuclideanDistanceTransformClass().New()
}



// Creates a Euclidean distance transform that uses a specified decoder for your data and runs on a specified device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/init(coder:device:)
func NewImageEuclideanDistanceTransformWithCoderDevice(aDecoder objc.IObject /* cross-framework: Coder */, device objectivec.IObject) ImageEuclideanDistanceTransform {
	instance := getImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[ImageEuclideanDistanceTransform](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Creates a Euclidean distance transform that runs on a specified device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/init(device:)
func NewImageEuclideanDistanceTransformWithDevice(device objectivec.IObject) ImageEuclideanDistanceTransform {
	instance := getImageEuclideanDistanceTransformClass().Alloc()
	rv := objc.Send[ImageEuclideanDistanceTransform](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/searchLimitRadius
func (i_ ImageEuclideanDistanceTransform) SearchLimitRadius() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](i_.ID, objc.Sel("searchLimitRadius"))
	return rv
}


// Limits the search in an image from a pixel to the closest nonzero pixel within a specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageEuclideanDistanceTransform/searchLimitRadius
func (i_ ImageEuclideanDistanceTransform) SetSearchLimitRadius(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchLimitRadius:"), value)
}


