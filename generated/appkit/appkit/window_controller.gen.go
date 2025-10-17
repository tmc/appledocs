// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowController] class.
var WindowControllerClass objc.Class

func init() {
	WindowControllerClass = objc.GetClass("NSWindowController")
}

type WindowController struct {
	objc.ID
}

func WindowControllerFrom(ptr unsafe.Pointer) WindowController {
	return WindowController{
		ID: objc.ID(ptr),
	}
}



