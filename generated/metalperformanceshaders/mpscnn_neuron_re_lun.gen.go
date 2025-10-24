// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNNeuronReLUN] class.
var (
	CNNNeuronReLUNClass     _CNNNeuronReLUNClass
	CNNNeuronReLUNClassOnce sync.Once
)

func getCNNNeuronReLUNClass() _CNNNeuronReLUNClass {
	CNNNeuronReLUNClassOnce.Do(func() {
		CNNNeuronReLUNClass = _CNNNeuronReLUNClass{objc.GetClass("MPSCNNNeuronReLUN")}
	})
	return CNNNeuronReLUNClass
}

type _CNNNeuronReLUNClass struct {
	class objc.Class
}





// An interface definition for the [CNNNeuronReLUN] class.
type ICNNNeuronReLUN interface {
	ICNNNeuron
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUNClass) Alloc() CNNNeuronReLUN {
	rv := objc.Send[CNNNeuronReLUN](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUNClass) New() CNNNeuronReLUN {
	rv := objc.Send[CNNNeuronReLUN](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLUN) Init() CNNNeuronReLUN {
	rv := objc.Send[CNNNeuronReLUN](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLUN) Autorelease() CNNNeuronReLUN {
	rv := objc.Send[CNNNeuronReLUN](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLUN creates a new CNNNeuronReLUN instance.
func NewCNNNeuronReLUN() CNNNeuronReLUN {
	return getCNNNeuronReLUNClass().New()
}





// A ReLUN neuron filter.
//
// For each pixel in an image, the filter applies the following function: The default value of is 1.0 and the default value of is 6.0.


// A ReLUN neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLUN
type CNNNeuronReLUN struct {
	CNNNeuron
}

// CNNNeuronReLUNFrom constructs a [CNNNeuronReLUN] from an unsafe.Pointer.
//
// A ReLUN neuron filter.
func CNNNeuronReLUNFrom(ptr unsafe.Pointer) CNNNeuronReLUN {
	return CNNNeuronReLUN{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelun/2921658-initwithdevice
func NewCNNNeuronReLUNWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronReLUN {
	instance := getCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[CNNNeuronReLUN](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}



























