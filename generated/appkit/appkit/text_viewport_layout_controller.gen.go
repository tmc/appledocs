// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextViewportLayoutController] class.
var TextViewportLayoutControllerClass objc.Class

func init() {
	TextViewportLayoutControllerClass = objc.GetClass("NSTextViewportLayoutController")
}

type TextViewportLayoutController struct {
	objc.ID
}

func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{
		ID: objc.ID(ptr),
	}
}




