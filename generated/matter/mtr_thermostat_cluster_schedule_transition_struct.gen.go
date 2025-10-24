// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterScheduleTransitionStruct */


/* debug [class_header]: Header for MTRThermostatClusterScheduleTransitionStruct */
// The class instance for the [MTRThermostatClusterScheduleTransitionStruct] class.
var (
	MTRThermostatClusterScheduleTransitionStructClass     _MTRThermostatClusterScheduleTransitionStructClass
	MTRThermostatClusterScheduleTransitionStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleTransitionStructClass() _MTRThermostatClusterScheduleTransitionStructClass {
	MTRThermostatClusterScheduleTransitionStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleTransitionStructClass = _MTRThermostatClusterScheduleTransitionStructClass{objc.GetClass("MTRThermostatClusterScheduleTransitionStruct")}
	})
	return MTRThermostatClusterScheduleTransitionStructClass
}

type _MTRThermostatClusterScheduleTransitionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterScheduleTransitionStruct */
// An interface definition for the [MTRThermostatClusterScheduleTransitionStruct] class.
type IMTRThermostatClusterScheduleTransitionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterScheduleTransitionStruct */
	// properties:
	CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	DayOfWeek() objc.IObject /* cross-framework: NSNumber */
	SetDayOfWeek(value objc.IObject /* cross-framework: NSNumber */)
	HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	PresetHandle() foundation.Data
	SetPresetHandle(value foundation.Data)
	SystemMode() objc.IObject /* cross-framework: NSNumber */
	SetSystemMode(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterScheduleTransitionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterScheduleTransitionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleTransitionStructClass) Alloc() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterScheduleTransitionStructClass) New() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleTransitionStruct) Init() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleTransitionStruct) Autorelease() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleTransitionStruct creates a new MTRThermostatClusterScheduleTransitionStruct instance.
func NewMTRThermostatClusterScheduleTransitionStruct() MTRThermostatClusterScheduleTransitionStruct {
	return getMTRThermostatClusterScheduleTransitionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterScheduleTransitionStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct
type MTRThermostatClusterScheduleTransitionStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleTransitionStructFrom constructs a [MTRThermostatClusterScheduleTransitionStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleTransitionStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleTransitionStruct {
	return MTRThermostatClusterScheduleTransitionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterScheduleTransitionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterScheduleTransitionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterScheduleTransitionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterScheduleTransitionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterScheduleTransitionStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/coolingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("coolingSetpoint"))
	return rv
}/* debug [instance_properties/getter]: coolingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/coolingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolingSetpoint:"), value)
}/* debug [instance_properties/setter]: coolingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/dayofweek
func (m_ MTRThermostatClusterScheduleTransitionStruct) DayOfWeek() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dayOfWeek"))
	return rv
}/* debug [instance_properties/getter]: dayOfWeek */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/dayofweek
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetDayOfWeek(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeek:"), value)
}/* debug [instance_properties/setter]: dayOfWeek */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/heatingsetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("heatingSetpoint"))
	return rv
}/* debug [instance_properties/getter]: heatingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/heatingsetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatingSetpoint:"), value)
}/* debug [instance_properties/setter]: heatingSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/presethandle
func (m_ MTRThermostatClusterScheduleTransitionStruct) PresetHandle() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("presetHandle"))
	return rv
}/* debug [instance_properties/getter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/presethandle
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetPresetHandle(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}/* debug [instance_properties/setter]: presetHandle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/systemmode
func (m_ MTRThermostatClusterScheduleTransitionStruct) SystemMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemMode"))
	return rv
}/* debug [instance_properties/getter]: systemMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/systemmode
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetSystemMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}/* debug [instance_properties/setter]: systemMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/transitiontime
func (m_ MTRThermostatClusterScheduleTransitionStruct) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletransitionstruct/transitiontime
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterScheduleTransitionStruct */



