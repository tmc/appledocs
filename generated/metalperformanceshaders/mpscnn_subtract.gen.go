// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSubtract */


/* debug [class_header]: Header for MPSCNNSubtract */
// The class instance for the [CNNSubtract] class.
var (
	CNNSubtractClass     _CNNSubtractClass
	CNNSubtractClassOnce sync.Once
)

func getCNNSubtractClass() _CNNSubtractClass {
	CNNSubtractClassOnce.Do(func() {
		CNNSubtractClass = _CNNSubtractClass{objc.GetClass("MPSCNNSubtract")}
	})
	return CNNSubtractClass
}

type _CNNSubtractClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSubtract */
// An interface definition for the [CNNSubtract] class.
type ICNNSubtract interface {
	ICNNArithmetic
	
/* debug [class_interface_properties]: Properties for CNNSubtract */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSubtract */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSubtract */
// Alloc allocates a new instance without initialization.
func (cc _CNNSubtractClass) Alloc() CNNSubtract {
	rv := objc.Send[CNNSubtract](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSubtractClass) New() CNNSubtract {
	rv := objc.Send[CNNSubtract](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSubtract) Init() CNNSubtract {
	rv := objc.Send[CNNSubtract](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSubtract) Autorelease() CNNSubtract {
	rv := objc.Send[CNNSubtract](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSubtract creates a new CNNSubtract instance.
func NewCNNSubtract() CNNSubtract {
	return getCNNSubtractClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSubtract */
// A subtraction operator.


// A subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtract
type CNNSubtract struct {
	CNNArithmetic
}

// CNNSubtractFrom constructs a [CNNSubtract] from an unsafe.Pointer.
//
// A subtraction operator.
func CNNSubtractFrom(ptr unsafe.Pointer) CNNSubtract {
	return CNNSubtract{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSubtract */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtract/2942503-initwithdevice
func NewCNNSubtractWithDevice(device unsafe.Pointer) CNNSubtract {
	instance := getCNNSubtractClass().Alloc()
	rv := objc.Send[CNNSubtract](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSubtractWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSubtract */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSubtract */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSubtract */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSubtract */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSubtract */


