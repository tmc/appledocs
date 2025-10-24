// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterLowBatteryEvent] class.
var (
	MTRSmokeCOAlarmClusterLowBatteryEventClass     _MTRSmokeCOAlarmClusterLowBatteryEventClass
	MTRSmokeCOAlarmClusterLowBatteryEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterLowBatteryEventClass() _MTRSmokeCOAlarmClusterLowBatteryEventClass {
	MTRSmokeCOAlarmClusterLowBatteryEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterLowBatteryEventClass = _MTRSmokeCOAlarmClusterLowBatteryEventClass{objc.GetClass("MTRSmokeCOAlarmClusterLowBatteryEvent")}
	})
	return MTRSmokeCOAlarmClusterLowBatteryEventClass
}

type _MTRSmokeCOAlarmClusterLowBatteryEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterLowBatteryEvent] class.
type IMTRSmokeCOAlarmClusterLowBatteryEvent interface {
	objectivec.IObject
	// properties:
	AlarmSeverityLevel() objc.IObject /* cross-framework: NSNumber */
	SetAlarmSeverityLevel(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterLowBatteryEvent
type MTRSmokeCOAlarmClusterLowBatteryEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterLowBatteryEventFrom constructs a [MTRSmokeCOAlarmClusterLowBatteryEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterLowBatteryEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterLowBatteryEvent {
	return MTRSmokeCOAlarmClusterLowBatteryEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterLowBatteryEventClass) Alloc() MTRSmokeCOAlarmClusterLowBatteryEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterLowBatteryEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterLowBatteryEventClass) New() MTRSmokeCOAlarmClusterLowBatteryEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterLowBatteryEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterLowBatteryEvent) Init() MTRSmokeCOAlarmClusterLowBatteryEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterLowBatteryEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterLowBatteryEvent) Autorelease() MTRSmokeCOAlarmClusterLowBatteryEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterLowBatteryEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterLowBatteryEvent creates a new MTRSmokeCOAlarmClusterLowBatteryEvent instance.
func NewMTRSmokeCOAlarmClusterLowBatteryEvent() MTRSmokeCOAlarmClusterLowBatteryEvent {
	return getMTRSmokeCOAlarmClusterLowBatteryEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterlowbatteryevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterLowBatteryEvent) AlarmSeverityLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmSeverityLevel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterlowbatteryevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterLowBatteryEvent) SetAlarmSeverityLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmSeverityLevel:"), value)
}



