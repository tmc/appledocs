// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRDeviceEnergyManagementClusterSlotAdjustmentStruct] class.
type IMTRDeviceEnergyManagementClusterSlotAdjustmentStruct interface {
	objectivec.IObject
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	NominalPower() objc.IObject /* cross-framework: NSNumber */
	SetNominalPower(value objc.IObject /* cross-framework: NSNumber */)
	SlotIndex() objc.IObject /* cross-framework: NSNumber */
	SetSlotIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct
type MTRDeviceEnergyManagementClusterSlotAdjustmentStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterSlotAdjustmentStructFrom constructs a [MTRDeviceEnergyManagementClusterSlotAdjustmentStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterSlotAdjustmentStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	return MTRDeviceEnergyManagementClusterSlotAdjustmentStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterSlotAdjustmentStructClass) Alloc() MTRDeviceEnergyManagementClusterSlotAdjustmentStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterSlotAdjustmentStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/duration
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/duration
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) NominalPower() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nominalPower"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetNominalPower(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/slotIndex
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SlotIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("slotIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterSlotAdjustmentStruct/slotIndex
func (m_ MTRDeviceEnergyManagementClusterSlotAdjustmentStruct) SetSlotIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlotIndex:"), value)
}



