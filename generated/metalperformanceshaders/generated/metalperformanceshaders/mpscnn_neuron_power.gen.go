// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronPower */


/* debug [class_header]: Header for MPSCNNNeuronPower */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronPower */
// An interface definition for the [CNNNeuronPower] class.
type ICNNNeuronPower interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronPower */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronPower */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronPower */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronPower */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronPower */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpower/2951868-initwithdevice
func NewCNNNeuronPowerWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronPower {
	instance := getCNNNeuronPowerClass().Alloc()
	rv := objc.Send[CNNNeuronPower](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronPowerWithDeviceABC */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronPower */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronPower */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronPower */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronPower */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronPower */


