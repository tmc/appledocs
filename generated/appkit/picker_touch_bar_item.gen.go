// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PickerTouchBarItem] class.
var (
	pickerTouchBarItemClass     _PickerTouchBarItemClass
	pickerTouchBarItemClassOnce sync.Once
)

func getPickerTouchBarItemClass() _PickerTouchBarItemClass {
	pickerTouchBarItemClassOnce.Do(func() {
		pickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}
	})
	return pickerTouchBarItemClass
}

type _PickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [PickerTouchBarItem] class.
type IPickerTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a picker control with multiple options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem

type PickerTouchBarItem struct {
	TouchBarItem
}

// PickerTouchBarItemFrom constructs a [PickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a picker control with multiple options.
func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PickerTouchBarItemClass) Alloc() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PickerTouchBarItemClass) New() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PickerTouchBarItem) Init() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PickerTouchBarItem) Autorelease() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPickerTouchBarItem creates a new PickerTouchBarItem instance.
func NewPickerTouchBarItem() PickerTouchBarItem {
	return getPickerTouchBarItemClass().New()
}




