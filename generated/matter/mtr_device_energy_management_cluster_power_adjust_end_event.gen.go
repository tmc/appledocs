// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] class.
var (
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass     _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPowerAdjustEndEventClass() _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass {
	MTRDeviceEnergyManagementClusterPowerAdjustEndEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass = _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterPowerAdjustEndEvent")}
	})
	return MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass
}

type _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] class.
type IMTRDeviceEnergyManagementClusterPowerAdjustEndEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent
type MTRDeviceEnergyManagementClusterPowerAdjustEndEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPowerAdjustEndEventFrom constructs a [MTRDeviceEnergyManagementClusterPowerAdjustEndEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPowerAdjustEndEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	return MTRDeviceEnergyManagementClusterPowerAdjustEndEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass) Alloc() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterPowerAdjustEndEventClass) New() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Init() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Autorelease() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPowerAdjustEndEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPowerAdjustEndEvent creates a new MTRDeviceEnergyManagementClusterPowerAdjustEndEvent instance.
func NewMTRDeviceEnergyManagementClusterPowerAdjustEndEvent() MTRDeviceEnergyManagementClusterPowerAdjustEndEvent {
	return getMTRDeviceEnergyManagementClusterPowerAdjustEndEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Cause() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cause"))
	return rv
}


// SetCause sets the value of the cause property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/cause
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetCause(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/duration
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/energyUse
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) EnergyUse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("energyUse"))
	return rv
}


// SetEnergyUse sets the value of the energyUse property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPowerAdjustEndEvent/energyUse
func (m_ MTRDeviceEnergyManagementClusterPowerAdjustEndEvent) SetEnergyUse(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnergyUse:"), value)
}


