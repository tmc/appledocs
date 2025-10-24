// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronSoftPlus */


/* debug [class_header]: Header for MPSCNNNeuronSoftPlus */
// The class instance for the [CNNNeuronSoftPlus] class.
var (
	CNNNeuronSoftPlusClass     _CNNNeuronSoftPlusClass
	CNNNeuronSoftPlusClassOnce sync.Once
)

func getCNNNeuronSoftPlusClass() _CNNNeuronSoftPlusClass {
	CNNNeuronSoftPlusClassOnce.Do(func() {
		CNNNeuronSoftPlusClass = _CNNNeuronSoftPlusClass{objc.GetClass("MPSCNNNeuronSoftPlus")}
	})
	return CNNNeuronSoftPlusClass
}

type _CNNNeuronSoftPlusClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronSoftPlus */
// An interface definition for the [CNNNeuronSoftPlus] class.
type ICNNNeuronSoftPlus interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronSoftPlus */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronSoftPlus */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronSoftPlus */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSoftPlusClass) Alloc() CNNNeuronSoftPlus {
	rv := objc.Send[CNNNeuronSoftPlus](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSoftPlusClass) New() CNNNeuronSoftPlus {
	rv := objc.Send[CNNNeuronSoftPlus](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSoftPlus) Init() CNNNeuronSoftPlus {
	rv := objc.Send[CNNNeuronSoftPlus](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSoftPlus) Autorelease() CNNNeuronSoftPlus {
	rv := objc.Send[CNNNeuronSoftPlus](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSoftPlus creates a new CNNNeuronSoftPlus instance.
func NewCNNNeuronSoftPlus() CNNNeuronSoftPlus {
	return getCNNNeuronSoftPlusClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronSoftPlus */
// A parametric softplus neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A parametric softplus neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftPlus
type CNNNeuronSoftPlus struct {
	CNNNeuron
}

// CNNNeuronSoftPlusFrom constructs a [CNNNeuronSoftPlus] from an unsafe.Pointer.
//
// A parametric softplus neuron filter.
func CNNNeuronSoftPlusFrom(ptr unsafe.Pointer) CNNNeuronSoftPlus {
	return CNNNeuronSoftPlus{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronSoftPlus */

// Initializes a parametric softplus neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplus/2866983-initwithdevice
func NewCNNNeuronSoftPlusWithDeviceAB(device unsafe.Pointer, a float32, b float32) CNNNeuronSoftPlus {
	instance := getCNNNeuronSoftPlusClass().Alloc()
	rv := objc.Send[CNNNeuronSoftPlus](instance.ID, objc.Sel("initWithDevice:a:b:"), device, a, b)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSoftPlusWithDeviceAB */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronSoftPlus */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronSoftPlus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronSoftPlus */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronSoftPlus */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronSoftPlus */


