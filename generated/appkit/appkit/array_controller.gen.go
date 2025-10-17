// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ArrayController] class.
var ArrayControllerClass objc.Class

func init() {
	ArrayControllerClass = objc.GetClass("NSArrayController")
}

type ArrayController struct {
	objc.ID
}

func ArrayControllerFrom(ptr unsafe.Pointer) ArrayController {
	return ArrayController{
		ID: objc.ID(ptr),
	}
}



