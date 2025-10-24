// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterDryRunningEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterDryRunningEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterDryRunningEvent] class.
var (
	MTRPumpConfigurationAndControlClusterDryRunningEventClass     _MTRPumpConfigurationAndControlClusterDryRunningEventClass
	MTRPumpConfigurationAndControlClusterDryRunningEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterDryRunningEventClass() _MTRPumpConfigurationAndControlClusterDryRunningEventClass {
	MTRPumpConfigurationAndControlClusterDryRunningEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterDryRunningEventClass = _MTRPumpConfigurationAndControlClusterDryRunningEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterDryRunningEvent")}
	})
	return MTRPumpConfigurationAndControlClusterDryRunningEventClass
}

type _MTRPumpConfigurationAndControlClusterDryRunningEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterDryRunningEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterDryRunningEvent] class.
type IMTRPumpConfigurationAndControlClusterDryRunningEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterDryRunningEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterDryRunningEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterDryRunningEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterDryRunningEventClass) Alloc() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterDryRunningEventClass) New() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterDryRunningEvent) Init() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterDryRunningEvent) Autorelease() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterDryRunningEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterDryRunningEvent creates a new MTRPumpConfigurationAndControlClusterDryRunningEvent instance.
func NewMTRPumpConfigurationAndControlClusterDryRunningEvent() MTRPumpConfigurationAndControlClusterDryRunningEvent {
	return getMTRPumpConfigurationAndControlClusterDryRunningEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterDryRunningEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterDryRunningEvent
type MTRPumpConfigurationAndControlClusterDryRunningEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterDryRunningEventFrom constructs a [MTRPumpConfigurationAndControlClusterDryRunningEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterDryRunningEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterDryRunningEvent {
	return MTRPumpConfigurationAndControlClusterDryRunningEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterDryRunningEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterDryRunningEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterDryRunningEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterDryRunningEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterDryRunningEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterDryRunningEvent */



