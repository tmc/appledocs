// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustStartEvent] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass     _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass
	MTRDeviceEnergyManagementClusterPowerAdjustStartEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustStartEventClass() _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass {
	MTRDeviceEnergyManagementClusterPowerAdjustStartEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass = _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustStartEvent")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustStartEvent] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustStartEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStartEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustStartEventClass) New() MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStartEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStartEvent) Init() MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStartEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStartEvent) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStartEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustStartEvent creates a new MTRDeviceEnergyManagementClusterPowerAdjustStartEvent instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustStartEvent() MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	return getMTRDeviceEnergyManagementClusterPowerAdjustStartEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStartEvent
type MTRDeviceEnergyManagementClusterPowerAdjustStartEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustStartEventFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustStartEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustStartEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustStartEvent {
	return MTRDeviceEnergyManagementClusterPowerAdjustStartEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPowerAdjustStartEvent */



