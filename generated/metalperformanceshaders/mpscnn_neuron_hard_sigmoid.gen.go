// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronHardSigmoid */


/* debug [class_header]: Header for MPSCNNNeuronHardSigmoid */
// The class instance for the [CNNNeuronHardSigmoid] class.
var (
	CNNNeuronHardSigmoidClass     _CNNNeuronHardSigmoidClass
	CNNNeuronHardSigmoidClassOnce sync.Once
)

func getCNNNeuronHardSigmoidClass() _CNNNeuronHardSigmoidClass {
	CNNNeuronHardSigmoidClassOnce.Do(func() {
		CNNNeuronHardSigmoidClass = _CNNNeuronHardSigmoidClass{objc.GetClass("MPSCNNNeuronHardSigmoid")}
	})
	return CNNNeuronHardSigmoidClass
}

type _CNNNeuronHardSigmoidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronHardSigmoid */
// An interface definition for the [CNNNeuronHardSigmoid] class.
type ICNNNeuronHardSigmoid interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronHardSigmoid */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronHardSigmoid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronHardSigmoid */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronHardSigmoidClass) Alloc() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronHardSigmoidClass) New() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronHardSigmoid) Init() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronHardSigmoid) Autorelease() CNNNeuronHardSigmoid {
	rv := objc.Send[CNNNeuronHardSigmoid](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronHardSigmoid creates a new CNNNeuronHardSigmoid instance.
func NewCNNNeuronHardSigmoid() CNNNeuronHardSigmoid {
	return getCNNNeuronHardSigmoidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronHardSigmoid */
// A hard sigmoid neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A hard sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronHardSigmoid
type CNNNeuronHardSigmoid struct {
	CNNNeuron
}

// CNNNeuronHardSigmoidFrom constructs a [CNNNeuronHardSigmoid] from an unsafe.Pointer.
//
// A hard sigmoid neuron filter.
func CNNNeuronHardSigmoidFrom(ptr unsafe.Pointer) CNNNeuronHardSigmoid {
	return CNNNeuronHardSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronHardSigmoid */

// Initializes a hard sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoid/2875217-initwithdevice
func NewCNNNeuronHardSigmoidWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronHardSigmoid {
	instance := getCNNNeuronHardSigmoidClass().Alloc()
	rv := objc.Send[CNNNeuronHardSigmoid](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronHardSigmoidWithDeviceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronHardSigmoid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronHardSigmoid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronHardSigmoid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronHardSigmoid */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronHardSigmoid */


