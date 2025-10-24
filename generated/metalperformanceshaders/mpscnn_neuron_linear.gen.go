// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronLinear */


/* debug [class_header]: Header for MPSCNNNeuronLinear */
// The class instance for the [CNNNeuronLinear] class.
var (
	CNNNeuronLinearClass     _CNNNeuronLinearClass
	CNNNeuronLinearClassOnce sync.Once
)

func getCNNNeuronLinearClass() _CNNNeuronLinearClass {
	CNNNeuronLinearClassOnce.Do(func() {
		CNNNeuronLinearClass = _CNNNeuronLinearClass{objc.GetClass("MPSCNNNeuronLinear")}
	})
	return CNNNeuronLinearClass
}

type _CNNNeuronLinearClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronLinear */
// An interface definition for the [CNNNeuronLinear] class.
type ICNNNeuronLinear interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronLinear */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronLinear */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronLinear */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronLinearClass) Alloc() CNNNeuronLinear {
	rv := objc.Send[CNNNeuronLinear](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronLinearClass) New() CNNNeuronLinear {
	rv := objc.Send[CNNNeuronLinear](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronLinear) Init() CNNNeuronLinear {
	rv := objc.Send[CNNNeuronLinear](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronLinear) Autorelease() CNNNeuronLinear {
	rv := objc.Send[CNNNeuronLinear](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronLinear creates a new CNNNeuronLinear instance.
func NewCNNNeuronLinear() CNNNeuronLinear {
	return getCNNNeuronLinearClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronLinear */
// A linear neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A linear neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinear
type CNNNeuronLinear struct {
	CNNNeuron
}

// CNNNeuronLinearFrom constructs a [CNNNeuronLinear] from an unsafe.Pointer.
//
// A linear neuron filter.
func CNNNeuronLinearFrom(ptr unsafe.Pointer) CNNNeuronLinear {
	return CNNNeuronLinear{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronLinear */

// Initializes a linear neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLinear/init(device:a:b:)
func NewCNNNeuronLinearWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronLinear {
	instance := getCNNNeuronLinearClass().Alloc()
	rv := objc.Send[CNNNeuronLinear](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLinearWithDeviceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronLinear */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronLinear */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronLinear */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronLinear */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronLinear */


