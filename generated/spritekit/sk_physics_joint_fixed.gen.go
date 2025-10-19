// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKPhysicsJointFixed] class.
var sKPhysicsJointFixedClass = _SKPhysicsJointFixedClass{objc.GetClass("SKPhysicsJointFixed")}

type _SKPhysicsJointFixedClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJointFixed] class.
type ISKPhysicsJointFixed interface {
	ISKPhysicsJoint
}

// A joint that fuses two physics bodies together at a reference point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointFixed

type SKPhysicsJointFixed struct {
	SKPhysicsJoint
}

// SKPhysicsJointFixedFrom constructs a [SKPhysicsJointFixed] from an unsafe.Pointer.
//
// A joint that fuses two physics bodies together at a reference point.
func SKPhysicsJointFixedFrom(ptr unsafe.Pointer) SKPhysicsJointFixed {
	return SKPhysicsJointFixed{
		SKPhysicsJoint: SKPhysicsJointFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointFixedClass) Alloc() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKPhysicsJointFixedClass) New() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJointFixed) Init() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJointFixed) Autorelease() SKPhysicsJointFixed {
	rv := objc.Send[SKPhysicsJointFixed](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointFixed creates a new SKPhysicsJointFixed instance.
func NewSKPhysicsJointFixed() SKPhysicsJointFixed {
	return sKPhysicsJointFixedClass.New()
}


// Creates a new fixed joint. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointFixed/joint(withBodyA:bodyB:anchor:)
func (sc _SKPhysicsJointFixedClass) JointWithBodyABodyBAnchor(bodyA unsafe.Pointer, bodyB unsafe.Pointer, anchor unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("jointWithBodyA:bodyB:anchor:"), bodyA, bodyB, anchor)
	return rv
}


