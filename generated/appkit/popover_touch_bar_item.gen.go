
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PopoverTouchBarItem] class.
var PopoverTouchBarItemClass _PopoverTouchBarItemClass

func init() {
	PopoverTouchBarItemClass = _PopoverTouchBarItemClass{objc.GetClass("NSPopoverTouchBarItem")}
}

type _PopoverTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [PopoverTouchBarItem] class.
type IPopoverTouchBarItem interface {
	ID() objc.ID
}

type PopoverTouchBarItem struct {
	id objc.ID
}

func PopoverTouchBarItemFrom(ptr unsafe.Pointer) PopoverTouchBarItem {
	return PopoverTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PopoverTouchBarItem) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PopoverTouchBarItemClass) Alloc() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PopoverTouchBarItemClass) New() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPopoverTouchBarItem creates and returns a new initialized instance.
func NewPopoverTouchBarItem() PopoverTouchBarItem {
	return PopoverTouchBarItemClass.New()
}

// Init initializes the instance.
func (p_ PopoverTouchBarItem) Init() PopoverTouchBarItem {
	rv := objc.Send[PopoverTouchBarItem](p_.ID(), selInit)
	return rv
}
// The bar displayed when this item is “popped.” [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) PopoverTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("popoverTouchBar"))
	return rv
}
// SetPopoverTouchBar sets the value of the popoverTouchBar property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/popoverTouchBar
func (p_ PopoverTouchBarItem) SetPopoverTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setPopoverTouchBar:"), value)
}
// The bar that is displayed when a user press-and-holds on the popover item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) PressAndHoldTouchBar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("pressAndHoldTouchBar"))
	return rv
}
// SetPressAndHoldTouchBar sets the value of the pressAndHoldTouchBar property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPopoverTouchBarItem/pressAndHoldTouchBar
func (p_ PopoverTouchBarItem) SetPressAndHoldTouchBar(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setPressAndHoldTouchBar:"), value)
}
