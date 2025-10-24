// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterSlotAdjustmentStruct] class.
var (
	MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass     _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass
	MTRDeviceEnergyManagementClusterSlotAdjustmentStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterSlotAdjustmentStructClass() _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass {
	MTRDeviceEnergyManagementClusterSlotAdjustmentStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass = _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterSlotAdjustmentStruct")}
	})
	return MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass
}

type _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterSlotAdjustmentStruct] class.
type IMTRDeviceEnergyManagementClusterSlotAdjustmentStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	NominalPower() objc.IObject /* cross-framework: NSNumber */
	SetNominalPower(value objc.IObject /* cross-framework: NSNumber */)
	SlotIndex() objc.IObject /* cross-framework: NSNumber */
	SetSlotIndex(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass) Alloc() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotAdjustmentStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass) New() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotAdjustmentStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) Init() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotAdjustmentStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) Autorelease() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotAdjustmentStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterSlotAdjustmentStruct creates a new MTRDeviceEnergyManagementClusterSlotAdjustmentStruct instance.
func NewMTRDeviceEnergyManagementClusterSlotAdjustmentStruct() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	return getMTRDeviceEnergyManagementClusterSlotAdjustmentStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct
type MTRDeviceEnergyManagementClusterSlotAdjustmentStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterSlotAdjustmentStructFrom constructs a [MTRDeviceEnergyManagementClusterSlotAdjustmentStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterSlotAdjustmentStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	return MTRDeviceEnergyManagementClusterSlotAdjustmentStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/duration
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/duration
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotadjustmentstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) NominalPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalPower"))
	return rv
}/* debug [instance_properties/getter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotadjustmentstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetNominalPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}/* debug [instance_properties/setter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotadjustmentstruct/slotindex
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SlotIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("slotIndex"))
	return rv
}/* debug [instance_properties/getter]: slotIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotadjustmentstruct/slotindex
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetSlotIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotIndex:"), value)
}/* debug [instance_properties/setter]: slotIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterSlotAdjustmentStruct */



