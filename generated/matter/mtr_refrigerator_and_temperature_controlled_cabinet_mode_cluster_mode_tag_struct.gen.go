// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */


/* debug [class_header]: Header for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/mfgCode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}/* debug [instance_properties/getter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/mfgCode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}/* debug [instance_properties/setter]: mfgCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclustermodetagstruct/value
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrefrigeratorandtemperaturecontrolledcabinetmodeclustermodetagstruct/value
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct */



