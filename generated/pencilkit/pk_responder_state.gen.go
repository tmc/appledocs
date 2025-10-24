// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKResponderState */


/* debug [class_header]: Header for PKResponderState */
// The class instance for the [ResponderState] class.
var (
	ResponderStateClass     _ResponderStateClass
	ResponderStateClassOnce sync.Once
)

func getResponderStateClass() _ResponderStateClass {
	ResponderStateClassOnce.Do(func() {
		ResponderStateClass = _ResponderStateClass{objc.GetClass("PKResponderState")}
	})
	return ResponderStateClass
}

type _ResponderStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ResponderState */
// An interface definition for the [ResponderState] class.
type IResponderState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ResponderState */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ResponderState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ResponderState */
// Alloc allocates a new instance without initialization.
func (rc _ResponderStateClass) Alloc() ResponderState {
	rv := objc.Send[ResponderState](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ResponderStateClass) New() ResponderState {
	rv := objc.Send[ResponderState](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResponderState) Init() ResponderState {
	rv := objc.Send[ResponderState](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResponderState) Autorelease() ResponderState {
	rv := objc.Send[ResponderState](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResponderState creates a new ResponderState instance.
func NewResponderState() ResponderState {
	return getResponderStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ResponderState */
// The state of PencilKit behavior related to a .
//
// Control the behavior of responders via the property.


// The state of PencilKit behavior related to a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKResponderState
type ResponderState struct {
	objectivec.Object
}

// ResponderStateFrom constructs a [ResponderState] from an unsafe.Pointer.
//
// The state of PencilKit behavior related to a .
func ResponderStateFrom(ptr unsafe.Pointer) ResponderState {
	return ResponderState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ResponderState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ResponderState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ResponderState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ResponderState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ResponderState */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKResponderState */


