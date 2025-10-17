// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PageController] class.
var PageControllerClass objc.Class

func init() {
	PageControllerClass = objc.GetClass("NSPageController")
}

type PageController struct {
	objc.ID
}

func PageControllerFrom(ptr unsafe.Pointer) PageController {
	return PageController{
		ID: objc.ID(ptr),
	}
}



