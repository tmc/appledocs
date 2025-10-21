// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [DialogController] class.
type IDialogController interface {
	appkit.IResponder
	Dismiss(sender objectivec.IObject)
}

// An object that provides the ability to present the dashboard in macOS games.
//
// For macOS games, use a object to present the dashboard from which players can browse and manage their Game Center data. Initialize a new object, as you would for an iOS game, specifying the state and setting its delegate. Then get the singleton dialog controller using the class method, or initialize a new object. To present the dashboard, set the property to the window that should display the dashboard and then call the method, passing the object. When the player closes the dashboard, GameKit calls the delegate method. Implement this method to dismiss the shared dialog controller using the method.
//
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

// Alloc allocates a new instance without initialization.
func (dc _DialogControllerClass) Alloc() DialogController {
	rv := objc.Send[DialogController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Dismisses the dashboard.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKDialogController/dismiss(_:)
func (d_ DialogController) Dismiss(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("dismiss:"), sender)
}

// The window that displays the dashboard.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkdialogcontroller/parentwindow
func (d_ DialogController) ParentWindow() appkit.Window {
	rv := objc.Send[appkit.Window](d_.ID, objc.Sel("parentWindow"))
	return rv
}


// SetParentWindow sets the value of the parentWindow property.
// The window that displays the dashboard.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkdialogcontroller/parentwindow
func (d_ DialogController) SetParentWindow(value appkit.IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParentWindow:"), value)
}



