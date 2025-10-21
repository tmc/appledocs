// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterMultiPressOngoingEvent] class.
var (
	MTRSwitchClusterMultiPressOngoingEventClass     _MTRSwitchClusterMultiPressOngoingEventClass
	MTRSwitchClusterMultiPressOngoingEventClassOnce sync.Once
)

func getMTRSwitchClusterMultiPressOngoingEventClass() _MTRSwitchClusterMultiPressOngoingEventClass {
	MTRSwitchClusterMultiPressOngoingEventClassOnce.Do(func() {
		MTRSwitchClusterMultiPressOngoingEventClass = _MTRSwitchClusterMultiPressOngoingEventClass{objc.GetClass("MTRSwitchClusterMultiPressOngoingEvent")}
	})
	return MTRSwitchClusterMultiPressOngoingEventClass
}

type _MTRSwitchClusterMultiPressOngoingEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterMultiPressOngoingEvent] class.
type IMTRSwitchClusterMultiPressOngoingEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent
type MTRSwitchClusterMultiPressOngoingEvent struct {
	objectivec.Object
}

// MTRSwitchClusterMultiPressOngoingEventFrom constructs a [MTRSwitchClusterMultiPressOngoingEvent] from an unsafe.Pointer.
func MTRSwitchClusterMultiPressOngoingEventFrom(ptr unsafe.Pointer) MTRSwitchClusterMultiPressOngoingEvent {
	return MTRSwitchClusterMultiPressOngoingEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterMultiPressOngoingEventClass) Alloc() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterMultiPressOngoingEventClass) New() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterMultiPressOngoingEvent) Init() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterMultiPressOngoingEvent) Autorelease() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterMultiPressOngoingEvent creates a new MTRSwitchClusterMultiPressOngoingEvent instance.
func NewMTRSwitchClusterMultiPressOngoingEvent() MTRSwitchClusterMultiPressOngoingEvent {
	return getMTRSwitchClusterMultiPressOngoingEventClass().New()
}




