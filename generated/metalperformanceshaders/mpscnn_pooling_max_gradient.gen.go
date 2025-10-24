// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingMaxGradient] class.
var (
	CNNPoolingMaxGradientClass     _CNNPoolingMaxGradientClass
	CNNPoolingMaxGradientClassOnce sync.Once
)

func getCNNPoolingMaxGradientClass() _CNNPoolingMaxGradientClass {
	CNNPoolingMaxGradientClassOnce.Do(func() {
		CNNPoolingMaxGradientClass = _CNNPoolingMaxGradientClass{objc.GetClass("MPSCNNPoolingMaxGradient")}
	})
	return CNNPoolingMaxGradientClass
}

type _CNNPoolingMaxGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingMaxGradient] class.
type ICNNPoolingMaxGradient interface {
	ICNNPoolingGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingMaxGradientClass) Alloc() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingMaxGradientClass) New() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingMaxGradient) Init() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingMaxGradient) Autorelease() CNNPoolingMaxGradient {
	rv := objc.Send[CNNPoolingMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingMaxGradient creates a new CNNPoolingMaxGradient instance.
func NewCNNPoolingMaxGradient() CNNPoolingMaxGradient {
	return getCNNPoolingMaxGradientClass().New()
}





// A gradient max pooling filter.


// A gradient max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingMaxGradient
type CNNPoolingMaxGradient struct {
	CNNPoolingGradient
}

// CNNPoolingMaxGradientFrom constructs a [CNNPoolingMaxGradient] from an unsafe.Pointer.
//
// A gradient max pooling filter.
func CNNPoolingMaxGradientFrom(ptr unsafe.Pointer) CNNPoolingMaxGradient {
	return CNNPoolingMaxGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient/2942342-initwithcoder
func NewCNNPoolingMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingMaxGradient {
	instance := getCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNPoolingMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient/2942348-initwithdevice
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMaxGradient {
	instance := getCNNPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNPoolingMaxGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}



























