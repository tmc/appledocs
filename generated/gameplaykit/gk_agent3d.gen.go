// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKAgent3D */


/* debug [class_header]: Header for GKAgent3D */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Agent3D */
// An interface definition for the [Agent3D] class.
type IAgent3D interface {
	IAgent
	
/* debug [class_interface_properties]: Properties for Agent3D */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
	RightHanded() bool
	SetRightHanded(value bool)
	Rotation() objectivec.IObject
	SetRotation(value objectivec.IObject)
	Velocity() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Agent3D */
	// methods:
	UpdateWithDeltaTime(seconds float64)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Agent3D */
// Alloc allocates a new instance without initialization.
func (ac _Agent3DClass) Alloc() Agent3D {
	rv := objc.Send[Agent3D](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Agent3D */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Agent3D *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Agent3D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Agent3D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Agent3D */

// Causes the agent to evaluate its goals and update its position, rotation, and velocity accordingly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/update(deltaTime:)
func (a_ Agent3D) UpdateWithDeltaTime(seconds float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithDeltaTime:"), seconds)
}/* debug [instance_methods/method]: UpdateWithDeltaTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Agent3D */

// The current position of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/position
func (a_ Agent3D) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The current position of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/position
func (a_ Agent3D) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rightHanded
func (a_ Agent3D) RightHanded() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("rightHanded"))
	return rv
}/* debug [instance_properties/getter]: rightHanded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rightHanded
func (a_ Agent3D) SetRightHanded(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRightHanded:"), value)
}/* debug [instance_properties/setter]: rightHanded */


// The orientation of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rotation
func (a_ Agent3D) Rotation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("rotation"))
	return rv
}/* debug [instance_properties/getter]: rotation */


// The orientation of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/rotation
func (a_ Agent3D) SetRotation(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRotation:"), value)
}/* debug [instance_properties/setter]: rotation */


// The current velocity of the agent in 3D space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKAgent3D/velocity
func (a_ Agent3D) Velocity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("velocity"))
	return rv
}/* debug [instance_properties/getter]: velocity */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAgent3D */



