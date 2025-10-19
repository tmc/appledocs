// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuItemCell] class.
var (
	menuItemCellClass     _MenuItemCellClass
	menuItemCellClassOnce sync.Once
)

func getMenuItemCellClass() _MenuItemCellClass {
	menuItemCellClassOnce.Do(func() {
		menuItemCellClass = _MenuItemCellClass{objc.GetClass("NSMenuItemCell")}
	})
	return menuItemCellClass
}

type _MenuItemCellClass struct {
	class objc.Class
}

// An interface definition for the [MenuItemCell] class.
type IMenuItemCell interface {
	IButtonCell
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

// Alloc allocates a new instance without initialization.
func (mc _MenuItemCellClass) Alloc() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemCellClass) New() MenuItemCell {
	rv := objc.Send[MenuItemCell](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItemCell) Init() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItemCell) Autorelease() MenuItemCell {
	rv := objc.Send[MenuItemCell](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItemCell creates a new MenuItemCell instance.
func NewMenuItemCell() MenuItemCell {
	return getMenuItemCellClass().New()
}




