// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKPhysicsJointSpring] class.
var (
	sKPhysicsJointSpringClass     _SKPhysicsJointSpringClass
	sKPhysicsJointSpringClassOnce sync.Once
)

func getSKPhysicsJointSpringClass() _SKPhysicsJointSpringClass {
	sKPhysicsJointSpringClassOnce.Do(func() {
		sKPhysicsJointSpringClass = _SKPhysicsJointSpringClass{objc.GetClass("SKPhysicsJointSpring")}
	})
	return sKPhysicsJointSpringClass
}

type _SKPhysicsJointSpringClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJointSpring] class.
type ISKPhysicsJointSpring interface {
	ISKPhysicsJoint
}

// A joint that simulates a spring connecting two physics bodies.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring
type SKPhysicsJointSpring struct {
	SKPhysicsJoint
}

// SKPhysicsJointSpringFrom constructs a [SKPhysicsJointSpring] from an unsafe.Pointer.
//
// A joint that simulates a spring connecting two physics bodies.
func SKPhysicsJointSpringFrom(ptr unsafe.Pointer) SKPhysicsJointSpring {
	return SKPhysicsJointSpring{
		SKPhysicsJoint: SKPhysicsJointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointSpringClass) Alloc() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKPhysicsJointSpringClass) New() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJointSpring) Init() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJointSpring) Autorelease() SKPhysicsJointSpring {
	rv := objc.Send[SKPhysicsJointSpring](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointSpring creates a new SKPhysicsJointSpring instance.
func NewSKPhysicsJointSpring() SKPhysicsJointSpring {
	return getSKPhysicsJointSpringClass().New()
}


// Creates a new spring joint.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSpring/joint(withBodyA:bodyB:anchorA:anchorB:)
func (sc _SKPhysicsJointSpringClass) JointWithBodyABodyBAnchorAAnchorB(bodyA unsafe.Pointer, bodyB unsafe.Pointer, anchorA unsafe.Pointer, anchorB unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("jointWithBodyA:bodyB:anchorA:anchorB:"), bodyA, bodyB, anchorA, anchorB)
	return rv
}


