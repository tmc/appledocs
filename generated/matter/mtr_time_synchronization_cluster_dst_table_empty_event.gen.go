// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterDSTTableEmptyEvent */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
// The class instance for the [MTRTimeSynchronizationClusterDSTTableEmptyEvent] class.
var (
	MTRTimeSynchronizationClusterDSTTableEmptyEventClass     _MTRTimeSynchronizationClusterDSTTableEmptyEventClass
	MTRTimeSynchronizationClusterDSTTableEmptyEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTTableEmptyEventClass() _MTRTimeSynchronizationClusterDSTTableEmptyEventClass {
	MTRTimeSynchronizationClusterDSTTableEmptyEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTTableEmptyEventClass = _MTRTimeSynchronizationClusterDSTTableEmptyEventClass{objc.GetClass("MTRTimeSynchronizationClusterDSTTableEmptyEvent")}
	})
	return MTRTimeSynchronizationClusterDSTTableEmptyEventClass
}

type _MTRTimeSynchronizationClusterDSTTableEmptyEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
// An interface definition for the [MTRTimeSynchronizationClusterDSTTableEmptyEvent] class.
type IMTRTimeSynchronizationClusterDSTTableEmptyEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTTableEmptyEventClass) Alloc() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterDSTTableEmptyEventClass) New() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTTableEmptyEvent) Init() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTTableEmptyEvent) Autorelease() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTTableEmptyEvent creates a new MTRTimeSynchronizationClusterDSTTableEmptyEvent instance.
func NewMTRTimeSynchronizationClusterDSTTableEmptyEvent() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	return getMTRTimeSynchronizationClusterDSTTableEmptyEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterDSTTableEmptyEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTTableEmptyEvent
type MTRTimeSynchronizationClusterDSTTableEmptyEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTTableEmptyEventFrom constructs a [MTRTimeSynchronizationClusterDSTTableEmptyEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTTableEmptyEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	return MTRTimeSynchronizationClusterDSTTableEmptyEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterDSTTableEmptyEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterDSTTableEmptyEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterDSTTableEmptyEvent */



