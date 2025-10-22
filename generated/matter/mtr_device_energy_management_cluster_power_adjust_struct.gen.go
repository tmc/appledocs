// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustStruct] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustStruct interface {
	objectivec.IObject
	MaxDuration() foundation.Number
	SetMaxDuration(value foundation.INumber)
	MaxPower() foundation.Number
	SetMaxPower(value foundation.INumber)
	MinDuration() foundation.Number
	SetMinDuration(value foundation.INumber)
	MinPower() foundation.Number
	SetMinPower(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct
type MTRDeviceEnergyManagementClusterPowerAdjustStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustStructFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	return MTRDeviceEnergyManagementClusterPowerAdjustStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustStructClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxDuration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxDuration"))
	return rv
}


// SetMaxDuration sets the value of the maxDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxPower() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("maxPower"))
	return rv
}


// SetMaxPower sets the value of the maxPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxPower(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinDuration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minDuration"))
	return rv
}


// SetMinDuration sets the value of the minDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinPower() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minPower"))
	return rv
}


// SetMinPower sets the value of the minPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinPower(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}



