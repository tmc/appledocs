// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageThresholdTruncate] class.
var (
	ImageThresholdTruncateClass     _ImageThresholdTruncateClass
	ImageThresholdTruncateClassOnce sync.Once
)

func getImageThresholdTruncateClass() _ImageThresholdTruncateClass {
	ImageThresholdTruncateClassOnce.Do(func() {
		ImageThresholdTruncateClass = _ImageThresholdTruncateClass{objc.GetClass("MPSImageThresholdTruncate")}
	})
	return ImageThresholdTruncateClass
}

type _ImageThresholdTruncateClass struct {
	class objc.Class
}





// An interface definition for the [ImageThresholdTruncate] class.
type IImageThresholdTruncate interface {
	IUnaryImageKernel
	

	// properties:
	Transform() objectivec.IObject
	SetTransform(value objectivec.IObject)
	ThresholdValue() objectivec.IObject
	SetThresholdValue(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageThresholdTruncateClass) Alloc() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageThresholdTruncateClass) New() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageThresholdTruncate) Init() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageThresholdTruncate) Autorelease() ImageThresholdTruncate {
	rv := objc.Send[ImageThresholdTruncate](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageThresholdTruncate creates a new ImageThresholdTruncate instance.
func NewImageThresholdTruncate() ImageThresholdTruncate {
	return getImageThresholdTruncateClass().New()
}





// A filter that clamps the return value to an upper specified value.
//
// An filter converts a single channel image to a binary image. If the input image is not a single channel image, the function first converts the input image into a single channel luminance image using the linear gray color transform, and then it applies the threshold. The following listing shows the threshold truncate function. Listing 1. Threshold truncate function


// A filter that clamps the return value to an upper specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageThresholdTruncate
type ImageThresholdTruncate struct {
	UnaryImageKernel
}

// ImageThresholdTruncateFrom constructs a [ImageThresholdTruncate] from an unsafe.Pointer.
//
// A filter that clamps the return value to an upper specified value.
func ImageThresholdTruncateFrom(ptr unsafe.Pointer) ImageThresholdTruncate {
	return ImageThresholdTruncate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/2865664-initwithcoder
func NewImageThresholdTruncateWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageThresholdTruncate {
	instance := getImageThresholdTruncateClass().Alloc()
	rv := objc.Send[ImageThresholdTruncate](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618818-initwithdevice
func NewImageThresholdTruncateWithDeviceThresholdValueLinearGrayColorTransform(device unsafe.Pointer, thresholdValue float32, transform objectivec.IObject) ImageThresholdTruncate {
	instance := getImageThresholdTruncateClass().Alloc()
	rv := objc.Send[ImageThresholdTruncate](instance.ID, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), device, thresholdValue, transform)
	rv.Autorelease()
	return rv
}






















// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618787-transform
func (i_ ImageThresholdTruncate) Transform() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("transform"))
	return rv
}


// The color transform used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618787-transform
func (i_ ImageThresholdTruncate) SetTransform(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransform:"), value)
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618882-thresholdvalue
func (i_ ImageThresholdTruncate) ThresholdValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("thresholdValue"))
	return rv
}


// The threshold value used to initialize the threshold filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtruncate/1618882-thresholdvalue
func (i_ ImageThresholdTruncate) SetThresholdValue(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setThresholdValue:"), value)
}







