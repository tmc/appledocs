// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) Costs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("costs"))
	return rv
}


// SetCosts sets the value of the costs property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/costs
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetCosts(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCosts:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/defaultDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) DefaultDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("defaultDuration"))
	return rv
}


// SetDefaultDuration sets the value of the defaultDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/defaultDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetDefaultDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/elapsedSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ElapsedSlotTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("elapsedSlotTime"))
	return rv
}


// SetElapsedSlotTime sets the value of the elapsedSlotTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/elapsedSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetElapsedSlotTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setElapsedSlotTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/manufacturerESAState
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) ManufacturerESAState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("manufacturerESAState"))
	return rv
}


// SetManufacturerESAState sets the value of the manufacturerESAState property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/manufacturerESAState
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetManufacturerESAState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setManufacturerESAState:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxDuration"))
	return rv
}


// SetMaxDuration sets the value of the maxDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxDurationAdjustment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxDurationAdjustment"))
	return rv
}


// SetMaxDurationAdjustment sets the value of the maxDurationAdjustment property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxDurationAdjustment(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDurationAdjustment:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPauseDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxPauseDuration"))
	return rv
}


// SetMaxPauseDuration sets the value of the maxPauseDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPauseDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPauseDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxPower"))
	return rv
}


// SetMaxPower sets the value of the maxPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MaxPowerAdjustment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxPowerAdjustment"))
	return rv
}


// SetMaxPowerAdjustment sets the value of the maxPowerAdjustment property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/maxPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMaxPowerAdjustment(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPowerAdjustment:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minDuration"))
	return rv
}


// SetMinDuration sets the value of the minDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinDurationAdjustment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minDurationAdjustment"))
	return rv
}


// SetMinDurationAdjustment sets the value of the minDurationAdjustment property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minDurationAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinDurationAdjustment(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDurationAdjustment:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPauseDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minPauseDuration"))
	return rv
}


// SetMinPauseDuration sets the value of the minPauseDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPauseDuration
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPauseDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPauseDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minPower"))
	return rv
}


// SetMinPower sets the value of the minPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) MinPowerAdjustment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minPowerAdjustment"))
	return rv
}


// SetMinPowerAdjustment sets the value of the minPowerAdjustment property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/minPowerAdjustment
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetMinPowerAdjustment(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPowerAdjustment:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalEnergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalEnergy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalEnergy"))
	return rv
}


// SetNominalEnergy sets the value of the nominalEnergy property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalEnergy
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalEnergy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalEnergy:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) NominalPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalPower"))
	return rv
}


// SetNominalPower sets the value of the nominalPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetNominalPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/remainingSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) RemainingSlotTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("remainingSlotTime"))
	return rv
}


// SetRemainingSlotTime sets the value of the remainingSlotTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/remainingSlotTime
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetRemainingSlotTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRemainingSlotTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/slotIsPausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SlotIsPausable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("slotIsPausable"))
	return rv
}


// SetSlotIsPausable sets the value of the slotIsPausable property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotStruct/slotIsPausable
func (m_ MTRDeviceEnergyManagementClusterSlotStruct) SetSlotIsPausable(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotIsPausable:"), value)
}


