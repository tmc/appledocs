// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKAgent2D */


/* debug [class_header]: Header for GKAgent2D */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Agent2D */
// An interface definition for the [Agent2D] class.
type IAgent2D interface {
	IAgent
	
/* debug [class_interface_properties]: Properties for Agent2D */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
	Rotation() float32
	SetRotation(value float32)
	Velocity() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Agent2D */
	// methods:
	UpdateWithDeltaTime(seconds float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Agent2D */
// Alloc allocates a new instance without initialization.
func (ac _Agent2DClass) Alloc() Agent2D {
	rv := objc.Send[Agent2D](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Agent2D */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Agent2D *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Agent2D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Agent2D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Agent2D */

// Causes the agent to evaluate its goals and update its position, rotation, and velocity accordingly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/update(deltaTime:)
func (a_ Agent2D) UpdateWithDeltaTime(seconds float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}/* debug [instance_methods/method]: UpdateWithDeltaTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Agent2D */

// The current position of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/position
func (a_ Agent2D) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The current position of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/position
func (a_ Agent2D) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// The rotation of the agent around the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/rotation
func (a_ Agent2D) Rotation() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("rotation"))
	return rv
}/* debug [instance_properties/getter]: rotation */


// The rotation of the agent around the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/rotation
func (a_ Agent2D) SetRotation(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRotation:"), value)
}/* debug [instance_properties/setter]: rotation */


// The current velocity of the agent in 2D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent2D/velocity
func (a_ Agent2D) Velocity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("velocity"))
	return rv
}/* debug [instance_properties/getter]: velocity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAgent2D */



