// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronExponential] class.
var (
	CNNNeuronExponentialClass     _CNNNeuronExponentialClass
	CNNNeuronExponentialClassOnce sync.Once
)

func getCNNNeuronExponentialClass() _CNNNeuronExponentialClass {
	CNNNeuronExponentialClassOnce.Do(func() {
		CNNNeuronExponentialClass = _CNNNeuronExponentialClass{objc.GetClass("MPSCNNNeuronExponential")}
	})
	return CNNNeuronExponentialClass
}

type _CNNNeuronExponentialClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronExponential] class.
type ICNNNeuronExponential interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronExponentialClass) Alloc() CNNNeuronExponential {
	rv := objc.Send[CNNNeuronExponential](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronExponentialClass) New() CNNNeuronExponential {
	rv := objc.Send[CNNNeuronExponential](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronExponential) Init() CNNNeuronExponential {
	rv := objc.Send[CNNNeuronExponential](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronExponential) Autorelease() CNNNeuronExponential {
	rv := objc.Send[CNNNeuronExponential](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronExponential creates a new CNNNeuronExponential instance.
func NewCNNNeuronExponential() CNNNeuronExponential {
	return getCNNNeuronExponentialClass().New()
}





// An exponential neuron filter.


// An exponential neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronExponential
type CNNNeuronExponential struct {
	CNNNeuron
}

// CNNNeuronExponentialFrom constructs a [CNNNeuronExponential] from an unsafe.Pointer.
//
// An exponential neuron filter.
func CNNNeuronExponentialFrom(ptr unsafe.Pointer) CNNNeuronExponential {
	return CNNNeuronExponential{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponential/2951870-initwithdevice
func NewCNNNeuronExponentialWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronExponential {
	instance := getCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[CNNNeuronExponential](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}



























