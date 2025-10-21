// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBridgedDeviceBasicClusterStartUpEvent] class.
var (
	MTRBridgedDeviceBasicClusterStartUpEventClass     _MTRBridgedDeviceBasicClusterStartUpEventClass
	MTRBridgedDeviceBasicClusterStartUpEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicClusterStartUpEventClass() _MTRBridgedDeviceBasicClusterStartUpEventClass {
	MTRBridgedDeviceBasicClusterStartUpEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicClusterStartUpEventClass = _MTRBridgedDeviceBasicClusterStartUpEventClass{objc.GetClass("MTRBridgedDeviceBasicClusterStartUpEvent")}
	})
	return MTRBridgedDeviceBasicClusterStartUpEventClass
}

type _MTRBridgedDeviceBasicClusterStartUpEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicClusterStartUpEvent] class.
type IMTRBridgedDeviceBasicClusterStartUpEvent interface {
	IMTRBridgedDeviceBasicInformationClusterStartUpEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicClusterStartUpEvent
type MTRBridgedDeviceBasicClusterStartUpEvent struct {
	MTRBridgedDeviceBasicInformationClusterStartUpEvent
}

// MTRBridgedDeviceBasicClusterStartUpEventFrom constructs a [MTRBridgedDeviceBasicClusterStartUpEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicClusterStartUpEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicClusterStartUpEvent {
	return MTRBridgedDeviceBasicClusterStartUpEvent{
		MTRBridgedDeviceBasicInformationClusterStartUpEvent: MTRBridgedDeviceBasicInformationClusterStartUpEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicClusterStartUpEventClass) Alloc() MTRBridgedDeviceBasicClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterStartUpEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicClusterStartUpEventClass) New() MTRBridgedDeviceBasicClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterStartUpEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicClusterStartUpEvent) Init() MTRBridgedDeviceBasicClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterStartUpEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicClusterStartUpEvent) Autorelease() MTRBridgedDeviceBasicClusterStartUpEvent {
	rv := objc.Send[MTRBridgedDeviceBasicClusterStartUpEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicClusterStartUpEvent creates a new MTRBridgedDeviceBasicClusterStartUpEvent instance.
func NewMTRBridgedDeviceBasicClusterStartUpEvent() MTRBridgedDeviceBasicClusterStartUpEvent {
	return getMTRBridgedDeviceBasicClusterStartUpEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicclusterstartupevent/softwareversion
func (m_ MTRBridgedDeviceBasicClusterStartUpEvent) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicclusterstartupevent/softwareversion
func (m_ MTRBridgedDeviceBasicClusterStartUpEvent) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}



