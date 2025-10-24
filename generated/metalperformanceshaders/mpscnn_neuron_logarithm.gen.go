// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronLogarithm */


/* debug [class_header]: Header for MPSCNNNeuronLogarithm */
// The class instance for the [CNNNeuronLogarithm] class.
var (
	CNNNeuronLogarithmClass     _CNNNeuronLogarithmClass
	CNNNeuronLogarithmClassOnce sync.Once
)

func getCNNNeuronLogarithmClass() _CNNNeuronLogarithmClass {
	CNNNeuronLogarithmClassOnce.Do(func() {
		CNNNeuronLogarithmClass = _CNNNeuronLogarithmClass{objc.GetClass("MPSCNNNeuronLogarithm")}
	})
	return CNNNeuronLogarithmClass
}

type _CNNNeuronLogarithmClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronLogarithm */
// An interface definition for the [CNNNeuronLogarithm] class.
type ICNNNeuronLogarithm interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronLogarithm */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronLogarithm */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronLogarithm */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronLogarithmClass) Alloc() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronLogarithmClass) New() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronLogarithm) Init() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronLogarithm) Autorelease() CNNNeuronLogarithm {
	rv := objc.Send[CNNNeuronLogarithm](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronLogarithm creates a new CNNNeuronLogarithm instance.
func NewCNNNeuronLogarithm() CNNNeuronLogarithm {
	return getCNNNeuronLogarithmClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronLogarithm */
// A logarithm neuron filter.


// A logarithm neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronLogarithm
type CNNNeuronLogarithm struct {
	CNNNeuron
}

// CNNNeuronLogarithmFrom constructs a [CNNNeuronLogarithm] from an unsafe.Pointer.
//
// A logarithm neuron filter.
func CNNNeuronLogarithmFrom(ptr unsafe.Pointer) CNNNeuronLogarithm {
	return CNNNeuronLogarithm{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronLogarithm */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithm/2951871-initwithdevice
func NewCNNNeuronLogarithmWithDeviceABC(device unsafe.Pointer, a float32, b float32, c float32) CNNNeuronLogarithm {
	instance := getCNNNeuronLogarithmClass().Alloc()
	rv := objc.Send[CNNNeuronLogarithm](instance.ID, objc.Sel("initWithDevice:a:b:c:"), device, a, b, c)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronLogarithmWithDeviceABC */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronLogarithm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronLogarithm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronLogarithm */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronLogarithm */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronLogarithm */


