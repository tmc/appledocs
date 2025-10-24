// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterPausedEvent] class.
var (
	MTRDeviceEnergyManagementClusterPausedEventClass     _MTRDeviceEnergyManagementClusterPausedEventClass
	MTRDeviceEnergyManagementClusterPausedEventClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterPausedEventClass() _MTRDeviceEnergyManagementClusterPausedEventClass {
	MTRDeviceEnergyManagementClusterPausedEventClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterPausedEventClass = _MTRDeviceEnergyManagementClusterPausedEventClass{objc.GetClass("MTRDeviceEnergyManagementClusterPausedEvent")}
	})
	return MTRDeviceEnergyManagementClusterPausedEventClass
}

type _MTRDeviceEnergyManagementClusterPausedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterPausedEvent] class.
type IMTRDeviceEnergyManagementClusterPausedEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterPausedEvent
type MTRDeviceEnergyManagementClusterPausedEvent struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterPausedEventFrom constructs a [MTRDeviceEnergyManagementClusterPausedEvent] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterPausedEventFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterPausedEvent {
	return MTRDeviceEnergyManagementClusterPausedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterPausedEventClass) Alloc() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterPausedEventClass) New() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterPausedEvent) Init() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterPausedEvent) Autorelease() MTRDeviceEnergyManagementClusterPausedEvent {
	rv := objc.Send[MTRDeviceEnergyManagementClusterPausedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterPausedEvent creates a new MTRDeviceEnergyManagementClusterPausedEvent instance.
func NewMTRDeviceEnergyManagementClusterPausedEvent() MTRDeviceEnergyManagementClusterPausedEvent {
	return getMTRDeviceEnergyManagementClusterPausedEventClass().New()
}




