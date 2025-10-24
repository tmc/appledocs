// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNNeuronSoftSign */


/* debug [class_header]: Header for MPSCNNNeuronSoftSign */
// The class instance for the [CNNNeuronSoftSign] class.
var (
	CNNNeuronSoftSignClass     _CNNNeuronSoftSignClass
	CNNNeuronSoftSignClassOnce sync.Once
)

func getCNNNeuronSoftSignClass() _CNNNeuronSoftSignClass {
	CNNNeuronSoftSignClassOnce.Do(func() {
		CNNNeuronSoftSignClass = _CNNNeuronSoftSignClass{objc.GetClass("MPSCNNNeuronSoftSign")}
	})
	return CNNNeuronSoftSignClass
}

type _CNNNeuronSoftSignClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNNeuronSoftSign */
// An interface definition for the [CNNNeuronSoftSign] class.
type ICNNNeuronSoftSign interface {
	ICNNNeuron
	
/* debug [class_interface_properties]: Properties for CNNNeuronSoftSign */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNNeuronSoftSign */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNNeuronSoftSign */
// Alloc allocates a new instance without initialization.
func (cc _CNNNeuronSoftSignClass) Alloc() CNNNeuronSoftSign {
	rv := objc.Send[CNNNeuronSoftSign](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNNeuronSoftSignClass) New() CNNNeuronSoftSign {
	rv := objc.Send[CNNNeuronSoftSign](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNNeuronSoftSign) Init() CNNNeuronSoftSign {
	rv := objc.Send[CNNNeuronSoftSign](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNNeuronSoftSign) Autorelease() CNNNeuronSoftSign {
	rv := objc.Send[CNNNeuronSoftSign](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNNeuronSoftSign creates a new CNNNeuronSoftSign instance.
func NewCNNNeuronSoftSign() CNNNeuronSoftSign {
	return getCNNNeuronSoftSignClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNNeuronSoftSign */
// A softsign neuron filter.
//
// For each pixel in an image, the filter applies the following function:


// A softsign neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNNeuronSoftSign
type CNNNeuronSoftSign struct {
	CNNNeuron
}

// CNNNeuronSoftSignFrom constructs a [CNNNeuronSoftSign] from an unsafe.Pointer.
//
// A softsign neuron filter.
func CNNNeuronSoftSignFrom(ptr unsafe.Pointer) CNNNeuronSoftSign {
	return CNNNeuronSoftSign{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNNeuronSoftSign */

// Initializes a softsign neuron filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsign/2866994-initwithdevice
func NewCNNNeuronSoftSignWithDevice(device unsafe.Pointer) CNNNeuronSoftSign {
	instance := getCNNNeuronSoftSignClass().Alloc()
	rv := objc.Send[CNNNeuronSoftSign](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNNeuronSoftSignWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNNeuronSoftSign */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNNeuronSoftSign */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNNeuronSoftSign */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNNeuronSoftSign */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNNeuronSoftSign */


