// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolPickerScribbleItem] class.
var (
	ToolPickerScribbleItemClass     _ToolPickerScribbleItemClass
	ToolPickerScribbleItemClassOnce sync.Once
)

func getToolPickerScribbleItemClass() _ToolPickerScribbleItemClass {
	ToolPickerScribbleItemClassOnce.Do(func() {
		ToolPickerScribbleItemClass = _ToolPickerScribbleItemClass{objc.GetClass("PKToolPickerScribbleItem")}
	})
	return ToolPickerScribbleItemClass
}

type _ToolPickerScribbleItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerScribbleItem] class.
type IToolPickerScribbleItem interface {
	IToolPickerItem
	// properties:
	// methods:
}

// An item that represents a Scribble tool in the tool picker.


// An item that represents a Scribble tool in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerScribbleItem
type ToolPickerScribbleItem struct {
	ToolPickerItem
}

// ToolPickerScribbleItemFrom constructs a [ToolPickerScribbleItem] from an unsafe.Pointer.
//
// An item that represents a Scribble tool in the tool picker.
func ToolPickerScribbleItemFrom(ptr unsafe.Pointer) ToolPickerScribbleItem {
	return ToolPickerScribbleItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerScribbleItemClass) Alloc() ToolPickerScribbleItem {
	rv := objc.Send[ToolPickerScribbleItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerScribbleItemClass) New() ToolPickerScribbleItem {
	rv := objc.Send[ToolPickerScribbleItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerScribbleItem) Init() ToolPickerScribbleItem {
	rv := objc.Send[ToolPickerScribbleItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerScribbleItem) Autorelease() ToolPickerScribbleItem {
	rv := objc.Send[ToolPickerScribbleItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerScribbleItem creates a new ToolPickerScribbleItem instance.
func NewToolPickerScribbleItem() ToolPickerScribbleItem {
	return getToolPickerScribbleItemClass().New()
}




