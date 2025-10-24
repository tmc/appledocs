// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNPoolingL2Norm] class.
var (
	CNNPoolingL2NormClass     _CNNPoolingL2NormClass
	CNNPoolingL2NormClassOnce sync.Once
)

func getCNNPoolingL2NormClass() _CNNPoolingL2NormClass {
	CNNPoolingL2NormClassOnce.Do(func() {
		CNNPoolingL2NormClass = _CNNPoolingL2NormClass{objc.GetClass("MPSCNNPoolingL2Norm")}
	})
	return CNNPoolingL2NormClass
}

type _CNNPoolingL2NormClass struct {
	class objc.Class
}





// An interface definition for the [CNNPoolingL2Norm] class.
type ICNNPoolingL2Norm interface {
	ICNNPooling
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNPoolingL2NormClass) Alloc() CNNPoolingL2Norm {
	rv := objc.Send[CNNPoolingL2Norm](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNPoolingL2NormClass) New() CNNPoolingL2Norm {
	rv := objc.Send[CNNPoolingL2Norm](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNPoolingL2Norm) Init() CNNPoolingL2Norm {
	rv := objc.Send[CNNPoolingL2Norm](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNPoolingL2Norm) Autorelease() CNNPoolingL2Norm {
	rv := objc.Send[CNNPoolingL2Norm](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNPoolingL2Norm creates a new CNNPoolingL2Norm instance.
func NewCNNPoolingL2Norm() CNNPoolingL2Norm {
	return getCNNPoolingL2NormClass().New()
}





// An L2-norm pooling filter.
//
// For each pixel, returns L2-Norm of pixels in the filter region:


// An L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNPoolingL2Norm
type CNNPoolingL2Norm struct {
	CNNPooling
}

// CNNPoolingL2NormFrom constructs a [CNNPoolingL2Norm] from an unsafe.Pointer.
//
// An L2-norm pooling filter.
func CNNPoolingL2NormFrom(ptr unsafe.Pointer) CNNPoolingL2Norm {
	return CNNPoolingL2Norm{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}






// Initializes an L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2norm/2867141-initwithcoder
func NewCNNPoolingL2NormWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNPoolingL2Norm {
	instance := getCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[CNNPoolingL2Norm](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// Initializes an L2-norm pooling filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2norm/2875162-initwithdevice
func NewCNNPoolingL2NormWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device unsafe.Pointer, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2Norm {
	instance := getCNNPoolingL2NormClass().Alloc()
	rv := objc.Send[CNNPoolingL2Norm](instance.ID, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	rv.Autorelease()
	return rv
}



























