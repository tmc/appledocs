// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ImageAreaMax] class.
var (
	ImageAreaMaxClass     _ImageAreaMaxClass
	ImageAreaMaxClassOnce sync.Once
)

func getImageAreaMaxClass() _ImageAreaMaxClass {
	ImageAreaMaxClassOnce.Do(func() {
		ImageAreaMaxClass = _ImageAreaMaxClass{objc.GetClass("MPSImageAreaMax")}
	})
	return ImageAreaMaxClass
}

type _ImageAreaMaxClass struct {
	class objc.Class
}





// An interface definition for the [ImageAreaMax] class.
type IImageAreaMax interface {
	IUnaryImageKernel
	

	// properties:
	KernelHeight() objectivec.IObject
	SetKernelHeight(value objectivec.IObject)
	KernelWidth() objectivec.IObject
	SetKernelWidth(value objectivec.IObject)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ic _ImageAreaMaxClass) Alloc() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageAreaMaxClass) New() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageAreaMax) Init() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageAreaMax) Autorelease() ImageAreaMax {
	rv := objc.Send[ImageAreaMax](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageAreaMax creates a new ImageAreaMax instance.
func NewImageAreaMax() ImageAreaMax {
	return getImageAreaMaxClass().New()
}





// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
//
// If there are multiple channels in the source image, each channel is processed independently. The property value is assumed to always be for this filter.


// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAreaMax
type ImageAreaMax struct {
	UnaryImageKernel
}

// ImageAreaMaxFrom constructs a [ImageAreaMax] from an unsafe.Pointer.
//
// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image.
func ImageAreaMaxFrom(ptr unsafe.Pointer) ImageAreaMax {
	return ImageAreaMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/2866327-initwithcoder
func NewImageAreaMaxWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) ImageAreaMax {
	instance := getImageAreaMaxClass().Alloc()
	rv := objc.Send[ImageAreaMax](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes the kernel with a specified width and height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618281-initwithdevice
func NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device unsafe.Pointer, kernelWidth uint, kernelHeight uint) ImageAreaMax {
	instance := getImageAreaMaxClass().Alloc()
	rv := objc.Send[ImageAreaMax](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), device, kernelWidth, kernelHeight)
	rv.Autorelease()
	return rv
}






















// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618277-kernelheight
func (i_ ImageAreaMax) KernelHeight() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelHeight"))
	return rv
}


// The height of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618277-kernelheight
func (i_ ImageAreaMax) SetKernelHeight(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelHeight:"), value)
}


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618282-kernelwidth
func (i_ ImageAreaMax) KernelWidth() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("kernelWidth"))
	return rv
}


// The width of the filter window. Must be an odd number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618282-kernelwidth
func (i_ ImageAreaMax) SetKernelWidth(value objectivec.IObject) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setKernelWidth:"), value)
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMax) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](i_.ID, objc.Sel("edgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/edgemode
func (i_ ImageAreaMax) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEdgeMode:"), value)
}







