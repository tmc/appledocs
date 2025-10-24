// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass     _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass() _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass {
	MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass = _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass) New() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent) Init() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent creates a new MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent() MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	return getMTRPumpConfigurationAndControlClusterMotorTemperatureHighEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent
type MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterMotorTemperatureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent {
	return MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterMotorTemperatureHighEvent */



