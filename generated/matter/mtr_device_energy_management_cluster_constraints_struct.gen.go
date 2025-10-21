// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterConstraintsStruct] class.
var (
	MTRDeviceEnergyManagementClusterConstraintsStructClass     _MTRDeviceEnergyManagementClusterConstraintsStructClass
	MTRDeviceEnergyManagementClusterConstraintsStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterConstraintsStructClass() _MTRDeviceEnergyManagementClusterConstraintsStructClass {
	MTRDeviceEnergyManagementClusterConstraintsStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterConstraintsStructClass = _MTRDeviceEnergyManagementClusterConstraintsStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterConstraintsStruct")}
	})
	return MTRDeviceEnergyManagementClusterConstraintsStructClass
}

type _MTRDeviceEnergyManagementClusterConstraintsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterConstraintsStruct] class.
type IMTRDeviceEnergyManagementClusterConstraintsStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct
type MTRDeviceEnergyManagementClusterConstraintsStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterConstraintsStructFrom constructs a [MTRDeviceEnergyManagementClusterConstraintsStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterConstraintsStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterConstraintsStruct {
	return MTRDeviceEnergyManagementClusterConstraintsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterConstraintsStructClass) Alloc() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterConstraintsStructClass) New() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Init() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Autorelease() MTRDeviceEnergyManagementClusterConstraintsStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterConstraintsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterConstraintsStruct creates a new MTRDeviceEnergyManagementClusterConstraintsStruct instance.
func NewMTRDeviceEnergyManagementClusterConstraintsStruct() MTRDeviceEnergyManagementClusterConstraintsStruct {
	return getMTRDeviceEnergyManagementClusterConstraintsStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/duration
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/duration
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/loadControl
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) LoadControl() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("loadControl"))
	return rv
}


// SetLoadControl sets the value of the loadControl property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/loadControl
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetLoadControl(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLoadControl:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/maximumEnergy
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) MaximumEnergy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("maximumEnergy"))
	return rv
}


// SetMaximumEnergy sets the value of the maximumEnergy property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/maximumEnergy
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetMaximumEnergy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumEnergy:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) NominalPower() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nominalPower"))
	return rv
}


// SetNominalPower sets the value of the nominalPower property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/nominalPower
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetNominalPower(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNominalPower:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) StartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterConstraintsStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterConstraintsStruct) SetStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}


