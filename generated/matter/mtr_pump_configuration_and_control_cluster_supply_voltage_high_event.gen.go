// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass     _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass() _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass {
	MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass = _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass
}

type _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] class.
type IMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass) New() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent) Init() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent creates a new MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent() MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	return getMTRPumpConfigurationAndControlClusterSupplyVoltageHighEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent
type MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSupplyVoltageHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent {
	return MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterSupplyVoltageHighEvent */



