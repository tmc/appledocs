// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterResumedEvent */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterResumedEvent */
// The class instance for the [MTRDeviceEnergyManagementClusterResumedEvent] class.
var (
	MTRDeviceEnergyManagementClusterResumedEventClass     _MTRDeviceEnergyManagementClusterResumedEventClass
	MTRDeviceEnergyManagementClusterResumedEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterResumedEventClass() _MTRDeviceEnergyManagementClusterResumedEventClass {
	MTRDeviceEnergyManagementClusterResumedEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterResumedEventClass = _MTRDeviceEnergyManagementClusterResumedEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterResumedEvent")}
	})
	return MTRDeviceEnergyManagementClusterResumedEventClass
}

type _MTRDeviceEnergyManagementClusterResumedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterResumedEvent */
// An interface definition for the [MTRDeviceEnergyManagementClusterResumedEvent] class.
type IMTRDeviceEnergyManagementClusterResumedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterResumedEvent */
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterResumedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterResumedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterResumedEventClass) Alloc() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterResumedEventClass) New() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Init() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Autorelease() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterResumedEvent creates a new MTRDeviceEnergyManagementClusterResumedEvent instance.
func NewMTRDeviceEnergyManagementClusterResumedEvent() MTRDeviceEnergyManagementClusterResumedEvent {
	return getMTRDeviceEnergyManagementClusterResumedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterResumedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent
type MTRDeviceEnergyManagementClusterResumedEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterResumedEventFrom constructs a [MTRDeviceEnergyManagementClusterResumedEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterResumedEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterResumedEvent {
	return MTRDeviceEnergyManagementClusterResumedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterResumedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterResumedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterResumedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterResumedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterResumedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent/cause
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent/cause
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterResumedEvent */



