// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingL2NormGradient] class.
var (
	CNNPoolingL2NormGradientClass     _CNNPoolingL2NormGradientClass
	CNNPoolingL2NormGradientClassOnce sync.Once
)

func getCNNPoolingL2NormGradientClass() _CNNPoolingL2NormGradientClass {
	CNNPoolingL2NormGradientClassOnce.Do(func() {
		CNNPoolingL2NormGradientClass = _CNNPoolingL2NormGradientClass{objc.GetClass("MPSCNNPoolingL2NormGradient")}
	})
	return CNNPoolingL2NormGradientClass
}

type _CNNPoolingL2NormGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingL2NormGradient] class.
type ICNNPoolingL2NormGradient interface {
	ICNNPoolingGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingL2NormGradientClass) Alloc() CNNPoolingL2NormGradient {
	rv := objc.Send[CNNPoolingL2NormGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingL2NormGradientClass) New() CNNPoolingL2NormGradient {
	rv := objc.Send[CNNPoolingL2NormGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingL2NormGradient) Init() CNNPoolingL2NormGradient {
	rv := objc.Send[CNNPoolingL2NormGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingL2NormGradient) Autorelease() CNNPoolingL2NormGradient {
	rv := objc.Send[CNNPoolingL2NormGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingL2NormGradient creates a new CNNPoolingL2NormGradient instance.
func NewCNNPoolingL2NormGradient() CNNPoolingL2NormGradient {
	return getCNNPoolingL2NormGradientClass().New()
}





// A gradient L2-norm pooling filter.


// A gradient L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2NormGradient
type CNNPoolingL2NormGradient struct {
	CNNPoolingGradient
}

// CNNPoolingL2NormGradientFrom constructs a [CNNPoolingL2NormGradient] from an unsafe.Pointer.
//
// A gradient L2-norm pooling filter.
func CNNPoolingL2NormGradientFrom(ptr unsafe.Pointer) CNNPoolingL2NormGradient {
	return CNNPoolingL2NormGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normgradient/2942352-initwithcoder
func NewCNNPoolingL2NormGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingL2NormGradient {
	instance := getCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[CNNPoolingL2NormGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normgradient/2942355-initwithdevice
func NewCNNPoolingL2NormGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2NormGradient {
	instance := getCNNPoolingL2NormGradientClass().Alloc()
	rv := objc.Send[CNNPoolingL2NormGradient](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}



























