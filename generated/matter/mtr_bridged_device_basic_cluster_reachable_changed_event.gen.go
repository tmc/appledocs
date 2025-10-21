// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBridgedDeviceBasicClusterReachableChangedEvent] class.
var (
	MTRBridgedDeviceBasicClusterReachableChangedEventClass     _MTRBridgedDeviceBasicClusterReachableChangedEventClass
	MTRBridgedDeviceBasicClusterReachableChangedEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicClusterReachableChangedEventClass() _MTRBridgedDeviceBasicClusterReachableChangedEventClass {
	MTRBridgedDeviceBasicClusterReachableChangedEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicClusterReachableChangedEventClass = _MTRBridgedDeviceBasicClusterReachableChangedEventClass{objc.GetClass("MTRBridgedDeviceBasicClusterReachableChangedEvent")}
	})
	return MTRBridgedDeviceBasicClusterReachableChangedEventClass
}

type _MTRBridgedDeviceBasicClusterReachableChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicClusterReachableChangedEvent] class.
type IMTRBridgedDeviceBasicClusterReachableChangedEvent interface {
	IMTRBridgedDeviceBasicInformationClusterReachableChangedEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicClusterReachableChangedEvent
type MTRBridgedDeviceBasicClusterReachableChangedEvent struct {
	MTRBridgedDeviceBasicInformationClusterReachableChangedEvent
}

// MTRBridgedDeviceBasicClusterReachableChangedEventFrom constructs a [MTRBridgedDeviceBasicClusterReachableChangedEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicClusterReachableChangedEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicClusterReachableChangedEvent {
	return MTRBridgedDeviceBasicClusterReachableChangedEvent{
		MTRBridgedDeviceBasicInformationClusterReachableChangedEvent: MTRBridgedDeviceBasicInformationClusterReachableChangedEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicClusterReachableChangedEventClass) Alloc() MTRBridgedDeviceBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicClusterReachableChangedEventClass) New() MTRBridgedDeviceBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicClusterReachableChangedEvent) Init() MTRBridgedDeviceBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterReachableChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicClusterReachableChangedEvent) Autorelease() MTRBridgedDeviceBasicClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterReachableChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicClusterReachableChangedEvent creates a new MTRBridgedDeviceBasicClusterReachableChangedEvent instance.
func NewMTRBridgedDeviceBasicClusterReachableChangedEvent() MTRBridgedDeviceBasicClusterReachableChangedEvent {
	return getMTRBridgedDeviceBasicClusterReachableChangedEventClass().New()
}




