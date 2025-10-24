// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterEndOfServiceEvent] class.
var (
	MTRSmokeCOAlarmClusterEndOfServiceEventClass     _MTRSmokeCOAlarmClusterEndOfServiceEventClass
	MTRSmokeCOAlarmClusterEndOfServiceEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterEndOfServiceEventClass() _MTRSmokeCOAlarmClusterEndOfServiceEventClass {
	MTRSmokeCOAlarmClusterEndOfServiceEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterEndOfServiceEventClass = _MTRSmokeCOAlarmClusterEndOfServiceEventClass{objc.GetClass("MTRSmokeCOAlarmClusterEndOfServiceEvent")}
	})
	return MTRSmokeCOAlarmClusterEndOfServiceEventClass
}

type _MTRSmokeCOAlarmClusterEndOfServiceEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterEndOfServiceEvent] class.
type IMTRSmokeCOAlarmClusterEndOfServiceEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterEndOfServiceEvent
type MTRSmokeCOAlarmClusterEndOfServiceEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterEndOfServiceEventFrom constructs a [MTRSmokeCOAlarmClusterEndOfServiceEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterEndOfServiceEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterEndOfServiceEvent {
	return MTRSmokeCOAlarmClusterEndOfServiceEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterEndOfServiceEventClass) Alloc() MTRSmokeCOAlarmClusterEndOfServiceEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterEndOfServiceEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterEndOfServiceEventClass) New() MTRSmokeCOAlarmClusterEndOfServiceEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterEndOfServiceEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterEndOfServiceEvent) Init() MTRSmokeCOAlarmClusterEndOfServiceEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterEndOfServiceEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterEndOfServiceEvent) Autorelease() MTRSmokeCOAlarmClusterEndOfServiceEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterEndOfServiceEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterEndOfServiceEvent creates a new MTRSmokeCOAlarmClusterEndOfServiceEvent instance.
func NewMTRSmokeCOAlarmClusterEndOfServiceEvent() MTRSmokeCOAlarmClusterEndOfServiceEvent {
	return getMTRSmokeCOAlarmClusterEndOfServiceEventClass().New()
}




