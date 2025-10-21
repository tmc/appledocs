// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterDSTStatusEvent] class.
var (
	MTRTimeSynchronizationClusterDSTStatusEventClass     _MTRTimeSynchronizationClusterDSTStatusEventClass
	MTRTimeSynchronizationClusterDSTStatusEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDSTStatusEventClass() _MTRTimeSynchronizationClusterDSTStatusEventClass {
	MTRTimeSynchronizationClusterDSTStatusEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDSTStatusEventClass = _MTRTimeSynchronizationClusterDSTStatusEventClass{objc.GetClass("MTRTimeSynchronizationClusterDSTStatusEvent")}
	})
	return MTRTimeSynchronizationClusterDSTStatusEventClass
}

type _MTRTimeSynchronizationClusterDSTStatusEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterDSTStatusEvent] class.
type IMTRTimeSynchronizationClusterDSTStatusEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent
type MTRTimeSynchronizationClusterDSTStatusEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterDSTStatusEventFrom constructs a [MTRTimeSynchronizationClusterDSTStatusEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDSTStatusEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDSTStatusEvent {
	return MTRTimeSynchronizationClusterDSTStatusEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDSTStatusEventClass) Alloc() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterDSTStatusEventClass) New() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) Init() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) Autorelease() MTRTimeSynchronizationClusterDSTStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterDSTStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDSTStatusEvent creates a new MTRTimeSynchronizationClusterDSTStatusEvent instance.
func NewMTRTimeSynchronizationClusterDSTStatusEvent() MTRTimeSynchronizationClusterDSTStatusEvent {
	return getMTRTimeSynchronizationClusterDSTStatusEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent/dstOffsetActive
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) DstOffsetActive() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dstOffsetActive"))
	return rv
}


// SetDstOffsetActive sets the value of the dstOffsetActive property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDSTStatusEvent/dstOffsetActive
func (m_ MTRTimeSynchronizationClusterDSTStatusEvent) SetDstOffsetActive(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffsetActive:"), value)
}



