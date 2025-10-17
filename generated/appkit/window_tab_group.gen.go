// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowTabGroup] class.
var WindowTabGroupClass objc.Class

func init() {
	WindowTabGroupClass = objc.GetClass("NSWindowTabGroup")
}

type WindowTabGroup struct {
	objc.ID
}

func WindowTabGroupFrom(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{
		ID: objc.ID(ptr),
	}
}


// Adds a window to the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w_ WindowTabGroup) AddWindow(window unsafe.Pointer) {
	sel := objc.RegisterName("addWindow:")
	w_.ID.Send(sel, window)
}
// Inserts a window at a specific location within the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w_ WindowTabGroup) InsertWindowAtIndex(window unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertWindow:atIndex:")
	w_.ID.Send(sel, window, index)
}
// Removes a window from the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w_ WindowTabGroup) RemoveWindow(window unsafe.Pointer) {
	sel := objc.RegisterName("removeWindow:")
	w_.ID.Send(sel, window)
}

