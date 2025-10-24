// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ToolPickerInkingItem] class.
var (
	ToolPickerInkingItemClass     _ToolPickerInkingItemClass
	ToolPickerInkingItemClassOnce sync.Once
)

func getToolPickerInkingItemClass() _ToolPickerInkingItemClass {
	ToolPickerInkingItemClassOnce.Do(func() {
		ToolPickerInkingItemClass = _ToolPickerInkingItemClass{objc.GetClass("PKToolPickerInkingItem")}
	})
	return ToolPickerInkingItemClass
}

type _ToolPickerInkingItemClass struct {
	class objc.Class
}

// An interface definition for the [ToolPickerInkingItem] class.
type IToolPickerInkingItem interface {
	IToolPickerItem
	// properties:
	AllowsColorSelection() bool
	SetAllowsColorSelection(value bool)
	InkingTool() IPKInkingTool
	// methods:
}

// An item that represents an inking tool in the tool picker.
//
// An inking item represents a — a tool for drawing marks in a canvas view — in a .


// An item that represents an inking tool in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem
type ToolPickerInkingItem struct {
	ToolPickerItem
}

// ToolPickerInkingItemFrom constructs a [ToolPickerInkingItem] from an unsafe.Pointer.
//
// An item that represents an inking tool in the tool picker.
func ToolPickerInkingItemFrom(ptr unsafe.Pointer) ToolPickerInkingItem {
	return ToolPickerInkingItem{
		ToolPickerItem: ToolPickerItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerInkingItemClass) Alloc() ToolPickerInkingItem {
	rv := objc.Send[ToolPickerInkingItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerInkingItemClass) New() ToolPickerInkingItem {
	rv := objc.Send[ToolPickerInkingItem](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPickerInkingItem) Init() ToolPickerInkingItem {
	rv := objc.Send[ToolPickerInkingItem](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPickerInkingItem) Autorelease() ToolPickerInkingItem {
	rv := objc.Send[ToolPickerInkingItem](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPickerInkingItem creates a new ToolPickerInkingItem instance.
func NewToolPickerInkingItem() ToolPickerInkingItem {
	return getToolPickerInkingItemClass().New()
}



// Create a new tool picker item with a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:
func NewToolPickerInkingItemWithInkType(inkType InkType /* not a class type */) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:"), inkType)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:color:
func NewToolPickerInkingItemWithInkTypeColor(inkType InkType /* not a class type */, color objc.IObject /* cross-framework: Color */) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:color:"), inkType, color)
	rv.Autorelease()
	return rv
}


// Creates a new inking item with the specified ink type, color, and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:color:width:
func NewToolPickerInkingItemWithInkTypeColorWidth(inkType InkType /* not a class type */, color objc.IObject /* cross-framework: Color */, width float64) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:color:width:"), inkType, color, width)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:color:width:azimuth:identifier:
func NewToolPickerInkingItemWithInkTypeColorWidthAzimuthIdentifier(inkType InkType /* not a class type */, color objc.IObject /* cross-framework: Color */, width float64, azimuth float64, identifier objc.IObject /* cross-framework: NSString */) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:color:width:azimuth:identifier:"), inkType, color, width, azimuth, identifier)
	rv.Autorelease()
	return rv
}


// Creates a new inking item with the specified ink type, color, width, and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:color:width:identifier:
func NewToolPickerInkingItemWithInkTypeColorWidthIdentifier(inkType InkType /* not a class type */, color objc.IObject /* cross-framework: Color */, width float64, identifier objc.IObject /* cross-framework: NSString */) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:color:width:identifier:"), inkType, color, width, identifier)
	rv.Autorelease()
	return rv
}


// Create a new tool picker item with a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/initWithInkType:width:
func NewToolPickerInkingItemWithInkTypeWidth(inkType InkType /* not a class type */, width float64) ToolPickerInkingItem {
	instance := getToolPickerInkingItemClass().Alloc()
	rv := objc.Send[ToolPickerInkingItem](instance.ID, objc.Sel("initWithInkType:width:"), inkType, width)
	rv.Autorelease()
	return rv
}



// Present color selection UI to the user. Default value is YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/allowsColorSelection
func (t_ ToolPickerInkingItem) AllowsColorSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("allowsColorSelection"))
	return rv
}


// Present color selection UI to the user. Default value is YES.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/allowsColorSelection
func (t_ ToolPickerInkingItem) SetAllowsColorSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowsColorSelection:"), value)
}


// A tool for drawing on a canvas view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerInkingItem/inkingTool-625y9
func (t_ ToolPickerInkingItem) InkingTool() IPKInkingTool {
	rv := objc.Send[InkingTool](t_.ID, objc.Sel("inkingTool"))
	return rv
}


