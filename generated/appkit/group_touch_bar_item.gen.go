// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [GroupTouchBarItem] class.
var GroupTouchBarItemClass objc.Class

func init() {
	GroupTouchBarItemClass = objc.GetClass("NSGroupTouchBarItem")
}

type GroupTouchBarItem struct {
	objc.ID
}

func GroupTouchBarItemFrom(ptr unsafe.Pointer) GroupTouchBarItem {
	return GroupTouchBarItem{
		ID: objc.ID(ptr),
	}
}



