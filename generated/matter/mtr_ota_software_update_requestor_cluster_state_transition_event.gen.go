// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent interface {
	IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent-1xzd5
type MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent struct {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEvent
}

// MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent{
		MTROTASoftwareUpdateRequestorClusterStateTransitionEvent: MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterStateTransitionEvent() MTROtaSoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROtaSoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}




