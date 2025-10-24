// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterGeneralFaultEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] class.
var (
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClass     _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterGeneralFaultEventClass() _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass {
	MTRPumpConfigurationAndControlClusterGeneralFaultEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterGeneralFaultEventClass = _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterGeneralFaultEvent")}
	})
	return MTRPumpConfigurationAndControlClusterGeneralFaultEventClass
}

type _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] class.
type IMTRPumpConfigurationAndControlClusterGeneralFaultEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass) Alloc() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterGeneralFaultEventClass) New() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterGeneralFaultEvent) Init() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterGeneralFaultEvent) Autorelease() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterGeneralFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterGeneralFaultEvent creates a new MTRPumpConfigurationAndControlClusterGeneralFaultEvent instance.
func NewMTRPumpConfigurationAndControlClusterGeneralFaultEvent() MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	return getMTRPumpConfigurationAndControlClusterGeneralFaultEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterGeneralFaultEvent
type MTRPumpConfigurationAndControlClusterGeneralFaultEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterGeneralFaultEventFrom constructs a [MTRPumpConfigurationAndControlClusterGeneralFaultEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterGeneralFaultEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterGeneralFaultEvent {
	return MTRPumpConfigurationAndControlClusterGeneralFaultEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterGeneralFaultEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterGeneralFaultEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterGeneralFaultEvent */



