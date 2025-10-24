// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterTurbineOperationEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] class.
var (
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClass     _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterTurbineOperationEventClass() _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass {
	MTRPumpConfigurationAndControlClusterTurbineOperationEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterTurbineOperationEventClass = _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterTurbineOperationEvent")}
	})
	return MTRPumpConfigurationAndControlClusterTurbineOperationEventClass
}

type _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] class.
type IMTRPumpConfigurationAndControlClusterTurbineOperationEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass) Alloc() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterTurbineOperationEventClass) New() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterTurbineOperationEvent) Init() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterTurbineOperationEvent) Autorelease() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterTurbineOperationEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterTurbineOperationEvent creates a new MTRPumpConfigurationAndControlClusterTurbineOperationEvent instance.
func NewMTRPumpConfigurationAndControlClusterTurbineOperationEvent() MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	return getMTRPumpConfigurationAndControlClusterTurbineOperationEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterTurbineOperationEvent
type MTRPumpConfigurationAndControlClusterTurbineOperationEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterTurbineOperationEventFrom constructs a [MTRPumpConfigurationAndControlClusterTurbineOperationEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterTurbineOperationEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterTurbineOperationEvent {
	return MTRPumpConfigurationAndControlClusterTurbineOperationEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterTurbineOperationEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterTurbineOperationEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterTurbineOperationEvent */



