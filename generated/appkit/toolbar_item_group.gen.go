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

// A group of subitems in a toolbar item.
//
// An represents a collection set of subitems in a toolbar that the system displays based on available space and settings that you specify. The system uses the views and labels of the subitems, but the parent’s attributes take precedence. This differs from other objects because they’re attached — the user drags them together as a single item rather than separately. If a subitem of the group has an action set on it, the group uses that action instead of its own when the user clicks or taps on that item. The system prefers the subitem’s action if it exists, otherwise it uses the group’s action. To configure an instance of , you first create the individual toolbar subitems: Then, you put them in a grouped item: In this configuration, you get two grouped items, and two labels. If you set a label on the parent item, you get two grouped items with one shared label: If instead you set a view on the parent item, you get two labels with one shared view:
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




