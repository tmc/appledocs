// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ObjectController] class.
var ObjectControllerClass objc.Class

func init() {
	ObjectControllerClass = objc.GetClass("NSObjectController")
}

type ObjectController struct {
	objc.ID
}

func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		ID: objc.ID(ptr),
	}
}




