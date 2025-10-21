// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolPickerLassoItem] class.
var (
	ToolPickerLassoItemClass     _ToolPickerLassoItemClass
	ToolPickerLassoItemClassOnce sync.Once
)

func getToolPickerLassoItemClass() _ToolPickerLassoItemClass {
	ToolPickerLassoItemClassOnce.Do(func() {
		ToolPickerLassoItemClass = _ToolPickerLassoItemClass{objc.GetClass("PKToolPickerLassoItem")}
	})
	return ToolPickerLassoItemClass
}

type _ToolPickerLassoItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerLassoItem] class.
type IToolPickerLassoItem interface {
	IToolPickerItem
}

// An item that represents a lasso tool in the tool picker.
//
// A lasso item represents a — a tool for selecting stroked lines and shapes in a canvas view — in a .
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerLassoItem
type ToolPickerLassoItem struct {
	ToolPickerItem
}

// ToolPickerLassoItemFrom constructs a [ToolPickerLassoItem] from an unsafe.Pointer.
//
// An item that represents a lasso tool in the tool picker.
func ToolPickerLassoItemFrom(ptr unsafe.Pointer) ToolPickerLassoItem {
	return ToolPickerLassoItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerLassoItemClass) Alloc() ToolPickerLassoItem {
	rv := objc.Send[ToolPickerLassoItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerLassoItemClass) New() ToolPickerLassoItem {
	rv := objc.Send[ToolPickerLassoItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerLassoItem) Init() ToolPickerLassoItem {
	rv := objc.Send[ToolPickerLassoItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerLassoItem) Autorelease() ToolPickerLassoItem {
	rv := objc.Send[ToolPickerLassoItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerLassoItem creates a new ToolPickerLassoItem instance.
func NewToolPickerLassoItem() ToolPickerLassoItem {
	return getToolPickerLassoItemClass().New()
}



// A lasso tool for selecting parts of a drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerLassoItem/lassoTool-1urgb
func (t_ ToolPickerLassoItem) LassoTool() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("lassoTool"))
	return rv
}


