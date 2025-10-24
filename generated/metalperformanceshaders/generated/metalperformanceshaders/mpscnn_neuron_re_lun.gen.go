// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronReLUN */


/* debug [class_header]: Header for MPSCNNNeuronReLUN */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronReLUN */
// An interface definition for the [CNNNeuronReLUN] class.
type ICNNNeuronReLUN interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronReLUN */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronReLUN */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronReLUN */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronReLUN */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronReLUN */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelun/2921658-initwithdevice
func NewCNNNeuronReLUNWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronReLUN {
	instance := getCNNNeuronReLUNClass().Alloc()
	rv := objc.Send[CNNNeuronReLUN](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronReLUNWithDeviceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronReLUN */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronReLUN */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronReLUN */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronReLUN */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronReLUN */


