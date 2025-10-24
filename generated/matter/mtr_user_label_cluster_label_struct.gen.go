// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUserLabelClusterLabelStruct */


/* debug [class_header]: Header for MTRUserLabelClusterLabelStruct */
// The class instance for the [MTRUserLabelClusterLabelStruct] class.
var (
	MTRUserLabelClusterLabelStructClass     _MTRUserLabelClusterLabelStructClass
	MTRUserLabelClusterLabelStructClassOnce sync.Once
)

func getMTRUserLabelClusterLabelStructClass() _MTRUserLabelClusterLabelStructClass {
	MTRUserLabelClusterLabelStructClassOnce.Do(func() {
		MTRUserLabelClusterLabelStructClass = _MTRUserLabelClusterLabelStructClass{objc.GetClass("MTRUserLabelClusterLabelStruct")}
	})
	return MTRUserLabelClusterLabelStructClass
}

type _MTRUserLabelClusterLabelStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUserLabelClusterLabelStruct */
// An interface definition for the [MTRUserLabelClusterLabelStruct] class.
type IMTRUserLabelClusterLabelStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUserLabelClusterLabelStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUserLabelClusterLabelStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUserLabelClusterLabelStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRUserLabelClusterLabelStructClass) Alloc() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUserLabelClusterLabelStructClass) New() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUserLabelClusterLabelStruct) Init() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUserLabelClusterLabelStruct) Autorelease() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUserLabelClusterLabelStruct creates a new MTRUserLabelClusterLabelStruct instance.
func NewMTRUserLabelClusterLabelStruct() MTRUserLabelClusterLabelStruct {
	return getMTRUserLabelClusterLabelStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUserLabelClusterLabelStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct
type MTRUserLabelClusterLabelStruct struct {
	objectivec.Object
}

// MTRUserLabelClusterLabelStructFrom constructs a [MTRUserLabelClusterLabelStruct] from an unsafe.Pointer.
func MTRUserLabelClusterLabelStructFrom(ptr unsafe.Pointer) MTRUserLabelClusterLabelStruct {
	return MTRUserLabelClusterLabelStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUserLabelClusterLabelStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUserLabelClusterLabelStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUserLabelClusterLabelStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUserLabelClusterLabelStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUserLabelClusterLabelStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct/label
func (m_ MTRUserLabelClusterLabelStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct/label
func (m_ MTRUserLabelClusterLabelStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct/value
func (m_ MTRUserLabelClusterLabelStruct) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct/value
func (m_ MTRUserLabelClusterLabelStruct) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUserLabelClusterLabelStruct */



