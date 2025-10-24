// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRThermostatClusterThermostatScheduleTransition */


/* debug [class_header]: Header for MTRThermostatClusterThermostatScheduleTransition */
// The class instance for the [MTRThermostatClusterThermostatScheduleTransition] class.
var (
	MTRThermostatClusterThermostatScheduleTransitionClass     _MTRThermostatClusterThermostatScheduleTransitionClass
	MTRThermostatClusterThermostatScheduleTransitionClassOnce sync.Once
)

func getMTRThermostatClusterThermostatScheduleTransitionClass() _MTRThermostatClusterThermostatScheduleTransitionClass {
	MTRThermostatClusterThermostatScheduleTransitionClassOnce.Do(func() {
		MTRThermostatClusterThermostatScheduleTransitionClass = _MTRThermostatClusterThermostatScheduleTransitionClass{objc.GetClass("MTRThermostatClusterThermostatScheduleTransition")}
	})
	return MTRThermostatClusterThermostatScheduleTransitionClass
}

type _MTRThermostatClusterThermostatScheduleTransitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThermostatClusterThermostatScheduleTransition */
// An interface definition for the [MTRThermostatClusterThermostatScheduleTransition] class.
type IMTRThermostatClusterThermostatScheduleTransition interface {
	IMTRThermostatClusterWeeklyScheduleTransitionStruct
	
/* debug [class_interface_properties]: Properties for MTRThermostatClusterThermostatScheduleTransition */
	// properties:
	CoolSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetCoolSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	HeatSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetHeatSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThermostatClusterThermostatScheduleTransition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThermostatClusterThermostatScheduleTransition */
// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterThermostatScheduleTransitionClass) Alloc() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThermostatClusterThermostatScheduleTransitionClass) New() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterThermostatScheduleTransition) Init() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterThermostatScheduleTransition) Autorelease() MTRThermostatClusterThermostatScheduleTransition {
	rv := objc.Send[MTRThermostatClusterThermostatScheduleTransition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterThermostatScheduleTransition creates a new MTRThermostatClusterThermostatScheduleTransition instance.
func NewMTRThermostatClusterThermostatScheduleTransition() MTRThermostatClusterThermostatScheduleTransition {
	return getMTRThermostatClusterThermostatScheduleTransitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThermostatClusterThermostatScheduleTransition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition
type MTRThermostatClusterThermostatScheduleTransition struct {
	MTRThermostatClusterWeeklyScheduleTransitionStruct
}

// MTRThermostatClusterThermostatScheduleTransitionFrom constructs a [MTRThermostatClusterThermostatScheduleTransition] from an unsafe.Pointer.
func MTRThermostatClusterThermostatScheduleTransitionFrom(ptr unsafe.Pointer) MTRThermostatClusterThermostatScheduleTransition {
	return MTRThermostatClusterThermostatScheduleTransition{
		MTRThermostatClusterWeeklyScheduleTransitionStruct: MTRThermostatClusterWeeklyScheduleTransitionStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThermostatClusterThermostatScheduleTransition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThermostatClusterThermostatScheduleTransition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThermostatClusterThermostatScheduleTransition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThermostatClusterThermostatScheduleTransition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThermostatClusterThermostatScheduleTransition */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/coolSetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) CoolSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("coolSetpoint"))
	return rv
}/* debug [instance_properties/getter]: coolSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/coolSetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetCoolSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolSetpoint:"), value)
}/* debug [instance_properties/setter]: coolSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/heatSetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) HeatSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("heatSetpoint"))
	return rv
}/* debug [instance_properties/getter]: heatSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/heatSetpoint
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetHeatSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatSetpoint:"), value)
}/* debug [instance_properties/setter]: heatSetpoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/transitionTime
func (m_ MTRThermostatClusterThermostatScheduleTransition) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}/* debug [instance_properties/getter]: transitionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterThermostatScheduleTransition/transitionTime
func (m_ MTRThermostatClusterThermostatScheduleTransition) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}/* debug [instance_properties/setter]: transitionTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThermostatClusterThermostatScheduleTransition */



