// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ArrayController] class.
var arrayControllerClass = _ArrayControllerClass{objc.GetClass("NSArrayController")}

type _ArrayControllerClass struct {
	class objc.Class
}

// A bindings-compatible controller that manages a collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSArrayController

type ArrayController struct {
	ObjectController
}

// ArrayControllerFrom constructs a [ArrayController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a collection of objects.
func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}



