// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingAverage] class.
var (
	CNNPoolingAverageClass     _CNNPoolingAverageClass
	CNNPoolingAverageClassOnce sync.Once
)

func getCNNPoolingAverageClass() _CNNPoolingAverageClass {
	CNNPoolingAverageClassOnce.Do(func() {
		CNNPoolingAverageClass = _CNNPoolingAverageClass{objc.GetClass("MPSCNNPoolingAverage")}
	})
	return CNNPoolingAverageClass
}

type _CNNPoolingAverageClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingAverage] class.
type ICNNPoolingAverage interface {
	ICNNPooling
	

	// properties:
	ZeroPadSizeX() objectivec.IObject
	SetZeroPadSizeX(value objectivec.IObject)
	ZeroPadSizeY() objectivec.IObject
	SetZeroPadSizeY(value objectivec.IObject)
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageClass) Alloc() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageClass) New() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverage) Init() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverage) Autorelease() CNNPoolingAverage {
	rv := objc.Send[CNNPoolingAverage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverage creates a new CNNPoolingAverage instance.
func NewCNNPoolingAverage() CNNPoolingAverage {
	return getCNNPoolingAverageClass().New()
}





// An average pooling filter.
//
// For each pixel in an image, the filter returns the average value of the pixels in the filter region defined by . When the value of the property is set to , the filtering window is shrunk to remain within the source image borders. For pixels close to the image borders, the filtering window will be smaller in order to fit inside the source image and less values will be used to compute the average value. In case the filtering window is entirely outside the source image border, the output value will be .


// An average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverage
type CNNPoolingAverage struct {
	CNNPooling
}

// CNNPoolingAverageFrom constructs a [CNNPoolingAverage] from an unsafe.Pointer.
//
// An average pooling filter.
func CNNPoolingAverageFrom(ptr unsafe.Pointer) CNNPoolingAverage {
	return CNNPoolingAverage{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}






// Initializes an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2866999-initwithcoder
func NewCNNPoolingAverageWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingAverage {
	instance := getCNNPoolingAverageClass().Alloc()
	rv := objc.Send[CNNPoolingAverage](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes an average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875216-initwithdevice
func NewCNNPoolingAverageWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverage {
	instance := getCNNPoolingAverageClass().Alloc()
	rv := objc.Send[CNNPoolingAverage](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex
func (c_ CNNPoolingAverage) ZeroPadSizeX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex
func (c_ CNNPoolingAverage) SetZeroPadSizeX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey
func (c_ CNNPoolingAverage) ZeroPadSizeY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey
func (c_ CNNPoolingAverage) SetZeroPadSizeY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeY:"), value)
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/edgemode
func (c_ CNNPoolingAverage) EdgeMode() ImageEdgeMode {
	rv := objc.Send[ImageEdgeMode](c_.ID, objc.Sel("edgeMode"))
	return rv
}


// The edge mode to use when texture reads stray off the edge of an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/edgemode
func (c_ CNNPoolingAverage) SetEdgeMode(value ImageEdgeMode) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEdgeMode:"), value)
}







