// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ObjectController] class.
var objectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}

type _ObjectControllerClass struct {
	class objc.Class
}

// An interface definition for the [ObjectController] class.
type IObjectController interface {
	IController
}

// A controller that can manage an object’s properties referenced by key-value paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSObjectController

type ObjectController struct {
	Controller
}

// ObjectControllerFrom constructs a [ObjectController] from an unsafe.Pointer.
//
// A controller that can manage an object’s properties referenced by key-value paths.
func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		Controller: ControllerFrom(ptr),
	}
}



