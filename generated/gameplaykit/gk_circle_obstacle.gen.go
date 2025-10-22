// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CircleObstacle] class.
var (
	CircleObstacleClass     _CircleObstacleClass
	CircleObstacleClassOnce sync.Once
)

func getCircleObstacleClass() _CircleObstacleClass {
	CircleObstacleClassOnce.Do(func() {
		CircleObstacleClass = _CircleObstacleClass{objc.GetClass("GKCircleObstacle")}
	})
	return CircleObstacleClass
}

type _CircleObstacleClass struct {
	class objc.Class
}

// An interface definition for the [CircleObstacle] class.
type ICircleObstacle interface {
	IObstacle
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	Radius() float32
	SetRadius(value float32)
}

// A circular impassable area to be avoided by agents.
//
// To make agents ( objects) avoid obstacles, create a goal with the method. Agents affected by an avoid-obstacles goal will attempt to move such that their radius never overlaps that of a circular obstacle. To learn more about using goals and agents, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle
type CircleObstacle struct {
	Obstacle
}

// CircleObstacleFrom constructs a [CircleObstacle] from an unsafe.Pointer.
//
// A circular impassable area to be avoided by agents.
func CircleObstacleFrom(ptr unsafe.Pointer) CircleObstacle {
	return CircleObstacle{
		Obstacle: ObstacleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CircleObstacleClass) Alloc() CircleObstacle {
	rv := objc.Send[CircleObstacle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CircleObstacleClass) New() CircleObstacle {
	rv := objc.Send[CircleObstacle](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CircleObstacle) Init() CircleObstacle {
	rv := objc.Send[CircleObstacle](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CircleObstacle) Autorelease() CircleObstacle {
	rv := objc.Send[CircleObstacle](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCircleObstacle creates a new CircleObstacle instance.
func NewCircleObstacle() CircleObstacle {
	return getCircleObstacleClass().New()
}




// Initializes a circular obstacle with the specified radius.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/init(radius:)
func NewCircleObstacleWithRadius(radius float32) CircleObstacle {
	instance := getCircleObstacleClass().Alloc()
	rv := objc.Send[CircleObstacle](instance.ID, objc.Sel("initWithRadius:"), radius)
	rv.Autorelease()
	return rv
}


// Creates a circular obstacle with the specified radius.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/obstacleWithRadius:
func (cc _CircleObstacleClass) ObstacleWithRadius(radius float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("obstacleWithRadius:"), radius)
	return rv
}

// The position of the obstacle.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/position
func (c_ CircleObstacle) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The position of the obstacle.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/position
func (c_ CircleObstacle) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPosition:"), value)
}

// The radius of the obstacle.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/radius
func (c_ CircleObstacle) Radius() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("radius"))
	return rv
}


// SetRadius sets the value of the radius property.
// The radius of the obstacle.

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/radius
func (c_ CircleObstacle) SetRadius(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRadius:"), value)
}


