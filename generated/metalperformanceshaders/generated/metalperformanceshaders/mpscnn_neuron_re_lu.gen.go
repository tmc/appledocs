// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronReLU */


/* debug [class_header]: Header for MPSCNNNeuronReLU */
// The class instance for the [CNNNeuronReLU] class.
var (
	CNNNeuronReLUClass     _CNNNeuronReLUClass
	CNNNeuronReLUClassOnce sync.Once
)

func getCNNNeuronReLUClass() _CNNNeuronReLUClass {
	CNNNeuronReLUClassOnce.Do(func() {
		CNNNeuronReLUClass = _CNNNeuronReLUClass{objc.GetClass("MPSCNNNeuronReLU")}
	})
	return CNNNeuronReLUClass
}

type _CNNNeuronReLUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronReLU */
// An interface definition for the [CNNNeuronReLU] class.
type ICNNNeuronReLU interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronReLU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronReLU */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronReLU */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronReLUClass) Alloc() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronReLUClass) New() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronReLU) Init() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronReLU) Autorelease() CNNNeuronReLU {
	rv := objc.Send[CNNNeuronReLU](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronReLU creates a new CNNNeuronReLU instance.
func NewCNNNeuronReLU() CNNNeuronReLU {
	return getCNNNeuronReLUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronReLU */
// A ReLU (Rectified Linear Unit) neuron filter.
//
// For each pixel in an image, the filter applies the following function: This filter is called in CNN literature. Some CNN literature defines as . If you want this behavior, simply set the property to .


// A ReLU (Rectified Linear Unit) neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU
type CNNNeuronReLU struct {
	CNNNeuron
}

// CNNNeuronReLUFrom constructs a [CNNNeuronReLU] from an unsafe.Pointer.
//
// A ReLU (Rectified Linear Unit) neuron filter.
func CNNNeuronReLUFrom(ptr unsafe.Pointer) CNNNeuronReLU {
	return CNNNeuronReLU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronReLU */

// Initializes a ReLU neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronReLU/init(device:a:)
func NewCNNNeuronReLUWithDeviceA(device unsafe.Pointer, a float32) CNNNeuronReLU {
	instance := getCNNNeuronReLUClass().Alloc()
	rv := objc.Send[CNNNeuronReLU](instance.ID, objc.Sel("initWithDevice:a:"), device, a)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronReLUWithDeviceA */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronReLU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronReLU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronReLU */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronReLU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronReLU */


