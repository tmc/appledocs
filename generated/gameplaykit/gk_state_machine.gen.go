// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKStateMachine */


/* debug [class_header]: Header for GKStateMachine */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StateMachine */
// An interface definition for the [StateMachine] class.
type IStateMachine interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StateMachine */
	// properties:
	CurrentState() IGKState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StateMachine */
	// methods:
	CanEnterState(stateClass objc.Class) bool
	EnterState(stateClass objc.Class) bool
	StateForClass(stateClass objc.Class) IState
	UpdateWithDeltaTime(sec float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StateMachine */
// Alloc allocates a new instance without initialization.
func (sc _StateMachineClass) Alloc() StateMachine {
	rv := objc.Send[StateMachine](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StateMachine */
// A finite-state machine—a collection of state objects that each define logic for a particular state of gameplay and rules for transitioning between states.
//
// In GameplayKit, you subclass to define each state and the rules for allowed transitions between states, and use a instance to manage a machine that combines several states. This system provides a way to organize code in your game by organizing state-dependent actions into methods that run when entering a state, when exiting a state, and periodically while in a state (for example, on every animation frame your game renders). You can use state machines to govern various aspects of a game. For example: An enemy character might use a state machine with Chase, Flee, Dead, and Respawn states, each of which drives the enemy’s behavior, with state transitions determined by player actions and elapsed time. An automated turret might use a state machine with Ready, Firing, and Cooldown states, controlling when it seeks out nearby targets and how often it fires. A game user interface might use Menu, Playing, Paused, and GameOver states, each of which determines what UI elements are shown and what other game elements are running. To build a state machine, first define a distinct subclass of for each possible state of the machine. In each state class, the method determines which other state classes the machine may transition into from that state. Then, create a state machine object by constructing instances of the state classes and passing them to one of the methods listed in Creating a State Machine below. Finally, set the machine in motion by choosing an initial state for it to enter with the method. To define state-dependent behavior, override the , , and methods in each subclass. The state machine notifies the current state whenever a state change happens. Use the and methods to perform actions in response to a state change. For example, an enemy character entering the Flee state might change its appearance to indicate that is has become vulnerable to attack by the player. When you call a state machine’s method, the state machine calls the method of its current state. Use this method to organize per-frame update code by state. For example, an enemy character in the Chase state can update its position to pursue the player, and an enemy in the Flee state can update its position to evade the player. For more information about state machines, read in .


// A finite-state machine—a collection of state objects that each define logic for a particular state of gameplay and rules for transitioning between states.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StateMachine */

// Initializes a state machine with the specified states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/init(states:)
func NewStateMachineWithStates(states []State) StateMachine {
	instance := getStateMachineClass().Alloc()
	rv := objc.Send[StateMachine](instance.ID, objc.Sel("initWithStates:"), states)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStateMachineWithStates */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StateMachine */

// Creates a state machine with the specified states.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/stateMachineWithStates:
func (sc _StateMachineClass) StateMachineWithStates(states []State) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("stateMachineWithStates:"), states)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StateMachineWithStates) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StateMachine */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StateMachine */

// Returns a Boolean value indicating whether it is valid for the state machine to transition from its current state to a state of the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/canEnterState(_:)
func (s_ StateMachine) CanEnterState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canEnterState:"), stateClass)
	return rv
}/* debug [instance_methods/method]: CanEnterState */


// Attempts to transition the state machine from its current state to a state of the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/enter(_:)
func (s_ StateMachine) EnterState(stateClass objc.Class) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enterState:"), stateClass)
	return rv
}/* debug [instance_methods/method]: EnterState */


// Returns the state object in the state machine corresponding to the specified class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/stateForClass:
func (s_ StateMachine) StateForClass(stateClass objc.Class) IState {
	rv := objc.Send[State](s_.ID, objc.Sel("stateForClass:"), stateClass)
	return rv
}/* debug [instance_methods/method]: StateForClass */


// Tells the current state object to perform per-frame updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/update(deltaTime:)
func (s_ StateMachine) UpdateWithDeltaTime(sec float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateWithDeltaTime:"), sec)
}/* debug [instance_methods/method]: UpdateWithDeltaTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StateMachine */

// The state machine’s current state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKStateMachine/currentState
func (s_ StateMachine) CurrentState() IGKState {
	rv := objc.Send[State](s_.ID, objc.Sel("currentState"))
	return rv
}/* debug [instance_properties/getter]: currentState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKStateMachine */


