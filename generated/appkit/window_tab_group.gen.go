// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowTabGroup] class.
var windowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}

type _WindowTabGroupClass struct {
	class objc.Class
}

// An interface definition for the [WindowTabGroup] class.
type IWindowTabGroup interface {
	objectivec.IObject
	AddWindow(window unsafe.Pointer)
	InsertWindowAtIndex(window unsafe.Pointer, index int)
	RemoveWindow(window unsafe.Pointer)
}

// A group of windows that display together as a single tabbed window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup

type WindowTabGroup struct {
	objectivec.Object
}

// WindowTabGroupFrom constructs a [WindowTabGroup] from an unsafe.Pointer.
//
// A group of windows that display together as a single tabbed window.
func WindowTabGroupFrom(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{objectivec.Object{objc.ID(ptr)}}
}

// Adds a window to the tab group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w_ WindowTabGroup) AddWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addWindow:"), window)
}
// Inserts a window at a specific location within the tab group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w_ WindowTabGroup) InsertWindowAtIndex(window unsafe.Pointer, index int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("insertWindow:atIndex:"), window, index)
}
// Removes a window from the tab group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w_ WindowTabGroup) RemoveWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeWindow:"), window)
}


