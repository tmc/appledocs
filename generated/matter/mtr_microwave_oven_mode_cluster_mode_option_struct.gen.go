// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMicrowaveOvenModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRMicrowaveOvenModeClusterModeOptionStruct */
// The class instance for the [MTRMicrowaveOvenModeClusterModeOptionStruct] class.
var (
	MTRMicrowaveOvenModeClusterModeOptionStructClass     _MTRMicrowaveOvenModeClusterModeOptionStructClass
	MTRMicrowaveOvenModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRMicrowaveOvenModeClusterModeOptionStructClass() _MTRMicrowaveOvenModeClusterModeOptionStructClass {
	MTRMicrowaveOvenModeClusterModeOptionStructClassOnce.Do(func() {
		MTRMicrowaveOvenModeClusterModeOptionStructClass = _MTRMicrowaveOvenModeClusterModeOptionStructClass{objc.GetClass("MTRMicrowaveOvenModeClusterModeOptionStruct")}
	})
	return MTRMicrowaveOvenModeClusterModeOptionStructClass
}

type _MTRMicrowaveOvenModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMicrowaveOvenModeClusterModeOptionStruct */
// An interface definition for the [MTRMicrowaveOvenModeClusterModeOptionStruct] class.
type IMTRMicrowaveOvenModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMicrowaveOvenModeClusterModeOptionStruct */
	// properties:
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMicrowaveOvenModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMicrowaveOvenModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenModeClusterModeOptionStructClass) Alloc() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMicrowaveOvenModeClusterModeOptionStructClass) New() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Init() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Autorelease() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenModeClusterModeOptionStruct creates a new MTRMicrowaveOvenModeClusterModeOptionStruct instance.
func NewMTRMicrowaveOvenModeClusterModeOptionStruct() MTRMicrowaveOvenModeClusterModeOptionStruct {
	return getMTRMicrowaveOvenModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMicrowaveOvenModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct
type MTRMicrowaveOvenModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRMicrowaveOvenModeClusterModeOptionStructFrom constructs a [MTRMicrowaveOvenModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRMicrowaveOvenModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRMicrowaveOvenModeClusterModeOptionStruct {
	return MTRMicrowaveOvenModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMicrowaveOvenModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMicrowaveOvenModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMicrowaveOvenModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMicrowaveOvenModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMicrowaveOvenModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovenmodeclustermodeoptionstruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmicrowaveovenmodeclustermodeoptionstruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMicrowaveOvenModeClusterModeOptionStruct */



