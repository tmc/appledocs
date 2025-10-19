// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItemBadge] class.
var (
	menuItemBadgeClass     _MenuItemBadgeClass
	menuItemBadgeClassOnce sync.Once
)

func getMenuItemBadgeClass() _MenuItemBadgeClass {
	menuItemBadgeClassOnce.Do(func() {
		menuItemBadgeClass = _MenuItemBadgeClass{objc.GetClass("NSMenuItemBadge")}
	})
	return menuItemBadgeClass
}

type _MenuItemBadgeClass struct {
	class objc.Class
}

// An interface definition for the [MenuItemBadge] class.
type IMenuItemBadge interface {
	objectivec.IObject
}

// A control that provides additional quantitative information specific to a menu item, such as the number of available updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge
type MenuItemBadge struct {
	objectivec.Object
}

// MenuItemBadgeFrom constructs a [MenuItemBadge] from an unsafe.Pointer.
//
// A control that provides additional quantitative information specific to a menu item, such as the number of available updates.
func MenuItemBadgeFrom(ptr unsafe.Pointer) MenuItemBadge {
	return MenuItemBadge{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemBadgeClass) Alloc() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemBadgeClass) New() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItemBadge) Init() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItemBadge) Autorelease() MenuItemBadge {
	rv := objc.Send[MenuItemBadge](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItemBadge creates a new MenuItemBadge instance.
func NewMenuItemBadge() MenuItemBadge {
	return getMenuItemBadgeClass().New()
}




