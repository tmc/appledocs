// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCDeviceLight */


/* debug [class_header]: Header for GCDeviceLight */
// The class instance for the [GCDeviceLight] class.
var (
	GCDeviceLightClass     _GCDeviceLightClass
	GCDeviceLightClassOnce sync.Once
)

func getGCDeviceLightClass() _GCDeviceLightClass {
	GCDeviceLightClassOnce.Do(func() {
		GCDeviceLightClass = _GCDeviceLightClass{objc.GetClass("GCDeviceLight")}
	})
	return GCDeviceLightClass
}

type _GCDeviceLightClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDeviceLight */
// An interface definition for the [GCDeviceLight] class.
type IGCDeviceLight interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCDeviceLight */
	// properties:
	Color() IGCColor
	SetColor(value IGCColor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCDeviceLight */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDeviceLight */
// Alloc allocates a new instance without initialization.
func (gc _GCDeviceLightClass) Alloc() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDeviceLightClass) New() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceLight) Init() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceLight) Autorelease() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceLight creates a new GCDeviceLight instance.
func NewGCDeviceLight() GCDeviceLight {
	return getGCDeviceLightClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDeviceLight */
// The colored light on a device.


// The colored light on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceLight
type GCDeviceLight struct {
	objectivec.Object
}

// GCDeviceLightFrom constructs a [GCDeviceLight] from an unsafe.Pointer.
//
// The colored light on a device.
func GCDeviceLightFrom(ptr unsafe.Pointer) GCDeviceLight {
	return GCDeviceLight{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDeviceLight *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDeviceLight */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDeviceLight */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDeviceLight */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDeviceLight */

// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceLight/color
func (g_ GCDeviceLight) Color() IGCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceLight/color
func (g_ GCDeviceLight) SetColor(value IGCColor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDeviceLight */



