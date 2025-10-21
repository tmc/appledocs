// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSmokeCOAlarmClusterSelfTestCompleteEvent] class.
var (
	MTRSmokeCOAlarmClusterSelfTestCompleteEventClass     _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass
	MTRSmokeCOAlarmClusterSelfTestCompleteEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterSelfTestCompleteEventClass() _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass {
	MTRSmokeCOAlarmClusterSelfTestCompleteEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterSelfTestCompleteEventClass = _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass{objc.GetClass("MTRSmokeCOAlarmClusterSelfTestCompleteEvent")}
	})
	return MTRSmokeCOAlarmClusterSelfTestCompleteEventClass
}

type _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterSelfTestCompleteEvent] class.
type IMTRSmokeCOAlarmClusterSelfTestCompleteEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterSelfTestCompleteEvent
type MTRSmokeCOAlarmClusterSelfTestCompleteEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterSelfTestCompleteEventFrom constructs a [MTRSmokeCOAlarmClusterSelfTestCompleteEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterSelfTestCompleteEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	return MTRSmokeCOAlarmClusterSelfTestCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass) Alloc() MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterSelfTestCompleteEventClass) New() MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterSelfTestCompleteEvent) Init() MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterSelfTestCompleteEvent) Autorelease() MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterSelfTestCompleteEvent creates a new MTRSmokeCOAlarmClusterSelfTestCompleteEvent instance.
func NewMTRSmokeCOAlarmClusterSelfTestCompleteEvent() MTRSmokeCOAlarmClusterSelfTestCompleteEvent {
	return getMTRSmokeCOAlarmClusterSelfTestCompleteEventClass().New()
}




