// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Agent2D] class.
var (
	Agent2DClass     _Agent2DClass
	Agent2DClassOnce sync.Once
)

func getAgent2DClass() _Agent2DClass {
	Agent2DClassOnce.Do(func() {
		Agent2DClass = _Agent2DClass{objc.GetClass("GKAgent2D")}
	})
	return Agent2DClass
}

type _Agent2DClass struct {
	class objc.Class
}

// An interface definition for the [Agent2D] class.
type IAgent2D interface {
	IAgent
	UpdateWithDeltaTime(seconds foundation.ITimeInterval)
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	Rotation() float32
	SetRotation(value float32)
	Velocity() unsafe.Pointer
}

// An agent that operates in a two-dimensional space.
//
// Agents are game entities that move according to realistic constraints and whose behavior is determined by goals that motivate movement. The general functionality of an agent is defined by the abstract superclass ; however, you use instances of the class to implement agent-based gameplay in a 2D game (or in a 3D game where gameplay-relevant movement is restricted to two dimensions). To learn more about using goals and agents, see in .


// An agent that operates in a two-dimensional space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D
type Agent2D struct {
	Agent
}

// Agent2DFrom constructs a [Agent2D] from an unsafe.Pointer.
//
// An agent that operates in a two-dimensional space.
func Agent2DFrom(ptr unsafe.Pointer) Agent2D {
	return Agent2D{
		Agent: AgentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _Agent2DClass) Alloc() Agent2D {
	rv := objc.Send[Agent2D](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _Agent2DClass) New() Agent2D {
	rv := objc.Send[Agent2D](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Agent2D) Init() Agent2D {
	rv := objc.Send[Agent2D](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Agent2D) Autorelease() Agent2D {
	rv := objc.Send[Agent2D](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAgent2D creates a new Agent2D instance.
func NewAgent2D() Agent2D {
	return getAgent2DClass().New()
}



// Causes the agent to evaluate its goals and update its position, rotation, and velocity accordingly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/update(deltaTime:)
func (a_ Agent2D) UpdateWithDeltaTime(seconds foundation.ITimeInterval) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}


// The current position of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/position
func (a_ Agent2D) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("position"))
	return rv
}


// The current position of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/position
func (a_ Agent2D) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPosition:"), value)
}


// The rotation of the agent around the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/rotation
func (a_ Agent2D) Rotation() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rotation"))
	return rv
}


// The rotation of the agent around the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/rotation
func (a_ Agent2D) SetRotation(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRotation:"), value)
}


// The current velocity of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/velocity
func (a_ Agent2D) Velocity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("velocity"))
	return rv
}



