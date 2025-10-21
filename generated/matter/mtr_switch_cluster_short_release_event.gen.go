// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSwitchClusterShortReleaseEvent] class.
var (
	MTRSwitchClusterShortReleaseEventClass     _MTRSwitchClusterShortReleaseEventClass
	MTRSwitchClusterShortReleaseEventClassOnce sync.Once
)

func getMTRSwitchClusterShortReleaseEventClass() _MTRSwitchClusterShortReleaseEventClass {
	MTRSwitchClusterShortReleaseEventClassOnce.Do(func() {
		MTRSwitchClusterShortReleaseEventClass = _MTRSwitchClusterShortReleaseEventClass{objc.GetClass("MTRSwitchClusterShortReleaseEvent")}
	})
	return MTRSwitchClusterShortReleaseEventClass
}

type _MTRSwitchClusterShortReleaseEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterShortReleaseEvent] class.
type IMTRSwitchClusterShortReleaseEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterShortReleaseEvent
type MTRSwitchClusterShortReleaseEvent struct {
	objectivec.Object
}

// MTRSwitchClusterShortReleaseEventFrom constructs a [MTRSwitchClusterShortReleaseEvent] from an unsafe.Pointer.
func MTRSwitchClusterShortReleaseEventFrom(ptr unsafe.Pointer) MTRSwitchClusterShortReleaseEvent {
	return MTRSwitchClusterShortReleaseEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterShortReleaseEventClass) Alloc() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterShortReleaseEventClass) New() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterShortReleaseEvent) Init() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterShortReleaseEvent) Autorelease() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterShortReleaseEvent creates a new MTRSwitchClusterShortReleaseEvent instance.
func NewMTRSwitchClusterShortReleaseEvent() MTRSwitchClusterShortReleaseEvent {
	return getMTRSwitchClusterShortReleaseEventClass().New()
}




