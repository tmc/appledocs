// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent] class.
var (
	MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass     _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass
	MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass() _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass {
	MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass = _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass{objc.GetClass("MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent")}
	})
	return MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass
}

type _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent] class.
type IMTRSmokeCOAlarmClusterInterconnectCOAlarmEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent
type MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterInterconnectCOAlarmEventFrom constructs a [MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterInterconnectCOAlarmEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	return MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass) Alloc() MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass) New() MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent) Init() MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent) Autorelease() MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterInterconnectCOAlarmEvent creates a new MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent instance.
func NewMTRSmokeCOAlarmClusterInterconnectCOAlarmEvent() MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent {
	return getMTRSmokeCOAlarmClusterInterconnectCOAlarmEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterinterconnectcoalarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent) AlarmSeverityLevel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("alarmSeverityLevel"))
	return rv
}


// SetAlarmSeverityLevel sets the value of the alarmSeverityLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterinterconnectcoalarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterInterconnectCOAlarmEvent) SetAlarmSeverityLevel(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmSeverityLevel:"), value)
}



