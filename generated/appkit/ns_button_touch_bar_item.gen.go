// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var (
	ButtonTouchBarItemClass     _ButtonTouchBarItemClass
	ButtonTouchBarItemClassOnce sync.Once
)

func getButtonTouchBarItemClass() _ButtonTouchBarItemClass {
	ButtonTouchBarItemClassOnce.Do(func() {
		ButtonTouchBarItemClass = _ButtonTouchBarItemClass{objc.GetClass("NSButtonTouchBarItem")}
	})
	return ButtonTouchBarItemClass
}

type _ButtonTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ITouchBarItem
}

// A bar item that provides a button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem
type ButtonTouchBarItem struct {
	TouchBarItem
}

// ButtonTouchBarItemFrom constructs a [ButtonTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a button.
func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonTouchBarItemClass) Alloc() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonTouchBarItemClass) New() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ButtonTouchBarItem) Init() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ButtonTouchBarItem) Autorelease() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButtonTouchBarItem creates a new ButtonTouchBarItem instance.
func NewButtonTouchBarItem() ButtonTouchBarItem {
	return getButtonTouchBarItemClass().New()
}




