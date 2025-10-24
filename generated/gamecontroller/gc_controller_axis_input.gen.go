// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCControllerAxisInput */


/* debug [class_header]: Header for GCControllerAxisInput */
// The class instance for the [GCControllerAxisInput] class.
var (
	GCControllerAxisInputClass     _GCControllerAxisInputClass
	GCControllerAxisInputClassOnce sync.Once
)

func getGCControllerAxisInputClass() _GCControllerAxisInputClass {
	GCControllerAxisInputClassOnce.Do(func() {
		GCControllerAxisInputClass = _GCControllerAxisInputClass{objc.GetClass("GCControllerAxisInput")}
	})
	return GCControllerAxisInputClass
}

type _GCControllerAxisInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerAxisInput */
// An interface definition for the [GCControllerAxisInput] class.
type IGCControllerAxisInput interface {
	IGCControllerElement
	
/* debug [class_interface_properties]: Properties for GCControllerAxisInput */
	// properties:
	Value() float32
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerAxisInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerAxisInput */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerAxisInputClass) Alloc() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCControllerAxisInputClass) New() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerAxisInput) Init() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerAxisInput) Autorelease() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerAxisInput creates a new GCControllerAxisInput instance.
func NewGCControllerAxisInput() GCControllerAxisInput {
	return getGCControllerAxisInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerAxisInput */
// A control element that tracks movement along an axis.
//
// A object represents the value of a physical controller’s axis. For example, a has x-axis and y-axis subelements.


// A control element that tracks movement along an axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerAxisInput
type GCControllerAxisInput struct {
	GCControllerElement
}

// GCControllerAxisInputFrom constructs a [GCControllerAxisInput] from an unsafe.Pointer.
//
// A control element that tracks movement along an axis.
func GCControllerAxisInputFrom(ptr unsafe.Pointer) GCControllerAxisInput {
	return GCControllerAxisInput{
		GCControllerElement: GCControllerElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerAxisInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerAxisInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerAxisInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerAxisInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerAxisInput */

// The current value of the axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerAxisInput/value
func (g_ GCControllerAxisInput) Value() float32 {
	rv := objc.Send[float32](g_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The block that the element calls when the user changes the axis value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerAxisInput/valueChangedHandler
func (g_ GCControllerAxisInput) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the element calls when the user changes the axis value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerAxisInput/valueChangedHandler
func (g_ GCControllerAxisInput) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerAxisInput */



