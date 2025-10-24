// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNDilatedPoolingMaxGradient] class.
var (
	CNNDilatedPoolingMaxGradientClass     _CNNDilatedPoolingMaxGradientClass
	CNNDilatedPoolingMaxGradientClassOnce sync.Once
)

func getCNNDilatedPoolingMaxGradientClass() _CNNDilatedPoolingMaxGradientClass {
	CNNDilatedPoolingMaxGradientClassOnce.Do(func() {
		CNNDilatedPoolingMaxGradientClass = _CNNDilatedPoolingMaxGradientClass{objc.GetClass("MPSCNNDilatedPoolingMaxGradient")}
	})
	return CNNDilatedPoolingMaxGradientClass
}

type _CNNDilatedPoolingMaxGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNDilatedPoolingMaxGradient] class.
type ICNNDilatedPoolingMaxGradient interface {
	ICNNPoolingGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNDilatedPoolingMaxGradientClass) Alloc() CNNDilatedPoolingMaxGradient {
	rv := objc.Send[CNNDilatedPoolingMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDilatedPoolingMaxGradientClass) New() CNNDilatedPoolingMaxGradient {
	rv := objc.Send[CNNDilatedPoolingMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDilatedPoolingMaxGradient) Init() CNNDilatedPoolingMaxGradient {
	rv := objc.Send[CNNDilatedPoolingMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDilatedPoolingMaxGradient) Autorelease() CNNDilatedPoolingMaxGradient {
	rv := objc.Send[CNNDilatedPoolingMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDilatedPoolingMaxGradient creates a new CNNDilatedPoolingMaxGradient instance.
func NewCNNDilatedPoolingMaxGradient() CNNDilatedPoolingMaxGradient {
	return getCNNDilatedPoolingMaxGradientClass().New()
}





// A gradient dilated max pooling filter.
//
// A gradient max pooling filter but the pixels selected in each “application” of the max pooling operation are exactly the same pixels that would be selected with dilated convolution


// A gradient dilated max pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDilatedPoolingMaxGradient
type CNNDilatedPoolingMaxGradient struct {
	CNNPoolingGradient
}

// CNNDilatedPoolingMaxGradientFrom constructs a [CNNDilatedPoolingMaxGradient] from an unsafe.Pointer.
//
// A gradient dilated max pooling filter.
func CNNDilatedPoolingMaxGradientFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxGradient {
	return CNNDilatedPoolingMaxGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradient/2942346-initwithcoder
func NewCNNDilatedPoolingMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNDilatedPoolingMaxGradient {
	instance := getCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradient/2942349-initwithdevice
func NewCNNDilatedPoolingMaxGradientWithDeviceKernelWidthKernelHeightDilationRateXDilationRateYStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, dilationRateX uint, dilationRateY uint, strideInPixelsX uint, strideInPixelsY uint) CNNDilatedPoolingMaxGradient {
	instance := getCNNDilatedPoolingMaxGradientClass().Alloc()
	rv := objc.Send[CNNDilatedPoolingMaxGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:dilationRateX:dilationRateY:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, dilationRateX, dilationRateY, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}



























