// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass     _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass() _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass {
	MTROTASoftwareUpdateRequestorClusterStateTransitionEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass = _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterStateTransitionEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass
}

type _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] class.
type IMTROTASoftwareUpdateRequestorClusterStateTransitionEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterStateTransitionEvent-3xhxb
type MTROTASoftwareUpdateRequestorClusterStateTransitionEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterStateTransitionEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterStateTransitionEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return MTROTASoftwareUpdateRequestorClusterStateTransitionEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateRequestorClusterStateTransitionEventClass) New() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Init() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterStateTransitionEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterStateTransitionEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent creates a new MTROTASoftwareUpdateRequestorClusterStateTransitionEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterStateTransitionEvent() MTROTASoftwareUpdateRequestorClusterStateTransitionEvent {
	return getMTROTASoftwareUpdateRequestorClusterStateTransitionEventClass().New()
}




