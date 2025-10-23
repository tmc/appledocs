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
	// properties:
	Identifier() objc.IObject /* cross-framework: WindowTabbingIdentifier */
	OverviewVisible() bool /* primitive/slice/pointer. */
	SetOverviewVisible(value bool /* primitive/slice/pointer. */)
	TabBarVisible() bool /* primitive/slice/pointer. */
	SelectedWindow() IWindow
	SetSelectedWindow(value IWindow)
	Windows() []Window /* primitive/slice/pointer. */
	TabGroup() IWindowTabGroup
	SetTabGroup(value IWindowTabGroup)
	IsOverviewVisible() bool /* primitive/slice/pointer. */
	SetIsOverviewVisible(value bool /* primitive/slice/pointer. */)
	IsTabBarVisible() bool /* primitive/slice/pointer. */
	SetIsTabBarVisible(value bool /* primitive/slice/pointer. */)
	// methods:
	AddWindow(window IWindow)
	InsertWindowAtIndex(window IWindow, index int /* primitive/slice/pointer. */)
	RemoveWindow(window IWindow)
}

// A group of windows that display together as a single tabbed window.
//
// AppKit automatically creates instances of to reflect the tabbing state of your windows. You can access a window’s current tab group using the property.


// A group of windows that display together as a single tabbed window.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w_ WindowTabGroup) AddWindow(window IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addWindow:"), window)
}


// Inserts a window at a specific location within the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w_ WindowTabGroup) InsertWindowAtIndex(window IWindow, index int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("insertWindow:atIndex:"), window, index)
}


// Removes a window from the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w_ WindowTabGroup) RemoveWindow(window IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeWindow:"), window)
}


// The unique identifier for a tabbed window group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/identifier
func (w_ WindowTabGroup) Identifier() objc.IObject /* cross-framework: WindowTabbingIdentifier */ {
	rv := objc.Send[WindowTabbingIdentifier](w_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) OverviewVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("overviewVisible"))
	return rv
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) SetOverviewVisible(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOverviewVisible:"), value)
}


// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/isTabBarVisible
func (w_ WindowTabGroup) TabBarVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("tabBarVisible"))
	return rv
}


// The selected, or frontmost, window in the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SelectedWindow() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("selectedWindow"))
	return rv
}


// The selected, or frontmost, window in the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedWindow:"), value)
}


// A collection of the windows that are currently grouped together by this window tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/windows
func (w_ WindowTabGroup) Windows() []Window /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("windows"))
	return rv
}


// A group of windows that display together as a tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tabgroup
func (w_ WindowTabGroup) TabGroup() IWindowTabGroup {
	rv := objc.Send[WindowTabGroup](w_.ID, objc.Sel("tabGroup"))
	return rv
}


// A group of windows that display together as a tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/tabgroup
func (w_ WindowTabGroup) SetTabGroup(value IWindowTabGroup) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabGroup:"), value)
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/isoverviewvisible
func (w_ WindowTabGroup) IsOverviewVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOverviewVisible"))
	return rv
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/isoverviewvisible
func (w_ WindowTabGroup) SetIsOverviewVisible(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOverviewVisible:"), value)
}


// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/istabbarvisible
func (w_ WindowTabGroup) IsTabBarVisible() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](w_.ID, objc.Sel("isTabBarVisible"))
	return rv
}


// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/istabbarvisible
func (w_ WindowTabGroup) SetIsTabBarVisible(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsTabBarVisible:"), value)
}



