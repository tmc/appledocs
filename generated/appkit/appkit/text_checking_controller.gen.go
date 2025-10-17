// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextCheckingController] class.
var TextCheckingControllerClass objc.Class

func init() {
	TextCheckingControllerClass = objc.GetClass("NSTextCheckingController")
}

type TextCheckingController struct {
	objc.ID
}

func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{
		ID: objc.ID(ptr),
	}
}




