// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GCDeviceCursor */


/* debug [class_header]: Header for GCDeviceCursor */
// The class instance for the [GCDeviceCursor] class.
var (
	GCDeviceCursorClass     _GCDeviceCursorClass
	GCDeviceCursorClassOnce sync.Once
)

func getGCDeviceCursorClass() _GCDeviceCursorClass {
	GCDeviceCursorClassOnce.Do(func() {
		GCDeviceCursorClass = _GCDeviceCursorClass{objc.GetClass("GCDeviceCursor")}
	})
	return GCDeviceCursorClass
}

type _GCDeviceCursorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCDeviceCursor */
// An interface definition for the [GCDeviceCursor] class.
type IGCDeviceCursor interface {
	IGCControllerDirectionPad
	
/* debug [class_interface_properties]: Properties for GCDeviceCursor */
	// properties:
	XAxis() IGCControllerAxisInput
	SetXAxis(value IGCControllerAxisInput)
	YAxis() IGCControllerAxisInput
	SetYAxis(value IGCControllerAxisInput)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCDeviceCursor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCDeviceCursor */
// Alloc allocates a new instance without initialization.
func (gc _GCDeviceCursorClass) Alloc() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCDeviceCursorClass) New() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCDeviceCursor) Init() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCDeviceCursor) Autorelease() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceCursor creates a new GCDeviceCursor instance.
func NewGCDeviceCursor() GCDeviceCursor {
	return getGCDeviceCursorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCDeviceCursor */
// A control element for the cursor used as a directional pad.
//
// This controller element allows you to use the cursor as a directional pad with the values of the and elements scaled to the width and height of the screen, not ranging from to .


// A control element for the cursor used as a directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceCursor
type GCDeviceCursor struct {
	GCControllerDirectionPad
}

// GCDeviceCursorFrom constructs a [GCDeviceCursor] from an unsafe.Pointer.
//
// A control element for the cursor used as a directional pad.
func GCDeviceCursorFrom(ptr unsafe.Pointer) GCDeviceCursor {
	return GCDeviceCursor{
		GCControllerDirectionPad: GCControllerDirectionPadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCDeviceCursor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCDeviceCursor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCDeviceCursor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCDeviceCursor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCDeviceCursor */

// The x-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/xaxis
func (g_ GCDeviceCursor) XAxis() IGCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("xAxis"))
	return rv
}/* debug [instance_properties/getter]: xAxis */


// The x-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/xaxis
func (g_ GCDeviceCursor) SetXAxis(value IGCControllerAxisInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setXAxis:"), value)
}/* debug [instance_properties/setter]: xAxis */


// The y-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/yaxis
func (g_ GCDeviceCursor) YAxis() IGCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g_.ID, objc.Sel("yAxis"))
	return rv
}/* debug [instance_properties/getter]: yAxis */


// The y-axis element of the directional pad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerdirectionpad/yaxis
func (g_ GCDeviceCursor) SetYAxis(value IGCControllerAxisInput) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setYAxis:"), value)
}/* debug [instance_properties/setter]: yAxis */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCDeviceCursor */



