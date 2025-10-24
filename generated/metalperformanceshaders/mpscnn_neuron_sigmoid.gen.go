// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronSigmoid] class.
var (
	CNNNeuronSigmoidClass     _CNNNeuronSigmoidClass
	CNNNeuronSigmoidClassOnce sync.Once
)

func getCNNNeuronSigmoidClass() _CNNNeuronSigmoidClass {
	CNNNeuronSigmoidClassOnce.Do(func() {
		CNNNeuronSigmoidClass = _CNNNeuronSigmoidClass{objc.GetClass("MPSCNNNeuronSigmoid")}
	})
	return CNNNeuronSigmoidClass
}

type _CNNNeuronSigmoidClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronSigmoid] class.
type ICNNNeuronSigmoid interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSigmoidClass) Alloc() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSigmoidClass) New() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSigmoid) Init() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSigmoid) Autorelease() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSigmoid creates a new CNNNeuronSigmoid instance.
func NewCNNNeuronSigmoid() CNNNeuronSigmoid {
	return getCNNNeuronSigmoidClass().New()
}





// A sigmoid neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoid
type CNNNeuronSigmoid struct {
	CNNNeuron
}

// CNNNeuronSigmoidFrom constructs a [CNNNeuronSigmoid] from an unsafe.Pointer.
//
// A sigmoid neuron filter.
func CNNNeuronSigmoidFrom(ptr unsafe.Pointer) CNNNeuronSigmoid {
	return CNNNeuronSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// Initializes a sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoid/1648890-initwithdevice
func NewCNNNeuronSigmoidWithDevice(device unsafe.Pointer) CNNNeuronSigmoid {
	instance := getCNNNeuronSigmoidClass().Alloc()
	rv := objc.Send[CNNNeuronSigmoid](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



























