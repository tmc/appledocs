// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UserDefaultsController] class.
var userDefaultsControllerClass = _UserDefaultsControllerClass{objc.GetClass("NSUserDefaultsController")}

type _UserDefaultsControllerClass struct {
	class objc.Class
}

// An interface definition for the [UserDefaultsController] class.
type IUserDefaultsController interface {
	IController
}

// A controller that accesses user preference information for your app from the user’s defaults database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController

type UserDefaultsController struct {
	Controller
}

// UserDefaultsControllerFrom constructs a [UserDefaultsController] from an unsafe.Pointer.
//
// A controller that accesses user preference information for your app from the user’s defaults database.
func UserDefaultsControllerFrom(ptr unsafe.Pointer) UserDefaultsController {
	return UserDefaultsController{
		Controller: ControllerFrom(ptr),
	}
}



