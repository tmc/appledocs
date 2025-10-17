// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItemCell] class.
var MenuItemCellClass objc.Class

func init() {
	MenuItemCellClass = objc.GetClass("NSMenuItemCell")
}

type MenuItemCell struct {
	objc.ID
}

func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{
		ID: objc.ID(ptr),
	}
}



