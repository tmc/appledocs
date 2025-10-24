// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronReLU] class.
var (
	CNNNeuronReLUClass     _CNNNeuronReLUClass
	CNNNeuronReLUClassOnce sync.Once
)

func getCNNNeuronReLUClass() _CNNNeuronReLUClass {
	CNNNeuronReLUClassOnce.Do(func() {
		CNNNeuronReLUClass = _CNNNeuronReLUClass{objc.GetClass("MPSCNNNeuronReLU")}
	})
	return CNNNeuronReLUClass
}

type _CNNNeuronReLUClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronReLU] class.
type ICNNNeuronReLU interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUClass) Alloc() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUClass) New() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLU) Init() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLU) Autorelease() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLU creates a new CNNNeuronReLU instance.
func NewCNNNeuronReLU() CNNNeuronReLU {
	return getCNNNeuronReLUClass().New()
}





// A ReLU (Rectified Linear Unit) neuron filter.
//
// For each pixel in an image, the filter applies the following function: This filter is called in CNN literature. Some CNN literature defines as . If you want this behavior, simply set the property to .


// A ReLU (Rectified Linear Unit) neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU
type CNNNeuronReLU struct {
	CNNNeuron
}

// CNNNeuronReLUFrom constructs a [CNNNeuronReLU] from an unsafe.Pointer.
//
// A ReLU (Rectified Linear Unit) neuron filter.
func CNNNeuronReLUFrom(ptr unsafe.Pointer) CNNNeuronReLU {
	return CNNNeuronReLU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// Initializes a ReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU/init(device:a:)
func NewCNNNeuronReLUWithDeviceA(device unsafe.Pointer, a float32) CNNNeuronReLU {
	instance := getCNNNeuronReLUClass().Alloc()
	rv := objc.Send[CNNNeuronReLU](instance.ID, objc.Sel("initWithDevice:a:"), device, a)
	rv.Autorelease()
	return rv
}



























