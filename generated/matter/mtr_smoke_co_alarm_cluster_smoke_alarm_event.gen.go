// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterSmokeAlarmEvent] class.
var (
	MTRSmokeCOAlarmClusterSmokeAlarmEventClass     _MTRSmokeCOAlarmClusterSmokeAlarmEventClass
	MTRSmokeCOAlarmClusterSmokeAlarmEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterSmokeAlarmEventClass() _MTRSmokeCOAlarmClusterSmokeAlarmEventClass {
	MTRSmokeCOAlarmClusterSmokeAlarmEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterSmokeAlarmEventClass = _MTRSmokeCOAlarmClusterSmokeAlarmEventClass{objc.GetClass("MTRSmokeCOAlarmClusterSmokeAlarmEvent")}
	})
	return MTRSmokeCOAlarmClusterSmokeAlarmEventClass
}

type _MTRSmokeCOAlarmClusterSmokeAlarmEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterSmokeAlarmEvent] class.
type IMTRSmokeCOAlarmClusterSmokeAlarmEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterSmokeAlarmEvent
type MTRSmokeCOAlarmClusterSmokeAlarmEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterSmokeAlarmEventFrom constructs a [MTRSmokeCOAlarmClusterSmokeAlarmEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterSmokeAlarmEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	return MTRSmokeCOAlarmClusterSmokeAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterSmokeAlarmEventClass) Alloc() MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSmokeAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterSmokeAlarmEventClass) New() MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSmokeAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterSmokeAlarmEvent) Init() MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSmokeAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterSmokeAlarmEvent) Autorelease() MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSmokeAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterSmokeAlarmEvent creates a new MTRSmokeCOAlarmClusterSmokeAlarmEvent instance.
func NewMTRSmokeCOAlarmClusterSmokeAlarmEvent() MTRSmokeCOAlarmClusterSmokeAlarmEvent {
	return getMTRSmokeCOAlarmClusterSmokeAlarmEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclustersmokealarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterSmokeAlarmEvent) AlarmSeverityLevel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("alarmSeverityLevel"))
	return rv
}


// SetAlarmSeverityLevel sets the value of the alarmSeverityLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclustersmokealarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterSmokeAlarmEvent) SetAlarmSeverityLevel(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmSeverityLevel:"), value)
}



