// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PasteboardItem] class.
var PasteboardItemClass objc.Class

func init() {
	PasteboardItemClass = objc.GetClass("NSPasteboardItem")
}

type PasteboardItem struct {
	objc.ID
}

func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{
		ID: objc.ID(ptr),
	}
}



