// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKPhysicsJoint] class.
var sKPhysicsJointClass = _SKPhysicsJointClass{objc.GetClass("SKPhysicsJoint")}

type _SKPhysicsJointClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsJoint] class.
type ISKPhysicsJoint interface {
	objectivec.IObject
}

// The abstract superclass for objects that connect physics bodies. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsJoint

type SKPhysicsJoint struct {
	objectivec.Object
}

// SKPhysicsJointFrom constructs a [SKPhysicsJoint] from an unsafe.Pointer.
//
// The abstract superclass for objects that connect physics bodies.
func SKPhysicsJointFrom(ptr unsafe.Pointer) SKPhysicsJoint {
	return SKPhysicsJoint{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsJointClass) Alloc() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKPhysicsJointClass) New() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsJoint) Init() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsJoint) Autorelease() SKPhysicsJoint {
	rv := objc.Send[SKPhysicsJoint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsJoint creates a new SKPhysicsJoint instance.
func NewSKPhysicsJoint() SKPhysicsJoint {
	return sKPhysicsJointClass.New()
}




