// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SKPhysicsJointSliding] class.
var (
	sKPhysicsJointSlidingClass     _SKPhysicsJointSlidingClass
	sKPhysicsJointSlidingClassOnce sync.Once
)

func getSKPhysicsJointSlidingClass() _SKPhysicsJointSlidingClass {
	sKPhysicsJointSlidingClassOnce.Do(func() {
		sKPhysicsJointSlidingClass = _SKPhysicsJointSlidingClass{objc.GetClass("SKPhysicsJointSliding")}
	})
	return sKPhysicsJointSlidingClass
}

type _SKPhysicsJointSlidingClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJointSliding] class.
type ISKPhysicsJointSliding interface {
	ISKPhysicsJoint
}

// A joint that allows two physics bodies to slide along an axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding
type SKPhysicsJointSliding struct {
	SKPhysicsJoint
}

// SKPhysicsJointSlidingFrom constructs a [SKPhysicsJointSliding] from an unsafe.Pointer.
//
// A joint that allows two physics bodies to slide along an axis.
func SKPhysicsJointSlidingFrom(ptr unsafe.Pointer) SKPhysicsJointSliding {
	return SKPhysicsJointSliding{
		SKPhysicsJoint: SKPhysicsJointFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointSlidingClass) Alloc() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKPhysicsJointSlidingClass) New() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJointSliding) Init() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJointSliding) Autorelease() SKPhysicsJointSliding {
	rv := objc.Send[SKPhysicsJointSliding](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJointSliding creates a new SKPhysicsJointSliding instance.
func NewSKPhysicsJointSliding() SKPhysicsJointSliding {
	return getSKPhysicsJointSlidingClass().New()
}


// Creates a new sliding joint. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJointSliding/joint(withBodyA:bodyB:anchor:axis:)
func (sc _SKPhysicsJointSlidingClass) JointWithBodyABodyBAnchorAxis(bodyA unsafe.Pointer, bodyB unsafe.Pointer, anchor unsafe.Pointer, axis coregraphics.CGVector) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("jointWithBodyA:bodyB:anchor:axis:"), bodyA, bodyB, anchor, axis)
	return rv
}


