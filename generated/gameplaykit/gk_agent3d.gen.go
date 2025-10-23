// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Agent3D] class.
var (
	Agent3DClass     _Agent3DClass
	Agent3DClassOnce sync.Once
)

func getAgent3DClass() _Agent3DClass {
	Agent3DClassOnce.Do(func() {
		Agent3DClass = _Agent3DClass{objc.GetClass("GKAgent3D")}
	})
	return Agent3DClass
}

type _Agent3DClass struct {
	class objc.Class
}

// An interface definition for the [Agent3D] class.
type IAgent3D interface {
	IAgent
	UpdateWithDeltaTime(seconds foundation.ITimeInterval)
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	RightHanded() bool
	SetRightHanded(value bool)
	Rotation() unsafe.Pointer
	SetRotation(value unsafe.Pointer)
	Velocity() unsafe.Pointer
}

// An agent that operates in a three-dimensional space.
//
// Agents are game entities that move according to realistic constraints and whose behavior is determined by goals that motivate movement. The general functionality of an agent is defined by the abstract superclass ; however, you use instances of the class to implement agent-based gameplay in a 3D game. To learn more about using goals and agents, see in .


// An agent that operates in a three-dimensional space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D
type Agent3D struct {
	Agent
}

// Agent3DFrom constructs a [Agent3D] from an unsafe.Pointer.
//
// An agent that operates in a three-dimensional space.
func Agent3DFrom(ptr unsafe.Pointer) Agent3D {
	return Agent3D{
		Agent: AgentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _Agent3DClass) Alloc() Agent3D {
	rv := objc.Send[Agent3D](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _Agent3DClass) New() Agent3D {
	rv := objc.Send[Agent3D](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Agent3D) Init() Agent3D {
	rv := objc.Send[Agent3D](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Agent3D) Autorelease() Agent3D {
	rv := objc.Send[Agent3D](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAgent3D creates a new Agent3D instance.
func NewAgent3D() Agent3D {
	return getAgent3DClass().New()
}



// Causes the agent to evaluate its goals and update its position, rotation, and velocity accordingly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/update(deltaTime:)
func (a_ Agent3D) UpdateWithDeltaTime(seconds foundation.ITimeInterval) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}


// The current position of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/position
func (a_ Agent3D) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("position"))
	return rv
}


// The current position of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/position
func (a_ Agent3D) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPosition:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rightHanded
func (a_ Agent3D) RightHanded() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("rightHanded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rightHanded
func (a_ Agent3D) SetRightHanded(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRightHanded:"), value)
}


// The orientation of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rotation
func (a_ Agent3D) Rotation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rotation"))
	return rv
}


// The orientation of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rotation
func (a_ Agent3D) SetRotation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRotation:"), value)
}


// The current velocity of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/velocity
func (a_ Agent3D) Velocity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("velocity"))
	return rv
}



