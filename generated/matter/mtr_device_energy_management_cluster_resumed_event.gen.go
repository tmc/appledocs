// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterResumedEvent] class.
var (
	MTRDeviceEnergyManagementClusterResumedEventClass     _MTRDeviceEnergyManagementClusterResumedEventClass
	MTRDeviceEnergyManagementClusterResumedEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterResumedEventClass() _MTRDeviceEnergyManagementClusterResumedEventClass {
	MTRDeviceEnergyManagementClusterResumedEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterResumedEventClass = _MTRDeviceEnergyManagementClusterResumedEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterResumedEvent")}
	})
	return MTRDeviceEnergyManagementClusterResumedEventClass
}

type _MTRDeviceEnergyManagementClusterResumedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterResumedEvent] class.
type IMTRDeviceEnergyManagementClusterResumedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent
type MTRDeviceEnergyManagementClusterResumedEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterResumedEventFrom constructs a [MTRDeviceEnergyManagementClusterResumedEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterResumedEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterResumedEvent {
	return MTRDeviceEnergyManagementClusterResumedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterResumedEventClass) Alloc() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterResumedEventClass) New() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Init() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Autorelease() MTRDeviceEnergyManagementClusterResumedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterResumedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterResumedEvent creates a new MTRDeviceEnergyManagementClusterResumedEvent instance.
func NewMTRDeviceEnergyManagementClusterResumedEvent() MTRDeviceEnergyManagementClusterResumedEvent {
	return getMTRDeviceEnergyManagementClusterResumedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent/cause
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) Cause() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("cause"))
	return rv
}


// SetCause sets the value of the cause property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterResumedEvent/cause
func (m_ MTRDeviceEnergyManagementClusterResumedEvent) SetCause(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCause:"), value)
}



