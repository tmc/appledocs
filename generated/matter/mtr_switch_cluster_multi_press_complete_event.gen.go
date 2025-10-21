// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRSwitchClusterMultiPressCompleteEvent] class.
var (
	MTRSwitchClusterMultiPressCompleteEventClass     _MTRSwitchClusterMultiPressCompleteEventClass
	MTRSwitchClusterMultiPressCompleteEventClassOnce sync.Once
)

func getMTRSwitchClusterMultiPressCompleteEventClass() _MTRSwitchClusterMultiPressCompleteEventClass {
	MTRSwitchClusterMultiPressCompleteEventClassOnce.Do(func() {
		MTRSwitchClusterMultiPressCompleteEventClass = _MTRSwitchClusterMultiPressCompleteEventClass{objc.GetClass("MTRSwitchClusterMultiPressCompleteEvent")}
	})
	return MTRSwitchClusterMultiPressCompleteEventClass
}

type _MTRSwitchClusterMultiPressCompleteEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRSwitchClusterMultiPressCompleteEvent] class.
type IMTRSwitchClusterMultiPressCompleteEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent
type MTRSwitchClusterMultiPressCompleteEvent struct {
	objectivec.Object
}

// MTRSwitchClusterMultiPressCompleteEventFrom constructs a [MTRSwitchClusterMultiPressCompleteEvent] from an unsafe.Pointer.
func MTRSwitchClusterMultiPressCompleteEventFrom(ptr unsafe.Pointer) MTRSwitchClusterMultiPressCompleteEvent {
	return MTRSwitchClusterMultiPressCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) Alloc() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) New() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Init() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Autorelease() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterMultiPressCompleteEvent creates a new MTRSwitchClusterMultiPressCompleteEvent instance.
func NewMTRSwitchClusterMultiPressCompleteEvent() MTRSwitchClusterMultiPressCompleteEvent {
	return getMTRSwitchClusterMultiPressCompleteEventClass().New()
}




