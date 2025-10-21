// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [Agent] class.
type IAgent interface {
	IComponent
}

// A component that moves a game entity according to a set of goals and realistic constraints.
//
// The class is abstract, defining only the general functionality of an agent—its movement constraints and the property containing its goals ( objects). To implement agent-based gameplay, choose a concrete subclass that fits your game. Use the class for 2D game worlds, or for 3D games where all gameplay-relevant movement is constrained to two dimensions. Use the class for game worlds that allow movement in three dimensions. To learn more about the agent simulation, see in .
//
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

// Alloc allocates a new instance without initialization.
func (ac _AgentClass) Alloc() Agent {
	rv := objc.Send[Agent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A weighted collection of goals that influence the agent’s movement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/behavior
func (a_ Agent) Behavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("behavior"))
	return rv
}


// SetBehavior sets the value of the behavior property.
// A weighted collection of goals that influence the agent’s movement.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/behavior
func (a_ Agent) SetBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBehavior:"), value)
}

// An object that prepares for or responds to updates in the agent simulation.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/delegate
func (a_ Agent) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that prepares for or responds to updates in the agent simulation.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/delegate
func (a_ Agent) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// The resistance of the agent to changes in speed or direction.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/mass
func (a_ Agent) Mass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mass"))
	return rv
}


// SetMass sets the value of the mass property.
// The resistance of the agent to changes in speed or direction.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/mass
func (a_ Agent) SetMass(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMass:"), value)
}

// The upper limit to changes in the agent’s speed or direction.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxAcceleration
func (a_ Agent) MaxAcceleration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("maxAcceleration"))
	return rv
}


// SetMaxAcceleration sets the value of the maxAcceleration property.
// The upper limit to changes in the agent’s speed or direction.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxAcceleration
func (a_ Agent) SetMaxAcceleration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxAcceleration:"), value)
}

// The agent’s maximum forward speed, in units per second.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxSpeed
func (a_ Agent) MaxSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("maxSpeed"))
	return rv
}


// SetMaxSpeed sets the value of the maxSpeed property.
// The agent’s maximum forward speed, in units per second.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/maxSpeed
func (a_ Agent) SetMaxSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaxSpeed:"), value)
}

// The agent’s radius.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/radius
func (a_ Agent) Radius() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("radius"))
	return rv
}


// SetRadius sets the value of the radius property.
// The agent’s radius.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/radius
func (a_ Agent) SetRadius(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRadius:"), value)
}

// The agent’s current forward speed, in units per second.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/speed
func (a_ Agent) Speed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("speed"))
	return rv
}


// SetSpeed sets the value of the speed property.
// The agent’s current forward speed, in units per second.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent/speed
func (a_ Agent) SetSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSpeed:"), value)
}



