// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterLockUserChangeEvent] class.
var (
	MTRDoorLockClusterLockUserChangeEventClass     _MTRDoorLockClusterLockUserChangeEventClass
	MTRDoorLockClusterLockUserChangeEventClassOnce sync.Once
)

func getMTRDoorLockClusterLockUserChangeEventClass() _MTRDoorLockClusterLockUserChangeEventClass {
	MTRDoorLockClusterLockUserChangeEventClassOnce.Do(func() {
		MTRDoorLockClusterLockUserChangeEventClass = _MTRDoorLockClusterLockUserChangeEventClass{objc.GetClass("MTRDoorLockClusterLockUserChangeEvent")}
	})
	return MTRDoorLockClusterLockUserChangeEventClass
}

type _MTRDoorLockClusterLockUserChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterLockUserChangeEvent] class.
type IMTRDoorLockClusterLockUserChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterLockUserChangeEvent
type MTRDoorLockClusterLockUserChangeEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterLockUserChangeEventFrom constructs a [MTRDoorLockClusterLockUserChangeEvent] from an unsafe.Pointer.
func MTRDoorLockClusterLockUserChangeEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterLockUserChangeEvent {
	return MTRDoorLockClusterLockUserChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterLockUserChangeEventClass) Alloc() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterLockUserChangeEventClass) New() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterLockUserChangeEvent) Init() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterLockUserChangeEvent) Autorelease() MTRDoorLockClusterLockUserChangeEvent {
	rv := objc.Send[MTRDoorLockClusterLockUserChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterLockUserChangeEvent creates a new MTRDoorLockClusterLockUserChangeEvent instance.
func NewMTRDoorLockClusterLockUserChangeEvent() MTRDoorLockClusterLockUserChangeEvent {
	return getMTRDoorLockClusterLockUserChangeEventClass().New()
}




