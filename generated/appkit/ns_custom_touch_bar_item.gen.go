// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomTouchBarItem] class.
var (
	CustomTouchBarItemClass     _CustomTouchBarItemClass
	CustomTouchBarItemClassOnce sync.Once
)

func getCustomTouchBarItemClass() _CustomTouchBarItemClass {
	CustomTouchBarItemClassOnce.Do(func() {
		CustomTouchBarItemClass = _CustomTouchBarItemClass{objc.GetClass("NSCustomTouchBarItem")}
	})
	return CustomTouchBarItemClass
}

type _CustomTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [CustomTouchBarItem] class.
type ICustomTouchBarItem interface {
	ITouchBarItem
}

// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem
type CustomTouchBarItem struct {
	TouchBarItem
}

// CustomTouchBarItemFrom constructs a [CustomTouchBarItem] from an unsafe.Pointer.
//
// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
func CustomTouchBarItemFrom(ptr unsafe.Pointer) CustomTouchBarItem {
	return CustomTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomTouchBarItemClass) Alloc() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomTouchBarItemClass) New() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomTouchBarItem) Init() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomTouchBarItem) Autorelease() CustomTouchBarItem {
	rv := objc.Send[CustomTouchBarItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomTouchBarItem creates a new CustomTouchBarItem instance.
func NewCustomTouchBarItem() CustomTouchBarItem {
	return getCustomTouchBarItemClass().New()
}

// A view controller whose view is displayed in the bar to represent this item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) ViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("viewController"))
	return rv
}

// SetViewController sets the value of the viewController property.
// A view controller whose view is displayed in the bar to represent this item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) SetViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setViewController:"), value)
}
