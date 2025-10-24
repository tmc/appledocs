// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingAverageGradient] class.
var (
	CNNPoolingAverageGradientClass     _CNNPoolingAverageGradientClass
	CNNPoolingAverageGradientClassOnce sync.Once
)

func getCNNPoolingAverageGradientClass() _CNNPoolingAverageGradientClass {
	CNNPoolingAverageGradientClassOnce.Do(func() {
		CNNPoolingAverageGradientClass = _CNNPoolingAverageGradientClass{objc.GetClass("MPSCNNPoolingAverageGradient")}
	})
	return CNNPoolingAverageGradientClass
}

type _CNNPoolingAverageGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingAverageGradient] class.
type ICNNPoolingAverageGradient interface {
	ICNNPoolingGradient
	

	// properties:
	ZeroPadSizeX() objectivec.IObject
	SetZeroPadSizeX(value objectivec.IObject)
	ZeroPadSizeY() objectivec.IObject
	SetZeroPadSizeY(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingAverageGradientClass) Alloc() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingAverageGradientClass) New() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingAverageGradient) Init() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingAverageGradient) Autorelease() CNNPoolingAverageGradient {
	rv := objc.Send[CNNPoolingAverageGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingAverageGradient creates a new CNNPoolingAverageGradient instance.
func NewCNNPoolingAverageGradient() CNNPoolingAverageGradient {
	return getCNNPoolingAverageGradientClass().New()
}





// A gradient average pooling filter.


// A gradient average pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingAverageGradient
type CNNPoolingAverageGradient struct {
	CNNPoolingGradient
}

// CNNPoolingAverageGradientFrom constructs a [CNNPoolingAverageGradient] from an unsafe.Pointer.
//
// A gradient average pooling filter.
func CNNPoolingAverageGradientFrom(ptr unsafe.Pointer) CNNPoolingAverageGradient {
	return CNNPoolingAverageGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942345-initwithcoder
func NewCNNPoolingAverageGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingAverageGradient {
	instance := getCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[CNNPoolingAverageGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942339-initwithdevice
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverageGradient {
	instance := getCNNPoolingAverageGradientClass().Alloc()
	rv := objc.Send[CNNPoolingAverageGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex
func (c_ CNNPoolingAverageGradient) ZeroPadSizeX() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeX"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeX(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeX:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey
func (c_ CNNPoolingAverageGradient) ZeroPadSizeY() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("zeroPadSizeY"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeY(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZeroPadSizeY:"), value)
}







