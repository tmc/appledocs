// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TreeController] class.
var treeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}

type _TreeControllerClass struct {
	class objc.Class
}

// An interface definition for the [TreeController] class.
type ITreeController interface {
	IObjectController
}

// A bindings-compatible controller that manages a tree of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController

type TreeController struct {
	ObjectController
}

// TreeControllerFrom constructs a [TreeController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a tree of objects.
func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}



