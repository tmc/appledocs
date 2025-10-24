// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronLogarithm] class.
var (
	CNNNeuronLogarithmClass     _CNNNeuronLogarithmClass
	CNNNeuronLogarithmClassOnce sync.Once
)

func getCNNNeuronLogarithmClass() _CNNNeuronLogarithmClass {
	CNNNeuronLogarithmClassOnce.Do(func() {
		CNNNeuronLogarithmClass = _CNNNeuronLogarithmClass{objc.GetClass("MPSCNNNeuronLogarithm")}
	})
	return CNNNeuronLogarithmClass
}

type _CNNNeuronLogarithmClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronLogarithm] class.
type ICNNNeuronLogarithm interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronLogarithmClass) Alloc() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronLogarithmClass) New() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronLogarithm) Init() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronLogarithm) Autorelease() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronLogarithm creates a new CNNNeuronLogarithm instance.
func NewCNNNeuronLogarithm() CNNNeuronLogarithm {
	return getCNNNeuronLogarithmClass().New()
}





// A logarithm neuron filter.


// A logarithm neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithm
type CNNNeuronLogarithm struct {
	CNNNeuron
}

// CNNNeuronLogarithmFrom constructs a [CNNNeuronLogarithm] from an unsafe.Pointer.
//
// A logarithm neuron filter.
func CNNNeuronLogarithmFrom(ptr unsafe.Pointer) CNNNeuronLogarithm {
	return CNNNeuronLogarithm{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithm/2951871-initwithdevice
func NewCNNNeuronLogarithmWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronLogarithm {
	instance := getCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[CNNNeuronLogarithm](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}



























