// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronPReLU */


/* debug [class_header]: Header for MPSCNNNeuronPReLU */
// The class instance for the [CNNNeuronPReLU] class.
var (
	CNNNeuronPReLUClass     _CNNNeuronPReLUClass
	CNNNeuronPReLUClassOnce sync.Once
)

func getCNNNeuronPReLUClass() _CNNNeuronPReLUClass {
	CNNNeuronPReLUClassOnce.Do(func() {
		CNNNeuronPReLUClass = _CNNNeuronPReLUClass{objc.GetClass("MPSCNNNeuronPReLU")}
	})
	return CNNNeuronPReLUClass
}

type _CNNNeuronPReLUClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronPReLU */
// An interface definition for the [CNNNeuronPReLU] class.
type ICNNNeuronPReLU interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronPReLU */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronPReLU */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronPReLU */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronPReLUClass) Alloc() CNNNeuronPReLU {
	rv := objc.Send[CNNNeuronPReLU](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronPReLUClass) New() CNNNeuronPReLU {
	rv := objc.Send[CNNNeuronPReLU](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronPReLU) Init() CNNNeuronPReLU {
	rv := objc.Send[CNNNeuronPReLU](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronPReLU) Autorelease() CNNNeuronPReLU {
	rv := objc.Send[CNNNeuronPReLU](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronPReLU creates a new CNNNeuronPReLU instance.
func NewCNNNeuronPReLU() CNNNeuronPReLU {
	return getCNNNeuronPReLUClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronPReLU */
// A parametric ReLU (Rectified Linear Unit) neuron filter.
//
// For each pixel in an image, the filter applies the following function: Where in . That is, parameters are learned and applied to each channel separately. Compare this to where parameter is shared across all channels.


// A parametric ReLU (Rectified Linear Unit) neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronPReLU
type CNNNeuronPReLU struct {
	CNNNeuron
}

// CNNNeuronPReLUFrom constructs a [CNNNeuronPReLU] from an unsafe.Pointer.
//
// A parametric ReLU (Rectified Linear Unit) neuron filter.
func CNNNeuronPReLUFrom(ptr unsafe.Pointer) CNNNeuronPReLU {
	return CNNNeuronPReLU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronPReLU */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelu/2921661-initwithdevice
func NewCNNNeuronPReLUWithDeviceACount(device unsafe.Pointer, a objectivec.IObject, count uint) CNNNeuronPReLU {
	instance := getCNNNeuronPReLUClass().Alloc()
	rv := objc.Send[CNNNeuronPReLU](instance.ID, objc.Sel("initWithDevice:a:count:"), device, a, count)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronPReLUWithDeviceACount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronPReLU */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronPReLU */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronPReLU */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronPReLU */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronPReLU */


