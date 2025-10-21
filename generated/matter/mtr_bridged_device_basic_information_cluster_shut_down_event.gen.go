// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterShutDownEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterShutDownEventClass     _MTRBridgedDeviceBasicInformationClusterShutDownEventClass
	MTRBridgedDeviceBasicInformationClusterShutDownEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterShutDownEventClass() _MTRBridgedDeviceBasicInformationClusterShutDownEventClass {
	MTRBridgedDeviceBasicInformationClusterShutDownEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterShutDownEventClass = _MTRBridgedDeviceBasicInformationClusterShutDownEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterShutDownEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterShutDownEventClass
}

type _MTRBridgedDeviceBasicInformationClusterShutDownEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterShutDownEvent] class.
type IMTRBridgedDeviceBasicInformationClusterShutDownEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterShutDownEvent
type MTRBridgedDeviceBasicInformationClusterShutDownEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterShutDownEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterShutDownEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterShutDownEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	return MTRBridgedDeviceBasicInformationClusterShutDownEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterShutDownEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterShutDownEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterShutDownEventClass) New() MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterShutDownEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterShutDownEvent) Init() MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterShutDownEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterShutDownEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterShutDownEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterShutDownEvent creates a new MTRBridgedDeviceBasicInformationClusterShutDownEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterShutDownEvent() MTRBridgedDeviceBasicInformationClusterShutDownEvent {
	return getMTRBridgedDeviceBasicInformationClusterShutDownEventClass().New()
}




