// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKDialogController */


/* debug [class_header]: Header for GKDialogController */
// The class instance for the [DialogController] class.
var (
	DialogControllerClass     _DialogControllerClass
	DialogControllerClassOnce sync.Once
)

func getDialogControllerClass() _DialogControllerClass {
	DialogControllerClassOnce.Do(func() {
		DialogControllerClass = _DialogControllerClass{objc.GetClass("GKDialogController")}
	})
	return DialogControllerClass
}

type _DialogControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DialogController */
// An interface definition for the [DialogController] class.
type IDialogController interface {
	appkit.IResponder
	
/* debug [class_interface_properties]: Properties for DialogController */
	// properties:
	ParentWindow() appkit.Window
	SetParentWindow(value appkit.Window)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DialogController */
	// methods:
	Dismiss(sender objc.IObject)
	PresentViewController(viewController unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DialogController */
// Alloc allocates a new instance without initialization.
func (dc _DialogControllerClass) Alloc() DialogController {
	rv := objc.Send[DialogController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DialogControllerClass) New() DialogController {
	rv := objc.Send[DialogController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DialogController) Init() DialogController {
	rv := objc.Send[DialogController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DialogController) Autorelease() DialogController {
	rv := objc.Send[DialogController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDialogController creates a new DialogController instance.
func NewDialogController() DialogController {
	return getDialogControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DialogController */
// An object that provides the ability to present the dashboard in macOS games.
//
// For macOS games, use a object to present the dashboard from which players can browse and manage their Game Center data. Initialize a new object, as you would for an iOS game, specifying the state and setting its delegate. Then get the singleton dialog controller using the class method, or initialize a new object. To present the dashboard, set the property to the window that should display the dashboard and then call the method, passing the object. When the player closes the dashboard, GameKit calls the delegate method. Implement this method to dismiss the shared dialog controller using the method.


// An object that provides the ability to present the dashboard in macOS games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController
type DialogController struct {
	appkit.Responder
}

// DialogControllerFrom constructs a [DialogController] from an unsafe.Pointer.
//
// An object that provides the ability to present the dashboard in macOS games.
func DialogControllerFrom(ptr unsafe.Pointer) DialogController {
	return DialogController{
		Responder: appkit.ResponderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DialogController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DialogController */

// Retrieves the shared instance of the dialog controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/shared()
func (dc _DialogControllerClass) SharedDialogController() IDialogController {
	rv := objc.Send[DialogController](objc.ID(dc.class), objc.Sel("sharedDialogController"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedDialogController) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DialogController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DialogController */

// Dismisses the dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/dismiss(_:)
func (d_ DialogController) Dismiss(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("dismiss:"), sender)
}/* debug [instance_methods/method]: Dismiss */


// Presents the dashboard in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/present(_:)
func (d_ DialogController) PresentViewController(viewController unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentViewController:"), viewController)
	return rv
}/* debug [instance_methods/method]: PresentViewController */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DialogController */

// The window that displays the dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/parentWindow
func (d_ DialogController) ParentWindow() appkit.Window {
	rv := objc.Send[appkit.Window](d_.ID, objc.Sel("parentWindow"))
	return rv
}/* debug [instance_properties/getter]: parentWindow */


// The window that displays the dashboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/parentWindow
func (d_ DialogController) SetParentWindow(value appkit.Window) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParentWindow:"), value)
}/* debug [instance_properties/setter]: parentWindow */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKDialogController */



