// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolbarItemGroup] class.
var (
	toolbarItemGroupClass     _ToolbarItemGroupClass
	toolbarItemGroupClassOnce sync.Once
)

func getToolbarItemGroupClass() _ToolbarItemGroupClass {
	toolbarItemGroupClassOnce.Do(func() {
		toolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}
	})
	return toolbarItemGroupClass
}

type _ToolbarItemGroupClass struct {
	class objc.Class
}

// An interface definition for the [ToolbarItemGroup] class.
type IToolbarItemGroup interface {
	IToolbarItem
}

// A group of subitems in a toolbar item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup
type ToolbarItemGroup struct {
	ToolbarItem
}

// ToolbarItemGroupFrom constructs a [ToolbarItemGroup] from an unsafe.Pointer.
//
// A group of subitems in a toolbar item.
func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolbarItemGroup) Autorelease() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolbarItemGroup creates a new ToolbarItemGroup instance.
func NewToolbarItemGroup() ToolbarItemGroup {
	return getToolbarItemGroupClass().New()
}




