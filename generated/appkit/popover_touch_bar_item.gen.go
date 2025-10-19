// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PopoverTouchBarItem] class.
var popoverTouchBarItemClass = _PopoverTouchBarItemClass{objc.GetClass("NSPopoverTouchBarItem")}

type _PopoverTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [PopoverTouchBarItem] class.
type IPopoverTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopoverTouchBarItem

type PopoverTouchBarItem struct {
	TouchBarItem
}

// PopoverTouchBarItemFrom constructs a [PopoverTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a two-state control that can expand into its second state, showing the contents of a bar that it owns.
func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PopoverTouchBarItemClass) Alloc() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PopoverTouchBarItemClass) New() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PopoverTouchBarItem) Init() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PopoverTouchBarItem) Autorelease() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPopoverTouchBarItem creates a new PopoverTouchBarItem instance.
func NewPopoverTouchBarItem() PopoverTouchBarItem {
	return popoverTouchBarItemClass.New()
}




