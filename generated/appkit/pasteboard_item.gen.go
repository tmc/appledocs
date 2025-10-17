// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasteboardItem] class.
var pasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	class objc.Class
}

// An item on a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem

type PasteboardItem struct {
	objectivec.Object
}

// PasteboardItemFrom constructs a [PasteboardItem] from an unsafe.Pointer.
//
// An item on a pasteboard.
func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{objectivec.Object{objc.ID(ptr)}}
}



