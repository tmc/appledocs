// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CustomTouchBarItem] class.
var CustomTouchBarItemClass objc.Class

func init() {
	CustomTouchBarItemClass = objc.GetClass("NSCustomTouchBarItem")
}

type CustomTouchBarItem struct {
	objc.ID
}

func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		ID: objc.ID(ptr),
	}
}



