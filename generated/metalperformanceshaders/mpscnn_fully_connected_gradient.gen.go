// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNFullyConnectedGradient] class.
var (
	CNNFullyConnectedGradientClass     _CNNFullyConnectedGradientClass
	CNNFullyConnectedGradientClassOnce sync.Once
)

func getCNNFullyConnectedGradientClass() _CNNFullyConnectedGradientClass {
	CNNFullyConnectedGradientClassOnce.Do(func() {
		CNNFullyConnectedGradientClass = _CNNFullyConnectedGradientClass{objc.GetClass("MPSCNNFullyConnectedGradient")}
	})
	return CNNFullyConnectedGradientClass
}

type _CNNFullyConnectedGradientClass struct {
	class objc.Class
}





// An interface definition for the [CNNFullyConnectedGradient] class.
type ICNNFullyConnectedGradient interface {
	ICNNConvolutionGradient
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNFullyConnectedGradientClass) Alloc() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNFullyConnectedGradientClass) New() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNFullyConnectedGradient) Init() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNFullyConnectedGradient) Autorelease() CNNFullyConnectedGradient {
	rv := objc.Send[CNNFullyConnectedGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNFullyConnectedGradient creates a new CNNFullyConnectedGradient instance.
func NewCNNFullyConnectedGradient() CNNFullyConnectedGradient {
	return getCNNFullyConnectedGradientClass().New()
}





// A gradient fully connected convolution layer.


// A gradient fully connected convolution layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNFullyConnectedGradient
type CNNFullyConnectedGradient struct {
	CNNConvolutionGradient
}

// CNNFullyConnectedGradientFrom constructs a [CNNFullyConnectedGradient] from an unsafe.Pointer.
//
// A gradient fully connected convolution layer.
func CNNFullyConnectedGradientFrom(ptr unsafe.Pointer) CNNFullyConnectedGradient {
	return CNNFullyConnectedGradient{
		CNNConvolutionGradient: CNNConvolutionGradientFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient/2951923-initwithcoder
func NewCNNFullyConnectedGradientWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) CNNFullyConnectedGradient {
	instance := getCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradient](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient/2951921-initwithdevice
func NewCNNFullyConnectedGradientWithDeviceWeights(device unsafe.Pointer, weights unsafe.Pointer) CNNFullyConnectedGradient {
	instance := getCNNFullyConnectedGradientClass().Alloc()
	rv := objc.Send[CNNFullyConnectedGradient](instance.ID, objc.Sel("initWithDevice:weights:"), device, weights)
	rv.Autorelease()
	return rv
}



























