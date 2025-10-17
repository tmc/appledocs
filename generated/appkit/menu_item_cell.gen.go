// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuItemCell] class.
var menuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}

type _MenuItemCellClass struct {
	class objc.Class
}

// An object that handles the measurement and display of a single menu item in its encompassing frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemCell

type MenuItemCell struct {
	ButtonCell
}

// MenuItemCellFrom constructs a [MenuItemCell] from an unsafe.Pointer.
//
// An object that handles the measurement and display of a single menu item in its encompassing frame.
func MenuItemCellFrom(ptr unsafe.Pointer) MenuItemCell {
	return MenuItemCell{
		ButtonCell: ButtonCellFrom(ptr),
	}
}



