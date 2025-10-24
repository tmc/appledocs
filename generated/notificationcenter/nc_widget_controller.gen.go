// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NCWidgetController */


/* debug [class_header]: Header for NCWidgetController */
// The class instance for the [NCWidgetController] class.
var (
	NCWidgetControllerClass     _NCWidgetControllerClass
	NCWidgetControllerClassOnce sync.Once
)

func getNCWidgetControllerClass() _NCWidgetControllerClass {
	NCWidgetControllerClassOnce.Do(func() {
		NCWidgetControllerClass = _NCWidgetControllerClass{objc.GetClass("NCWidgetController")}
	})
	return NCWidgetControllerClass
}

type _NCWidgetControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NCWidgetController */
// An interface definition for the [NCWidgetController] class.
type INCWidgetController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NCWidgetController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NCWidgetController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NCWidgetController */
// Alloc allocates a new instance without initialization.
func (nc _NCWidgetControllerClass) Alloc() NCWidgetController {
	rv := objc.Send[NCWidgetController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NCWidgetControllerClass) New() NCWidgetController {
	rv := objc.Send[NCWidgetController](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NCWidgetController) Init() NCWidgetController {
	rv := objc.Send[NCWidgetController](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NCWidgetController) Autorelease() NCWidgetController {
	rv := objc.Send[NCWidgetController](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNCWidgetController creates a new NCWidgetController instance.
func NewNCWidgetController() NCWidgetController {
	return getNCWidgetControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NCWidgetController */
// An object used to specify whether a Today widget has content to display.
//
// The class defines an object that both a Today widget and the containing app that delivers the widget can use to specify whether the widget has content to display. Because this class helps a widget and its containing app coordinate the display of the widget’s content, a widget that doesn’t communicate with its containing app is unlikely to use this class. Typically, a widget appears in the Today view when it has content to display. If a currently running widget no longer has content to display, it can get a widget controller and set the flag in the method to . If the containing app later determines that there is content this widget should display, the app can get a widget controller and update the flag, even while the widget isn’t running. The class should not be subclassed.


// An object used to specify whether a Today widget has content to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController
type NCWidgetController struct {
	objectivec.Object
}

// NCWidgetControllerFrom constructs a [NCWidgetController] from an unsafe.Pointer.
//
// An object used to specify whether a Today widget has content to display.
func NCWidgetControllerFrom(ptr unsafe.Pointer) NCWidgetController {
	return NCWidgetController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NCWidgetController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NCWidgetController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController/default()
func (nc _NCWidgetControllerClass) DefaultWidgetController() NCWidgetController {
	rv := objc.Send[NCWidgetController](objc.ID(nc.class), objc.Sel("defaultWidgetController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultWidgetController) */


// Returns a widget controller used to specify whether a widget has content to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController/widgetController()
func (nc _NCWidgetControllerClass) WidgetController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("widgetController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WidgetController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NCWidgetController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NCWidgetController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NCWidgetController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NCWidgetController */



