// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MenuToolbarItem] class.
var (
	menuToolbarItemClass     _MenuToolbarItemClass
	menuToolbarItemClassOnce sync.Once
)

func getMenuToolbarItemClass() _MenuToolbarItemClass {
	menuToolbarItemClassOnce.Do(func() {
		menuToolbarItemClass = _MenuToolbarItemClass{objc.GetClass("NSMenuToolbarItem")}
	})
	return menuToolbarItemClass
}

type _MenuToolbarItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuToolbarItem] class.
type IMenuToolbarItem interface {
	IToolbarItem
}

// A control that presents a menu in a window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuToolbarItem

type MenuToolbarItem struct {
	ToolbarItem
}

// MenuToolbarItemFrom constructs a [MenuToolbarItem] from an unsafe.Pointer.
//
// A control that presents a menu in a window’s toolbar.
func MenuToolbarItemFrom(ptr unsafe.Pointer) MenuToolbarItem {
	return MenuToolbarItem{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MenuToolbarItemClass) Alloc() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuToolbarItemClass) New() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuToolbarItem) Init() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuToolbarItem) Autorelease() MenuToolbarItem {
	rv := objc.Send[MenuToolbarItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuToolbarItem creates a new MenuToolbarItem instance.
func NewMenuToolbarItem() MenuToolbarItem {
	return getMenuToolbarItemClass().New()
}




