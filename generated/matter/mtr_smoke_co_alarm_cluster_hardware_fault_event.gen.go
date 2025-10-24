// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterHardwareFaultEvent] class.
var (
	MTRSmokeCOAlarmClusterHardwareFaultEventClass     _MTRSmokeCOAlarmClusterHardwareFaultEventClass
	MTRSmokeCOAlarmClusterHardwareFaultEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterHardwareFaultEventClass() _MTRSmokeCOAlarmClusterHardwareFaultEventClass {
	MTRSmokeCOAlarmClusterHardwareFaultEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterHardwareFaultEventClass = _MTRSmokeCOAlarmClusterHardwareFaultEventClass{objc.GetClass("MTRSmokeCOAlarmClusterHardwareFaultEvent")}
	})
	return MTRSmokeCOAlarmClusterHardwareFaultEventClass
}

type _MTRSmokeCOAlarmClusterHardwareFaultEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterHardwareFaultEvent] class.
type IMTRSmokeCOAlarmClusterHardwareFaultEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterHardwareFaultEvent
type MTRSmokeCOAlarmClusterHardwareFaultEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterHardwareFaultEventFrom constructs a [MTRSmokeCOAlarmClusterHardwareFaultEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterHardwareFaultEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterHardwareFaultEvent {
	return MTRSmokeCOAlarmClusterHardwareFaultEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterHardwareFaultEventClass) Alloc() MTRSmokeCOAlarmClusterHardwareFaultEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterHardwareFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterHardwareFaultEventClass) New() MTRSmokeCOAlarmClusterHardwareFaultEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterHardwareFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterHardwareFaultEvent) Init() MTRSmokeCOAlarmClusterHardwareFaultEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterHardwareFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterHardwareFaultEvent) Autorelease() MTRSmokeCOAlarmClusterHardwareFaultEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterHardwareFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterHardwareFaultEvent creates a new MTRSmokeCOAlarmClusterHardwareFaultEvent instance.
func NewMTRSmokeCOAlarmClusterHardwareFaultEvent() MTRSmokeCOAlarmClusterHardwareFaultEvent {
	return getMTRSmokeCOAlarmClusterHardwareFaultEventClass().New()
}




