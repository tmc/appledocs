// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTROvenModeClusterModeOptionStruct */
// The class instance for the [MTROvenModeClusterModeOptionStruct] class.
var (
	MTROvenModeClusterModeOptionStructClass     _MTROvenModeClusterModeOptionStructClass
	MTROvenModeClusterModeOptionStructClassOnce sync.Once
)

func getMTROvenModeClusterModeOptionStructClass() _MTROvenModeClusterModeOptionStructClass {
	MTROvenModeClusterModeOptionStructClassOnce.Do(func() {
		MTROvenModeClusterModeOptionStructClass = _MTROvenModeClusterModeOptionStructClass{objc.GetClass("MTROvenModeClusterModeOptionStruct")}
	})
	return MTROvenModeClusterModeOptionStructClass
}

type _MTROvenModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenModeClusterModeOptionStruct */
// An interface definition for the [MTROvenModeClusterModeOptionStruct] class.
type IMTROvenModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterModeOptionStructClass) Alloc() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenModeClusterModeOptionStructClass) New() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterModeOptionStruct) Init() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterModeOptionStruct) Autorelease() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterModeOptionStruct creates a new MTROvenModeClusterModeOptionStruct instance.
func NewMTROvenModeClusterModeOptionStruct() MTROvenModeClusterModeOptionStruct {
	return getMTROvenModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct
type MTROvenModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTROvenModeClusterModeOptionStructFrom constructs a [MTROvenModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTROvenModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTROvenModeClusterModeOptionStruct {
	return MTROvenModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/label
func (m_ MTROvenModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/label
func (m_ MTROvenModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclustermodeoptionstruct/mode
func (m_ MTROvenModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovenmodeclustermodeoptionstruct/mode
func (m_ MTROvenModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenModeClusterModeOptionStruct */



