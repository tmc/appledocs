// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDishwasherModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRDishwasherModeClusterModeOptionStruct */
// The class instance for the [MTRDishwasherModeClusterModeOptionStruct] class.
var (
	MTRDishwasherModeClusterModeOptionStructClass     _MTRDishwasherModeClusterModeOptionStructClass
	MTRDishwasherModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRDishwasherModeClusterModeOptionStructClass() _MTRDishwasherModeClusterModeOptionStructClass {
	MTRDishwasherModeClusterModeOptionStructClassOnce.Do(func() {
		MTRDishwasherModeClusterModeOptionStructClass = _MTRDishwasherModeClusterModeOptionStructClass{objc.GetClass("MTRDishwasherModeClusterModeOptionStruct")}
	})
	return MTRDishwasherModeClusterModeOptionStructClass
}

type _MTRDishwasherModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDishwasherModeClusterModeOptionStruct */
// An interface definition for the [MTRDishwasherModeClusterModeOptionStruct] class.
type IMTRDishwasherModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDishwasherModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDishwasherModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDishwasherModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterModeOptionStructClass) Alloc() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDishwasherModeClusterModeOptionStructClass) New() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterModeOptionStruct) Init() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterModeOptionStruct) Autorelease() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterModeOptionStruct creates a new MTRDishwasherModeClusterModeOptionStruct instance.
func NewMTRDishwasherModeClusterModeOptionStruct() MTRDishwasherModeClusterModeOptionStruct {
	return getMTRDishwasherModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDishwasherModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct
type MTRDishwasherModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRDishwasherModeClusterModeOptionStructFrom constructs a [MTRDishwasherModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRDishwasherModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterModeOptionStruct {
	return MTRDishwasherModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDishwasherModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDishwasherModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDishwasherModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDishwasherModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDishwasherModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclustermodeoptionstruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdishwashermodeclustermodeoptionstruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDishwasherModeClusterModeOptionStruct */



