// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKPhysicsJointLimit] class.
var sKPhysicsJointLimitClass = _SKPhysicsJointLimitClass{objc.GetClass("SKPhysicsJointLimit")}

type _SKPhysicsJointLimitClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJointLimit] class.
type ISKPhysicsJointLimit interface {
	ISKPhysicsJoint
}

// A joint that imposes a maximum distance between two physics bodies, as if they were connected by a rope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit

type SKPhysicsJointLimit struct {
	SKPhysicsJoint
}

// SKPhysicsJointLimitFrom constructs a [SKPhysicsJointLimit] from an unsafe.Pointer.
//
// A joint that imposes a maximum distance between two physics bodies, as if they were connected by a rope.
func SKPhysicsJointLimitFrom(ptr unsafe.Pointer) SKPhysicsJointLimit {
	return SKPhysicsJointLimit{
		SKPhysicsJoint: SKPhysicsJointFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointLimitClass) Alloc() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKPhysicsJointLimitClass) New() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJointLimit) Init() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJointLimit) Autorelease() SKPhysicsJointLimit {
	rv := objc.Send[SKPhysicsJointLimit](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointLimit creates a new SKPhysicsJointLimit instance.
func NewSKPhysicsJointLimit() SKPhysicsJointLimit {
	return sKPhysicsJointLimitClass.New()
}


// Creates a new limit joint. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointLimit/joint(withBodyA:bodyB:anchorA:anchorB:)
func (sc _SKPhysicsJointLimitClass) JointWithBodyABodyBAnchorAAnchorB(bodyA unsafe.Pointer, bodyB unsafe.Pointer, anchorA unsafe.Pointer, anchorB unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("jointWithBodyA:bodyB:anchorA:anchorB:"), bodyA, bodyB, anchorA, anchorB)
	return rv
}


