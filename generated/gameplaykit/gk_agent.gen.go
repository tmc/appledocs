// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKAgent */


/* debug [class_header]: Header for GKAgent */
// The class instance for the [Agent] class.
var (
	AgentClass     _AgentClass
	AgentClassOnce sync.Once
)

func getAgentClass() _AgentClass {
	AgentClassOnce.Do(func() {
		AgentClass = _AgentClass{objc.GetClass("GKAgent")}
	})
	return AgentClass
}

type _AgentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Agent */
// An interface definition for the [Agent] class.
type IAgent interface {
	IComponent
	
/* debug [class_interface_properties]: Properties for Agent */
	// properties:
	Behavior() IGKBehavior
	SetBehavior(value IGKBehavior)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Mass() float32
	SetMass(value float32)
	MaxAcceleration() float32
	SetMaxAcceleration(value float32)
	MaxSpeed() float32
	SetMaxSpeed(value float32)
	Radius() float32
	SetRadius(value float32)
	Speed() float32
	SetSpeed(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Agent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Agent */
// Alloc allocates a new instance without initialization.
func (ac _AgentClass) Alloc() Agent {
	rv := objc.Send[Agent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AgentClass) New() Agent {
	rv := objc.Send[Agent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Agent) Init() Agent {
	rv := objc.Send[Agent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Agent) Autorelease() Agent {
	rv := objc.Send[Agent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAgent creates a new Agent instance.
func NewAgent() Agent {
	return getAgentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Agent */
// A component that moves a game entity according to a set of goals and realistic constraints.
//
// The class is abstract, defining only the general functionality of an agent—its movement constraints and the property containing its goals ( objects). To implement agent-based gameplay, choose a concrete subclass that fits your game. Use the class for 2D game worlds, or for 3D games where all gameplay-relevant movement is constrained to two dimensions. Use the class for game worlds that allow movement in three dimensions. To learn more about the agent simulation, see in .


// A component that moves a game entity according to a set of goals and realistic constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent
type Agent struct {
	Component
}

// AgentFrom constructs a [Agent] from an unsafe.Pointer.
//
// A component that moves a game entity according to a set of goals and realistic constraints.
func AgentFrom(ptr unsafe.Pointer) Agent {
	return Agent{
		Component: ComponentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Agent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Agent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Agent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Agent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Agent */

// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/behavior
func (a_ Agent) Behavior() IGKBehavior {
	rv := objc.Send[Behavior](a_.ID, objc.Sel("behavior"))
	return rv
}/* debug [instance_properties/getter]: behavior */


// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/behavior
func (a_ Agent) SetBehavior(value IGKBehavior) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBehavior:"), value)
}/* debug [instance_properties/setter]: behavior */


// An object that prepares for or responds to updates in the agent simulation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/delegate
func (a_ Agent) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// An object that prepares for or responds to updates in the agent simulation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/delegate
func (a_ Agent) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The resistance of the agent to changes in speed or direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/mass
func (a_ Agent) Mass() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("mass"))
	return rv
}/* debug [instance_properties/getter]: mass */


// The resistance of the agent to changes in speed or direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/mass
func (a_ Agent) SetMass(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMass:"), value)
}/* debug [instance_properties/setter]: mass */


// The upper limit to changes in the agent’s speed or direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxAcceleration
func (a_ Agent) MaxAcceleration() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("maxAcceleration"))
	return rv
}/* debug [instance_properties/getter]: maxAcceleration */


// The upper limit to changes in the agent’s speed or direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxAcceleration
func (a_ Agent) SetMaxAcceleration(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxAcceleration:"), value)
}/* debug [instance_properties/setter]: maxAcceleration */


// The agent’s maximum forward speed, in units per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxSpeed
func (a_ Agent) MaxSpeed() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("maxSpeed"))
	return rv
}/* debug [instance_properties/getter]: maxSpeed */


// The agent’s maximum forward speed, in units per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxSpeed
func (a_ Agent) SetMaxSpeed(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxSpeed:"), value)
}/* debug [instance_properties/setter]: maxSpeed */


// The agent’s radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/radius
func (a_ Agent) Radius() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// The agent’s radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/radius
func (a_ Agent) SetRadius(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadius:"), value)
}/* debug [instance_properties/setter]: radius */


// The agent’s current forward speed, in units per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/speed
func (a_ Agent) Speed() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("speed"))
	return rv
}/* debug [instance_properties/getter]: speed */


// The agent’s current forward speed, in units per second.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/speed
func (a_ Agent) SetSpeed(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSpeed:"), value)
}/* debug [instance_properties/setter]: speed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAgent */



