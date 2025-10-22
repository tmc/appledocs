// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ToolPickerEraserItem] class.
var (
	ToolPickerEraserItemClass     _ToolPickerEraserItemClass
	ToolPickerEraserItemClassOnce sync.Once
)

func getToolPickerEraserItemClass() _ToolPickerEraserItemClass {
	ToolPickerEraserItemClassOnce.Do(func() {
		ToolPickerEraserItemClass = _ToolPickerEraserItemClass{objc.GetClass("PKToolPickerEraserItem")}
	})
	return ToolPickerEraserItemClass
}

type _ToolPickerEraserItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerEraserItem] class.
type IToolPickerEraserItem interface {
	IToolPickerItem
	EraserTool() PKEraserTool
}

// An item that represents an eraser tool in the tool picker.
//
// An eraser item represents a  — a tool for erasing content in a canvas view — in a .
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem
type ToolPickerEraserItem struct {
	ToolPickerItem
}

// ToolPickerEraserItemFrom constructs a [ToolPickerEraserItem] from an unsafe.Pointer.
//
// An item that represents an eraser tool in the tool picker.
func ToolPickerEraserItemFrom(ptr unsafe.Pointer) ToolPickerEraserItem {
	return ToolPickerEraserItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerEraserItemClass) Alloc() ToolPickerEraserItem {
	rv := objc.Send[ToolPickerEraserItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerEraserItemClass) New() ToolPickerEraserItem {
	rv := objc.Send[ToolPickerEraserItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerEraserItem) Init() ToolPickerEraserItem {
	rv := objc.Send[ToolPickerEraserItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerEraserItem) Autorelease() ToolPickerEraserItem {
	rv := objc.Send[ToolPickerEraserItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerEraserItem creates a new ToolPickerEraserItem instance.
func NewToolPickerEraserItem() ToolPickerEraserItem {
	return getToolPickerEraserItemClass().New()
}




// Creates a new eraser item.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/initWithEraserType:
func NewToolPickerEraserItemWithEraserType(eraserType EraserType) ToolPickerEraserItem {
	instance := getToolPickerEraserItemClass().Alloc()
	rv := objc.Send[ToolPickerEraserItem](instance.ID, objc.Sel("initWithEraserType:"), eraserType)
	rv.Autorelease()
	return rv
}



// Creates a new eraser item with the specified width.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/initWithEraserType:width:
func NewToolPickerEraserItemWithEraserTypeWidth(eraserType EraserType, width float64) ToolPickerEraserItem {
	instance := getToolPickerEraserItemClass().Alloc()
	rv := objc.Send[ToolPickerEraserItem](instance.ID, objc.Sel("initWithEraserType:width:"), eraserType, width)
	rv.Autorelease()
	return rv
}


// An eraser tool for erasing parts of a drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/eraserTool-4q3hp
func (t_ ToolPickerEraserItem) EraserTool() PKEraserTool {
	rv := objc.Send[PKEraserTool](t_.ID, objc.Sel("eraserTool"))
	return rv
}


