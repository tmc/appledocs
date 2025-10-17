// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var SharingServicePickerToolbarItemClass objc.Class

func init() {
	SharingServicePickerToolbarItemClass = objc.GetClass("NSSharingServicePickerToolbarItem")
}

type SharingServicePickerToolbarItem struct {
	objc.ID
}

func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		ID: objc.ID(ptr),
	}
}




