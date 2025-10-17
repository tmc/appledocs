// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ToolbarItem] class.
var ToolbarItemClass objc.Class

func init() {
	ToolbarItemClass = objc.GetClass("NSToolbarItem")
}

type ToolbarItem struct {
	objc.ID
}

func ToolbarItemFrom(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{
		ID: objc.ID(ptr),
	}
}




