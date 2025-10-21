// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct
type MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	return MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStructClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) Cause() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cause"))
	return rv
}


// SetCause sets the value of the cause property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) SetCause(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/powerAdjustCapability
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) PowerAdjustCapability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("powerAdjustCapability"))
	return rv
}


// SetPowerAdjustCapability sets the value of the powerAdjustCapability property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct/powerAdjustCapability
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustCapabilityStruct) SetPowerAdjustCapability(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPowerAdjustCapability:"), value)
}


