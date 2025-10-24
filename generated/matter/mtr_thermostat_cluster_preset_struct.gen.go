// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterPresetStruct */


/* debug [class_header]: Header for MTRThermostatClusterPresetStruct */
// The class instance for the [MTRThermostatClusterPresetStruct] class.
var (
	MTRThermostatClusterPresetStructClass     _MTRThermostatClusterPresetStructClass
	MTRThermostatClusterPresetStructClassOnce sync.Once
)

func getMTRThermostatClusterPresetStructClass() _MTRThermostatClusterPresetStructClass {
	MTRThermostatClusterPresetStructClassOnce.Do(func() {
		MTRThermostatClusterPresetStructClass = _MTRThermostatClusterPresetStructClass{objc.GetClass("MTRThermostatClusterPresetStruct")}
	})
	return MTRThermostatClusterPresetStructClass
}

type _MTRThermostatClusterPresetStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterPresetStruct */
// An interface definition for the [MTRThermostatClusterPresetStruct] class.
type IMTRThermostatClusterPresetStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterPresetStruct */
	// properties:
	BuiltIn() objc.IObject /* cross-framework: NSNumber */
	SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */)
	CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PresetHandle() foundation.Data
	SetPresetHandle(value foundation.Data)
	PresetScenario() objc.IObject /* cross-framework: NSNumber */
	SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterPresetStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterPresetStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterPresetStructClass) Alloc() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterPresetStructClass) New() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterPresetStruct) Init() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterPresetStruct) Autorelease() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterPresetStruct creates a new MTRThermostatClusterPresetStruct instance.
func NewMTRThermostatClusterPresetStruct() MTRThermostatClusterPresetStruct {
	return getMTRThermostatClusterPresetStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterPresetStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct
type MTRThermostatClusterPresetStruct struct {
	objectivec.Object
}

// MTRThermostatClusterPresetStructFrom constructs a [MTRThermostatClusterPresetStruct] from an unsafe.Pointer.
func MTRThermostatClusterPresetStructFrom(ptr unsafe.Pointer) MTRThermostatClusterPresetStruct {
	return MTRThermostatClusterPresetStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterPresetStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterPresetStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterPresetStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterPresetStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterPresetStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) BuiltIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("builtIn"))
	return rv
}/* debug [instance_properties/getter]: builtIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}/* debug [instance_properties/setter]: builtIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/coolingsetpoint
func (m_ MTRThermostatClusterPresetStruct) CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("coolingSetpoint"))
	return rv
}/* debug [instance_properties/getter]: coolingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/coolingsetpoint
func (m_ MTRThermostatClusterPresetStruct) SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolingSetpoint:"), value)
}/* debug [instance_properties/setter]: coolingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/heatingsetpoint
func (m_ MTRThermostatClusterPresetStruct) HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("heatingSetpoint"))
	return rv
}/* debug [instance_properties/getter]: heatingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/heatingsetpoint
func (m_ MTRThermostatClusterPresetStruct) SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatingSetpoint:"), value)
}/* debug [instance_properties/setter]: heatingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/name
func (m_ MTRThermostatClusterPresetStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/name
func (m_ MTRThermostatClusterPresetStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/presethandle
func (m_ MTRThermostatClusterPresetStruct) PresetHandle() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("presetHandle"))
	return rv
}/* debug [instance_properties/getter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/presethandle
func (m_ MTRThermostatClusterPresetStruct) SetPresetHandle(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}/* debug [instance_properties/setter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/presetscenario
func (m_ MTRThermostatClusterPresetStruct) PresetScenario() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetScenario"))
	return rv
}/* debug [instance_properties/getter]: presetScenario */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterpresetstruct/presetscenario
func (m_ MTRThermostatClusterPresetStruct) SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}/* debug [instance_properties/setter]: presetScenario */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterPresetStruct */



