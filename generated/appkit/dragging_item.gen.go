// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DraggingItem] class.
var DraggingItemClass objc.Class

func init() {
	DraggingItemClass = objc.GetClass("NSDraggingItem")
}

type DraggingItem struct {
	objc.ID
}

func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		ID: objc.ID(ptr),
	}
}



