// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UserDefaultsController] class.
var UserDefaultsControllerClass objc.Class

func init() {
	UserDefaultsControllerClass = objc.GetClass("NSUserDefaultsController")
}

type UserDefaultsController struct {
	objc.ID
}

func UserDefaultsControllerFrom(ptr unsafe.Pointer) UserDefaultsController {
	return UserDefaultsController{
		ID: objc.ID(ptr),
	}
}



