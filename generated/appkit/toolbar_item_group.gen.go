// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolbarItemGroup] class.
var toolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}

type _ToolbarItemGroupClass struct {
	class objc.Class
}

// A group of subitems in a toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup

type ToolbarItemGroup struct {
	ToolbarItem
}

// ToolbarItemGroupFrom constructs a [ToolbarItemGroup] from an unsafe.Pointer.
//
// A group of subitems in a toolbar item.
func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}



