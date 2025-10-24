// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKToolPickerItem */


/* debug [class_header]: Header for PKToolPickerItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolPickerItem */
// An interface definition for the [ToolPickerItem] class.
type IToolPickerItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ToolPickerItem */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Tool() IPKTool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolPickerItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolPickerItem */
// Alloc allocates a new instance without initialization.
func (tc _ToolPickerItemClass) Alloc() ToolPickerItem {
	rv := objc.Send[ToolPickerItem](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolPickerItem */
// The base class for an item in the tool picker.


// The base class for an item in the tool picker.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolPickerItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolPickerItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolPickerItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolPickerItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolPickerItem */

// A string that identifies the item in the tool picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerItem/identifier
func (t_ ToolPickerItem) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The this tool picker item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerItem/tool-918ln
func (t_ ToolPickerItem) Tool() IPKTool {
	rv := objc.Send[Tool](t_.ID, objc.Sel("tool"))
	return rv
}/* debug [instance_properties/getter]: tool */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKToolPickerItem */



