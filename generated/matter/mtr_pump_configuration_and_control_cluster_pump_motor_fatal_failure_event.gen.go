// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent() MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent
type MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterPumpMotorFatalFailureEvent */



