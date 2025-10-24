// Code generated from Apple documentation for AVRouting. DO NOT EDIT.

package avrouting

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCustomRoutingController */


/* debug [class_header]: Header for AVCustomRoutingController */
// The class instance for the [CustomRoutingController] class.
var (
	CustomRoutingControllerClass     _CustomRoutingControllerClass
	CustomRoutingControllerClassOnce sync.Once
)

func getCustomRoutingControllerClass() _CustomRoutingControllerClass {
	CustomRoutingControllerClassOnce.Do(func() {
		CustomRoutingControllerClass = _CustomRoutingControllerClass{objc.GetClass("AVCustomRoutingController")}
	})
	return CustomRoutingControllerClass
}

type _CustomRoutingControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomRoutingController */
// An interface definition for the [CustomRoutingController] class.
type ICustomRoutingController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CustomRoutingController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomRoutingController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomRoutingController */
// Alloc allocates a new instance without initialization.
func (cc _CustomRoutingControllerClass) Alloc() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomRoutingControllerClass) New() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomRoutingController) Init() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomRoutingController) Autorelease() CustomRoutingController {
	rv := objc.Send[CustomRoutingController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomRoutingController creates a new CustomRoutingController instance.
func NewCustomRoutingController() CustomRoutingController {
	return getCustomRoutingControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomRoutingController */
// An object that manages the connection from a device to a destination.
//
// A routing controller also informs its object about which routes the user previously authorized, so it can reconnect, if appropriate.


// An object that manages the connection from a device to a destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController
type CustomRoutingController struct {
	objectivec.Object
}

// CustomRoutingControllerFrom constructs a [CustomRoutingController] from an unsafe.Pointer.
//
// An object that manages the connection from a device to a destination.
func CustomRoutingControllerFrom(ptr unsafe.Pointer) CustomRoutingController {
	return CustomRoutingController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomRoutingController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomRoutingController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomRoutingController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomRoutingController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomRoutingController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCustomRoutingController */


