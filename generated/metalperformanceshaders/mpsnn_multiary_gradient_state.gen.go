// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNMultiaryGradientState */


/* debug [class_header]: Header for MPSNNMultiaryGradientState */
// The class instance for the [MultiaryGradientState] class.
var (
	MultiaryGradientStateClass     _MultiaryGradientStateClass
	MultiaryGradientStateClassOnce sync.Once
)

func getMultiaryGradientStateClass() _MultiaryGradientStateClass {
	MultiaryGradientStateClassOnce.Do(func() {
		MultiaryGradientStateClass = _MultiaryGradientStateClass{objc.GetClass("MPSNNMultiaryGradientState")}
	})
	return MultiaryGradientStateClass
}

type _MultiaryGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MultiaryGradientState */
// An interface definition for the [MultiaryGradientState] class.
type IMultiaryGradientState interface {
	IState
	
/* debug [class_interface_properties]: Properties for MultiaryGradientState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MultiaryGradientState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MultiaryGradientState */
// Alloc allocates a new instance without initialization.
func (mc _MultiaryGradientStateClass) Alloc() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MultiaryGradientStateClass) New() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MultiaryGradientState) Init() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MultiaryGradientState) Autorelease() MultiaryGradientState {
	rv := objc.Send[MultiaryGradientState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMultiaryGradientState creates a new MultiaryGradientState instance.
func NewMultiaryGradientState() MultiaryGradientState {
	return getMultiaryGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MultiaryGradientState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNMultiaryGradientState
type MultiaryGradientState struct {
	State
}

// MultiaryGradientStateFrom constructs a [MultiaryGradientState] from an unsafe.Pointer.
func MultiaryGradientStateFrom(ptr unsafe.Pointer) MultiaryGradientState {
	return MultiaryGradientState{
		State: StateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MultiaryGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MultiaryGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MultiaryGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MultiaryGradientState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MultiaryGradientState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNMultiaryGradientState */



