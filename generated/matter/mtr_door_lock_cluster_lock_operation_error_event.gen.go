// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterLockOperationErrorEvent] class.
var (
	MTRDoorLockClusterLockOperationErrorEventClass     _MTRDoorLockClusterLockOperationErrorEventClass
	MTRDoorLockClusterLockOperationErrorEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockOperationErrorEventClass() _MTRDoorLockClusterLockOperationErrorEventClass {
	MTRDoorLockClusterLockOperationErrorEventClassOnce.Do(func() {
		MTRDoorLockClusterLockOperationErrorEventClass = _MTRDoorLockClusterLockOperationErrorEventClass{objc.GetClass("MTRDoorLockClusterLockOperationErrorEvent")}
	})
	return MTRDoorLockClusterLockOperationErrorEventClass
}

type _MTRDoorLockClusterLockOperationErrorEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterLockOperationErrorEvent] class.
type IMTRDoorLockClusterLockOperationErrorEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockOperationErrorEvent
type MTRDoorLockClusterLockOperationErrorEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockOperationErrorEventFrom constructs a [MTRDoorLockClusterLockOperationErrorEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockOperationErrorEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockOperationErrorEvent {
	return MTRDoorLockClusterLockOperationErrorEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockOperationErrorEventClass) Alloc() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterLockOperationErrorEventClass) New() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Init() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockOperationErrorEvent) Autorelease() MTRDoorLockClusterLockOperationErrorEvent {
	rv := objc.Send[MTRDoorLockClusterLockOperationErrorEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockOperationErrorEvent creates a new MTRDoorLockClusterLockOperationErrorEvent instance.
func NewMTRDoorLockClusterLockOperationErrorEvent() MTRDoorLockClusterLockOperationErrorEvent {
	return getMTRDoorLockClusterLockOperationErrorEventClass().New()
}




