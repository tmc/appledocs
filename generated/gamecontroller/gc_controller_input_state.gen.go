// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCControllerInputState */


/* debug [class_header]: Header for GCControllerInputState */
// The class instance for the [GCControllerInputState] class.
var (
	GCControllerInputStateClass     _GCControllerInputStateClass
	GCControllerInputStateClassOnce sync.Once
)

func getGCControllerInputStateClass() _GCControllerInputStateClass {
	GCControllerInputStateClassOnce.Do(func() {
		GCControllerInputStateClass = _GCControllerInputStateClass{objc.GetClass("GCControllerInputState")}
	})
	return GCControllerInputStateClass
}

type _GCControllerInputStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerInputState */
// An interface definition for the [GCControllerInputState] class.
type IGCControllerInputState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCControllerInputState */
	// properties:
	Input() IGCControllerLiveInput
	SetInput(value IGCControllerLiveInput)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerInputState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerInputState */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerInputStateClass) Alloc() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerInputStateClass) New() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerInputState) Init() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerInputState) Autorelease() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerInputState creates a new GCControllerInputState instance.
func NewGCControllerInputState() GCControllerInputState {
	return getGCControllerInputStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerInputState */
// A class that represents an input state for gamepads and arcade sticks.
//
// This class implements the protocol for gamepads and arcade sticks. Instances of this class represent the state of the controller’s inputs at a moment in time, which can be the current time.


// A class that represents an input state for gamepads and arcade sticks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerInputState
type GCControllerInputState struct {
	objectivec.Object
}

// GCControllerInputStateFrom constructs a [GCControllerInputState] from an unsafe.Pointer.
//
// A class that represents an input state for gamepads and arcade sticks.
func GCControllerInputStateFrom(ptr unsafe.Pointer) GCControllerInputState {
	return GCControllerInputState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerInputState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerInputState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerInputState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerInputState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerInputState */

// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerInputState) Input() IGCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("input"))
	return rv
}/* debug [instance_properties/getter]: input */


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCControllerInputState) SetInput(value IGCControllerLiveInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInput:"), value)
}/* debug [instance_properties/setter]: input */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerInputState */



