// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWaterHeaterManagementClusterBoostEndedEvent */


/* debug [class_header]: Header for MTRWaterHeaterManagementClusterBoostEndedEvent */
// The class instance for the [MTRWaterHeaterManagementClusterBoostEndedEvent] class.
var (
	MTRWaterHeaterManagementClusterBoostEndedEventClass     _MTRWaterHeaterManagementClusterBoostEndedEventClass
	MTRWaterHeaterManagementClusterBoostEndedEventClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostEndedEventClass() _MTRWaterHeaterManagementClusterBoostEndedEventClass {
	MTRWaterHeaterManagementClusterBoostEndedEventClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostEndedEventClass = _MTRWaterHeaterManagementClusterBoostEndedEventClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostEndedEvent")}
	})
	return MTRWaterHeaterManagementClusterBoostEndedEventClass
}

type _MTRWaterHeaterManagementClusterBoostEndedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWaterHeaterManagementClusterBoostEndedEvent */
// An interface definition for the [MTRWaterHeaterManagementClusterBoostEndedEvent] class.
type IMTRWaterHeaterManagementClusterBoostEndedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWaterHeaterManagementClusterBoostEndedEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWaterHeaterManagementClusterBoostEndedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWaterHeaterManagementClusterBoostEndedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostEndedEventClass) Alloc() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWaterHeaterManagementClusterBoostEndedEventClass) New() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostEndedEvent) Init() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostEndedEvent) Autorelease() MTRWaterHeaterManagementClusterBoostEndedEvent {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostEndedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostEndedEvent creates a new MTRWaterHeaterManagementClusterBoostEndedEvent instance.
func NewMTRWaterHeaterManagementClusterBoostEndedEvent() MTRWaterHeaterManagementClusterBoostEndedEvent {
	return getMTRWaterHeaterManagementClusterBoostEndedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWaterHeaterManagementClusterBoostEndedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostEndedEvent
type MTRWaterHeaterManagementClusterBoostEndedEvent struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostEndedEventFrom constructs a [MTRWaterHeaterManagementClusterBoostEndedEvent] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostEndedEventFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostEndedEvent {
	return MTRWaterHeaterManagementClusterBoostEndedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWaterHeaterManagementClusterBoostEndedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWaterHeaterManagementClusterBoostEndedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWaterHeaterManagementClusterBoostEndedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWaterHeaterManagementClusterBoostEndedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWaterHeaterManagementClusterBoostEndedEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWaterHeaterManagementClusterBoostEndedEvent */



