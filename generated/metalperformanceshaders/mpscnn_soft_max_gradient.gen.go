// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNSoftMaxGradient] class.
var (
	CNNSoftMaxGradientClass     _CNNSoftMaxGradientClass
	CNNSoftMaxGradientClassOnce sync.Once
)

func getCNNSoftMaxGradientClass() _CNNSoftMaxGradientClass {
	CNNSoftMaxGradientClassOnce.Do(func() {
		CNNSoftMaxGradientClass = _CNNSoftMaxGradientClass{objc.GetClass("MPSCNNSoftMaxGradient")}
	})
	return CNNSoftMaxGradientClass
}

type _CNNSoftMaxGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNSoftMaxGradient] class.
type ICNNSoftMaxGradient interface {
	ICNNGradientKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNSoftMaxGradientClass) Alloc() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSoftMaxGradientClass) New() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSoftMaxGradient) Init() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSoftMaxGradient) Autorelease() CNNSoftMaxGradient {
	rv := objc.Send[CNNSoftMaxGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSoftMaxGradient creates a new CNNSoftMaxGradient instance.
func NewCNNSoftMaxGradient() CNNSoftMaxGradient {
	return getCNNSoftMaxGradientClass().New()
}





// A gradient softmax filter.


// A gradient softmax filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSoftMaxGradient
type CNNSoftMaxGradient struct {
	CNNGradientKernel
}

// CNNSoftMaxGradientFrom constructs a [CNNSoftMaxGradient] from an unsafe.Pointer.
//
// A gradient softmax filter.
func CNNSoftMaxGradientFrom(ptr unsafe.Pointer) CNNSoftMaxGradient {
	return CNNSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient/2942618-initwithcoder
func NewCNNSoftMaxGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNSoftMaxGradient {
	instance := getCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient/2942620-initwithdevice
func NewCNNSoftMaxGradientWithDevice(device unsafe.Pointer) CNNSoftMaxGradient {
	instance := getCNNSoftMaxGradientClass().Alloc()
	rv := objc.Send[CNNSoftMaxGradient](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























