// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronHardSigmoid] class.
var (
	CNNNeuronHardSigmoidClass     _CNNNeuronHardSigmoidClass
	CNNNeuronHardSigmoidClassOnce sync.Once
)

func getCNNNeuronHardSigmoidClass() _CNNNeuronHardSigmoidClass {
	CNNNeuronHardSigmoidClassOnce.Do(func() {
		CNNNeuronHardSigmoidClass = _CNNNeuronHardSigmoidClass{objc.GetClass("MPSCNNNeuronHardSigmoid")}
	})
	return CNNNeuronHardSigmoidClass
}

type _CNNNeuronHardSigmoidClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronHardSigmoid] class.
type ICNNNeuronHardSigmoid interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronHardSigmoidClass) Alloc() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronHardSigmoidClass) New() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronHardSigmoid) Init() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronHardSigmoid) Autorelease() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronHardSigmoid creates a new CNNNeuronHardSigmoid instance.
func NewCNNNeuronHardSigmoid() CNNNeuronHardSigmoid {
	return getCNNNeuronHardSigmoidClass().New()
}





// A hard sigmoid neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A hard sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoid
type CNNNeuronHardSigmoid struct {
	CNNNeuron
}

// CNNNeuronHardSigmoidFrom constructs a [CNNNeuronHardSigmoid] from an unsafe.Pointer.
//
// A hard sigmoid neuron filter.
func CNNNeuronHardSigmoidFrom(ptr unsafe.Pointer) CNNNeuronHardSigmoid {
	return CNNNeuronHardSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// Initializes a hard sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoid/2875217-initwithdevice
func NewCNNNeuronHardSigmoidWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronHardSigmoid {
	instance := getCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[CNNNeuronHardSigmoid](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}



























