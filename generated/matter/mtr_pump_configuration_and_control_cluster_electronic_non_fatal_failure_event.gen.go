// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] class.
var (
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass     _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass() _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass {
	MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass = _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent")}
	})
	return MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass
}

type _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] class.
type IMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass) Alloc() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass) New() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent) Init() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent) Autorelease() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent creates a new MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent instance.
func NewMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent() MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	return getMTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent
type MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventFrom constructs a [MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent {
	return MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterElectronicNonFatalFailureEvent */



