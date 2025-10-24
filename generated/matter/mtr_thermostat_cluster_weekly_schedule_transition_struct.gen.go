// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterWeeklyScheduleTransitionStruct */


/* debug [class_header]: Header for MTRThermostatClusterWeeklyScheduleTransitionStruct */
// The class instance for the [MTRThermostatClusterWeeklyScheduleTransitionStruct] class.
var (
	MTRThermostatClusterWeeklyScheduleTransitionStructClass     _MTRThermostatClusterWeeklyScheduleTransitionStructClass
	MTRThermostatClusterWeeklyScheduleTransitionStructClassOnce sync.Once
)

func getMTRThermostatClusterWeeklyScheduleTransitionStructClass() _MTRThermostatClusterWeeklyScheduleTransitionStructClass {
	MTRThermostatClusterWeeklyScheduleTransitionStructClassOnce.Do(func() {
		MTRThermostatClusterWeeklyScheduleTransitionStructClass = _MTRThermostatClusterWeeklyScheduleTransitionStructClass{objc.GetClass("MTRThermostatClusterWeeklyScheduleTransitionStruct")}
	})
	return MTRThermostatClusterWeeklyScheduleTransitionStructClass
}

type _MTRThermostatClusterWeeklyScheduleTransitionStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterWeeklyScheduleTransitionStruct */
// An interface definition for the [MTRThermostatClusterWeeklyScheduleTransitionStruct] class.
type IMTRThermostatClusterWeeklyScheduleTransitionStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterWeeklyScheduleTransitionStruct */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterWeeklyScheduleTransitionStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterWeeklyScheduleTransitionStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterWeeklyScheduleTransitionStructClass) Alloc() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterWeeklyScheduleTransitionStructClass) New() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterWeeklyScheduleTransitionStruct) Init() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterWeeklyScheduleTransitionStruct) Autorelease() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterWeeklyScheduleTransitionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterWeeklyScheduleTransitionStruct creates a new MTRThermostatClusterWeeklyScheduleTransitionStruct instance.
func NewMTRThermostatClusterWeeklyScheduleTransitionStruct() MTRThermostatClusterWeeklyScheduleTransitionStruct {
	return getMTRThermostatClusterWeeklyScheduleTransitionStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterWeeklyScheduleTransitionStruct */
// A parent class referenced by other Matter classes.


// A parent class referenced by other Matter classes. [Full Topic]
type MTRThermostatClusterWeeklyScheduleTransitionStruct struct {
	objectivec.Object
}

// MTRThermostatClusterWeeklyScheduleTransitionStructFrom constructs a [MTRThermostatClusterWeeklyScheduleTransitionStruct] from an unsafe.Pointer.
//
// A parent class referenced by other Matter classes.
func MTRThermostatClusterWeeklyScheduleTransitionStructFrom(ptr unsafe.Pointer) MTRThermostatClusterWeeklyScheduleTransitionStruct {
	return MTRThermostatClusterWeeklyScheduleTransitionStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterWeeklyScheduleTransitionStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterWeeklyScheduleTransitionStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterWeeklyScheduleTransitionStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterWeeklyScheduleTransitionStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterWeeklyScheduleTransitionStruct */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterWeeklyScheduleTransitionStruct */



