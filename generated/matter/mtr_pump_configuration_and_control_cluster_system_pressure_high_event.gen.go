// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass     _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSystemPressureHighEventClass() _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass {
	MTRPumpConfigurationAndControlClusterSystemPressureHighEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass = _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSystemPressureHighEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass
}

type _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] class.
type IMTRPumpConfigurationAndControlClusterSystemPressureHighEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass) Alloc() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureHighEventClass) New() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureHighEvent) Init() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureHighEvent) Autorelease() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureHighEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSystemPressureHighEvent creates a new MTRPumpConfigurationAndControlClusterSystemPressureHighEvent instance.
func NewMTRPumpConfigurationAndControlClusterSystemPressureHighEvent() MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	return getMTRPumpConfigurationAndControlClusterSystemPressureHighEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSystemPressureHighEvent
type MTRPumpConfigurationAndControlClusterSystemPressureHighEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSystemPressureHighEventFrom constructs a [MTRPumpConfigurationAndControlClusterSystemPressureHighEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSystemPressureHighEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSystemPressureHighEvent {
	return MTRPumpConfigurationAndControlClusterSystemPressureHighEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterSystemPressureHighEvent */



