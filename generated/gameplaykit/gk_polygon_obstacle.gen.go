// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PolygonObstacle] class.
var (
	PolygonObstacleClass     _PolygonObstacleClass
	PolygonObstacleClassOnce sync.Once
)

func getPolygonObstacleClass() _PolygonObstacleClass {
	PolygonObstacleClassOnce.Do(func() {
		PolygonObstacleClass = _PolygonObstacleClass{objc.GetClass("GKPolygonObstacle")}
	})
	return PolygonObstacleClass
}

type _PolygonObstacleClass struct {
	class objc.Class
}

// An interface definition for the [PolygonObstacle] class.
type IPolygonObstacle interface {
	IObstacle
	VertexAtIndex(index uint) unsafe.Pointer
	VertexCount() uint
}

// A polygon-shaped impassable area in a 2D game world.
//
// Polygon obstacles serve two purposes in GameplayKit: You can use polygon obstacles to construct a navigability graph of your game world (a object) for use in pathfinding. You can also use polygon obstacles to define regions for agents ( objects) to avoid, using the method . To easily create obstacles for use with a SpriteKit game, create and arrange a set of nodes that define the non-navigable regions of your game world. You can create such nodes programmatically, or use the SpriteKit Scene Editor in Xcode. If you’re already using nodes with physics bodies to keep sprites from entering those regions, you can reuse those nodes. Then, use the , , or method to generate a set of objects. To learn more about both ways of using polygon obstacles, see and in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle
type PolygonObstacle struct {
	Obstacle
}

// PolygonObstacleFrom constructs a [PolygonObstacle] from an unsafe.Pointer.
//
// A polygon-shaped impassable area in a 2D game world.
func PolygonObstacleFrom(ptr unsafe.Pointer) PolygonObstacle {
	return PolygonObstacle{
		Obstacle: ObstacleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PolygonObstacleClass) Alloc() PolygonObstacle {
	rv := objc.Send[PolygonObstacle](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PolygonObstacleClass) New() PolygonObstacle {
	rv := objc.Send[PolygonObstacle](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PolygonObstacle) Init() PolygonObstacle {
	rv := objc.Send[PolygonObstacle](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PolygonObstacle) Autorelease() PolygonObstacle {
	rv := objc.Send[PolygonObstacle](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPolygonObstacle creates a new PolygonObstacle instance.
func NewPolygonObstacle() PolygonObstacle {
	return getPolygonObstacleClass().New()
}




// Initializes a polygon obstacle with the specified list of vertices.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle/initWithPoints:count:
func NewPolygonObstacleWithPointsCount(points unsafe.Pointer, numPoints Iuintptr) PolygonObstacle {
	instance := getPolygonObstacleClass().Alloc()
	rv := objc.Send[PolygonObstacle](instance.ID, objc.Sel("initWithPoints:count:"), points, numPoints)
	rv.Autorelease()
	return rv
}


// Creates a polygon obstacle with the specified list of vertices.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle/obstacleWithPoints:count:
func (pc _PolygonObstacleClass) ObstacleWithPointsCount(points unsafe.Pointer, numPoints Iuintptr) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("obstacleWithPoints:count:"), points, numPoints)
	return rv
}

// Returns the point coordinates of the specified vertex.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle/vertex(at:)
func (p_ PolygonObstacle) VertexAtIndex(index uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("vertexAtIndex:"), index)
	return rv
}

// The number of vertices that define the polygon-shaped area of the obstacle.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle/vertexCount
func (p_ PolygonObstacle) VertexCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("vertexCount"))
	return rv
}


