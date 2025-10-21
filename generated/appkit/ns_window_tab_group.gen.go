// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WindowTabGroup] class.
var (
	WindowTabGroupClass     _WindowTabGroupClass
	WindowTabGroupClassOnce sync.Once
)

func getWindowTabGroupClass() _WindowTabGroupClass {
	WindowTabGroupClassOnce.Do(func() {
		WindowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}
	})
	return WindowTabGroupClass
}

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

// A group of windows that display together as a single tabbed window.
//
// AppKit automatically creates instances of to reflect the tabbing state of your windows. You can access a window’s current tab group using the property.
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

// Alloc allocates a new instance without initialization.
func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WindowTabGroup) Autorelease() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindowTabGroup creates a new WindowTabGroup instance.
func NewWindowTabGroup() WindowTabGroup {
	return getWindowTabGroupClass().New()
}


// Adds a window to the tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w_ WindowTabGroup) AddWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addWindow:"), window)
}

// Inserts a window at a specific location within the tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w_ WindowTabGroup) InsertWindowAtIndex(window unsafe.Pointer, index int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("insertWindow:atIndex:"), window, index)
}

// Removes a window from the tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w_ WindowTabGroup) RemoveWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeWindow:"), window)
}

// The unique identifier for a tabbed window group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/identifier
func (w_ WindowTabGroup) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("identifier"))
	return rv
}

// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) OverviewVisible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("overviewVisible"))
	return rv
}


// SetOverviewVisible sets the value of the overviewVisible property.
// A Boolean value indicating if the tab overview is currently displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOverviewVisible:"), value)
}

// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isTabBarVisible
func (w_ WindowTabGroup) TabBarVisible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("tabBarVisible"))
	return rv
}

// The selected, or frontmost, window in the tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SelectedWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("selectedWindow"))
	return rv
}


// SetSelectedWindow sets the value of the selectedWindow property.
// The selected, or frontmost, window in the tab group.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SetSelectedWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedWindow:"), value)
}

// A collection of the windows that are currently grouped together by this window tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/windows
func (w_ WindowTabGroup) Windows() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("windows"))
	return rv
}



