// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKLassoTool */


/* debug [class_header]: Header for PKLassoTool */
// The class instance for the [LassoTool] class.
var (
	LassoToolClass     _LassoToolClass
	LassoToolClassOnce sync.Once
)

func getLassoToolClass() _LassoToolClass {
	LassoToolClassOnce.Do(func() {
		LassoToolClass = _LassoToolClass{objc.GetClass("PKLassoTool")}
	})
	return LassoToolClass
}

type _LassoToolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LassoTool */
// An interface definition for the [LassoTool] class.
type ILassoTool interface {
	ITool
	
/* debug [class_interface_properties]: Properties for LassoTool */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LassoTool */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LassoTool */
// Alloc allocates a new instance without initialization.
func (lc _LassoToolClass) Alloc() LassoTool {
	rv := objc.Send[LassoTool](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LassoToolClass) New() LassoTool {
	rv := objc.Send[LassoTool](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LassoTool) Init() LassoTool {
	rv := objc.Send[LassoTool](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LassoTool) Autorelease() LassoTool {
	rv := objc.Send[LassoTool](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLassoTool creates a new LassoTool instance.
func NewLassoTool() LassoTool {
	return getLassoToolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LassoTool */
// A tool for selecting stroked lines and shapes in a canvas view.
//
// A object supports the selection of content on a . When active, the canvas uses incoming touch events to determine what content to add to the selection. Create a lasso tool programmatically or display a object from which the user selects the tool. Assign the resulting object to the property of your object. The canvas uses any subsequent touch sequences to select content on the canvas.


// A tool for selecting stroked lines and shapes in a canvas view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKLassoToolReference
type LassoTool struct {
	Tool
}

// LassoToolFrom constructs a [LassoTool] from an unsafe.Pointer.
//
// A tool for selecting stroked lines and shapes in a canvas view.
func LassoToolFrom(ptr unsafe.Pointer) LassoTool {
	return LassoTool{
		Tool: ToolFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LassoTool */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LassoTool */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LassoTool */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LassoTool */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LassoTool */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKLassoTool */


