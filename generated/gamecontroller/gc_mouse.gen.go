// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCMouse */


/* debug [class_header]: Header for GCMouse */
// The class instance for the [GCMouse] class.
var (
	GCMouseClass     _GCMouseClass
	GCMouseClassOnce sync.Once
)

func getGCMouseClass() _GCMouseClass {
	GCMouseClassOnce.Do(func() {
		GCMouseClass = _GCMouseClass{objc.GetClass("GCMouse")}
	})
	return GCMouseClass
}

type _GCMouseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCMouse */
// An interface definition for the [GCMouse] class.
type IGCMouse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCMouse */
	// properties:
	MouseInput() IGCMouseInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCMouse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCMouse */
// Alloc allocates a new instance without initialization.
func (gc _GCMouseClass) Alloc() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCMouseClass) New() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMouse) Init() GCMouse {
	rv := objc.Send[GCMouse](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMouse) Autorelease() GCMouse {
	rv := objc.Send[GCMouse](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMouse creates a new GCMouse instance.
func NewGCMouse() GCMouse {
	return getGCMouseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCMouse */
// An object that represents a physical mouse connected to a device.
//
// To get a mouse object and its input values, register for the (Swift) or (Objective-C) notification for when a mouse connects to the device. Then register for the (Swift) or (Objective-C) notification for when it becomes the mouse. Alternatively, use the class property or the class method to get a mouse object. Then get the current input values from the mouse object’s controller profile.


// An object that represents a physical mouse connected to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse
type GCMouse struct {
	objectivec.Object
}

// GCMouseFrom constructs a [GCMouse] from an unsafe.Pointer.
//
// An object that represents a physical mouse connected to a device.
func GCMouseFrom(ptr unsafe.Pointer) GCMouse {
	return GCMouse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCMouse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCMouse */

// Returns any mice that the user connects to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/mice()
func (gc _GCMouseClass) Mice() []GCMouse {
	rv := objc.Send[[]GCMouse](objc.ID(gc.class), objc.Sel("mice"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Mice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCMouse */

// The most recent mouse that the user connects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/current
func (gc _GCMouseClass) Current() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("current"))
	return rv
}/* debug [class_properties_class/property]: current */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCMouse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCMouse */

// The most recent mouse that the user connects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/current
func (g_ GCMouse) Current() IGCMouse {
	rv := objc.Send[GCMouse](g_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// The controller profile for the mouse device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMouse/mouseInput
func (g_ GCMouse) MouseInput() IGCMouseInput {
	rv := objc.Send[GCMouseInput](g_.ID, objc.Sel("mouseInput"))
	return rv
}/* debug [instance_properties/getter]: mouseInput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCMouse */



