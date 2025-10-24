// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronExponential */


/* debug [class_header]: Header for MPSCNNNeuronExponential */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronExponential */
// An interface definition for the [CNNNeuronExponential] class.
type ICNNNeuronExponential interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronExponential */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronExponential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronExponential */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronExponential */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronExponential */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponential/2951870-initwithdevice
func NewCNNNeuronExponentialWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronExponential {
	instance := getCNNNeuronExponentialClass().Alloc()
	rv := objc.Send[CNNNeuronExponential](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronExponentialWithDeviceABC */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronExponential */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronExponential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronExponential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronExponential */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronExponential */


