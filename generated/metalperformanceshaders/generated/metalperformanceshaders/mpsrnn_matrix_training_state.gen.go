// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSRNNMatrixTrainingState */


/* debug [class_header]: Header for MPSRNNMatrixTrainingState */
// The class instance for the [RNNMatrixTrainingState] class.
var (
	RNNMatrixTrainingStateClass     _RNNMatrixTrainingStateClass
	RNNMatrixTrainingStateClassOnce sync.Once
)

func getRNNMatrixTrainingStateClass() _RNNMatrixTrainingStateClass {
	RNNMatrixTrainingStateClassOnce.Do(func() {
		RNNMatrixTrainingStateClass = _RNNMatrixTrainingStateClass{objc.GetClass("MPSRNNMatrixTrainingState")}
	})
	return RNNMatrixTrainingStateClass
}

type _RNNMatrixTrainingStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RNNMatrixTrainingState */
// An interface definition for the [RNNMatrixTrainingState] class.
type IRNNMatrixTrainingState interface {
	IState
	
/* debug [class_interface_properties]: Properties for RNNMatrixTrainingState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RNNMatrixTrainingState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RNNMatrixTrainingState */
// Alloc allocates a new instance without initialization.
func (rc _RNNMatrixTrainingStateClass) Alloc() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RNNMatrixTrainingStateClass) New() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RNNMatrixTrainingState) Init() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RNNMatrixTrainingState) Autorelease() RNNMatrixTrainingState {
	rv := objc.Send[RNNMatrixTrainingState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRNNMatrixTrainingState creates a new RNNMatrixTrainingState instance.
func NewRNNMatrixTrainingState() RNNMatrixTrainingState {
	return getRNNMatrixTrainingStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RNNMatrixTrainingState */
// A class that holds data from a forward pass to be used in a backward pass.


// A class that holds data from a forward pass to be used in a backward pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRNNMatrixTrainingState
type RNNMatrixTrainingState struct {
	State
}

// RNNMatrixTrainingStateFrom constructs a [RNNMatrixTrainingState] from an unsafe.Pointer.
//
// A class that holds data from a forward pass to be used in a backward pass.
func RNNMatrixTrainingStateFrom(ptr unsafe.Pointer) RNNMatrixTrainingState {
	return RNNMatrixTrainingState{
		State: StateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RNNMatrixTrainingState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RNNMatrixTrainingState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RNNMatrixTrainingState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RNNMatrixTrainingState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RNNMatrixTrainingState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSRNNMatrixTrainingState */



