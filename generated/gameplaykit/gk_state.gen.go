// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [State] class.
type IState interface {
	objectivec.IObject
	// properties:
	StateMachine() IGKStateMachine
	// methods:
	DidEnterWithPreviousState(previousState IGKState)
	IsValidNextState(stateClass objc.Class) bool
	UpdateWithDeltaTime(seconds float64)
	WillExitWithNextState(nextState IGKState)
}

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

// Alloc allocates a new instance without initialization.
func (sc _StateClass) Alloc() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a state object with the specified list of valid next states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/state
func (sc _StateClass) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("state"))
	return rv
}


// Performs custom actions when a state machine transitions into this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/didEnter(from:)
func (s_ State) DidEnterWithPreviousState(previousState IGKState) {
	objc.Send[objc.ID](s_.ID, objc.Sel("didEnterWithPreviousState:"), previousState)
}


// Returns a Boolean value indicating whether a state machine currently in this state is allowed to transition into the specified state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/isValidNextState(_:)
func (s_ State) IsValidNextState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isValidNextState:"), stateClass)
	return rv
}


// Performs custom actions when a state machine updates while in this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/update(deltaTime:)
func (s_ State) UpdateWithDeltaTime(seconds float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}


// Performs custom actions when a state machine transitions out of this state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/willExit(to:)
func (s_ State) WillExitWithNextState(nextState IGKState) {
	objc.Send[objc.ID](s_.ID, objc.Sel("willExitWithNextState:"), nextState)
}


// The state machine that owns this state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKState/stateMachine
func (s_ State) StateMachine() IGKStateMachine {
	rv := objc.Send[StateMachine](s_.ID, objc.Sel("stateMachine"))
	return rv
}


