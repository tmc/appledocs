// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKToolPickerEraserItem */


/* debug [class_header]: Header for PKToolPickerEraserItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerEraserItem */
// An interface definition for the [ToolPickerEraserItem] class.
type IToolPickerEraserItem interface {
	IToolPickerItem
	
/* debug [class_interface_properties]: Properties for ToolPickerEraserItem */
	// properties:
	EraserTool() IPKEraserTool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerEraserItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerEraserItem */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerEraserItemClass) Alloc() ToolPickerEraserItem {
	rv := objc.Send[ToolPickerEraserItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerEraserItem */
// An item that represents an eraser tool in the tool picker.
//
// An eraser item represents a  — a tool for erasing content in a canvas view — in a .


// An item that represents an eraser tool in the tool picker.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerEraserItem */

// Creates a new eraser item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/initWithEraserType:
func NewToolPickerEraserItemWithEraserType(eraserType EraserType) ToolPickerEraserItem {
	instance := getToolPickerEraserItemClass().Alloc()
	rv := objc.Send[ToolPickerEraserItem](instance.ID, objc.Sel("initWithEraserType:"), eraserType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewToolPickerEraserItemWithEraserType */


// Creates a new eraser item with the specified width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/initWithEraserType:width:
func NewToolPickerEraserItemWithEraserTypeWidth(eraserType EraserType, width float64) ToolPickerEraserItem {
	instance := getToolPickerEraserItemClass().Alloc()
	rv := objc.Send[ToolPickerEraserItem](instance.ID, objc.Sel("initWithEraserType:width:"), eraserType, width)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewToolPickerEraserItemWithEraserTypeWidth */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerEraserItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerEraserItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerEraserItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerEraserItem */

// An eraser tool for erasing parts of a drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerEraserItem/eraserTool-4q3hp
func (t_ ToolPickerEraserItem) EraserTool() IPKEraserTool {
	rv := objc.Send[EraserTool](t_.ID, objc.Sel("eraserTool"))
	return rv
}/* debug [instance_properties/getter]: eraserTool */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerEraserItem */


