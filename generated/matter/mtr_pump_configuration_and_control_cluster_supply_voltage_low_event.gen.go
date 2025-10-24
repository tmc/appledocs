// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass     _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass() _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass {
	MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass = _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass
}

type _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] class.
type IMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass) Alloc() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass) New() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent) Init() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent) Autorelease() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent creates a new MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent instance.
func NewMTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent() MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	return getMTRPumpConfigurationAndControlClusterSupplyVoltageLowEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent
type MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventFrom constructs a [MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSupplyVoltageLowEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent {
	return MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterSupplyVoltageLowEvent */



