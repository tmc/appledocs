// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKObstacle */


/* debug [class_header]: Header for GKObstacle */
// The class instance for the [Obstacle] class.
var (
	ObstacleClass     _ObstacleClass
	ObstacleClassOnce sync.Once
)

func getObstacleClass() _ObstacleClass {
	ObstacleClassOnce.Do(func() {
		ObstacleClass = _ObstacleClass{objc.GetClass("GKObstacle")}
	})
	return ObstacleClass
}

type _ObstacleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Obstacle */
// An interface definition for the [Obstacle] class.
type IObstacle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Obstacle */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Obstacle */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Obstacle */
// Alloc allocates a new instance without initialization.
func (oc _ObstacleClass) Alloc() Obstacle {
	rv := objc.Send[Obstacle](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _ObstacleClass) New() Obstacle {
	rv := objc.Send[Obstacle](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Obstacle) Init() Obstacle {
	rv := objc.Send[Obstacle](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Obstacle) Autorelease() Obstacle {
	rv := objc.Send[Obstacle](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObstacle creates a new Obstacle instance.
func NewObstacle() Obstacle {
	return getObstacleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Obstacle */
// The abstract base class for objects representing impassable areas in a game world.
//
// You do not use this class directly; instead, create instances of its concrete subclasses , , and . To make agents ( objects) avoid obstacles, create a goal with the goalToAvoidObstacles:timeBeforeCollisionToAvoid: method. To learn more about using goals and agents, see in . For more information, see .


// The abstract base class for objects representing impassable areas in a game world.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKObstacle
type Obstacle struct {
	objectivec.Object
}

// ObstacleFrom constructs a [Obstacle] from an unsafe.Pointer.
//
// The abstract base class for objects representing impassable areas in a game world.
func ObstacleFrom(ptr unsafe.Pointer) Obstacle {
	return Obstacle{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Obstacle *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Obstacle */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Obstacle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Obstacle */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Obstacle */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKObstacle */



