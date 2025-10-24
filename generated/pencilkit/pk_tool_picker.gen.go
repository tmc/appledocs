// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ToolPicker] class.
var (
	ToolPickerClass     _ToolPickerClass
	ToolPickerClassOnce sync.Once
)

func getToolPickerClass() _ToolPickerClass {
	ToolPickerClassOnce.Do(func() {
		ToolPickerClass = _ToolPickerClass{objc.GetClass("PKToolPicker")}
	})
	return ToolPickerClass
}

type _ToolPickerClass struct {
	class objc.Class
}

// An interface definition for the [ToolPicker] class.
type IToolPicker interface {
	objectivec.IObject
	// properties:
	IsRulerActive() bool
	SetIsRulerActive(value bool)
	// methods:
}

// A tool palette that displays a selection of drawing tools and colors for tools that a person can choose from.
//
// A manages a draggable palette that displays drawing tools, colors, and additional options. You add a tool picker to your interface and configure it to display its palette at appropriate times. While the palette is onscreen, a person may reposition it anywhere within the current window. When a person interacts with the palette, the tool picker notifies registered observers of the changes so that they can respond. When configuring your interface, call the method to associate the tool picker with one or more views in your interface. Each window manages its own tool picker, and the window’s first responder determines the visibility of that tool picker. When one of the registered objects becomes first responder, the tool picker automatically adds its palette view to the current window. When there isn’t a registered object as first responder, the tool picker hides its palette view. implements the observer protocol for detecting tool picker changes. Adding your canvas view as an observer to a tool picker automatically updates the current drawing tools. For more information about implementing custom observer objects, see .


// A tool palette that displays a selection of drawing tools and colors for tools that a person can choose from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker
type ToolPicker struct {
	objectivec.Object
}

// ToolPickerFrom constructs a [ToolPicker] from an unsafe.Pointer.
//
// A tool palette that displays a selection of drawing tools and colors for tools that a person can choose from.
func ToolPickerFrom(ptr unsafe.Pointer) ToolPicker {
	return ToolPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolPickerClass) Alloc() ToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolPickerClass) New() ToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolPicker) Init() ToolPicker {
	rv := objc.Send[ToolPicker](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolPicker) Autorelease() ToolPicker {
	rv := objc.Send[ToolPicker](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolPicker creates a new ToolPicker instance.
func NewToolPicker() ToolPicker {
	return getToolPickerClass().New()
}



// Creates a new tool picker with the tools you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/init(toolItems:)
func NewToolPickerWithToolItems(items []IToolPickerItem) ToolPicker {
	instance := getToolPickerClass().Alloc()
	rv := objc.Send[ToolPicker](instance.ID, objc.Sel("initWithToolItems:"), items)
	rv.Autorelease()
	return rv
}



// Returns the tool picker object to use for the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/shared(for:)
func (tc _ToolPickerClass) SharedToolPickerForWindow(window objc.IObject /* cross-framework: Window */) IToolPicker {
	rv := objc.Send[ToolPicker](objc.ID(tc.class), objc.Sel("sharedToolPickerForWindow:"), window)
	return rv
}


// The default tool items for new tool pickers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPicker/defaultToolItems
func (tc _ToolPickerClass) DefaultToolItems() []IToolPickerItem {
	rv := objc.Send[[]ToolPickerItem](objc.ID(tc.class), objc.Sel("defaultToolItems"))
	return rv
}

// A Boolean value that indicates whether the ruler is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpicker/isruleractive
func (t_ ToolPicker) IsRulerActive() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isRulerActive"))
	return rv
}


// A Boolean value that indicates whether the ruler is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pktoolpicker/isruleractive
func (t_ ToolPicker) SetIsRulerActive(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsRulerActive:"), value)
}


