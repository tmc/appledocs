
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WindowTabGroup] class.
var WindowTabGroupClass _WindowTabGroupClass

func init() {
	WindowTabGroupClass = _WindowTabGroupClass{objc.GetClass("NSWindowTabGroup")}
}

type _WindowTabGroupClass struct {
	objc.Class
}

// An interface definition for the [WindowTabGroup] class.
type IWindowTabGroup interface {
	ID() objc.ID
	AddWindow(window unsafe.Pointer)
	InsertWindowAtIndex(window unsafe.Pointer, index int)
	RemoveWindow(window unsafe.Pointer)
}

type WindowTabGroup struct {
	id objc.ID
}

func WindowTabGroupFrom(ptr unsafe.Pointer) WindowTabGroup {
	return WindowTabGroup{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ WindowTabGroup) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WindowTabGroupClass) Alloc() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WindowTabGroupClass) New() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWindowTabGroup creates and returns a new initialized instance.
func NewWindowTabGroup() WindowTabGroup {
	return WindowTabGroupClass.New()
}

// Init initializes the instance.
func (w_ WindowTabGroup) Init() WindowTabGroup {
	rv := objc.Send[WindowTabGroup](w_.ID(), selInit)
	return rv
}
// Adds a window to the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/addWindow(_:)
func (w_ WindowTabGroup) AddWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("addWindow:"), window)
}
// Inserts a window at a specific location within the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/insertWindow(_:at:)
func (w_ WindowTabGroup) InsertWindowAtIndex(window unsafe.Pointer, index int) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("insertWindow:atIndex:"), window, index)
}
// Removes a window from the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/removeWindow(_:)
func (w_ WindowTabGroup) RemoveWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("removeWindow:"), window)
}
// The unique identifier for a tabbed window group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/identifier
func (w_ WindowTabGroup) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("identifier"))
	return rv
}
// A Boolean value indicating if the tab overview is currently displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) OverviewVisible() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("overviewVisible"))
	return rv
}
// SetOverviewVisible sets the value of the overviewVisible property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/isOverviewVisible
func (w_ WindowTabGroup) SetOverviewVisible(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setOverviewVisible:"), value)
}
// A Boolean value indicating whether the tabbed window group currently displays a tab bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/isTabBarVisible
func (w_ WindowTabGroup) TabBarVisible() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("tabBarVisible"))
	return rv
}
// The selected, or frontmost, window in the tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SelectedWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("selectedWindow"))
	return rv
}
// SetSelectedWindow sets the value of the selectedWindow property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/selectedWindow
func (w_ WindowTabGroup) SetSelectedWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setSelectedWindow:"), value)
}
// A collection of the windows that are currently grouped together by this window tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindowTabGroup/windows
func (w_ WindowTabGroup) Windows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("windows"))
	return rv
}
