// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [CNNLogSoftMax] class.
var (
	CNNLogSoftMaxClass     _CNNLogSoftMaxClass
	CNNLogSoftMaxClassOnce sync.Once
)

func getCNNLogSoftMaxClass() _CNNLogSoftMaxClass {
	CNNLogSoftMaxClassOnce.Do(func() {
		CNNLogSoftMaxClass = _CNNLogSoftMaxClass{objc.GetClass("MPSCNNLogSoftMax")}
	})
	return CNNLogSoftMaxClass
}

type _CNNLogSoftMaxClass struct {
	class objc.Class
}





// An interface definition for the [CNNLogSoftMax] class.
type ICNNLogSoftMax interface {
	ICNNKernel
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLogSoftMaxClass) Alloc() CNNLogSoftMax {
	rv := objc.Send[CNNLogSoftMax](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLogSoftMaxClass) New() CNNLogSoftMax {
	rv := objc.Send[CNNLogSoftMax](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLogSoftMax) Init() CNNLogSoftMax {
	rv := objc.Send[CNNLogSoftMax](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLogSoftMax) Autorelease() CNNLogSoftMax {
	rv := objc.Send[CNNLogSoftMax](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLogSoftMax creates a new CNNLogSoftMax instance.
func NewCNNLogSoftMax() CNNLogSoftMax {
	return getCNNLogSoftMaxClass().New()
}





// A neural transfer function that is useful for constructing a loss function to be minimized when training neural networks.
//
// The logarithmic softmax filter is calculated by taking the natural logarithm of the result of a softmax filter. For each feature channel per pixel in an image in a feature map, the logarithmic softmax filter computes the following: Where is the result channel in the pixel, is the number of feature channels, and satisfies .


// A neural transfer function that is useful for constructing a loss function to be minimized when training neural networks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLogSoftMax
type CNNLogSoftMax struct {
	CNNKernel
}

// CNNLogSoftMaxFrom constructs a [CNNLogSoftMax] from an unsafe.Pointer.
//
// A neural transfer function that is useful for constructing a loss function to be minimized when training neural networks.
func CNNLogSoftMaxFrom(ptr unsafe.Pointer) CNNLogSoftMax {
	return CNNLogSoftMax{
		CNNKernel: CNNKernelFrom(ptr),
	}
}































