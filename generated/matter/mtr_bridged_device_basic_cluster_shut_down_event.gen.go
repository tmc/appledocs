// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBridgedDeviceBasicClusterShutDownEvent] class.
var (
	MTRBridgedDeviceBasicClusterShutDownEventClass     _MTRBridgedDeviceBasicClusterShutDownEventClass
	MTRBridgedDeviceBasicClusterShutDownEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicClusterShutDownEventClass() _MTRBridgedDeviceBasicClusterShutDownEventClass {
	MTRBridgedDeviceBasicClusterShutDownEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicClusterShutDownEventClass = _MTRBridgedDeviceBasicClusterShutDownEventClass{objc.GetClass("MTRBridgedDeviceBasicClusterShutDownEvent")}
	})
	return MTRBridgedDeviceBasicClusterShutDownEventClass
}

type _MTRBridgedDeviceBasicClusterShutDownEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicClusterShutDownEvent] class.
type IMTRBridgedDeviceBasicClusterShutDownEvent interface {
	IMTRBridgedDeviceBasicInformationClusterShutDownEvent
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicClusterShutDownEvent
type MTRBridgedDeviceBasicClusterShutDownEvent struct {
	MTRBridgedDeviceBasicInformationClusterShutDownEvent
}

// MTRBridgedDeviceBasicClusterShutDownEventFrom constructs a [MTRBridgedDeviceBasicClusterShutDownEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicClusterShutDownEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicClusterShutDownEvent {
	return MTRBridgedDeviceBasicClusterShutDownEvent{
		MTRBridgedDeviceBasicInformationClusterShutDownEvent: MTRBridgedDeviceBasicInformationClusterShutDownEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicClusterShutDownEventClass) Alloc() MTRBridgedDeviceBasicClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterShutDownEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicClusterShutDownEventClass) New() MTRBridgedDeviceBasicClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterShutDownEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicClusterShutDownEvent) Init() MTRBridgedDeviceBasicClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterShutDownEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicClusterShutDownEvent) Autorelease() MTRBridgedDeviceBasicClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterShutDownEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicClusterShutDownEvent creates a new MTRBridgedDeviceBasicClusterShutDownEvent instance.
func NewMTRBridgedDeviceBasicClusterShutDownEvent() MTRBridgedDeviceBasicClusterShutDownEvent {
	return getMTRBridgedDeviceBasicClusterShutDownEventClass().New()
}




