// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItem] class.
var (
	MenuItemClass     _MenuItemClass
	MenuItemClassOnce sync.Once
)

func getMenuItemClass() _MenuItemClass {
	MenuItemClassOnce.Do(func() {
		MenuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}
	})
	return MenuItemClass
}

type _MenuItemClass struct {
	class objc.Class
}

// An interface definition for the [MenuItem] class.
type IMenuItem interface {
	objectivec.IObject
}

// A command item in an app menu.
//
// The class includes some private functionality needed to maintain binary compatibility with other components of Cocoa. Because of this fact, you can’t replace the class with a different class, but you can subclass it if necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem
type MenuItem struct {
	objectivec.Object
}

// MenuItemFrom constructs a [MenuItem] from an unsafe.Pointer.
//
// A command item in an app menu.
func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MenuItemClass) New() MenuItem {
	rv := objc.Send[MenuItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MenuItem) Init() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MenuItem) Autorelease() MenuItem {
	rv := objc.Send[MenuItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMenuItem creates a new MenuItem instance.
func NewMenuItem() MenuItem {
	return getMenuItemClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) Badge() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("badge"))
	return rv
}


// SetBadge sets the value of the badge property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m_ MenuItem) SetBadge(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBadge:"), value)
}

// A Boolean value that indicates whether the menu item has a submenu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItem/hasSubmenu
func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasSubmenu"))
	return rv
}



