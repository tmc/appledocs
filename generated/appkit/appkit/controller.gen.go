// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Controller] class.
var ControllerClass objc.Class

func init() {
	ControllerClass = objc.GetClass("NSController")
}

type Controller struct {
	objc.ID
}

func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{
		ID: objc.ID(ptr),
	}
}



