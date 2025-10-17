// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuToolbarItem] class.
var menuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}

type _MenuToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	IToolbarItem
}

// A control that presents a menu in a window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem

type MenuToolbarItem struct {
	ToolbarItem
}

// MenuToolbarItemFrom constructs a [MenuToolbarItem] from an unsafe.Pointer.
//
// A control that presents a menu in a window’s toolbar.
func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}



