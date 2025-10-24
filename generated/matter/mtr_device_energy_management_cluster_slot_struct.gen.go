// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceEnergyManagementClusterSlotStruct */


/* debug [class_header]: Header for MTRDeviceEnergyManagementClusterSlotStruct */
// The class instance for the [MTRDeviceEnergyManagementClusterSlotStruct] class.
var (
	MTRDeviceEnergyManagementClusterSlotStructClass     _MTRDeviceEnergyManagementClusterSlotStructClass
	MTRDeviceEnergyManagementClusterSlotStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterSlotStructClass() _MTRDeviceEnergyManagementClusterSlotStructClass {
	MTRDeviceEnergyManagementClusterSlotStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterSlotStructClass = _MTRDeviceEnergyManagementClusterSlotStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterSlotStruct")}
	})
	return MTRDeviceEnergyManagementClusterSlotStructClass
}

type _MTRDeviceEnergyManagementClusterSlotStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceEnergyManagementClusterSlotStruct */
// An interface definition for the [MTRDeviceEnergyManagementClusterSlotStruct] class.
type IMTRDeviceEnergyManagementClusterSlotStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceEnergyManagementClusterSlotStruct */
	// properties:
	Costs() objc.IObject /* cross-framework: NSArray */
	SetCosts(value objc.IObject /* cross-framework: NSArray */)
	DefaultDuration() objc.IObject /* cross-framework: NSNumber */
	SetDefaultDuration(value objc.IObject /* cross-framework: NSNumber */)
	ElapsedSlotTime() objc.IObject /* cross-framework: NSNumber */
	SetElapsedSlotTime(value objc.IObject /* cross-framework: NSNumber */)
	ManufacturerESAState() objc.IObject /* cross-framework: NSNumber */
	SetManufacturerESAState(value objc.IObject /* cross-framework: NSNumber */)
	MaxDuration() objc.IObject /* cross-framework: NSNumber */
	SetMaxDuration(value objc.IObject /* cross-framework: NSNumber */)
	MaxDurationAdjustment() objc.IObject /* cross-framework: NSNumber */
	SetMaxDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */)
	MaxPauseDuration() objc.IObject /* cross-framework: NSNumber */
	SetMaxPauseDuration(value objc.IObject /* cross-framework: NSNumber */)
	MaxPower() objc.IObject /* cross-framework: NSNumber */
	SetMaxPower(value objc.IObject /* cross-framework: NSNumber */)
	MaxPowerAdjustment() objc.IObject /* cross-framework: NSNumber */
	SetMaxPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */)
	MinDuration() objc.IObject /* cross-framework: NSNumber */
	SetMinDuration(value objc.IObject /* cross-framework: NSNumber */)
	MinDurationAdjustment() objc.IObject /* cross-framework: NSNumber */
	SetMinDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */)
	MinPauseDuration() objc.IObject /* cross-framework: NSNumber */
	SetMinPauseDuration(value objc.IObject /* cross-framework: NSNumber */)
	MinPower() objc.IObject /* cross-framework: NSNumber */
	SetMinPower(value objc.IObject /* cross-framework: NSNumber */)
	MinPowerAdjustment() objc.IObject /* cross-framework: NSNumber */
	SetMinPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */)
	NominalEnergy() objc.IObject /* cross-framework: NSNumber */
	SetNominalEnergy(value objc.IObject /* cross-framework: NSNumber */)
	NominalPower() objc.IObject /* cross-framework: NSNumber */
	SetNominalPower(value objc.IObject /* cross-framework: NSNumber */)
	RemainingSlotTime() objc.IObject /* cross-framework: NSNumber */
	SetRemainingSlotTime(value objc.IObject /* cross-framework: NSNumber */)
	SlotIsPausable() objc.IObject /* cross-framework: NSNumber */
	SetSlotIsPausable(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceEnergyManagementClusterSlotStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceEnergyManagementClusterSlotStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterSlotStructClass) Alloc() MTRDeviceEnergyManagementClusterSlotStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceEnergyManagementClusterSlotStructClass) New() MTRDeviceEnergyManagementClusterSlotStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) Init() MTRDeviceEnergyManagementClusterSlotStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) Autorelease() MTRDeviceEnergyManagementClusterSlotStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterSlotStruct creates a new MTRDeviceEnergyManagementClusterSlotStruct instance.
func NewMTRDeviceEnergyManagementClusterSlotStruct() MTRDeviceEnergyManagementClusterSlotStruct {
	return getMTRDeviceEnergyManagementClusterSlotStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceEnergyManagementClusterSlotStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct
type MTRDeviceEnergyManagementClusterSlotStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterSlotStructFrom constructs a [MTRDeviceEnergyManagementClusterSlotStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterSlotStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterSlotStruct {
	return MTRDeviceEnergyManagementClusterSlotStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceEnergyManagementClusterSlotStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceEnergyManagementClusterSlotStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceEnergyManagementClusterSlotStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceEnergyManagementClusterSlotStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceEnergyManagementClusterSlotStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) Costs() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("costs"))
	return rv
}/* debug [instance_properties/getter]: costs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetCosts(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCosts:"), value)
}/* debug [instance_properties/setter]: costs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/defaultduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) DefaultDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("defaultDuration"))
	return rv
}/* debug [instance_properties/getter]: defaultDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/defaultduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetDefaultDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultDuration:"), value)
}/* debug [instance_properties/setter]: defaultDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/elapsedslottime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ElapsedSlotTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("elapsedSlotTime"))
	return rv
}/* debug [instance_properties/getter]: elapsedSlotTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/elapsedslottime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetElapsedSlotTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElapsedSlotTime:"), value)
}/* debug [instance_properties/setter]: elapsedSlotTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/manufactureresastate
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ManufacturerESAState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("manufacturerESAState"))
	return rv
}/* debug [instance_properties/getter]: manufacturerESAState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/manufactureresastate
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetManufacturerESAState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setManufacturerESAState:"), value)
}/* debug [instance_properties/setter]: manufacturerESAState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxDuration"))
	return rv
}/* debug [instance_properties/getter]: maxDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}/* debug [instance_properties/setter]: maxDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxdurationadjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDurationAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxDurationAdjustment"))
	return rv
}/* debug [instance_properties/getter]: maxDurationAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxdurationadjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDurationAdjustment:"), value)
}/* debug [instance_properties/setter]: maxDurationAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpauseduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPauseDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPauseDuration"))
	return rv
}/* debug [instance_properties/getter]: maxPauseDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpauseduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPauseDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPauseDuration:"), value)
}/* debug [instance_properties/setter]: maxPauseDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPower"))
	return rv
}/* debug [instance_properties/getter]: maxPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}/* debug [instance_properties/setter]: maxPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpoweradjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPowerAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPowerAdjustment"))
	return rv
}/* debug [instance_properties/getter]: maxPowerAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/maxpoweradjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPowerAdjustment:"), value)
}/* debug [instance_properties/setter]: maxPowerAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minDuration"))
	return rv
}/* debug [instance_properties/getter]: minDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}/* debug [instance_properties/setter]: minDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/mindurationadjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDurationAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minDurationAdjustment"))
	return rv
}/* debug [instance_properties/getter]: minDurationAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/mindurationadjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDurationAdjustment:"), value)
}/* debug [instance_properties/setter]: minDurationAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpauseduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPauseDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPauseDuration"))
	return rv
}/* debug [instance_properties/getter]: minPauseDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpauseduration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPauseDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPauseDuration:"), value)
}/* debug [instance_properties/setter]: minPauseDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPower"))
	return rv
}/* debug [instance_properties/getter]: minPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}/* debug [instance_properties/setter]: minPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpoweradjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPowerAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPowerAdjustment"))
	return rv
}/* debug [instance_properties/getter]: minPowerAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/minpoweradjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPowerAdjustment:"), value)
}/* debug [instance_properties/setter]: minPowerAdjustment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/nominalenergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalEnergy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalEnergy"))
	return rv
}/* debug [instance_properties/getter]: nominalEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/nominalenergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalEnergy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalEnergy:"), value)
}/* debug [instance_properties/setter]: nominalEnergy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalPower"))
	return rv
}/* debug [instance_properties/getter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/nominalpower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}/* debug [instance_properties/setter]: nominalPower */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/remainingslottime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) RemainingSlotTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("remainingSlotTime"))
	return rv
}/* debug [instance_properties/getter]: remainingSlotTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/remainingslottime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetRemainingSlotTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRemainingSlotTime:"), value)
}/* debug [instance_properties/setter]: remainingSlotTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/slotispausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SlotIsPausable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("slotIsPausable"))
	return rv
}/* debug [instance_properties/getter]: slotIsPausable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdeviceenergymanagementclusterslotstruct/slotispausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetSlotIsPausable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotIsPausable:"), value)
}/* debug [instance_properties/setter]: slotIsPausable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceEnergyManagementClusterSlotStruct */



