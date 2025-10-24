// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNSubtractGradient */


/* debug [class_header]: Header for MPSCNNSubtractGradient */
// The class instance for the [CNNSubtractGradient] class.
var (
	CNNSubtractGradientClass     _CNNSubtractGradientClass
	CNNSubtractGradientClassOnce sync.Once
)

func getCNNSubtractGradientClass() _CNNSubtractGradientClass {
	CNNSubtractGradientClassOnce.Do(func() {
		CNNSubtractGradientClass = _CNNSubtractGradientClass{objc.GetClass("MPSCNNSubtractGradient")}
	})
	return CNNSubtractGradientClass
}

type _CNNSubtractGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNSubtractGradient */
// An interface definition for the [CNNSubtractGradient] class.
type ICNNSubtractGradient interface {
	ICNNArithmeticGradient
	
/* debug [class_interface_properties]: Properties for CNNSubtractGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNSubtractGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNSubtractGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNSubtractGradientClass) Alloc() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNSubtractGradientClass) New() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNSubtractGradient) Init() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNSubtractGradient) Autorelease() CNNSubtractGradient {
	rv := objc.Send[CNNSubtractGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNSubtractGradient creates a new CNNSubtractGradient instance.
func NewCNNSubtractGradient() CNNSubtractGradient {
	return getCNNSubtractGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNSubtractGradient */
// A gradient subtraction operator.


// A gradient subtraction operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNSubtractGradient
type CNNSubtractGradient struct {
	CNNArithmeticGradient
}

// CNNSubtractGradientFrom constructs a [CNNSubtractGradient] from an unsafe.Pointer.
//
// A gradient subtraction operator.
func CNNSubtractGradientFrom(ptr unsafe.Pointer) CNNSubtractGradient {
	return CNNSubtractGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNSubtractGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtractgradient/2956165-initwithdevice
func NewCNNSubtractGradientWithDeviceIsSecondarySourceFilter(device unsafe.Pointer, isSecondarySourceFilter bool) CNNSubtractGradient {
	instance := getCNNSubtractGradientClass().Alloc()
	rv := objc.Send[CNNSubtractGradient](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNSubtractGradientWithDeviceIsSecondarySourceFilter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNSubtractGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNSubtractGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNSubtractGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNSubtractGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNSubtractGradient */


