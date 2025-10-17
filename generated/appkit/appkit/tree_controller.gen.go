// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TreeController] class.
var TreeControllerClass objc.Class

func init() {
	TreeControllerClass = objc.GetClass("NSTreeController")
}

type TreeController struct {
	objc.ID
}

func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ID: objc.ID(ptr),
	}
}




