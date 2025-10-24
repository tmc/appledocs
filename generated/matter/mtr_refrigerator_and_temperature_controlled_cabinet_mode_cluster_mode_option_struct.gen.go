// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */


/* debug [class_header]: Header for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/label
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/label
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclustermodeoptionstruct/mode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}/* debug [instance_properties/getter]: mode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclustermodeoptionstruct/mode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}/* debug [instance_properties/setter]: mode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct */



