// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuToolbarItem] class.
var MenuToolbarItemClass objc.Class

func init() {
	MenuToolbarItemClass = objc.GetClass("NSMenuToolbarItem")
}

type MenuToolbarItem struct {
	objc.ID
}

func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ID: objc.ID(ptr),
	}
}




