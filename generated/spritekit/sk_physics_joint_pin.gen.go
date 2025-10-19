// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKPhysicsJointPin] class.
var sKPhysicsJointPinClass = _SKPhysicsJointPinClass{objc.GetClass("SKPhysicsJointPin")}

type _SKPhysicsJointPinClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJointPin] class.
type ISKPhysicsJointPin interface {
	ISKPhysicsJoint
}

// A joint that pins together two physics bodies, allowing independent rotation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin

type SKPhysicsJointPin struct {
	SKPhysicsJoint
}

// SKPhysicsJointPinFrom constructs a [SKPhysicsJointPin] from an unsafe.Pointer.
//
// A joint that pins together two physics bodies, allowing independent rotation.
func SKPhysicsJointPinFrom(ptr unsafe.Pointer) SKPhysicsJointPin {
	return SKPhysicsJointPin{
		SKPhysicsJoint: SKPhysicsJointFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointPinClass) Alloc() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKPhysicsJointPinClass) New() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJointPin) Init() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJointPin) Autorelease() SKPhysicsJointPin {
	rv := objc.Send[SKPhysicsJointPin](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointPin creates a new SKPhysicsJointPin instance.
func NewSKPhysicsJointPin() SKPhysicsJointPin {
	return sKPhysicsJointPinClass.New()
}


// Creates a new pin joint. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointPin/joint(withBodyA:bodyB:anchor:)
func (sc _SKPhysicsJointPinClass) JointWithBodyABodyBAnchor(bodyA unsafe.Pointer, bodyB unsafe.Pointer, anchor unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("jointWithBodyA:bodyB:anchor:"), bodyA, bodyB, anchor)
	return rv
}


