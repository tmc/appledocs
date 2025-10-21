// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementModeClusterModeTagStruct] class.
var (
	MTRDeviceEnergyManagementModeClusterModeTagStructClass     _MTRDeviceEnergyManagementModeClusterModeTagStructClass
	MTRDeviceEnergyManagementModeClusterModeTagStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterModeTagStructClass() _MTRDeviceEnergyManagementModeClusterModeTagStructClass {
	MTRDeviceEnergyManagementModeClusterModeTagStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterModeTagStructClass = _MTRDeviceEnergyManagementModeClusterModeTagStructClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterModeTagStruct")}
	})
	return MTRDeviceEnergyManagementModeClusterModeTagStructClass
}

type _MTRDeviceEnergyManagementModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementModeClusterModeTagStruct] class.
type IMTRDeviceEnergyManagementModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct
type MTRDeviceEnergyManagementModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterModeTagStructFrom constructs a [MTRDeviceEnergyManagementModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterModeTagStruct {
	return MTRDeviceEnergyManagementModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterModeTagStructClass) Alloc() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementModeClusterModeTagStructClass) New() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Init() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Autorelease() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterModeTagStruct creates a new MTRDeviceEnergyManagementModeClusterModeTagStruct instance.
func NewMTRDeviceEnergyManagementModeClusterModeTagStruct() MTRDeviceEnergyManagementModeClusterModeTagStruct {
	return getMTRDeviceEnergyManagementModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/mfgCode
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) MfgCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/mfgCode
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) SetMfgCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/value
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeTagStruct/value
func (m_ MTRDeviceEnergyManagementModeClusterModeTagStruct) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



