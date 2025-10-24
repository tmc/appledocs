// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCSteeringWheelElement */


/* debug [class_header]: Header for GCSteeringWheelElement */
// The class instance for the [GCSteeringWheelElement] class.
var (
	GCSteeringWheelElementClass     _GCSteeringWheelElementClass
	GCSteeringWheelElementClassOnce sync.Once
)

func getGCSteeringWheelElementClass() _GCSteeringWheelElementClass {
	GCSteeringWheelElementClassOnce.Do(func() {
		GCSteeringWheelElementClass = _GCSteeringWheelElementClass{objc.GetClass("GCSteeringWheelElement")}
	})
	return GCSteeringWheelElementClass
}

type _GCSteeringWheelElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCSteeringWheelElement */
// An interface definition for the [GCSteeringWheelElement] class.
type IGCSteeringWheelElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCSteeringWheelElement */
	// properties:
	MaximumDegreesOfRotation() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCSteeringWheelElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCSteeringWheelElement */
// Alloc allocates a new instance without initialization.
func (gc _GCSteeringWheelElementClass) Alloc() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCSteeringWheelElementClass) New() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCSteeringWheelElement) Init() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCSteeringWheelElement) Autorelease() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCSteeringWheelElement creates a new GCSteeringWheelElement instance.
func NewGCSteeringWheelElement() GCSteeringWheelElement {
	return getGCSteeringWheelElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCSteeringWheelElement */
// The element that represents the wheel of a racing wheel controller.


// The element that represents the wheel of a racing wheel controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCSteeringWheelElement
type GCSteeringWheelElement struct {
	objectivec.Object
}

// GCSteeringWheelElementFrom constructs a [GCSteeringWheelElement] from an unsafe.Pointer.
//
// The element that represents the wheel of a racing wheel controller.
func GCSteeringWheelElementFrom(ptr unsafe.Pointer) GCSteeringWheelElement {
	return GCSteeringWheelElement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCSteeringWheelElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCSteeringWheelElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCSteeringWheelElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCSteeringWheelElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCSteeringWheelElement */

// The maximum number of degrees that the user can rotate the wheel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCSteeringWheelElement/maximumDegreesOfRotation
func (g_ GCSteeringWheelElement) MaximumDegreesOfRotation() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("maximumDegreesOfRotation"))
	return rv
}/* debug [instance_properties/getter]: maximumDegreesOfRotation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCSteeringWheelElement */



