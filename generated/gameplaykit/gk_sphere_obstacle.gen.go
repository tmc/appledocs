// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SphereObstacle] class.
var (
	SphereObstacleClass     _SphereObstacleClass
	SphereObstacleClassOnce sync.Once
)

func getSphereObstacleClass() _SphereObstacleClass {
	SphereObstacleClassOnce.Do(func() {
		SphereObstacleClass = _SphereObstacleClass{objc.GetClass("GKSphereObstacle")}
	})
	return SphereObstacleClass
}

type _SphereObstacleClass struct {
	class objc.Class
}

// An interface definition for the [SphereObstacle] class.
type ISphereObstacle interface {
	IObstacle
	// properties:
	Position() unsafe.Pointer
	SetPosition(value unsafe.Pointer)
	Radius() float32
	SetRadius(value float32)
	// methods:
}

// A spherical impassable volume to be avoided by agents.
//
// To make agents ( objects) avoid obstacles, create a goal with the method. Agents affected by an avoid-obstacles goal will attempt to move such that their radius never overlaps that of a spherical obstacle. To learn more about using goals and agents, see in .


// A spherical impassable volume to be avoided by agents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle
type SphereObstacle struct {
	Obstacle
}

// SphereObstacleFrom constructs a [SphereObstacle] from an unsafe.Pointer.
//
// A spherical impassable volume to be avoided by agents.
func SphereObstacleFrom(ptr unsafe.Pointer) SphereObstacle {
	return SphereObstacle{
		Obstacle: ObstacleFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SphereObstacleClass) Alloc() SphereObstacle {
	rv := objc.Send[SphereObstacle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SphereObstacleClass) New() SphereObstacle {
	rv := objc.Send[SphereObstacle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SphereObstacle) Init() SphereObstacle {
	rv := objc.Send[SphereObstacle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SphereObstacle) Autorelease() SphereObstacle {
	rv := objc.Send[SphereObstacle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSphereObstacle creates a new SphereObstacle instance.
func NewSphereObstacle() SphereObstacle {
	return getSphereObstacleClass().New()
}



// Initializes a spherical obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/init(radius:)
func NewSphereObstacleWithRadius(radius float32) SphereObstacle {
	instance := getSphereObstacleClass().Alloc()
	rv := objc.Send[SphereObstacle](instance.ID, objc.Sel("initWithRadius:"), radius)
	rv.Autorelease()
	return rv
}



// Creates a spherical obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/obstacleWithRadius:
func (sc _SphereObstacleClass) ObstacleWithRadius(radius float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("obstacleWithRadius:"), radius)
	return rv
}


// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/position
func (s_ SphereObstacle) Position() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("position"))
	return rv
}


// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/position
func (s_ SphereObstacle) SetPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPosition:"), value)
}


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/radius
func (s_ SphereObstacle) Radius() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("radius"))
	return rv
}


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/radius
func (s_ SphereObstacle) SetRadius(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRadius:"), value)
}


