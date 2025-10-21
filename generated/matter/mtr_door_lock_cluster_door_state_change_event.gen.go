// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterDoorStateChangeEvent] class.
var (
	MTRDoorLockClusterDoorStateChangeEventClass     _MTRDoorLockClusterDoorStateChangeEventClass
	MTRDoorLockClusterDoorStateChangeEventClassOnce sync.Once
)

func getMTRDoorLockClusterDoorStateChangeEventClass() _MTRDoorLockClusterDoorStateChangeEventClass {
	MTRDoorLockClusterDoorStateChangeEventClassOnce.Do(func() {
		MTRDoorLockClusterDoorStateChangeEventClass = _MTRDoorLockClusterDoorStateChangeEventClass{objc.GetClass("MTRDoorLockClusterDoorStateChangeEvent")}
	})
	return MTRDoorLockClusterDoorStateChangeEventClass
}

type _MTRDoorLockClusterDoorStateChangeEventClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterDoorStateChangeEvent] class.
type IMTRDoorLockClusterDoorStateChangeEvent interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorStateChangeEvent
type MTRDoorLockClusterDoorStateChangeEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterDoorStateChangeEventFrom constructs a [MTRDoorLockClusterDoorStateChangeEvent] from an unsafe.Pointer.
func MTRDoorLockClusterDoorStateChangeEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterDoorStateChangeEvent {
	return MTRDoorLockClusterDoorStateChangeEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterDoorStateChangeEventClass) Alloc() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterDoorStateChangeEventClass) New() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterDoorStateChangeEvent) Init() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterDoorStateChangeEvent) Autorelease() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterDoorStateChangeEvent creates a new MTRDoorLockClusterDoorStateChangeEvent instance.
func NewMTRDoorLockClusterDoorStateChangeEvent() MTRDoorLockClusterDoorStateChangeEvent {
	return getMTRDoorLockClusterDoorStateChangeEventClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterdoorstatechangeevent/doorstate
func (m_ MTRDoorLockClusterDoorStateChangeEvent) DoorState() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("doorState"))
	return rv
}


// SetDoorState sets the value of the doorState property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdoorlockclusterdoorstatechangeevent/doorstate
func (m_ MTRDoorLockClusterDoorStateChangeEvent) SetDoorState(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoorState:"), value)
}



