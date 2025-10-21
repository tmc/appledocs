// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ToolPickerItem] class.
var (
	ToolPickerItemClass     _ToolPickerItemClass
	ToolPickerItemClassOnce sync.Once
)

func getToolPickerItemClass() _ToolPickerItemClass {
	ToolPickerItemClassOnce.Do(func() {
		ToolPickerItemClass = _ToolPickerItemClass{objc.GetClass("PKToolPickerItem")}
	})
	return ToolPickerItemClass
}

type _ToolPickerItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerItem] class.
type IToolPickerItem interface {
	objectivec.IObject
}

// The base class for an item in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerItem
type ToolPickerItem struct {
	objectivec.Object
}

// ToolPickerItemFrom constructs a [ToolPickerItem] from an unsafe.Pointer.
//
// The base class for an item in the tool picker.
func ToolPickerItemFrom(ptr unsafe.Pointer) ToolPickerItem {
	return ToolPickerItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerItemClass) Alloc() ToolPickerItem {
	rv := objc.Send[ToolPickerItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerItemClass) New() ToolPickerItem {
	rv := objc.Send[ToolPickerItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerItem) Init() ToolPickerItem {
	rv := objc.Send[ToolPickerItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerItem) Autorelease() ToolPickerItem {
	rv := objc.Send[ToolPickerItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerItem creates a new ToolPickerItem instance.
func NewToolPickerItem() ToolPickerItem {
	return getToolPickerItemClass().New()
}


// A string that identifies the item in the tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerItem/identifier
func (t_ ToolPickerItem) Identifier() string {
	rv := objc.Send[string](t_.ID, objc.Sel("identifier"))
	return rv
}

// The this tool picker item represents.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerItem/tool-918ln
func (t_ ToolPickerItem) Tool() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("tool"))
	return rv
}



