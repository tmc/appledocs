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
	ViewController() NSViewController
	SetViewController(value IViewController)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	View() NSView
	SetView(value IView)
}

// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.


// A bar item that contains a responder of your choice, such as a view, a button, or a scrubber.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) ViewController() NSViewController {
	rv := objc.Send[NSViewController](c_.ID, objc.Sel("viewController"))
	return rv
}


// A view controller whose view is displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomTouchBarItem/viewController
func (c_ CustomTouchBarItem) SetViewController(value IViewController) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setViewController:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/customizationlabel
func (c_ CustomTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/customizationlabel
func (c_ CustomTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}


// The view displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/view
func (c_ CustomTouchBarItem) View() NSView {
	rv := objc.Send[NSView](c_.ID, objc.Sel("view"))
	return rv
}


// The view displayed in the bar to represent this item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomtouchbaritem/view
func (c_ CustomTouchBarItem) SetView(value IView) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setView:"), value)
}



