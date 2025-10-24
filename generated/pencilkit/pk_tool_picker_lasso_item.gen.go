// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKToolPickerLassoItem */


/* debug [class_header]: Header for PKToolPickerLassoItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerLassoItem */
// An interface definition for the [ToolPickerLassoItem] class.
type IToolPickerLassoItem interface {
	IToolPickerItem
	
/* debug [class_interface_properties]: Properties for ToolPickerLassoItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerLassoItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerLassoItem */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerLassoItemClass) Alloc() ToolPickerLassoItem {
	rv := objc.Send[ToolPickerLassoItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerLassoItem */
// An item that represents a lasso tool in the tool picker.
//
// A lasso item represents a — a tool for selecting stroked lines and shapes in a canvas view — in a .


// An item that represents a lasso tool in the tool picker.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerLassoItem */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerLassoItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerLassoItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerLassoItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerLassoItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerLassoItem */


