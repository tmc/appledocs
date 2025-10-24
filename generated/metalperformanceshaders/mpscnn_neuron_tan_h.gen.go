// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronTanH] class.
var (
	CNNNeuronTanHClass     _CNNNeuronTanHClass
	CNNNeuronTanHClassOnce sync.Once
)

func getCNNNeuronTanHClass() _CNNNeuronTanHClass {
	CNNNeuronTanHClassOnce.Do(func() {
		CNNNeuronTanHClass = _CNNNeuronTanHClass{objc.GetClass("MPSCNNNeuronTanH")}
	})
	return CNNNeuronTanHClass
}

type _CNNNeuronTanHClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronTanH] class.
type ICNNNeuronTanH interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronTanHClass) Alloc() CNNNeuronTanH {
	rv := objc.Send[CNNNeuronTanH](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronTanHClass) New() CNNNeuronTanH {
	rv := objc.Send[CNNNeuronTanH](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronTanH) Init() CNNNeuronTanH {
	rv := objc.Send[CNNNeuronTanH](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronTanH) Autorelease() CNNNeuronTanH {
	rv := objc.Send[CNNNeuronTanH](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronTanH creates a new CNNNeuronTanH instance.
func NewCNNNeuronTanH() CNNNeuronTanH {
	return getCNNNeuronTanHClass().New()
}





// A hyperbolic tangent neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A hyperbolic tangent neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanH
type CNNNeuronTanH struct {
	CNNNeuron
}

// CNNNeuronTanHFrom constructs a [CNNNeuronTanH] from an unsafe.Pointer.
//
// A hyperbolic tangent neuron filter.
func CNNNeuronTanHFrom(ptr unsafe.Pointer) CNNNeuronTanH {
	return CNNNeuronTanH{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// Initializes a hyperbolic tangent neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronTanH/init(device:a:b:)
func NewCNNNeuronTanHWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronTanH {
	instance := getCNNNeuronTanHClass().Alloc()
	rv := objc.Send[CNNNeuronTanH](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}



























