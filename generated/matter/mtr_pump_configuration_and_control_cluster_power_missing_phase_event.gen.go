// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] class.
var (
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass     _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass() _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass {
	MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass = _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent")}
	})
	return MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass
}

type _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] class.
type IMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass) Alloc() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass) New() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent) Init() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent) Autorelease() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent creates a new MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent instance.
func NewMTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent() MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	return getMTRPumpConfigurationAndControlClusterPowerMissingPhaseEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent
type MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventFrom constructs a [MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterPowerMissingPhaseEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent {
	return MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterPowerMissingPhaseEvent */



