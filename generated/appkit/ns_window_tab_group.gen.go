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
	Windows() []Window
	TabGroup() IWindowTabGroup
	SetTabGroup(value IWindowTabGroup)
	Identifier() objectivec.IObject
	SetIdentifier(value objectivec.IObject)
	IsOverviewVisible() bool
	SetIsOverviewVisible(value bool)
	IsTabBarVisible() bool
	SetIsTabBarVisible(value bool)
	SelectedWindow() IWindow
	SetSelectedWindow(value IWindow)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// A collection of the windows that are currently grouped together by this window tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowTabGroup/windows
func (w_ WindowTabGroup) Windows() []Window {
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


// The unique identifier for a tabbed window group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/identifier
func (w_ WindowTabGroup) Identifier() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier for a tabbed window group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/identifier
func (w_ WindowTabGroup) SetIdentifier(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIdentifier:"), value)
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/isoverviewvisible
func (w_ WindowTabGroup) IsOverviewVisible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOverviewVisible"))
	return rv
}


// A Boolean value indicating if the tab overview is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/isoverviewvisible
func (w_ WindowTabGroup) SetIsOverviewVisible(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOverviewVisible:"), value)
}


// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/istabbarvisible
func (w_ WindowTabGroup) IsTabBarVisible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isTabBarVisible"))
	return rv
}


// A Boolean value indicating whether the tabbed window group currently displays a tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/istabbarvisible
func (w_ WindowTabGroup) SetIsTabBarVisible(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsTabBarVisible:"), value)
}


// The selected, or frontmost, window in the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/selectedwindow
func (w_ WindowTabGroup) SelectedWindow() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("selectedWindow"))
	return rv
}


// The selected, or frontmost, window in the tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowtabgroup/selectedwindow
func (w_ WindowTabGroup) SetSelectedWindow(value IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSelectedWindow:"), value)
}








