// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
// The class instance for the [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] class.
var (
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass     _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass() _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass {
	MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass = _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass{objc.GetClass("MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent")}
	})
	return MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass
}

type _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
// An interface definition for the [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] class.
type IMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass) Alloc() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass) New() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent) Init() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent) Autorelease() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent creates a new MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent instance.
func NewMTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent() MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	return getMTRTimeSynchronizationClusterMissingTrustedTimeSourceEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent
type MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventFrom constructs a [MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterMissingTrustedTimeSourceEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent {
	return MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterMissingTrustedTimeSourceEvent */



