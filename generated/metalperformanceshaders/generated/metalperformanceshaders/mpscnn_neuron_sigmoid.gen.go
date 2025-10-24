// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronSigmoid */


/* debug [class_header]: Header for MPSCNNNeuronSigmoid */
// The class instance for the [CNNNeuronSigmoid] class.
var (
	CNNNeuronSigmoidClass     _CNNNeuronSigmoidClass
	CNNNeuronSigmoidClassOnce sync.Once
)

func getCNNNeuronSigmoidClass() _CNNNeuronSigmoidClass {
	CNNNeuronSigmoidClassOnce.Do(func() {
		CNNNeuronSigmoidClass = _CNNNeuronSigmoidClass{objc.GetClass("MPSCNNNeuronSigmoid")}
	})
	return CNNNeuronSigmoidClass
}

type _CNNNeuronSigmoidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronSigmoid */
// An interface definition for the [CNNNeuronSigmoid] class.
type ICNNNeuronSigmoid interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronSigmoid */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronSigmoid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronSigmoid */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSigmoidClass) Alloc() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSigmoidClass) New() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSigmoid) Init() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSigmoid) Autorelease() CNNNeuronSigmoid {
	rv := objc.Send[CNNNeuronSigmoid](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSigmoid creates a new CNNNeuronSigmoid instance.
func NewCNNNeuronSigmoid() CNNNeuronSigmoid {
	return getCNNNeuronSigmoidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronSigmoid */
// A sigmoid neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSigmoid
type CNNNeuronSigmoid struct {
	CNNNeuron
}

// CNNNeuronSigmoidFrom constructs a [CNNNeuronSigmoid] from an unsafe.Pointer.
//
// A sigmoid neuron filter.
func CNNNeuronSigmoidFrom(ptr unsafe.Pointer) CNNNeuronSigmoid {
	return CNNNeuronSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronSigmoid */

// Initializes a sigmoid neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoid/1648890-initwithdevice
func NewCNNNeuronSigmoidWithDevice(device unsafe.Pointer) CNNNeuronSigmoid {
	instance := getCNNNeuronSigmoidClass().Alloc()
	rv := objc.Send[CNNNeuronSigmoid](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSigmoidWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronSigmoid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronSigmoid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronSigmoid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronSigmoid */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronSigmoid */


