// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRBridgedDeviceBasicInformationClusterReachableChangedEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass     _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass
	MTRBridgedDeviceBasicInformationClusterReachableChangedEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterReachableChangedEventClass() _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass {
	MTRBridgedDeviceBasicInformationClusterReachableChangedEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass = _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterReachableChangedEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass
}

type _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRBridgedDeviceBasicInformationClusterReachableChangedEvent] class.
type IMTRBridgedDeviceBasicInformationClusterReachableChangedEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterReachableChangedEvent
type MTRBridgedDeviceBasicInformationClusterReachableChangedEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterReachableChangedEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterReachableChangedEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterReachableChangedEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	return MTRBridgedDeviceBasicInformationClusterReachableChangedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBridgedDeviceBasicInformationClusterReachableChangedEventClass) New() MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterReachableChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterReachableChangedEvent) Init() MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterReachableChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterReachableChangedEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterReachableChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterReachableChangedEvent creates a new MTRBridgedDeviceBasicInformationClusterReachableChangedEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterReachableChangedEvent() MTRBridgedDeviceBasicInformationClusterReachableChangedEvent {
	return getMTRBridgedDeviceBasicInformationClusterReachableChangedEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterreachablechangedevent/reachablenewvalue
func (m_ MTRBridgedDeviceBasicInformationClusterReachableChangedEvent) ReachableNewValue() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("reachableNewValue"))
	return rv
}


// SetReachableNewValue sets the value of the reachableNewValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrbridgeddevicebasicinformationclusterreachablechangedevent/reachablenewvalue
func (m_ MTRBridgedDeviceBasicInformationClusterReachableChangedEvent) SetReachableNewValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setReachableNewValue:"), value)
}



