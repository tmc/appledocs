// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DraggingItem] class.
var draggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}

type _DraggingItemClass struct {
	class objc.Class
}

// A single dragged item within a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItem

type DraggingItem struct {
	objectivec.Object
}

// DraggingItemFrom constructs a [DraggingItem] from an unsafe.Pointer.
//
// A single dragged item within a dragging session.
func DraggingItemFrom(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{objectivec.Object{objc.ID(ptr)}}
}



