// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPumpConfigurationAndControlClusterAirDetectionEvent */


/* debug [class_header]: Header for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
// The class instance for the [MTRPumpConfigurationAndControlClusterAirDetectionEvent] class.
var (
	MTRPumpConfigurationAndControlClusterAirDetectionEventClass     _MTRPumpConfigurationAndControlClusterAirDetectionEventClass
	MTRPumpConfigurationAndControlClusterAirDetectionEventClassOnce sync.Once
)

func getMTRPumpConfigurationAndControlClusterAirDetectionEventClass() _MTRPumpConfigurationAndControlClusterAirDetectionEventClass {
	MTRPumpConfigurationAndControlClusterAirDetectionEventClassOnce.Do(func() {
		MTRPumpConfigurationAndControlClusterAirDetectionEventClass = _MTRPumpConfigurationAndControlClusterAirDetectionEventClass{objc.GetClass("MTRPumpConfigurationAndControlClusterAirDetectionEvent")}
	})
	return MTRPumpConfigurationAndControlClusterAirDetectionEventClass
}

type _MTRPumpConfigurationAndControlClusterAirDetectionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
// An interface definition for the [MTRPumpConfigurationAndControlClusterAirDetectionEvent] class.
type IMTRPumpConfigurationAndControlClusterAirDetectionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPumpConfigurationAndControlClusterAirDetectionEventClass) Alloc() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPumpConfigurationAndControlClusterAirDetectionEventClass) New() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPumpConfigurationAndControlClusterAirDetectionEvent) Init() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPumpConfigurationAndControlClusterAirDetectionEvent) Autorelease() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	rv := objc.Send[MTRPumpConfigurationAndControlClusterAirDetectionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPumpConfigurationAndControlClusterAirDetectionEvent creates a new MTRPumpConfigurationAndControlClusterAirDetectionEvent instance.
func NewMTRPumpConfigurationAndControlClusterAirDetectionEvent() MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	return getMTRPumpConfigurationAndControlClusterAirDetectionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPumpConfigurationAndControlClusterAirDetectionEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPumpConfigurationAndControlClusterAirDetectionEvent
type MTRPumpConfigurationAndControlClusterAirDetectionEvent struct {
	objectivec.Object
}

// MTRPumpConfigurationAndControlClusterAirDetectionEventFrom constructs a [MTRPumpConfigurationAndControlClusterAirDetectionEvent] from an unsafe.Pointer.
func MTRPumpConfigurationAndControlClusterAirDetectionEventFrom(ptr unsafe.Pointer) MTRPumpConfigurationAndControlClusterAirDetectionEvent {
	return MTRPumpConfigurationAndControlClusterAirDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPumpConfigurationAndControlClusterAirDetectionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPumpConfigurationAndControlClusterAirDetectionEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPumpConfigurationAndControlClusterAirDetectionEvent */



