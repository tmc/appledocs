// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var ButtonTouchBarItemClass objc.Class

func init() {
	ButtonTouchBarItemClass = objc.GetClass("NSButtonTouchBarItem")
}

type ButtonTouchBarItem struct {
	objc.ID
}

func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		ID: objc.ID(ptr),
	}
}




