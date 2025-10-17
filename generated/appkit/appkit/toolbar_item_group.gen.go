// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ToolbarItemGroup] class.
var ToolbarItemGroupClass objc.Class

func init() {
	ToolbarItemGroupClass = objc.GetClass("NSToolbarItemGroup")
}

type ToolbarItemGroup struct {
	objc.ID
}

func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ID: objc.ID(ptr),
	}
}




