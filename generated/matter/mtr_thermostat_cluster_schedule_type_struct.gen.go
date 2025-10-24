// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterScheduleTypeStruct */


/* debug [class_header]: Header for MTRThermostatClusterScheduleTypeStruct */
// The class instance for the [MTRThermostatClusterScheduleTypeStruct] class.
var (
	MTRThermostatClusterScheduleTypeStructClass     _MTRThermostatClusterScheduleTypeStructClass
	MTRThermostatClusterScheduleTypeStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleTypeStructClass() _MTRThermostatClusterScheduleTypeStructClass {
	MTRThermostatClusterScheduleTypeStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleTypeStructClass = _MTRThermostatClusterScheduleTypeStructClass{objc.GetClass("MTRThermostatClusterScheduleTypeStruct")}
	})
	return MTRThermostatClusterScheduleTypeStructClass
}

type _MTRThermostatClusterScheduleTypeStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterScheduleTypeStruct */
// An interface definition for the [MTRThermostatClusterScheduleTypeStruct] class.
type IMTRThermostatClusterScheduleTypeStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterScheduleTypeStruct */
	// properties:
	NumberOfSchedules() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfSchedules(value objc.IObject /* cross-framework: NSNumber */)
	ScheduleTypeFeatures() objc.IObject /* cross-framework: NSNumber */
	SetScheduleTypeFeatures(value objc.IObject /* cross-framework: NSNumber */)
	SystemMode() objc.IObject /* cross-framework: NSNumber */
	SetSystemMode(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterScheduleTypeStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterScheduleTypeStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleTypeStructClass) Alloc() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterScheduleTypeStructClass) New() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleTypeStruct) Init() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleTypeStruct) Autorelease() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleTypeStruct creates a new MTRThermostatClusterScheduleTypeStruct instance.
func NewMTRThermostatClusterScheduleTypeStruct() MTRThermostatClusterScheduleTypeStruct {
	return getMTRThermostatClusterScheduleTypeStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterScheduleTypeStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct
type MTRThermostatClusterScheduleTypeStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleTypeStructFrom constructs a [MTRThermostatClusterScheduleTypeStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleTypeStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleTypeStruct {
	return MTRThermostatClusterScheduleTypeStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterScheduleTypeStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterScheduleTypeStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterScheduleTypeStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterScheduleTypeStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterScheduleTypeStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/numberOfSchedules
func (m_ MTRThermostatClusterScheduleTypeStruct) NumberOfSchedules() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfSchedules"))
	return rv
}/* debug [instance_properties/getter]: numberOfSchedules */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/numberOfSchedules
func (m_ MTRThermostatClusterScheduleTypeStruct) SetNumberOfSchedules(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfSchedules:"), value)
}/* debug [instance_properties/setter]: numberOfSchedules */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletypestruct/scheduletypefeatures
func (m_ MTRThermostatClusterScheduleTypeStruct) ScheduleTypeFeatures() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("scheduleTypeFeatures"))
	return rv
}/* debug [instance_properties/getter]: scheduleTypeFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletypestruct/scheduletypefeatures
func (m_ MTRThermostatClusterScheduleTypeStruct) SetScheduleTypeFeatures(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleTypeFeatures:"), value)
}/* debug [instance_properties/setter]: scheduleTypeFeatures */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletypestruct/systemmode
func (m_ MTRThermostatClusterScheduleTypeStruct) SystemMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemMode"))
	return rv
}/* debug [instance_properties/getter]: systemMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthermostatclusterscheduletypestruct/systemmode
func (m_ MTRThermostatClusterScheduleTypeStruct) SetSystemMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}/* debug [instance_properties/setter]: systemMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterScheduleTypeStruct */



