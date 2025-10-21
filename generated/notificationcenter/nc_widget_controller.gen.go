// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NCWidgetController] class.
type INCWidgetController interface {
	objectivec.IObject
	SetHasContentForWidgetWithBundleIdentifier(flag bool, bundleID string)
}

// An object used to specify whether a Today widget has content to display.
//
// The class defines an object that both a Today widget and the containing app that delivers the widget can use to specify whether the widget has content to display. Because this class helps a widget and its containing app coordinate the display of the widget’s content, a widget that doesn’t communicate with its containing app is unlikely to use this class. Typically, a widget appears in the Today view when it has content to display. If a currently running widget no longer has content to display, it can get a widget controller and set the flag in the method to . If the containing app later determines that there is content this widget should display, the app can get a widget controller and update the flag, even while the widget isn’t running. The class should not be subclassed.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NCWidgetControllerClass) Alloc() NCWidgetController {
	rv := objc.Send[NCWidgetController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController/default()
func (nc _NCWidgetControllerClass) DefaultWidgetController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("defaultWidgetController"))
	return rv
}

// Returns a widget controller used to specify whether a widget has content to display.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController/widgetController()
func (nc _NCWidgetControllerClass) WidgetController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("widgetController"))
	return rv
}

// Sets whether the specified widget has content to display.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetController/setHasContent(_:forWidgetWithBundleIdentifier:)
func (n_ NCWidgetController) SetHasContentForWidgetWithBundleIdentifier(flag bool, bundleID string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasContent:forWidgetWithBundleIdentifier:"), flag, objc.String(bundleID))
}



