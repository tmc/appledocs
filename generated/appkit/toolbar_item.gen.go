
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ToolbarItem] class.
var ToolbarItemClass _ToolbarItemClass

func init() {
	ToolbarItemClass = _ToolbarItemClass{objc.GetClass("NSToolbarItem")}
}

type _ToolbarItemClass struct {
	objc.Class
}

// An interface definition for the [ToolbarItem] class.
type IToolbarItem interface {
	ID() objc.ID
}

type ToolbarItem struct {
	id objc.ID
}

func ToolbarItemFrom(ptr unsafe.Pointer) ToolbarItem {
	return ToolbarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ ToolbarItem) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemClass) Alloc() ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _ToolbarItemClass) New() ToolbarItem {
	rv := objc.Send[ToolbarItem](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewToolbarItem creates and returns a new initialized instance.
func NewToolbarItem() ToolbarItem {
	return ToolbarItemClass.New()
}

// Init initializes the instance.
func (t_ ToolbarItem) Init() ToolbarItem {
	rv := objc.Send[ToolbarItem](t_.ID(), selInit)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) BackgroundTintColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("backgroundTintColor"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t_ ToolbarItem) SetBackgroundTintColor(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setBackgroundTintColor:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) Hidden() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("hidden"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/isHidden
func (t_ ToolbarItem) SetHidden(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setHidden:"), value)
}
// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) Navigational() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("navigational"))
	return rv
}
// SetNavigational sets the value of the navigational property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSToolbarItem/isNavigational
func (t_ ToolbarItem) SetNavigational(value bool) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setNavigational:"), value)
}
