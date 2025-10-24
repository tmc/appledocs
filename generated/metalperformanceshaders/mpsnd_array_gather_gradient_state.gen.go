// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNDArrayGatherGradientState */


/* debug [class_header]: Header for MPSNDArrayGatherGradientState */
// The class instance for the [NDArrayGatherGradientState] class.
var (
	NDArrayGatherGradientStateClass     _NDArrayGatherGradientStateClass
	NDArrayGatherGradientStateClassOnce sync.Once
)

func getNDArrayGatherGradientStateClass() _NDArrayGatherGradientStateClass {
	NDArrayGatherGradientStateClassOnce.Do(func() {
		NDArrayGatherGradientStateClass = _NDArrayGatherGradientStateClass{objc.GetClass("MPSNDArrayGatherGradientState")}
	})
	return NDArrayGatherGradientStateClass
}

type _NDArrayGatherGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayGatherGradientState */
// An interface definition for the [NDArrayGatherGradientState] class.
type INDArrayGatherGradientState interface {
	INDArrayGradientState
	
/* debug [class_interface_properties]: Properties for NDArrayGatherGradientState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayGatherGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayGatherGradientState */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayGatherGradientStateClass) Alloc() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayGatherGradientStateClass) New() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayGatherGradientState) Init() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayGatherGradientState) Autorelease() NDArrayGatherGradientState {
	rv := objc.Send[NDArrayGatherGradientState](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayGatherGradientState creates a new NDArrayGatherGradientState instance.
func NewNDArrayGatherGradientState() NDArrayGatherGradientState {
	return getNDArrayGatherGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayGatherGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayGatherGradientState
type NDArrayGatherGradientState struct {
	NDArrayGradientState
}

// NDArrayGatherGradientStateFrom constructs a [NDArrayGatherGradientState] from an unsafe.Pointer.
func NDArrayGatherGradientStateFrom(ptr unsafe.Pointer) NDArrayGatherGradientState {
	return NDArrayGatherGradientState{
		NDArrayGradientState: NDArrayGradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayGatherGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayGatherGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayGatherGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayGatherGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayGatherGradientState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayGatherGradientState */



