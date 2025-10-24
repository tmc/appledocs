// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKSphereObstacle */


/* debug [class_header]: Header for GKSphereObstacle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SphereObstacle */
// An interface definition for the [SphereObstacle] class.
type ISphereObstacle interface {
	IObstacle
	
/* debug [class_interface_properties]: Properties for SphereObstacle */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
	Radius() float32
	SetRadius(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SphereObstacle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SphereObstacle */
// Alloc allocates a new instance without initialization.
func (sc _SphereObstacleClass) Alloc() SphereObstacle {
	rv := objc.Send[SphereObstacle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SphereObstacle */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SphereObstacle */

// Initializes a spherical obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/init(radius:)
func NewSphereObstacleWithRadius(radius float32) SphereObstacle {
	instance := getSphereObstacleClass().Alloc()
	rv := objc.Send[SphereObstacle](instance.ID, objc.Sel("initWithRadius:"), radius)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSphereObstacleWithRadius */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SphereObstacle */

// Creates a spherical obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/obstacleWithRadius:
func (sc _SphereObstacleClass) ObstacleWithRadius(radius float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("obstacleWithRadius:"), radius)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ObstacleWithRadius) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SphereObstacle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SphereObstacle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SphereObstacle */

// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/position
func (s_ SphereObstacle) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/position
func (s_ SphereObstacle) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/radius
func (s_ SphereObstacle) Radius() float32 {
	rv := objc.Send[float32](s_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSphereObstacle/radius
func (s_ SphereObstacle) SetRadius(value float32) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRadius:"), value)
}/* debug [instance_properties/setter]: radius */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKSphereObstacle */


