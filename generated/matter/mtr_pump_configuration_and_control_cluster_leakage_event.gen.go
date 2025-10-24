// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterLeakageEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterLeakageEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterLeakageEvent] class.
var (
	MTRPumpConfigurationAndControlClusterLeakageEventClass     _MTRPumpConfigurationAndControlClusterLeakageEventClass
	MTRPumpConfigurationAndControlClusterLeakageEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterLeakageEventClass() _MTRPumpConfigurationAndControlClusterLeakageEventClass {
	MTRPumpConfigurationAndControlClusterLeakageEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterLeakageEventClass = _MTRPumpConfigurationAndControlClusterLeakageEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterLeakageEvent")}
	})
	return MTRPumpConfigurationAndControlClusterLeakageEventClass
}

type _MTRPumpConfigurationAndControlClusterLeakageEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterLeakageEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterLeakageEvent] class.
type IMTRPumpConfigurationAndControlClusterLeakageEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterLeakageEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterLeakageEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterLeakageEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterLeakageEventClass) Alloc() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterLeakageEventClass) New() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterLeakageEvent) Init() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterLeakageEvent) Autorelease() MTRPumpConfigurationAndControlClusterLeakageEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterLeakageEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterLeakageEvent creates a new MTRPumpConfigurationAndControlClusterLeakageEvent instance.
func NewMTRPumpConfigurationAndControlClusterLeakageEvent() MTRPumpConfigurationAndControlClusterLeakageEvent {
	return getMTRPumpConfigurationAndControlClusterLeakageEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterLeakageEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterLeakageEvent
type MTRPumpConfigurationAndControlClusterLeakageEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterLeakageEventFrom constructs a [MTRPumpConfigurationAndControlClusterLeakageEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterLeakageEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterLeakageEvent {
	return MTRPumpConfigurationAndControlClusterLeakageEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterLeakageEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterLeakageEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterLeakageEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterLeakageEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterLeakageEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterLeakageEvent */



