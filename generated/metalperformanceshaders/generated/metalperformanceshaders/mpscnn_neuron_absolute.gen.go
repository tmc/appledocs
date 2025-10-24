// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronAbsolute */


/* debug [class_header]: Header for MPSCNNNeuronAbsolute */
// The class instance for the [CNNNeuronAbsolute] class.
var (
	CNNNeuronAbsoluteClass     _CNNNeuronAbsoluteClass
	CNNNeuronAbsoluteClassOnce sync.Once
)

func getCNNNeuronAbsoluteClass() _CNNNeuronAbsoluteClass {
	CNNNeuronAbsoluteClassOnce.Do(func() {
		CNNNeuronAbsoluteClass = _CNNNeuronAbsoluteClass{objc.GetClass("MPSCNNNeuronAbsolute")}
	})
	return CNNNeuronAbsoluteClass
}

type _CNNNeuronAbsoluteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronAbsolute */
// An interface definition for the [CNNNeuronAbsolute] class.
type ICNNNeuronAbsolute interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronAbsolute */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronAbsolute */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronAbsolute */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronAbsoluteClass) Alloc() CNNNeuronAbsolute {
	rv := objc.Send[CNNNeuronAbsolute](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronAbsoluteClass) New() CNNNeuronAbsolute {
	rv := objc.Send[CNNNeuronAbsolute](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronAbsolute) Init() CNNNeuronAbsolute {
	rv := objc.Send[CNNNeuronAbsolute](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronAbsolute) Autorelease() CNNNeuronAbsolute {
	rv := objc.Send[CNNNeuronAbsolute](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronAbsolute creates a new CNNNeuronAbsolute instance.
func NewCNNNeuronAbsolute() CNNNeuronAbsolute {
	return getCNNNeuronAbsoluteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronAbsolute */
// An absolute neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// An absolute neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsolute
type CNNNeuronAbsolute struct {
	CNNNeuron
}

// CNNNeuronAbsoluteFrom constructs a [CNNNeuronAbsolute] from an unsafe.Pointer.
//
// An absolute neuron filter.
func CNNNeuronAbsoluteFrom(ptr unsafe.Pointer) CNNNeuronAbsolute {
	return CNNNeuronAbsolute{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronAbsolute */

// Initializes an absolute neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronAbsolute/init(device:)
func NewCNNNeuronAbsoluteWithDevice(device unsafe.Pointer) CNNNeuronAbsolute {
	instance := getCNNNeuronAbsoluteClass().Alloc()
	rv := objc.Send[CNNNeuronAbsolute](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronAbsoluteWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronAbsolute */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronAbsolute */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronAbsolute */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronAbsolute */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronAbsolute */


