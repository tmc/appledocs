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
// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsControllerClass) Alloc() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (uc _UserDefaultsControllerClass) New() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserDefaultsController) Init() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserDefaultsController) Autorelease() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserDefaultsController creates a new UserDefaultsController instance.
func NewUserDefaultsController() UserDefaultsController {
	return userDefaultsControllerClass.New()
}




