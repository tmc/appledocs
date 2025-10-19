// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKPhysicsWorld] class.
var sKPhysicsWorldClass = _SKPhysicsWorldClass{objc.GetClass("SKPhysicsWorld")}

type _SKPhysicsWorldClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsWorld] class.
type ISKPhysicsWorld interface {
	objectivec.IObject
	AddJoint(joint unsafe.Pointer)
	BodyAlongRayStartEnd(start unsafe.Pointer, end unsafe.Pointer) unsafe.Pointer
	BodyAtPoint(point unsafe.Pointer) unsafe.Pointer
	BodyInRect(rect unsafe.Pointer) unsafe.Pointer
	EnumerateBodiesAlongRayStartEndUsingBlock(start unsafe.Pointer, end unsafe.Pointer, block unsafe.Pointer)
	EnumerateBodiesAtPointUsingBlock(point unsafe.Pointer, block unsafe.Pointer)
	EnumerateBodiesInRectUsingBlock(rect unsafe.Pointer, block unsafe.Pointer)
	RemoveJoint(joint unsafe.Pointer)
	RemoveAllJoints()
	SampleFieldsAt(position unsafe.Pointer) unsafe.Pointer
}

// The driver of the physics engine in a scene; it exposes the ability for you to configure and query the physics system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld

type SKPhysicsWorld struct {
	objectivec.Object
}

// SKPhysicsWorldFrom constructs a [SKPhysicsWorld] from an unsafe.Pointer.
//
// The driver of the physics engine in a scene; it exposes the ability for you to configure and query the physics system.
func SKPhysicsWorldFrom(ptr unsafe.Pointer) SKPhysicsWorld {
	return SKPhysicsWorld{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsWorldClass) Alloc() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKPhysicsWorldClass) New() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsWorld) Init() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsWorld) Autorelease() SKPhysicsWorld {
	rv := objc.Send[SKPhysicsWorld](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsWorld creates a new SKPhysicsWorld instance.
func NewSKPhysicsWorld() SKPhysicsWorld {
	return sKPhysicsWorldClass.New()
}


// Adds a joint to the physics world. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/add(_:)
func (s_ SKPhysicsWorld) AddJoint(joint unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addJoint:"), joint)
}
// Searches for the first physics body that intersects a ray. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(alongRayStart:end:)
func (s_ SKPhysicsWorld) BodyAlongRayStartEnd(start unsafe.Pointer, end unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("bodyAlongRayStart:end:"), start, end)
	return rv
}
// Searches for the first physics body that contains a point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(at:)
func (s_ SKPhysicsWorld) BodyAtPoint(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("bodyAtPoint:"), point)
	return rv
}
// Searches for the first physics body that intersects the specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/body(in:)
func (s_ SKPhysicsWorld) BodyInRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("bodyInRect:"), rect)
	return rv
}
// Enumerates all the physics bodies in the scene that intersect a ray. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(alongRayStart:end:using:)
func (s_ SKPhysicsWorld) EnumerateBodiesAlongRayStartEndUsingBlock(start unsafe.Pointer, end unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateBodiesAlongRayStart:end:usingBlock:"), start, end, block)
}
// Enumerates all the physics bodies in the scene that contain a point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(at:using:)
func (s_ SKPhysicsWorld) EnumerateBodiesAtPointUsingBlock(point unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateBodiesAtPoint:usingBlock:"), point, block)
}
// Enumerates all the physics bodies in the scene that intersect the specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/enumerateBodies(in:using:)
func (s_ SKPhysicsWorld) EnumerateBodiesInRectUsingBlock(rect unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("enumerateBodiesInRect:usingBlock:"), rect, block)
}
// Removes a specific joint from the physics world. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/remove(_:)
func (s_ SKPhysicsWorld) RemoveJoint(joint unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeJoint:"), joint)
}
// Removes all joints from the physics world. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/removeAllJoints()
func (s_ SKPhysicsWorld) RemoveAllJoints() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeAllJoints"))
}
// Samples all of the field nodes in the scene and returns the summation of their forces at that point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsWorld/sampleFields(at:)
func (s_ SKPhysicsWorld) SampleFieldsAt(position unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("sampleFieldsAt:"), position)
	return rv
}


