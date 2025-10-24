// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterDSTStatusEvent */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterDSTStatusEvent */
// The class instance for the [MTRTimeSynchronizationClusterDSTStatusEvent] class.
var (
	MTRTimeSynchronizationClusterDSTStatusEventClass     _MTRTimeSynchronizationClusterDSTStatusEventClass
	MTRTimeSynchronizationClusterDSTStatusEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTStatusEventClass() _MTRTimeSynchronizationClusterDSTStatusEventClass {
	MTRTimeSynchronizationClusterDSTStatusEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTStatusEventClass = _MTRTimeSynchronizationClusterDSTStatusEventClass{objc.GetClass("MTRTimeSynchronizationClusterDSTStatusEvent")}
	})
	return MTRTimeSynchronizationClusterDSTStatusEventClass
}

type _MTRTimeSynchronizationClusterDSTStatusEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterDSTStatusEvent */
// An interface definition for the [MTRTimeSynchronizationClusterDSTStatusEvent] class.
type IMTRTimeSynchronizationClusterDSTStatusEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterDSTStatusEvent */
	// properties:
	DstOffsetActive() objc.IObject /* cross-framework: NSNumber */
	SetDstOffsetActive(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterDSTStatusEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterDSTStatusEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTStatusEventClass) Alloc() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterDSTStatusEventClass) New() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) Init() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) Autorelease() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTStatusEvent creates a new MTRTimeSynchronizationClusterDSTStatusEvent instance.
func NewMTRTimeSynchronizationClusterDSTStatusEvent() MTRTimeSynchronizationClusterDSTStatusEvent {
	return getMTRTimeSynchronizationClusterDSTStatusEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterDSTStatusEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent
type MTRTimeSynchronizationClusterDSTStatusEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTStatusEventFrom constructs a [MTRTimeSynchronizationClusterDSTStatusEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTStatusEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTStatusEvent {
	return MTRTimeSynchronizationClusterDSTStatusEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterDSTStatusEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterDSTStatusEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterDSTStatusEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterDSTStatusEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterDSTStatusEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent/dstOffsetActive
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) DstOffsetActive() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dstOffsetActive"))
	return rv
}/* debug [instance_properties/getter]: dstOffsetActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent/dstOffsetActive
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) SetDstOffsetActive(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffsetActive:"), value)
}/* debug [instance_properties/setter]: dstOffsetActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterDSTStatusEvent */



