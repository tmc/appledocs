// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterTimeFailureEvent] class.
var (
	MTRTimeSynchronizationClusterTimeFailureEventClass     _MTRTimeSynchronizationClusterTimeFailureEventClass
	MTRTimeSynchronizationClusterTimeFailureEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeFailureEventClass() _MTRTimeSynchronizationClusterTimeFailureEventClass {
	MTRTimeSynchronizationClusterTimeFailureEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeFailureEventClass = _MTRTimeSynchronizationClusterTimeFailureEventClass{objc.GetClass("MTRTimeSynchronizationClusterTimeFailureEvent")}
	})
	return MTRTimeSynchronizationClusterTimeFailureEventClass
}

type _MTRTimeSynchronizationClusterTimeFailureEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTimeFailureEvent] class.
type IMTRTimeSynchronizationClusterTimeFailureEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeFailureEvent
type MTRTimeSynchronizationClusterTimeFailureEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTimeFailureEventFrom constructs a [MTRTimeSynchronizationClusterTimeFailureEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeFailureEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeFailureEvent {
	return MTRTimeSynchronizationClusterTimeFailureEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeFailureEventClass) Alloc() MTRTimeSynchronizationClusterTimeFailureEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTimeFailureEventClass) New() MTRTimeSynchronizationClusterTimeFailureEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeFailureEvent) Init() MTRTimeSynchronizationClusterTimeFailureEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeFailureEvent) Autorelease() MTRTimeSynchronizationClusterTimeFailureEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeFailureEvent creates a new MTRTimeSynchronizationClusterTimeFailureEvent instance.
func NewMTRTimeSynchronizationClusterTimeFailureEvent() MTRTimeSynchronizationClusterTimeFailureEvent {
	return getMTRTimeSynchronizationClusterTimeFailureEventClass().New()
}




