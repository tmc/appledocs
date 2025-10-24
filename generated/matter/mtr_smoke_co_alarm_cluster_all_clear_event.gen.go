// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterAllClearEvent] class.
var (
	MTRSmokeCOAlarmClusterAllClearEventClass     _MTRSmokeCOAlarmClusterAllClearEventClass
	MTRSmokeCOAlarmClusterAllClearEventClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterAllClearEventClass() _MTRSmokeCOAlarmClusterAllClearEventClass {
	MTRSmokeCOAlarmClusterAllClearEventClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterAllClearEventClass = _MTRSmokeCOAlarmClusterAllClearEventClass{objc.GetClass("MTRSmokeCOAlarmClusterAllClearEvent")}
	})
	return MTRSmokeCOAlarmClusterAllClearEventClass
}

type _MTRSmokeCOAlarmClusterAllClearEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterAllClearEvent] class.
type IMTRSmokeCOAlarmClusterAllClearEvent interface {
	objectivec.IObject
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterAllClearEvent
type MTRSmokeCOAlarmClusterAllClearEvent struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterAllClearEventFrom constructs a [MTRSmokeCOAlarmClusterAllClearEvent] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterAllClearEventFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterAllClearEvent {
	return MTRSmokeCOAlarmClusterAllClearEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterAllClearEventClass) Alloc() MTRSmokeCOAlarmClusterAllClearEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAllClearEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterAllClearEventClass) New() MTRSmokeCOAlarmClusterAllClearEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAllClearEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterAllClearEvent) Init() MTRSmokeCOAlarmClusterAllClearEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAllClearEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterAllClearEvent) Autorelease() MTRSmokeCOAlarmClusterAllClearEvent {
	rv := objc.Send[MTRSmokeCOAlarmClusterAllClearEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterAllClearEvent creates a new MTRSmokeCOAlarmClusterAllClearEvent instance.
func NewMTRSmokeCOAlarmClusterAllClearEvent() MTRSmokeCOAlarmClusterAllClearEvent {
	return getMTRSmokeCOAlarmClusterAllClearEventClass().New()
}
