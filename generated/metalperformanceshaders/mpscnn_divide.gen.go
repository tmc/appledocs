// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNDivide */


/* debug [class_header]: Header for MPSCNNDivide */
// The class instance for the [CNNDivide] class.
var (
	CNNDivideClass     _CNNDivideClass
	CNNDivideClassOnce sync.Once
)

func getCNNDivideClass() _CNNDivideClass {
	CNNDivideClassOnce.Do(func() {
		CNNDivideClass = _CNNDivideClass{objc.GetClass("MPSCNNDivide")}
	})
	return CNNDivideClass
}

type _CNNDivideClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDivide */
// An interface definition for the [CNNDivide] class.
type ICNNDivide interface {
	ICNNArithmetic
	
/* debug [class_interface_properties]: Properties for CNNDivide */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDivide */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDivide */
// Alloc allocates a new instance without initialization.
func (cc _CNNDivideClass) Alloc() CNNDivide {
	rv := objc.Send[CNNDivide](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDivideClass) New() CNNDivide {
	rv := objc.Send[CNNDivide](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDivide) Init() CNNDivide {
	rv := objc.Send[CNNDivide](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDivide) Autorelease() CNNDivide {
	rv := objc.Send[CNNDivide](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDivide creates a new CNNDivide instance.
func NewCNNDivide() CNNDivide {
	return getCNNDivideClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDivide */
// A division operator.


// A division operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDivide
type CNNDivide struct {
	CNNArithmetic
}

// CNNDivideFrom constructs a [CNNDivide] from an unsafe.Pointer.
//
// A division operator.
func CNNDivideFrom(ptr unsafe.Pointer) CNNDivide {
	return CNNDivide{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDivide */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndivide/2942508-initwithdevice
func NewCNNDivideWithDevice(device unsafe.Pointer) CNNDivide {
	instance := getCNNDivideClass().Alloc()
	rv := objc.Send[CNNDivide](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNDivideWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDivide */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDivide */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDivide */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDivide */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDivide */


