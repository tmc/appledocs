// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNBinaryGradientState */


/* debug [class_header]: Header for MPSNNBinaryGradientState */
// The class instance for the [BinaryGradientState] class.
var (
	BinaryGradientStateClass     _BinaryGradientStateClass
	BinaryGradientStateClassOnce sync.Once
)

func getBinaryGradientStateClass() _BinaryGradientStateClass {
	BinaryGradientStateClassOnce.Do(func() {
		BinaryGradientStateClass = _BinaryGradientStateClass{objc.GetClass("MPSNNBinaryGradientState")}
	})
	return BinaryGradientStateClass
}

type _BinaryGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BinaryGradientState */
// An interface definition for the [BinaryGradientState] class.
type IBinaryGradientState interface {
	IState
	
/* debug [class_interface_properties]: Properties for BinaryGradientState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BinaryGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BinaryGradientState */
// Alloc allocates a new instance without initialization.
func (bc _BinaryGradientStateClass) Alloc() BinaryGradientState {
	rv := objc.Send[BinaryGradientState](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BinaryGradientStateClass) New() BinaryGradientState {
	rv := objc.Send[BinaryGradientState](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryGradientState) Init() BinaryGradientState {
	rv := objc.Send[BinaryGradientState](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryGradientState) Autorelease() BinaryGradientState {
	rv := objc.Send[BinaryGradientState](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryGradientState creates a new BinaryGradientState instance.
func NewBinaryGradientState() BinaryGradientState {
	return getBinaryGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BinaryGradientState */
// A class representing the state of a gradient binary kernel when it was encoded.


// A class representing the state of a gradient binary kernel when it was encoded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNBinaryGradientState
type BinaryGradientState struct {
	State
}

// BinaryGradientStateFrom constructs a [BinaryGradientState] from an unsafe.Pointer.
//
// A class representing the state of a gradient binary kernel when it was encoded.
func BinaryGradientStateFrom(ptr unsafe.Pointer) BinaryGradientState {
	return BinaryGradientState{
		State: StateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BinaryGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BinaryGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BinaryGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BinaryGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BinaryGradientState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNBinaryGradientState */



