// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCColor */


/* debug [class_header]: Header for GCColor */
// The class instance for the [GCColor] class.
var (
	GCColorClass     _GCColorClass
	GCColorClassOnce sync.Once
)

func getGCColorClass() _GCColorClass {
	GCColorClassOnce.Do(func() {
		GCColorClass = _GCColorClass{objc.GetClass("GCColor")}
	})
	return GCColorClass
}

type _GCColorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCColor */
// An interface definition for the [GCColor] class.
type IGCColor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCColor */
	// properties:
	Blue() float32
	Green() float32
	Red() float32
	Color() IGCColor
	SetColor(value IGCColor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCColor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCColor */
// Alloc allocates a new instance without initialization.
func (gc _GCColorClass) Alloc() GCColor {
	rv := objc.Send[GCColor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCColorClass) New() GCColor {
	rv := objc.Send[GCColor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCColor) Init() GCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCColor) Autorelease() GCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCColor creates a new GCColor instance.
func NewGCColor() GCColor {
	return getGCColorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCColor */
// The color of a device light.


// The color of a device light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor
type GCColor struct {
	objectivec.Object
}

// GCColorFrom constructs a [GCColor] from an unsafe.Pointer.
//
// The color of a device light.
func GCColorFrom(ptr unsafe.Pointer) GCColor {
	return GCColor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCColor */

// Creates a color with the specified red, green, and blue values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/init(red:green:blue:)
func NewGCColorWithRedGreenBlue(red float32, green float32, blue float32) GCColor {
	instance := getGCColorClass().Alloc()
	rv := objc.Send[GCColor](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCColorWithRedGreenBlue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCColor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCColor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCColor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCColor */

// The normalized value of the blue component ranging from 0 to 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/blue
func (g_ GCColor) Blue() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("blue"))
	return rv
}/* debug [instance_properties/getter]: blue */


// The normalized value of the green component ranging from 0 to 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/green
func (g_ GCColor) Green() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("green"))
	return rv
}/* debug [instance_properties/getter]: green */


// The normalized value of the red component ranging from 0 to 1.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCColor/red
func (g_ GCColor) Red() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("red"))
	return rv
}/* debug [instance_properties/getter]: red */


// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicelight/color
func (g_ GCColor) Color() IGCColor {
	rv := objc.Send[GCColor](g_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The color of a device’s light.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcdevicelight/color
func (g_ GCColor) SetColor(value IGCColor) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCColor */


