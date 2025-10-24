// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNArithmeticGradientState */


/* debug [class_header]: Header for MPSCNNArithmeticGradientState */
// The class instance for the [CNNArithmeticGradientState] class.
var (
	CNNArithmeticGradientStateClass     _CNNArithmeticGradientStateClass
	CNNArithmeticGradientStateClassOnce sync.Once
)

func getCNNArithmeticGradientStateClass() _CNNArithmeticGradientStateClass {
	CNNArithmeticGradientStateClassOnce.Do(func() {
		CNNArithmeticGradientStateClass = _CNNArithmeticGradientStateClass{objc.GetClass("MPSCNNArithmeticGradientState")}
	})
	return CNNArithmeticGradientStateClass
}

type _CNNArithmeticGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNArithmeticGradientState */
// An interface definition for the [CNNArithmeticGradientState] class.
type ICNNArithmeticGradientState interface {
	IBinaryGradientState
	
/* debug [class_interface_properties]: Properties for CNNArithmeticGradientState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNArithmeticGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNArithmeticGradientState */
// Alloc allocates a new instance without initialization.
func (cc _CNNArithmeticGradientStateClass) Alloc() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNArithmeticGradientStateClass) New() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNArithmeticGradientState) Init() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNArithmeticGradientState) Autorelease() CNNArithmeticGradientState {
	rv := objc.Send[CNNArithmeticGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNArithmeticGradientState creates a new CNNArithmeticGradientState instance.
func NewCNNArithmeticGradientState() CNNArithmeticGradientState {
	return getCNNArithmeticGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNArithmeticGradientState */
// An object that stores the clamp mask used by gradient arithmetic operators.


// An object that stores the clamp mask used by gradient arithmetic operators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNArithmeticGradientState
type CNNArithmeticGradientState struct {
	BinaryGradientState
}

// CNNArithmeticGradientStateFrom constructs a [CNNArithmeticGradientState] from an unsafe.Pointer.
//
// An object that stores the clamp mask used by gradient arithmetic operators.
func CNNArithmeticGradientStateFrom(ptr unsafe.Pointer) CNNArithmeticGradientState {
	return CNNArithmeticGradientState{
		BinaryGradientState: BinaryGradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNArithmeticGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNArithmeticGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNArithmeticGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNArithmeticGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNArithmeticGradientState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNArithmeticGradientState */



