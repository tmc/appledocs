// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBridgedDeviceBasicClusterLeaveEvent] class.
var (
	MTRBridgedDeviceBasicClusterLeaveEventClass     _MTRBridgedDeviceBasicClusterLeaveEventClass
	MTRBridgedDeviceBasicClusterLeaveEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicClusterLeaveEventClass() _MTRBridgedDeviceBasicClusterLeaveEventClass {
	MTRBridgedDeviceBasicClusterLeaveEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicClusterLeaveEventClass = _MTRBridgedDeviceBasicClusterLeaveEventClass{objc.GetClass("MTRBridgedDeviceBasicClusterLeaveEvent")}
	})
	return MTRBridgedDeviceBasicClusterLeaveEventClass
}

type _MTRBridgedDeviceBasicClusterLeaveEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicClusterLeaveEvent] class.
type IMTRBridgedDeviceBasicClusterLeaveEvent interface {
	IMTRBridgedDeviceBasicInformationClusterLeaveEvent
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicClusterLeaveEvent
type MTRBridgedDeviceBasicClusterLeaveEvent struct {
	MTRBridgedDeviceBasicInformationClusterLeaveEvent
}

// MTRBridgedDeviceBasicClusterLeaveEventFrom constructs a [MTRBridgedDeviceBasicClusterLeaveEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicClusterLeaveEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicClusterLeaveEvent {
	return MTRBridgedDeviceBasicClusterLeaveEvent{
		MTRBridgedDeviceBasicInformationClusterLeaveEvent: MTRBridgedDeviceBasicInformationClusterLeaveEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicClusterLeaveEventClass) Alloc() MTRBridgedDeviceBasicClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterLeaveEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicClusterLeaveEventClass) New() MTRBridgedDeviceBasicClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterLeaveEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicClusterLeaveEvent) Init() MTRBridgedDeviceBasicClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterLeaveEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicClusterLeaveEvent) Autorelease() MTRBridgedDeviceBasicClusterLeaveEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterLeaveEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicClusterLeaveEvent creates a new MTRBridgedDeviceBasicClusterLeaveEvent instance.
func NewMTRBridgedDeviceBasicClusterLeaveEvent() MTRBridgedDeviceBasicClusterLeaveEvent {
	return getMTRBridgedDeviceBasicClusterLeaveEventClass().New()
}




