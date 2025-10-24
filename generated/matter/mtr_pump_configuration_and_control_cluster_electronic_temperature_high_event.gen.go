// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass     _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass() _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass {
	MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass = _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass) New() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent) Init() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent creates a new MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent() MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent
type MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent {
	return MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterElectronicTemperatureHighEvent */



