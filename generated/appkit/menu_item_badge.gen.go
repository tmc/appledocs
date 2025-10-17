// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MenuItemBadge] class.
var menuItemBadgeClass = _MenuItemBadgeClass{objc.GetClass("NSMenuItemBadge")}

type _MenuItemBadgeClass struct {
	class objc.Class
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



