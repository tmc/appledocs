// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKState */


/* debug [class_header]: Header for GKState */
// The class instance for the [State] class.
var (
	StateClass     _StateClass
	StateClassOnce sync.Once
)

func getStateClass() _StateClass {
	StateClassOnce.Do(func() {
		StateClass = _StateClass{objc.GetClass("GKState")}
	})
	return StateClass
}

type _StateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for State */
// An interface definition for the [State] class.
type IState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for State */
	// properties:
	StateMachine() IGKStateMachine
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for State */
	// methods:
	DidEnterWithPreviousState(previousState IGKState)
	IsValidNextState(stateClass objc.Class) bool
	UpdateWithDeltaTime(seconds float64)
	WillExitWithNextState(nextState IGKState)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for State */
// Alloc allocates a new instance without initialization.
func (sc _StateClass) Alloc() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StateClass) New() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ State) Init() State {
	rv := objc.Send[State](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ State) Autorelease() State {
	rv := objc.Send[State](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewState creates a new State instance.
func NewState() State {
	return getStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for State */
// The abstract superclass for defining state-specific logic as part of a state machine.
//
// The abstract class defines the features of state classes to be used with a state machine (a object). You build a state machine by defining a separate subclass for each state of the machine. In each state class, you use the method to define which other states are valid for a machine to transition into. State classes provide a place to put state-dependent game logic, such as actions that should happen when entering or exiting a specific state, or per-frame update code that is valid only when in a specific state. For more information about state machines, read in .


// The abstract superclass for defining state-specific logic as part of a state machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState
type State struct {
	objectivec.Object
}

// StateFrom constructs a [State] from an unsafe.Pointer.
//
// The abstract superclass for defining state-specific logic as part of a state machine.
func StateFrom(ptr unsafe.Pointer) State {
	return State{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for State */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for State */

// Creates a state object with the specified list of valid next states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/state
func (sc _StateClass) State() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("state"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=State) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for State */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for State */

// Performs custom actions when a state machine transitions into this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/didEnter(from:)
func (s_ State) DidEnterWithPreviousState(previousState IGKState) {
	objc.Send[objc.ID](s_.ID, objc.Sel("didEnterWithPreviousState:"), previousState)
}/* debug [instance_methods/method]: DidEnterWithPreviousState */


// Returns a Boolean value indicating whether a state machine currently in this state is allowed to transition into the specified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/isValidNextState(_:)
func (s_ State) IsValidNextState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isValidNextState:"), stateClass)
	return rv
}/* debug [instance_methods/method]: IsValidNextState */


// Performs custom actions when a state machine updates while in this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/update(deltaTime:)
func (s_ State) UpdateWithDeltaTime(seconds float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}/* debug [instance_methods/method]: UpdateWithDeltaTime */


// Performs custom actions when a state machine transitions out of this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/willExit(to:)
func (s_ State) WillExitWithNextState(nextState IGKState) {
	objc.Send[objc.ID](s_.ID, objc.Sel("willExitWithNextState:"), nextState)
}/* debug [instance_methods/method]: WillExitWithNextState */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for State */

// The state machine that owns this state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/stateMachine
func (s_ State) StateMachine() IGKStateMachine {
	rv := objc.Send[StateMachine](s_.ID, objc.Sel("stateMachine"))
	return rv
}/* debug [instance_properties/getter]: stateMachine */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKState */


