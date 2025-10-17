// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Menu] class.
var MenuClass objc.Class

func init() {
	MenuClass = objc.GetClass("NSMenu")
}

type Menu struct {
	objc.ID
}

func MenuFrom(ptr unsafe.Pointer) Menu {
	return Menu{
		ID: objc.ID(ptr),
	}
}


// Displays a contextual menu over a view for an event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (mc Menu) PopUpContextMenuWithEventForView(menu unsafe.Pointer, event unsafe.Pointer, view unsafe.Pointer) {
	sel := objc.RegisterName("popUpContextMenu:withEvent:forView:")
	objc.ID(MenuClass).Send(sel, menu, event, view)
}

