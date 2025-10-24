// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterPowerAdjustStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustStruct] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustStructClass     _MTRDeviceEnergyManagementClusterPowerAdjustStructClass
	MTRDeviceEnergyManagementClusterPowerAdjustStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustStructClass() _MTRDeviceEnergyManagementClusterPowerAdjustStructClass {
	MTRDeviceEnergyManagementClusterPowerAdjustStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustStructClass = _MTRDeviceEnergyManagementClusterPowerAdjustStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustStruct")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustStructClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustStruct] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
	// properties:
	MaxDuration() objc.IObject /* cross-framework: NSNumber */
	SetMaxDuration(value objc.IObject /* cross-framework: NSNumber */)
	MaxPower() objc.IObject /* cross-framework: NSNumber */
	SetMaxPower(value objc.IObject /* cross-framework: NSNumber */)
	MinDuration() objc.IObject /* cross-framework: NSNumber */
	SetMinDuration(value objc.IObject /* cross-framework: NSNumber */)
	MinPower() objc.IObject /* cross-framework: NSNumber */
	SetMinPower(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustStructClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustStructClass) New() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) Init() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustStruct creates a new MTRDeviceEnergyManagementClusterPowerAdjustStruct instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustStruct() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	return getMTRDeviceEnergyManagementClusterPowerAdjustStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterPowerAdjustStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct
type MTRDeviceEnergyManagementClusterPowerAdjustStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustStructFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	return MTRDeviceEnergyManagementClusterPowerAdjustStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterPowerAdjustStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterPowerAdjustStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterPowerAdjustStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxDuration"))
	return rv
}/* debug [instance_properties/getter]: maxDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}/* debug [instance_properties/setter]: maxDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/maxpower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPower"))
	return rv
}/* debug [instance_properties/getter]: maxPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/maxpower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}/* debug [instance_properties/setter]: maxPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/minduration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minDuration"))
	return rv
}/* debug [instance_properties/getter]: minDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/minduration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}/* debug [instance_properties/setter]: minDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/minpower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPower"))
	return rv
}/* debug [instance_properties/getter]: minPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterpoweradjuststruct/minpower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}/* debug [instance_properties/setter]: minPower */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterPowerAdjustStruct */



