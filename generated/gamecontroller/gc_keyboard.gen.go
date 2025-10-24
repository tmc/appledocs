// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCKeyboard */


/* debug [class_header]: Header for GCKeyboard */
// The class instance for the [GCKeyboard] class.
var (
	GCKeyboardClass     _GCKeyboardClass
	GCKeyboardClassOnce sync.Once
)

func getGCKeyboardClass() _GCKeyboardClass {
	GCKeyboardClassOnce.Do(func() {
		GCKeyboardClass = _GCKeyboardClass{objc.GetClass("GCKeyboard")}
	})
	return GCKeyboardClass
}

type _GCKeyboardClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCKeyboard */
// An interface definition for the [GCKeyboard] class.
type IGCKeyboard interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCKeyboard */
	// properties:
	KeyboardInput() IGCKeyboardInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCKeyboard */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCKeyboard */
// Alloc allocates a new instance without initialization.
func (gc _GCKeyboardClass) Alloc() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCKeyboardClass) New() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCKeyboard) Init() GCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCKeyboard) Autorelease() GCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCKeyboard creates a new GCKeyboard instance.
func NewGCKeyboard() GCKeyboard {
	return getGCKeyboardClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCKeyboard */
// An object that represents a physical keyboard connected to a device.
//
// To get the keyboard object and its input values, register for the (Swift) or (Objective-C) notification for when a keyboard connects to the device, or use the class property. Then get the input values from the keyboard object’s controller profile.


// An object that represents a physical keyboard connected to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard
type GCKeyboard struct {
	objectivec.Object
}

// GCKeyboardFrom constructs a [GCKeyboard] from an unsafe.Pointer.
//
// An object that represents a physical keyboard connected to a device.
func GCKeyboardFrom(ptr unsafe.Pointer) GCKeyboard {
	return GCKeyboard{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCKeyboard *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCKeyboard */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCKeyboard */

// The keyboard currently connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard/coalesced
func (gc _GCKeyboardClass) CoalescedKeyboard() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("coalescedKeyboard"))
	return rv
}/* debug [class_properties_class/property]: coalescedKeyboard */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCKeyboard */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCKeyboard */

// The keyboard currently connected to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard/coalesced
func (g_ GCKeyboard) CoalescedKeyboard() IGCKeyboard {
	rv := objc.Send[GCKeyboard](g_.ID, objc.Sel("coalescedKeyboard"))
	return rv
}/* debug [instance_properties/getter]: coalescedKeyboard */


// The controller profile for the keyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCKeyboard/keyboardInput
func (g_ GCKeyboard) KeyboardInput() IGCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](g_.ID, objc.Sel("keyboardInput"))
	return rv
}/* debug [instance_properties/getter]: keyboardInput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCKeyboard */



