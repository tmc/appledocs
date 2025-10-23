// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineStateUpdateEvent] class.
var (
	CKSyncEngineStateUpdateEventClass     _CKSyncEngineStateUpdateEventClass
	CKSyncEngineStateUpdateEventClassOnce sync.Once
)

func getCKSyncEngineStateUpdateEventClass() _CKSyncEngineStateUpdateEventClass {
	CKSyncEngineStateUpdateEventClassOnce.Do(func() {
		CKSyncEngineStateUpdateEventClass = _CKSyncEngineStateUpdateEventClass{objc.GetClass("CKSyncEngineStateUpdateEvent")}
	})
	return CKSyncEngineStateUpdateEventClass
}

type _CKSyncEngineStateUpdateEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineStateUpdateEvent] class.
type ICKSyncEngineStateUpdateEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about an update to the sync engine’s state.


// An object that provides information about an update to the sync engine’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent
type CKSyncEngineStateUpdateEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineStateUpdateEventFrom constructs a [CKSyncEngineStateUpdateEvent] from an unsafe.Pointer.
//
// An object that provides information about an update to the sync engine’s state.
func CKSyncEngineStateUpdateEventFrom(ptr unsafe.Pointer) CKSyncEngineStateUpdateEvent {
	return CKSyncEngineStateUpdateEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateUpdateEventClass) Alloc() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineStateUpdateEventClass) New() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineStateUpdateEvent) Init() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineStateUpdateEvent) Autorelease() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateUpdateEvent creates a new CKSyncEngineStateUpdateEvent instance.
func NewCKSyncEngineStateUpdateEvent() CKSyncEngineStateUpdateEvent {
	return getCKSyncEngineStateUpdateEventClass().New()
}




