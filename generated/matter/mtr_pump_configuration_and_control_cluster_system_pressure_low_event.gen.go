// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] class.
var (
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass     _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterSystemPressureLowEventClass() _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass {
	MTRPumpConfigurationAndControlClusterSystemPressureLowEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass = _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterSystemPressureLowEvent")}
	})
	return MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass
}

type _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] class.
type IMTRPumpConfigurationAndControlClusterSystemPressureLowEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass) Alloc() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterSystemPressureLowEventClass) New() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureLowEvent) Init() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterSystemPressureLowEvent) Autorelease() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterSystemPressureLowEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterSystemPressureLowEvent creates a new MTRPumpConfigurationAndControlClusterSystemPressureLowEvent instance.
func NewMTRPumpConfigurationAndControlClusterSystemPressureLowEvent() MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	return getMTRPumpConfigurationAndControlClusterSystemPressureLowEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterSystemPressureLowEvent
type MTRPumpConfigurationAndControlClusterSystemPressureLowEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterSystemPressureLowEventFrom constructs a [MTRPumpConfigurationAndControlClusterSystemPressureLowEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterSystemPressureLowEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterSystemPressureLowEvent {
	return MTRPumpConfigurationAndControlClusterSystemPressureLowEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterSystemPressureLowEvent */



