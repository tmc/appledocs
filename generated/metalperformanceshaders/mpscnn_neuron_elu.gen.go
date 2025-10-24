// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronELU */


/* debug [class_header]: Header for MPSCNNNeuronELU */
// The class instance for the [CNNNeuronELU] class.
var (
	CNNNeuronELUClass     _CNNNeuronELUClass
	CNNNeuronELUClassOnce sync.Once
)

func getCNNNeuronELUClass() _CNNNeuronELUClass {
	CNNNeuronELUClassOnce.Do(func() {
		CNNNeuronELUClass = _CNNNeuronELUClass{objc.GetClass("MPSCNNNeuronELU")}
	})
	return CNNNeuronELUClass
}

type _CNNNeuronELUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronELU */
// An interface definition for the [CNNNeuronELU] class.
type ICNNNeuronELU interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronELU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronELU */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronELU */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronELUClass) Alloc() CNNNeuronELU {
	rv := objc.Send[CNNNeuronELU](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronELUClass) New() CNNNeuronELU {
	rv := objc.Send[CNNNeuronELU](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronELU) Init() CNNNeuronELU {
	rv := objc.Send[CNNNeuronELU](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronELU) Autorelease() CNNNeuronELU {
	rv := objc.Send[CNNNeuronELU](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronELU creates a new CNNNeuronELU instance.
func NewCNNNeuronELU() CNNNeuronELU {
	return getCNNNeuronELUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronELU */
// A parametric ELU neuron filter.
//
// For each pixel in an image, the filter applies the following function: ![f(x) = a * (exp(x) - 1) if x < 0 | f(x) =


// A parametric ELU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronELU
type CNNNeuronELU struct {
	CNNNeuron
}

// CNNNeuronELUFrom constructs a [CNNNeuronELU] from an unsafe.Pointer.
//
// A parametric ELU neuron filter.
func CNNNeuronELUFrom(ptr unsafe.Pointer) CNNNeuronELU {
	return CNNNeuronELU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronELU */

// Initializes a parametric ELU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelu/2867048-initwithdevice
func NewCNNNeuronELUWithDeviceA(device unsafe.Pointer, a float32) CNNNeuronELU {
	instance := getCNNNeuronELUClass().Alloc()
	rv := objc.Send[CNNNeuronELU](instance.ID, objc.Sel("initWithDevice:a:"), device, a)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronELUWithDeviceA */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronELU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronELU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronELU */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronELU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronELU */


