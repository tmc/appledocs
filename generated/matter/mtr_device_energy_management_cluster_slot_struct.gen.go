// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDeviceEnergyManagementClusterSlotStruct] class.
type IMTRDeviceEnergyManagementClusterSlotStruct interface {
	objectivec.IObject
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
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct
type MTRDeviceEnergyManagementClusterSlotStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterSlotStructFrom constructs a [MTRDeviceEnergyManagementClusterSlotStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterSlotStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterSlotStruct {
	return MTRDeviceEnergyManagementClusterSlotStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterSlotStructClass) Alloc() MTRDeviceEnergyManagementClusterSlotStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) Costs() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("costs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetCosts(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCosts:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/defaultDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) DefaultDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("defaultDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/defaultDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetDefaultDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/elapsedSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ElapsedSlotTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("elapsedSlotTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/elapsedSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetElapsedSlotTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElapsedSlotTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/manufacturerESAState
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ManufacturerESAState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("manufacturerESAState"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/manufacturerESAState
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetManufacturerESAState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setManufacturerESAState:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDurationAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxDurationAdjustment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDurationAdjustment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPauseDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPauseDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPauseDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPauseDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPower"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPowerAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("maxPowerAdjustment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPowerAdjustment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDurationAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minDurationAdjustment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDurationAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDurationAdjustment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPauseDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPauseDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPauseDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPauseDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPower"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPowerAdjustment() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minPowerAdjustment"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPowerAdjustment(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPowerAdjustment:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalEnergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalEnergy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalEnergy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalEnergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalEnergy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalEnergy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalPower"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/remainingSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) RemainingSlotTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("remainingSlotTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/remainingSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetRemainingSlotTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRemainingSlotTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/slotIsPausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SlotIsPausable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("slotIsPausable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/slotIsPausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetSlotIsPausable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotIsPausable:"), value)
}



