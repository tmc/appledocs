// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass     _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass
	MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass() _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass {
	MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass = _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
	// properties:
	Cause() objc.IObject /* cross-framework: NSNumber */
	SetCause(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass) New() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) Init() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct creates a new MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	return getMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct
type MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	return MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) Cause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cause"))
	return rv
}/* debug [instance_properties/getter]: cause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) SetCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}/* debug [instance_properties/setter]: cause */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct */



