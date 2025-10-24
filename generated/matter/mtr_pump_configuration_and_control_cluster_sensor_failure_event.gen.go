// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterSensorFailureEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterSensorFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSensorFailureEventClass     _MTRPumpConfigurationAndControlClusterSensorFailureEventClass
	MTRPumpConfigurationAndControlClusterSensorFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSensorFailureEventClass() _MTRPumpConfigurationAndControlClusterSensorFailureEventClass {
	MTRPumpConfigurationAndControlClusterSensorFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSensorFailureEventClass = _MTRPumpConfigurationAndControlClusterSensorFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSensorFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSensorFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterSensorFailureEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterSensorFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterSensorFailureEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSensorFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterSensorFailureEventClass) New() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSensorFailureEvent) Init() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSensorFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSensorFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSensorFailureEvent creates a new MTRPumpConfigurationAndControlClusterSensorFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterSensorFailureEvent() MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	return getMTRPumpConfigurationAndControlClusterSensorFailureEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterSensorFailureEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSensorFailureEvent
type MTRPumpConfigurationAndControlClusterSensorFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSensorFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterSensorFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSensorFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSensorFailureEvent {
	return MTRPumpConfigurationAndControlClusterSensorFailureEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterSensorFailureEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterSensorFailureEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterSensorFailureEvent */



