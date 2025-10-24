// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKToolPickerRulerItem */


/* debug [class_header]: Header for PKToolPickerRulerItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerRulerItem */
// An interface definition for the [ToolPickerRulerItem] class.
type IToolPickerRulerItem interface {
	IToolPickerItem
	
/* debug [class_interface_properties]: Properties for ToolPickerRulerItem */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerRulerItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerRulerItem */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerRulerItemClass) Alloc() ToolPickerRulerItem {
	rv := objc.Send[ToolPickerRulerItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerRulerItem */
// An item that represents a ruler tool in the tool picker.


// An item that represents a ruler tool in the tool picker.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerRulerItem */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerRulerItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerRulerItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerRulerItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerRulerItem */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerRulerItem */


