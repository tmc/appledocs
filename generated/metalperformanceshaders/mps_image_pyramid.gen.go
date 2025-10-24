// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImagePyramidClass) Alloc() ImagePyramid {
	rv := objc.Send[ImagePyramid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/2866151-initwithcoder
func NewImagePyramidWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes a downwards 5-tap image pyramid with the default filter kernel and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice
func NewImagePyramidWithDevice(device unsafe.Pointer) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Initialize a downwards 5-tap image pyramid with a central weight parameter and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648889-initwithdevice
func NewImagePyramidWithDeviceCenterWeight(device unsafe.Pointer, centerWeight float32) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:centerWeight:"), device, centerWeight)
	rv.Autorelease()
	return rv
}


// Initialize a downwards n-tap image pyramid with a custom filter kernel and device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648821-initwithdevice
func NewImagePyramidWithDeviceKernelWidthKernelHeightWeights(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, kernelWeights objectivec.IObject) ImagePyramid {
	instance := getImagePyramidClass().Alloc()
	rv := objc.Send[ImagePyramid](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), device, kernelWidth, kernelHeight, kernelWeights)
	rv.Autorelease()
	return rv
}






















// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648842-kernelwidth
func (i_ ImagePyramid) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648842-kernelwidth
func (i_ ImagePyramid) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648863-kernelheight
func (i_ ImagePyramid) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648863-kernelheight
func (i_ ImagePyramid) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}







