
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserDefaultsController] class.
var UserDefaultsControllerClass _UserDefaultsControllerClass

func init() {
	UserDefaultsControllerClass = _UserDefaultsControllerClass{objc.GetClass("NSUserDefaultsController")}
}

type _UserDefaultsControllerClass struct {
	objc.Class
}

// An interface definition for the [UserDefaultsController] class.
type IUserDefaultsController interface {
	ID() objc.ID
}

type UserDefaultsController struct {
	id objc.ID
}

func UserDefaultsControllerFrom(ptr unsafe.Pointer) UserDefaultsController {
	return UserDefaultsController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (u_ UserDefaultsController) ID() objc.ID {
	return u_.id
}

// Alloc allocates a new instance without initialization.
func (uc _UserDefaultsControllerClass) Alloc() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (uc _UserDefaultsControllerClass) New() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](objc.ID(uc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewUserDefaultsController creates and returns a new initialized instance.
func NewUserDefaultsController() UserDefaultsController {
	return UserDefaultsControllerClass.New()
}

// Init initializes the instance.
func (u_ UserDefaultsController) Init() UserDefaultsController {
	rv := objc.Send[UserDefaultsController](u_.ID(), selInit)
	return rv
}
