// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass     _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustEndEventClass() _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass {
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass = _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustEndEvent")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustEndEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	EnergyUse() objc.IObject /* cross-framework: NSNumber */
	SetEnergyUse(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass) New() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Init() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustEndEvent creates a new MTRDeviceEnergyManagementClusterPowerAdjustEndEvent instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustEndEvent() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	return getMTRDeviceEnergyManagementClusterPowerAdjustEndEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent
type MTRDeviceEnergyManagementClusterPowerAdjustEndEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustEndEventFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustEndEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	return MTRDeviceEnergyManagementClusterPowerAdjustEndEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustendevent/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustendevent/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustendevent/energyuse
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) EnergyUse() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("energyUse"))
	return rv
}/* debug [instance_properties/getter]: energyUse */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjustendevent/energyuse
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetEnergyUse(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyUse:"), value)
}/* debug [instance_properties/setter]: energyUse */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPowerAdjustEndEvent */



