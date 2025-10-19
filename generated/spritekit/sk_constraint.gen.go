// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKConstraint] class.
var (
	sKConstraintClass     _SKConstraintClass
	sKConstraintClassOnce sync.Once
)

func getSKConstraintClass() _SKConstraintClass {
	sKConstraintClassOnce.Do(func() {
		sKConstraintClass = _SKConstraintClass{objc.GetClass("SKConstraint")}
	})
	return sKConstraintClass
}

type _SKConstraintClass struct {
	class objc.Class
}

// An interface definition for the [SKConstraint] class.
type ISKConstraint interface {
	objectivec.IObject
}

// A specification for constraining a node’s position or rotation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint
type SKConstraint struct {
	objectivec.Object
}

// SKConstraintFrom constructs a [SKConstraint] from an unsafe.Pointer.
//
// A specification for constraining a node’s position or rotation.
func SKConstraintFrom(ptr unsafe.Pointer) SKConstraint {
	return SKConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKConstraintClass) Alloc() SKConstraint {
	rv := objc.Send[SKConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKConstraintClass) New() SKConstraint {
	rv := objc.Send[SKConstraint](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKConstraint) Init() SKConstraint {
	rv := objc.Send[SKConstraint](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKConstraint) Autorelease() SKConstraint {
	rv := objc.Send[SKConstraint](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKConstraint creates a new SKConstraint instance.
func NewSKConstraint() SKConstraint {
	return getSKConstraintClass().New()
}


// Creates a constraint that keeps a node within a certain distance of another node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:)-6507j
func (sc _SKConstraintClass) DistanceToNode(range_ unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("distance:toNode:"), range_, node)
	return rv
}
// Creates a constraint that keeps a node within a certain distance of a point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:)-7yk7n
func (sc _SKConstraintClass) DistanceToPoint(range_ unsafe.Pointer, point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("distance:toPoint:"), range_, point)
	return rv
}
// Creates a constraint that keeps a node within a certain distance of a point in another node’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/distance(_:to:in:)
func (sc _SKConstraintClass) DistanceToPointInNode(range_ unsafe.Pointer, point unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("distance:toPoint:inNode:"), range_, point, node)
	return rv
}
// Creates a constraint that forces a node to rotate to face a point in another node’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:in:offset:)
func (sc _SKConstraintClass) OrientToPointInNodeOffset(point unsafe.Pointer, node unsafe.Pointer, radians unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("orientToPoint:inNode:offset:"), point, node, radians)
	return rv
}
// Creates a constraint that forces a node to rotate to face another node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:offset:)-1h1tw
func (sc _SKConstraintClass) OrientToNodeOffset(node unsafe.Pointer, radians unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("orientToNode:offset:"), node, radians)
	return rv
}
// Creates a constraint that forces a node to rotate to face a fixed point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/orient(to:offset:)-9lq3h
func (sc _SKConstraintClass) OrientToPointOffset(point unsafe.Pointer, radians unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("orientToPoint:offset:"), point, radians)
	return rv
}
// Creates a constraint that restricts the x-coordinate of a node’s position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionX(_:)
func (sc _SKConstraintClass) PositionX(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("positionX:"), range_)
	return rv
}
// Creates a constraint that restricts both coordinates of a node’s position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionX(_:y:)
func (sc _SKConstraintClass) PositionXY(xRange unsafe.Pointer, yRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("positionX:Y:"), xRange, yRange)
	return rv
}
// Creates a constraint that restricts the y-coordinate of a node’s position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/positionY(_:)
func (sc _SKConstraintClass) PositionY(range_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("positionY:"), range_)
	return rv
}
// Creates a constraint that limits the orientation of a node. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKConstraint/zRotation(_:)
func (sc _SKConstraintClass) ZRotation(zRange unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("zRotation:"), zRange)
	return rv
}


