// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronPower] class.
var (
	CNNNeuronPowerClass     _CNNNeuronPowerClass
	CNNNeuronPowerClassOnce sync.Once
)

func getCNNNeuronPowerClass() _CNNNeuronPowerClass {
	CNNNeuronPowerClassOnce.Do(func() {
		CNNNeuronPowerClass = _CNNNeuronPowerClass{objc.GetClass("MPSCNNNeuronPower")}
	})
	return CNNNeuronPowerClass
}

type _CNNNeuronPowerClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronPower] class.
type ICNNNeuronPower interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronPowerClass) Alloc() CNNNeuronPower {
	rv := objc.Send[CNNNeuronPower](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronPowerClass) New() CNNNeuronPower {
	rv := objc.Send[CNNNeuronPower](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronPower) Init() CNNNeuronPower {
	rv := objc.Send[CNNNeuronPower](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronPower) Autorelease() CNNNeuronPower {
	rv := objc.Send[CNNNeuronPower](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronPower creates a new CNNNeuronPower instance.
func NewCNNNeuronPower() CNNNeuronPower {
	return getCNNNeuronPowerClass().New()
}





// A power neuron filter.


// A power neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPower
type CNNNeuronPower struct {
	CNNNeuron
}

// CNNNeuronPowerFrom constructs a [CNNNeuronPower] from an unsafe.Pointer.
//
// A power neuron filter.
func CNNNeuronPowerFrom(ptr unsafe.Pointer) CNNNeuronPower {
	return CNNNeuronPower{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpower/2951868-initwithdevice
func NewCNNNeuronPowerWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronPower {
	instance := getCNNNeuronPowerClass().Alloc()
	rv := objc.Send[CNNNeuronPower](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}



























