//go:build darwin && ios

// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MenuToolbarItem


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem/itemMenu
func (m_ MenuToolbarItem) ItemMenu() IMenu {
	rv := objc.Send[Menu](m_.ID, objc.Sel("itemMenu"))
	return rv
}
func (m_ MenuToolbarItem) SetItemMenu(value IMenu) {
	m_.ID.Send(objc.RegisterName("setItemMenu:"), value)
}





