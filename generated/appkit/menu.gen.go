// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Menu] class.
var menuClass = _MenuClass{objc.GetClass("NSMenu")}

type _MenuClass struct {
	class objc.Class
}

// An object that manages an app’s menus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu

type Menu struct {
	objectivec.Object
}

// MenuFrom constructs a [Menu] from an unsafe.Pointer.
//
// An object that manages an app’s menus.
func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{objectivec.Object{objc.ID(ptr)}}
}

// Displays a contextual menu over a view for an event. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc _MenuClass) PopUpContextMenuWithEventForView(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}


