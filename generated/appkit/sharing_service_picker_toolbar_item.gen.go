// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SharingServicePickerToolbarItem] class.
var sharingServicePickerToolbarItemClass = _SharingServicePickerToolbarItemClass{objc.GetClass("NSSharingServicePickerToolbarItem")}

type _SharingServicePickerToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePickerToolbarItem] class.
type ISharingServicePickerToolbarItem interface {
	IToolbarItem
}

// A toolbar item that displays the macOS share sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem

type SharingServicePickerToolbarItem struct {
	ToolbarItem
}

// SharingServicePickerToolbarItemFrom constructs a [SharingServicePickerToolbarItem] from an unsafe.Pointer.
//
// A toolbar item that displays the macOS share sheet.
func SharingServicePickerToolbarItemFrom(ptr unsafe.Pointer) SharingServicePickerToolbarItem {
	return SharingServicePickerToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}



