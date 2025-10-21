// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StateMachine] class.
var (
	StateMachineClass     _StateMachineClass
	StateMachineClassOnce sync.Once
)

func getStateMachineClass() _StateMachineClass {
	StateMachineClassOnce.Do(func() {
		StateMachineClass = _StateMachineClass{objc.GetClass("GKStateMachine")}
	})
	return StateMachineClass
}

type _StateMachineClass struct {
	class objc.Class
}

// An interface definition for the [StateMachine] class.
type IStateMachine interface {
	objectivec.IObject
	CanEnterState(stateClass objc.Class) bool
	EnterState(stateClass objc.Class) bool
	StateForClass(stateClass objc.Class) unsafe.Pointer
	UpdateWithDeltaTime(sec foundation.TimeInterval)
}

// A finite-state machine—a collection of state objects that each define logic for a particular state of gameplay and rules for transitioning between states.
//
// In GameplayKit, you subclass to define each state and the rules for allowed transitions between states, and use a instance to manage a machine that combines several states. This system provides a way to organize code in your game by organizing state-dependent actions into methods that run when entering a state, when exiting a state, and periodically while in a state (for example, on every animation frame your game renders). You can use state machines to govern various aspects of a game. For example: An enemy character might use a state machine with Chase, Flee, Dead, and Respawn states, each of which drives the enemy’s behavior, with state transitions determined by player actions and elapsed time. An automated turret might use a state machine with Ready, Firing, and Cooldown states, controlling when it seeks out nearby targets and how often it fires. A game user interface might use Menu, Playing, Paused, and GameOver states, each of which determines what UI elements are shown and what other game elements are running. To build a state machine, first define a distinct subclass of for each possible state of the machine. In each state class, the method determines which other state classes the machine may transition into from that state. Then, create a state machine object by constructing instances of the state classes and passing them to one of the methods listed in Creating a State Machine below. Finally, set the machine in motion by choosing an initial state for it to enter with the method. To define state-dependent behavior, override the , , and methods in each subclass. The state machine notifies the current state whenever a state change happens. Use the and methods to perform actions in response to a state change. For example, an enemy character entering the Flee state might change its appearance to indicate that is has become vulnerable to attack by the player. When you call a state machine’s method, the state machine calls the method of its current state. Use this method to organize per-frame update code by state. For example, an enemy character in the Chase state can update its position to pursue the player, and an enemy in the Flee state can update its position to evade the player. For more information about state machines, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine
type StateMachine struct {
	objectivec.Object
}

// StateMachineFrom constructs a [StateMachine] from an unsafe.Pointer.
//
// A finite-state machine—a collection of state objects that each define logic for a particular state of gameplay and rules for transitioning between states.
func StateMachineFrom(ptr unsafe.Pointer) StateMachine {
	return StateMachine{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StateMachineClass) Alloc() StateMachine {
	rv := objc.Send[StateMachine](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StateMachineClass) New() StateMachine {
	rv := objc.Send[StateMachine](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateMachine) Init() StateMachine {
	rv := objc.Send[StateMachine](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateMachine) Autorelease() StateMachine {
	rv := objc.Send[StateMachine](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateMachine creates a new StateMachine instance.
func NewStateMachine() StateMachine {
	return getStateMachineClass().New()
}




// Initializes a state machine with the specified states.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/init(states:)
func NewStateMachineWithStates(states unsafe.Pointer) StateMachine {
	instance := getStateMachineClass().Alloc()
	rv := objc.Send[StateMachine](instance.ID, objc.Sel("initWithStates:"), states)
	rv.Autorelease()
	return rv
}


// Creates a state machine with the specified states.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/stateMachineWithStates:
func (sc _StateMachineClass) StateMachineWithStates(states unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("stateMachineWithStates:"), states)
	return rv
}

// Returns a Boolean value indicating whether it is valid for the state machine to transition from its current state to a state of the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/canEnterState(_:)
func (s_ StateMachine) CanEnterState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canEnterState:"), stateClass)
	return rv
}

// Attempts to transition the state machine from its current state to a state of the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/enter(_:)
func (s_ StateMachine) EnterState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enterState:"), stateClass)
	return rv
}

// Returns the state object in the state machine corresponding to the specified class.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/stateForClass:
func (s_ StateMachine) StateForClass(stateClass objc.Class) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("stateForClass:"), stateClass)
	return rv
}

// Tells the current state object to perform per-frame updates.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/update(deltaTime:)
func (s_ StateMachine) UpdateWithDeltaTime(sec foundation.TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateWithDeltaTime:"), sec)
}

// The state machine’s current state.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/currentState
func (s_ StateMachine) CurrentState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentState"))
	return rv
}


