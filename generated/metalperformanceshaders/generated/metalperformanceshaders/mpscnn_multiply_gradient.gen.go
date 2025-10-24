// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSCNNMultiplyGradient */


/* debug [class_header]: Header for MPSCNNMultiplyGradient */
// The class instance for the [CNNMultiplyGradient] class.
var (
	CNNMultiplyGradientClass     _CNNMultiplyGradientClass
	CNNMultiplyGradientClassOnce sync.Once
)

func getCNNMultiplyGradientClass() _CNNMultiplyGradientClass {
	CNNMultiplyGradientClassOnce.Do(func() {
		CNNMultiplyGradientClass = _CNNMultiplyGradientClass{objc.GetClass("MPSCNNMultiplyGradient")}
	})
	return CNNMultiplyGradientClass
}

type _CNNMultiplyGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNMultiplyGradient */
// An interface definition for the [CNNMultiplyGradient] class.
type ICNNMultiplyGradient interface {
	ICNNArithmeticGradient
	
/* debug [class_interface_properties]: Properties for CNNMultiplyGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNMultiplyGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNMultiplyGradient */
// Alloc allocates a new instance without initialization.
func (cc _CNNMultiplyGradientClass) Alloc() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNMultiplyGradientClass) New() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNMultiplyGradient) Init() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNMultiplyGradient) Autorelease() CNNMultiplyGradient {
	rv := objc.Send[CNNMultiplyGradient](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNMultiplyGradient creates a new CNNMultiplyGradient instance.
func NewCNNMultiplyGradient() CNNMultiplyGradient {
	return getCNNMultiplyGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNMultiplyGradient */
// A gradient multiply operator.


// A gradient multiply operator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNMultiplyGradient
type CNNMultiplyGradient struct {
	CNNArithmeticGradient
}

// CNNMultiplyGradientFrom constructs a [CNNMultiplyGradient] from an unsafe.Pointer.
//
// A gradient multiply operator.
func CNNMultiplyGradientFrom(ptr unsafe.Pointer) CNNMultiplyGradient {
	return CNNMultiplyGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNMultiplyGradient */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiplygradient/2956164-initwithdevice
func NewCNNMultiplyGradientWithDeviceIsSecondarySourceFilter(device unsafe.Pointer, isSecondarySourceFilter bool) CNNMultiplyGradient {
	instance := getCNNMultiplyGradientClass().Alloc()
	rv := objc.Send[CNNMultiplyGradient](instance.ID, objc.Sel("initWithDevice:isSecondarySourceFilter:"), device, isSecondarySourceFilter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNNMultiplyGradientWithDeviceIsSecondarySourceFilter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNMultiplyGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNMultiplyGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNMultiplyGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNMultiplyGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNMultiplyGradient */


