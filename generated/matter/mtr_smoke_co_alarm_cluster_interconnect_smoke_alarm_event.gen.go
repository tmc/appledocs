// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent] class.
var (
	MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass     _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass
	MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass() _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass {
	MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass = _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass{objc.GetClass("MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent")}
	})
	return MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass
}

type _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent] class.
type IMTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent interface {
	objectivec.IObject
	AlarmSeverityLevel() foundation.Number
	SetAlarmSeverityLevel(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent
type MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventFrom constructs a [MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	return MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass) Alloc() MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass) New() MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent) Init() MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent) Autorelease() MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent creates a new MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent instance.
func NewMTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent() MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent {
	return getMTRSmokeCOAlarmClusterInterconnectSmokeAlarmEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterinterconnectsmokealarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent) AlarmSeverityLevel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("alarmSeverityLevel"))
	return rv
}


// SetAlarmSeverityLevel sets the value of the alarmSeverityLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterinterconnectsmokealarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterInterconnectSmokeAlarmEvent) SetAlarmSeverityLevel(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmSeverityLevel:"), value)
}



