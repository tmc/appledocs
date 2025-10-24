// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKCircleObstacle */


/* debug [class_header]: Header for GKCircleObstacle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CircleObstacle */
// An interface definition for the [CircleObstacle] class.
type ICircleObstacle interface {
	IObstacle
	
/* debug [class_interface_properties]: Properties for CircleObstacle */
	// properties:
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
	Radius() float32
	SetRadius(value float32)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CircleObstacle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CircleObstacle */
// Alloc allocates a new instance without initialization.
func (cc _CircleObstacleClass) Alloc() CircleObstacle {
	rv := objc.Send[CircleObstacle](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CircleObstacle */
// A circular impassable area to be avoided by agents.
//
// To make agents ( objects) avoid obstacles, create a goal with the method. Agents affected by an avoid-obstacles goal will attempt to move such that their radius never overlaps that of a circular obstacle. To learn more about using goals and agents, see in .


// A circular impassable area to be avoided by agents.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CircleObstacle */

// Initializes a circular obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/init(radius:)
func NewCircleObstacleWithRadius(radius float32) CircleObstacle {
	instance := getCircleObstacleClass().Alloc()
	rv := objc.Send[CircleObstacle](instance.ID, objc.Sel("initWithRadius:"), radius)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCircleObstacleWithRadius */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CircleObstacle */

// Creates a circular obstacle with the specified radius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/obstacleWithRadius:
func (cc _CircleObstacleClass) ObstacleWithRadius(radius float32) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("obstacleWithRadius:"), radius)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ObstacleWithRadius) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CircleObstacle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CircleObstacle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CircleObstacle */

// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/position
func (c_ CircleObstacle) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The position of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/position
func (c_ CircleObstacle) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/radius
func (c_ CircleObstacle) Radius() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("radius"))
	return rv
}/* debug [instance_properties/getter]: radius */


// The radius of the obstacle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKCircleObstacle/radius
func (c_ CircleObstacle) SetRadius(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRadius:"), value)
}/* debug [instance_properties/setter]: radius */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKCircleObstacle */


