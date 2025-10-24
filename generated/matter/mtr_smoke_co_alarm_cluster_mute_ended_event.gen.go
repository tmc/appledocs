// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterMuteEndedEvent] class.
var (
	MTRSmokeCOAlarmClusterMuteEndedEventClass     _MTRSmokeCOAlarmClusterMuteEndedEventClass
	MTRSmokeCOAlarmClusterMuteEndedEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterMuteEndedEventClass() _MTRSmokeCOAlarmClusterMuteEndedEventClass {
	MTRSmokeCOAlarmClusterMuteEndedEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterMuteEndedEventClass = _MTRSmokeCOAlarmClusterMuteEndedEventClass{objc.GetClass("MTRSmokeCOAlarmClusterMuteEndedEvent")}
	})
	return MTRSmokeCOAlarmClusterMuteEndedEventClass
}

type _MTRSmokeCOAlarmClusterMuteEndedEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterMuteEndedEvent] class.
type IMTRSmokeCOAlarmClusterMuteEndedEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterMuteEndedEvent
type MTRSmokeCOAlarmClusterMuteEndedEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterMuteEndedEventFrom constructs a [MTRSmokeCOAlarmClusterMuteEndedEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterMuteEndedEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterMuteEndedEvent {
	return MTRSmokeCOAlarmClusterMuteEndedEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterMuteEndedEventClass) Alloc() MTRSmokeCOAlarmClusterMuteEndedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterMuteEndedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterMuteEndedEventClass) New() MTRSmokeCOAlarmClusterMuteEndedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterMuteEndedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterMuteEndedEvent) Init() MTRSmokeCOAlarmClusterMuteEndedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterMuteEndedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterMuteEndedEvent) Autorelease() MTRSmokeCOAlarmClusterMuteEndedEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterMuteEndedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterMuteEndedEvent creates a new MTRSmokeCOAlarmClusterMuteEndedEvent instance.
func NewMTRSmokeCOAlarmClusterMuteEndedEvent() MTRSmokeCOAlarmClusterMuteEndedEvent {
	return getMTRSmokeCOAlarmClusterMuteEndedEventClass().New()
}




