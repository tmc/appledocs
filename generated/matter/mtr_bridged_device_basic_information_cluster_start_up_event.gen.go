// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterStartUpEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterStartUpEventClass     _MTRBridgedDeviceBasicInformationClusterStartUpEventClass
	MTRBridgedDeviceBasicInformationClusterStartUpEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterStartUpEventClass() _MTRBridgedDeviceBasicInformationClusterStartUpEventClass {
	MTRBridgedDeviceBasicInformationClusterStartUpEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterStartUpEventClass = _MTRBridgedDeviceBasicInformationClusterStartUpEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterStartUpEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterStartUpEventClass
}

type _MTRBridgedDeviceBasicInformationClusterStartUpEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterStartUpEvent] class.
type IMTRBridgedDeviceBasicInformationClusterStartUpEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterStartUpEvent
type MTRBridgedDeviceBasicInformationClusterStartUpEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterStartUpEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterStartUpEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterStartUpEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	return MTRBridgedDeviceBasicInformationClusterStartUpEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterStartUpEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterStartUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterStartUpEventClass) New() MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterStartUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterStartUpEvent) Init() MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterStartUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterStartUpEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterStartUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterStartUpEvent creates a new MTRBridgedDeviceBasicInformationClusterStartUpEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterStartUpEvent() MTRBridgedDeviceBasicInformationClusterStartUpEvent {
	return getMTRBridgedDeviceBasicInformationClusterStartUpEventClass().New()
}




