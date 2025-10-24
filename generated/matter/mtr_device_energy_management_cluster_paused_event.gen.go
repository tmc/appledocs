// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPausedEvent */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPausedEvent */
// The class instance for the [MTRDeviceEnergyManagementClusterPausedEvent] class.
var (
	MTRDeviceEnergyManagementClusterPausedEventClass     _MTRDeviceEnergyManagementClusterPausedEventClass
	MTRDeviceEnergyManagementClusterPausedEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPausedEventClass() _MTRDeviceEnergyManagementClusterPausedEventClass {
	MTRDeviceEnergyManagementClusterPausedEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPausedEventClass = _MTRDeviceEnergyManagementClusterPausedEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterPausedEvent")}
	})
	return MTRDeviceEnergyManagementClusterPausedEventClass
}

type _MTRDeviceEnergyManagementClusterPausedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPausedEvent */
// An interface definition for the [MTRDeviceEnergyManagementClusterPausedEvent] class.
type IMTRDeviceEnergyManagementClusterPausedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPausedEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPausedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPausedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPausedEventClass) Alloc() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPausedEventClass) New() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPausedEvent) Init() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPausedEvent) Autorelease() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPausedEvent creates a new MTRDeviceEnergyManagementClusterPausedEvent instance.
func NewMTRDeviceEnergyManagementClusterPausedEvent() MTRDeviceEnergyManagementClusterPausedEvent {
	return getMTRDeviceEnergyManagementClusterPausedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPausedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPausedEvent
type MTRDeviceEnergyManagementClusterPausedEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPausedEventFrom constructs a [MTRDeviceEnergyManagementClusterPausedEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPausedEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPausedEvent {
	return MTRDeviceEnergyManagementClusterPausedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPausedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPausedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPausedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPausedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPausedEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPausedEvent */



