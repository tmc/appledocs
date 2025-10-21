// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxDuration"))
	return rv
}


// SetMaxDuration sets the value of the maxDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MaxPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maxPower"))
	return rv
}


// SetMaxPower sets the value of the maxPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/maxPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMaxPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxPower:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minDuration"))
	return rv
}


// SetMinDuration sets the value of the minDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minDuration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) MinPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("minPower"))
	return rv
}


// SetMinPower sets the value of the minPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustStruct/minPower
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustStruct) SetMinPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinPower:"), value)
}


