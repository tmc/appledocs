// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLogSoftMaxGradient] class.
var (
	CNNLogSoftMaxGradientClass     _CNNLogSoftMaxGradientClass
	CNNLogSoftMaxGradientClassOnce sync.Once
)

func getCNNLogSoftMaxGradientClass() _CNNLogSoftMaxGradientClass {
	CNNLogSoftMaxGradientClassOnce.Do(func() {
		CNNLogSoftMaxGradientClass = _CNNLogSoftMaxGradientClass{objc.GetClass("MPSCNNLogSoftMaxGradient")}
	})
	return CNNLogSoftMaxGradientClass
}

type _CNNLogSoftMaxGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNLogSoftMaxGradient] class.
type ICNNLogSoftMaxGradient interface {
	ICNNGradientKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxGradientClass) Alloc() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxGradientClass) New() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMaxGradient) Init() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMaxGradient) Autorelease() CNNLogSoftMaxGradient {
	rv := objc.Send[CNNLogSoftMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMaxGradient creates a new CNNLogSoftMaxGradient instance.
func NewCNNLogSoftMaxGradient() CNNLogSoftMaxGradient {
	return getCNNLogSoftMaxGradientClass().New()
}





// A gradient logarithmic softmax filter.


// A gradient logarithmic softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMaxGradient
type CNNLogSoftMaxGradient struct {
	CNNGradientKernel
}

// CNNLogSoftMaxGradientFrom constructs a [CNNLogSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient logarithmic softmax filter.
func CNNLogSoftMaxGradientFrom(ptr unsafe.Pointer) CNNLogSoftMaxGradient {
	return CNNLogSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient/2942619-initwithcoder
func NewCNNLogSoftMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNLogSoftMaxGradient {
	instance := getCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient/2942622-initwithdevice
func NewCNNLogSoftMaxGradientWithDevice(device unsafe.Pointer) CNNLogSoftMaxGradient {
	instance := getCNNLogSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNLogSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























