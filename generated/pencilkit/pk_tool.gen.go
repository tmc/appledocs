// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKTool */


/* debug [class_header]: Header for PKTool */
// The class instance for the [Tool] class.
var (
	ToolClass     _ToolClass
	ToolClassOnce sync.Once
)

func getToolClass() _ToolClass {
	ToolClassOnce.Do(func() {
		ToolClass = _ToolClass{objc.GetClass("PKTool")}
	})
	return ToolClass
}

type _ToolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Tool */
// An interface definition for the [Tool] class.
type ITool interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Tool */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Tool */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Tool */
// Alloc allocates a new instance without initialization.
func (tc _ToolClass) Alloc() Tool {
	rv := objc.Send[Tool](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _ToolClass) New() Tool {
	rv := objc.Send[Tool](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ Tool) Init() Tool {
	rv := objc.Send[Tool](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ Tool) Autorelease() Tool {
	rv := objc.Send[Tool](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTool creates a new Tool instance.
func NewTool() Tool {
	return getToolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Tool */
// An abstract base class for tools used by a canvas view.
//
// A object is an abstract base class for tool types associated with a . Tools are user-facing, and the selected tool determines how the canvas interprets incoming gestures. Don’t create objects directly. Instead, create one of its subclasses to provide users with the desired the tool behavior.


// An abstract base class for tools used by a canvas view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKTool-c.class
type Tool struct {
	objectivec.Object
}

// ToolFrom constructs a [Tool] from an unsafe.Pointer.
//
// An abstract base class for tools used by a canvas view.
func ToolFrom(ptr unsafe.Pointer) Tool {
	return Tool{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Tool *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Tool */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Tool */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Tool */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Tool */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKTool */



