// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterCOAlarmEvent] class.
var (
	MTRSmokeCOAlarmClusterCOAlarmEventClass     _MTRSmokeCOAlarmClusterCOAlarmEventClass
	MTRSmokeCOAlarmClusterCOAlarmEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterCOAlarmEventClass() _MTRSmokeCOAlarmClusterCOAlarmEventClass {
	MTRSmokeCOAlarmClusterCOAlarmEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterCOAlarmEventClass = _MTRSmokeCOAlarmClusterCOAlarmEventClass{objc.GetClass("MTRSmokeCOAlarmClusterCOAlarmEvent")}
	})
	return MTRSmokeCOAlarmClusterCOAlarmEventClass
}

type _MTRSmokeCOAlarmClusterCOAlarmEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterCOAlarmEvent] class.
type IMTRSmokeCOAlarmClusterCOAlarmEvent interface {
	objectivec.IObject
	// properties:
	AlarmSeverityLevel() objc.IObject /* cross-framework: NSNumber */
	SetAlarmSeverityLevel(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterCOAlarmEvent
type MTRSmokeCOAlarmClusterCOAlarmEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterCOAlarmEventFrom constructs a [MTRSmokeCOAlarmClusterCOAlarmEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterCOAlarmEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterCOAlarmEvent {
	return MTRSmokeCOAlarmClusterCOAlarmEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterCOAlarmEventClass) Alloc() MTRSmokeCOAlarmClusterCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterCOAlarmEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterCOAlarmEventClass) New() MTRSmokeCOAlarmClusterCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterCOAlarmEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterCOAlarmEvent) Init() MTRSmokeCOAlarmClusterCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterCOAlarmEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterCOAlarmEvent) Autorelease() MTRSmokeCOAlarmClusterCOAlarmEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterCOAlarmEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterCOAlarmEvent creates a new MTRSmokeCOAlarmClusterCOAlarmEvent instance.
func NewMTRSmokeCOAlarmClusterCOAlarmEvent() MTRSmokeCOAlarmClusterCOAlarmEvent {
	return getMTRSmokeCOAlarmClusterCOAlarmEventClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclustercoalarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterCOAlarmEvent) AlarmSeverityLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("alarmSeverityLevel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclustercoalarmevent/alarmseveritylevel
func (m_ MTRSmokeCOAlarmClusterCOAlarmEvent) SetAlarmSeverityLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAlarmSeverityLevel:"), value)
}



