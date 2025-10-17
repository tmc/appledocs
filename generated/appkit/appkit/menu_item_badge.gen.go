// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MenuItemBadge] class.
var MenuItemBadgeClass objc.Class

func init() {
	MenuItemBadgeClass = objc.GetClass("NSMenuItemBadge")
}

type MenuItemBadge struct {
	objc.ID
}

func MenuItemBadgeFrom(ptr unsafe.Pointer) MenuItemBadge {
	return MenuItemBadge{
		ID: objc.ID(ptr),
	}
}



