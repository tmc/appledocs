// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSCNNDropoutGradientState */


/* debug [class_header]: Header for MPSCNNDropoutGradientState */
// The class instance for the [CNNDropoutGradientState] class.
var (
	CNNDropoutGradientStateClass     _CNNDropoutGradientStateClass
	CNNDropoutGradientStateClassOnce sync.Once
)

func getCNNDropoutGradientStateClass() _CNNDropoutGradientStateClass {
	CNNDropoutGradientStateClassOnce.Do(func() {
		CNNDropoutGradientStateClass = _CNNDropoutGradientStateClass{objc.GetClass("MPSCNNDropoutGradientState")}
	})
	return CNNDropoutGradientStateClass
}

type _CNNDropoutGradientStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNNDropoutGradientState */
// An interface definition for the [CNNDropoutGradientState] class.
type ICNNDropoutGradientState interface {
	IGradientState
	
/* debug [class_interface_properties]: Properties for CNNDropoutGradientState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNNDropoutGradientState */
	// methods:
	MaskData()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNNDropoutGradientState */
// Alloc allocates a new instance without initialization.
func (cc _CNNDropoutGradientStateClass) Alloc() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNDropoutGradientStateClass) New() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNDropoutGradientState) Init() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNDropoutGradientState) Autorelease() CNNDropoutGradientState {
	rv := objc.Send[CNNDropoutGradientState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNDropoutGradientState creates a new CNNDropoutGradientState instance.
func NewCNNDropoutGradientState() CNNDropoutGradientState {
	return getCNNDropoutGradientStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNNDropoutGradientState */
// A class that stores the mask used by dropout and gradient dropout filters.


// A class that stores the mask used by dropout and gradient dropout filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNDropoutGradientState
type CNNDropoutGradientState struct {
	GradientState
}

// CNNDropoutGradientStateFrom constructs a [CNNDropoutGradientState] from an unsafe.Pointer.
//
// A class that stores the mask used by dropout and gradient dropout filters.
func CNNDropoutGradientStateFrom(ptr unsafe.Pointer) CNNDropoutGradientState {
	return CNNDropoutGradientState{
		GradientState: GradientStateFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNNDropoutGradientState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNNDropoutGradientState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNNDropoutGradientState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNNDropoutGradientState */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate/2942527-maskdata
func (c_ CNNDropoutGradientState) MaskData() {
	objc.Send[objc.ID](c_.ID, objc.Sel("maskData"))
}/* debug [instance_methods/method]: MaskData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNNDropoutGradientState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSCNNDropoutGradientState */



