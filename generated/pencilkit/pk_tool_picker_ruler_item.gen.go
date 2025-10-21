// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolPickerRulerItem] class.
var (
	ToolPickerRulerItemClass     _ToolPickerRulerItemClass
	ToolPickerRulerItemClassOnce sync.Once
)

func getToolPickerRulerItemClass() _ToolPickerRulerItemClass {
	ToolPickerRulerItemClassOnce.Do(func() {
		ToolPickerRulerItemClass = _ToolPickerRulerItemClass{objc.GetClass("PKToolPickerRulerItem")}
	})
	return ToolPickerRulerItemClass
}

type _ToolPickerRulerItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerRulerItem] class.
type IToolPickerRulerItem interface {
	IToolPickerItem
}

// An item that represents a ruler tool in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerRulerItem
type ToolPickerRulerItem struct {
	ToolPickerItem
}

// ToolPickerRulerItemFrom constructs a [ToolPickerRulerItem] from an unsafe.Pointer.
//
// An item that represents a ruler tool in the tool picker.
func ToolPickerRulerItemFrom(ptr unsafe.Pointer) ToolPickerRulerItem {
	return ToolPickerRulerItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerRulerItemClass) Alloc() ToolPickerRulerItem {
	rv := objc.Send[ToolPickerRulerItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerRulerItemClass) New() ToolPickerRulerItem {
	rv := objc.Send[ToolPickerRulerItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerRulerItem) Init() ToolPickerRulerItem {
	rv := objc.Send[ToolPickerRulerItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerRulerItem) Autorelease() ToolPickerRulerItem {
	rv := objc.Send[ToolPickerRulerItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerRulerItem creates a new ToolPickerRulerItem instance.
func NewToolPickerRulerItem() ToolPickerRulerItem {
	return getToolPickerRulerItemClass().New()
}




