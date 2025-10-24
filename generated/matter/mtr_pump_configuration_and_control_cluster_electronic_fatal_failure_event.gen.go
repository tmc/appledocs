// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent() MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicFatalFailureEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent
type MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterElectronicFatalFailureEvent */



