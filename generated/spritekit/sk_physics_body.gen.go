// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [SKPhysicsBody] class.
var (
	sKPhysicsBodyClass     _SKPhysicsBodyClass
	sKPhysicsBodyClassOnce sync.Once
)

func getSKPhysicsBodyClass() _SKPhysicsBodyClass {
	sKPhysicsBodyClassOnce.Do(func() {
		sKPhysicsBodyClass = _SKPhysicsBodyClass{objc.GetClass("SKPhysicsBody")}
	})
	return sKPhysicsBodyClass
}

type _SKPhysicsBodyClass struct {
	class objc.Class
}

// An interface definition for the [SKPhysicsBody] class.
type ISKPhysicsBody interface {
	objectivec.IObject
	AllContactedBodies() unsafe.Pointer
	ApplyAngularImpulse(impulse float64)
	ApplyForce(force coregraphics.CGVector)
	ApplyForceAtPoint(force coregraphics.CGVector, point unsafe.Pointer)
	ApplyImpulse(impulse coregraphics.CGVector)
	ApplyImpulseAtPoint(impulse coregraphics.CGVector, point unsafe.Pointer)
	ApplyTorque(torque float64)
}

// An object that adds physics simulation to a node. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody
type SKPhysicsBody struct {
	objectivec.Object
}

// SKPhysicsBodyFrom constructs a [SKPhysicsBody] from an unsafe.Pointer.
//
// An object that adds physics simulation to a node.
func SKPhysicsBodyFrom(ptr unsafe.Pointer) SKPhysicsBody {
	return SKPhysicsBody{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKPhysicsBodyClass) Alloc() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKPhysicsBodyClass) New() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKPhysicsBody) Init() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKPhysicsBody) Autorelease() SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKPhysicsBody creates a new SKPhysicsBody instance.
func NewSKPhysicsBody() SKPhysicsBody {
	return getSKPhysicsBodyClass().New()
}


// Creates an edge loop from a rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-8sqfy
func NewSKPhysicsBodyWithEdgeLoopFromRect(rect unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeLoopFromRect:"), rect)
	return rv
}
// Creates a physics body from the contents of a texture, capturing only the texels that exceed a specified transparency value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:alphaThreshold:size:)
func NewSKPhysicsBodyWithTextureAlphaThresholdSize(texture unsafe.Pointer, alphaThreshold float32, size unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithTexture:alphaThreshold:size:"), texture, alphaThreshold, size)
	return rv
}
// Creates a physics body from the contents of a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:size:)
func NewSKPhysicsBodyWithTextureSize(texture unsafe.Pointer, size unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithTexture:size:"), texture, size)
	return rv
}
// Creates a physics body that’s shaped like a union of the argument physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(bodies:)
func NewSKPhysicsBodyWithBodies(bodies unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithBodies:"), bodies)
	return rv
}
// Creates a circular physics body centered on an arbitrary point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:center:)
func NewSKPhysicsBodyWithCircleOfRadiusCenter(r float64, center unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithCircleOfRadius:center:"), r, center)
	return rv
}
// Creates an edge between two points. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeFrom:to:)
func NewSKPhysicsBodyWithEdgeFromPointToPoint(p1 unsafe.Pointer, p2 unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeFromPoint:toPoint:"), p1, p2)
	return rv
}
// Creates an edge loop from a path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-5grxu
func NewSKPhysicsBodyWithEdgeLoopFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeLoopFromPath:"), path)
	return rv
}
// Creates a polygonal physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(polygonFrom:)
func NewSKPhysicsBodyWithPolygonFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithPolygonFromPath:"), path)
	return rv
}
// Creates a rectangular physics body centered on the owning node’s origin. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:)
func NewSKPhysicsBodyWithRectangleOfSize(s unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithRectangleOfSize:"), s)
	return rv
}
// Creates a rectangular physics body centered on an arbitrary point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:center:)
func NewSKPhysicsBodyWithRectangleOfSizeCenter(s unsafe.Pointer, center unsafe.Pointer) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithRectangleOfSize:center:"), s, center)
	return rv
}
// Creates a circular physics body centered on the owning node’s origin. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:)
func NewSKPhysicsBodyWithCircleOfRadius(r float64) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithCircleOfRadius:"), r)
	return rv
}
// Creates an edge chain from a path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeChainFrom:)
func NewSKPhysicsBodyWithEdgeChainFromPath(path coregraphics.CGPathRef) SKPhysicsBody {
	rv := objc.Send[SKPhysicsBody](objc.ID(getSKPhysicsBodyClass().class), objc.Sel("bodyWithEdgeChainFromPath:"), path)
	return rv
}


// Creates a physics body that’s shaped like a union of the argument physics bodies. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(bodies:)
func (sc _SKPhysicsBodyClass) BodyWithBodies(bodies unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithBodies:"), bodies)
	return rv
}
// Creates a circular physics body centered on the owning node’s origin. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:)
func (sc _SKPhysicsBodyClass) BodyWithCircleOfRadius(r float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithCircleOfRadius:"), r)
	return rv
}
// Creates a circular physics body centered on an arbitrary point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(circleOfRadius:center:)
func (sc _SKPhysicsBodyClass) BodyWithCircleOfRadiusCenter(r float64, center unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithCircleOfRadius:center:"), r, center)
	return rv
}
// Creates an edge chain from a path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeChainFrom:)
func (sc _SKPhysicsBodyClass) BodyWithEdgeChainFromPath(path coregraphics.CGPathRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithEdgeChainFromPath:"), path)
	return rv
}
// Creates an edge between two points. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeFrom:to:)
func (sc _SKPhysicsBodyClass) BodyWithEdgeFromPointToPoint(p1 unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithEdgeFromPoint:toPoint:"), p1, p2)
	return rv
}
// Creates an edge loop from a path. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-5grxu
func (sc _SKPhysicsBodyClass) BodyWithEdgeLoopFromPath(path coregraphics.CGPathRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithEdgeLoopFromPath:"), path)
	return rv
}
// Creates an edge loop from a rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(edgeLoopFrom:)-8sqfy
func (sc _SKPhysicsBodyClass) BodyWithEdgeLoopFromRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithEdgeLoopFromRect:"), rect)
	return rv
}
// Creates a polygonal physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(polygonFrom:)
func (sc _SKPhysicsBodyClass) BodyWithPolygonFromPath(path coregraphics.CGPathRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithPolygonFromPath:"), path)
	return rv
}
// Creates a rectangular physics body centered on the owning node’s origin. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:)
func (sc _SKPhysicsBodyClass) BodyWithRectangleOfSize(s unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithRectangleOfSize:"), s)
	return rv
}
// Creates a rectangular physics body centered on an arbitrary point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(rectangleOf:center:)
func (sc _SKPhysicsBodyClass) BodyWithRectangleOfSizeCenter(s unsafe.Pointer, center unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithRectangleOfSize:center:"), s, center)
	return rv
}
// Creates a physics body from the contents of a texture, capturing only the texels that exceed a specified transparency value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:alphaThreshold:size:)
func (sc _SKPhysicsBodyClass) BodyWithTextureAlphaThresholdSize(texture unsafe.Pointer, alphaThreshold float32, size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithTexture:alphaThreshold:size:"), texture, alphaThreshold, size)
	return rv
}
// Creates a physics body from the contents of a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/init(texture:size:)
func (sc _SKPhysicsBodyClass) BodyWithTextureSize(texture unsafe.Pointer, size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("bodyWithTexture:size:"), texture, size)
	return rv
}
// The physics bodies that this physics body is in contact with. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/allContactedBodies()
func (s_ SKPhysicsBody) AllContactedBodies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("allContactedBodies"))
	return rv
}
// Applies an impulse that imparts angular momentum to an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyAngularImpulse(_:)
func (s_ SKPhysicsBody) ApplyAngularImpulse(impulse float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyAngularImpulse:"), impulse)
}
// Applies a force to the center of gravity of a physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyForce(_:)
func (s_ SKPhysicsBody) ApplyForce(force coregraphics.CGVector) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyForce:"), force)
}
// Applies a force to a specific point of a physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyForce(_:at:)
func (s_ SKPhysicsBody) ApplyForceAtPoint(force coregraphics.CGVector, point unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyForce:atPoint:"), force, point)
}
// Applies an impulse to the center of gravity of a physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyImpulse(_:)
func (s_ SKPhysicsBody) ApplyImpulse(impulse coregraphics.CGVector) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyImpulse:"), impulse)
}
// Applies an impulse to a specific point of a physics body. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyImpulse(_:at:)
func (s_ SKPhysicsBody) ApplyImpulseAtPoint(impulse coregraphics.CGVector, point unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyImpulse:atPoint:"), impulse, point)
}
// Applies torque to an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKPhysicsBody/applyTorque(_:)
func (s_ SKPhysicsBody) ApplyTorque(torque float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("applyTorque:"), torque)
}

