// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterDSTTableEmptyEvent] class.
var (
	MTRTimeSynchronizationClusterDSTTableEmptyEventClass     _MTRTimeSynchronizationClusterDSTTableEmptyEventClass
	MTRTimeSynchronizationClusterDSTTableEmptyEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTTableEmptyEventClass() _MTRTimeSynchronizationClusterDSTTableEmptyEventClass {
	MTRTimeSynchronizationClusterDSTTableEmptyEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTTableEmptyEventClass = _MTRTimeSynchronizationClusterDSTTableEmptyEventClass{objc.GetClass("MTRTimeSynchronizationClusterDSTTableEmptyEvent")}
	})
	return MTRTimeSynchronizationClusterDSTTableEmptyEventClass
}

type _MTRTimeSynchronizationClusterDSTTableEmptyEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterDSTTableEmptyEvent] class.
type IMTRTimeSynchronizationClusterDSTTableEmptyEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTTableEmptyEvent
type MTRTimeSynchronizationClusterDSTTableEmptyEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTTableEmptyEventFrom constructs a [MTRTimeSynchronizationClusterDSTTableEmptyEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTTableEmptyEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	return MTRTimeSynchronizationClusterDSTTableEmptyEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTTableEmptyEventClass) Alloc() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterDSTTableEmptyEventClass) New() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTTableEmptyEvent) Init() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTTableEmptyEvent) Autorelease() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTTableEmptyEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTTableEmptyEvent creates a new MTRTimeSynchronizationClusterDSTTableEmptyEvent instance.
func NewMTRTimeSynchronizationClusterDSTTableEmptyEvent() MTRTimeSynchronizationClusterDSTTableEmptyEvent {
	return getMTRTimeSynchronizationClusterDSTTableEmptyEventClass().New()
}




