// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNMultiply */


/* debug [class_header]: Header for MPSCNNMultiply */
// The class instance for the [CNNMultiply] class.
var (
	CNNMultiplyClass     _CNNMultiplyClass
	CNNMultiplyClassOnce sync.Once
)

func getCNNMultiplyClass() _CNNMultiplyClass {
	CNNMultiplyClassOnce.Do(func() {
		CNNMultiplyClass = _CNNMultiplyClass{objc.GetClass("MPSCNNMultiply")}
	})
	return CNNMultiplyClass
}

type _CNNMultiplyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNMultiply */
// An interface definition for the [CNNMultiply] class.
type ICNNMultiply interface {
	ICNNArithmetic
	
/* debug [class_interface_properties]: Properties for CNNMultiply */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNMultiply */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNMultiply */
// Alloc allocates a new instance without initialization.
func (cc _CNNMultiplyClass) Alloc() CNNMultiply {
	rv := objc.Send[CNNMultiply](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNMultiplyClass) New() CNNMultiply {
	rv := objc.Send[CNNMultiply](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiply) Init() CNNMultiply {
	rv := objc.Send[CNNMultiply](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiply) Autorelease() CNNMultiply {
	rv := objc.Send[CNNMultiply](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiply creates a new CNNMultiply instance.
func NewCNNMultiply() CNNMultiply {
	return getCNNMultiplyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNMultiply */
// A multiply operator.


// A multiply operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiply
type CNNMultiply struct {
	CNNArithmetic
}

// CNNMultiplyFrom constructs a [CNNMultiply] from an unsafe.Pointer.
//
// A multiply operator.
func CNNMultiplyFrom(ptr unsafe.Pointer) CNNMultiply {
	return CNNMultiply{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNMultiply */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiply/2942507-initwithdevice
func NewCNNMultiplyWithDevice(device unsafe.Pointer) CNNMultiply {
	instance := getCNNMultiplyClass().Alloc()
	rv := objc.Send[CNNMultiply](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNMultiplyWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNMultiply */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNMultiply */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNMultiply */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNMultiply */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNMultiply */


