// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterLockOperationEvent] class.
var (
	MTRDoorLockClusterLockOperationEventClass     _MTRDoorLockClusterLockOperationEventClass
	MTRDoorLockClusterLockOperationEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockOperationEventClass() _MTRDoorLockClusterLockOperationEventClass {
	MTRDoorLockClusterLockOperationEventClassOnce.Do(func() {
		MTRDoorLockClusterLockOperationEventClass = _MTRDoorLockClusterLockOperationEventClass{objc.GetClass("MTRDoorLockClusterLockOperationEvent")}
	})
	return MTRDoorLockClusterLockOperationEventClass
}

type _MTRDoorLockClusterLockOperationEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterLockOperationEvent] class.
type IMTRDoorLockClusterLockOperationEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationEvent
type MTRDoorLockClusterLockOperationEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationEventFrom constructs a [MTRDoorLockClusterLockOperationEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationEvent {
	return MTRDoorLockClusterLockOperationEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationEventClass) Alloc() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterLockOperationEventClass) New() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockOperationEvent) Init() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockOperationEvent) Autorelease() MTRDoorLockClusterLockOperationEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockOperationEvent creates a new MTRDoorLockClusterLockOperationEvent instance.
func NewMTRDoorLockClusterLockOperationEvent() MTRDoorLockClusterLockOperationEvent {
	return getMTRDoorLockClusterLockOperationEventClass().New()
}




