// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKToolPickerScribbleItem */


/* debug [class_header]: Header for PKToolPickerScribbleItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerScribbleItem */
// An interface definition for the [ToolPickerScribbleItem] class.
type IToolPickerScribbleItem interface {
	IToolPickerItem
	
/* debug [class_interface_properties]: Properties for ToolPickerScribbleItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerScribbleItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerScribbleItem */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerScribbleItemClass) Alloc() ToolPickerScribbleItem {
	rv := objc.Send[ToolPickerScribbleItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerScribbleItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerScribbleItem */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerScribbleItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerScribbleItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerScribbleItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerScribbleItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerScribbleItem */


