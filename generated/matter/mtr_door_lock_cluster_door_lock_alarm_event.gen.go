// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterDoorLockAlarmEvent] class.
var (
	MTRDoorLockClusterDoorLockAlarmEventClass     _MTRDoorLockClusterDoorLockAlarmEventClass
	MTRDoorLockClusterDoorLockAlarmEventClassOnce sync.Once
)

func getMTRDoorLockClusterDoorLockAlarmEventClass() _MTRDoorLockClusterDoorLockAlarmEventClass {
	MTRDoorLockClusterDoorLockAlarmEventClassOnce.Do(func() {
		MTRDoorLockClusterDoorLockAlarmEventClass = _MTRDoorLockClusterDoorLockAlarmEventClass{objc.GetClass("MTRDoorLockClusterDoorLockAlarmEvent")}
	})
	return MTRDoorLockClusterDoorLockAlarmEventClass
}

type _MTRDoorLockClusterDoorLockAlarmEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterDoorLockAlarmEvent] class.
type IMTRDoorLockClusterDoorLockAlarmEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorLockAlarmEvent
type MTRDoorLockClusterDoorLockAlarmEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterDoorLockAlarmEventFrom constructs a [MTRDoorLockClusterDoorLockAlarmEvent] from an unsafe.Pointer.
func MTRDoorLockClusterDoorLockAlarmEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterDoorLockAlarmEvent {
	return MTRDoorLockClusterDoorLockAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterDoorLockAlarmEventClass) Alloc() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterDoorLockAlarmEventClass) New() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) Init() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) Autorelease() MTRDoorLockClusterDoorLockAlarmEvent {
	rv := objc.Send[MTRDoorLockClusterDoorLockAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterDoorLockAlarmEvent creates a new MTRDoorLockClusterDoorLockAlarmEvent instance.
func NewMTRDoorLockClusterDoorLockAlarmEvent() MTRDoorLockClusterDoorLockAlarmEvent {
	return getMTRDoorLockClusterDoorLockAlarmEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterdoorlockalarmevent/alarmcode
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) AlarmCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("alarmCode"))
	return rv
}


// SetAlarmCode sets the value of the alarmCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterdoorlockalarmevent/alarmcode
func (m_ MTRDoorLockClusterDoorLockAlarmEvent) SetAlarmCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmCode:"), value)
}



