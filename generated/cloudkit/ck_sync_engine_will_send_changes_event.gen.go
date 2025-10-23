// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineWillSendChangesEvent] class.
var (
	CKSyncEngineWillSendChangesEventClass     _CKSyncEngineWillSendChangesEventClass
	CKSyncEngineWillSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillSendChangesEventClass() _CKSyncEngineWillSendChangesEventClass {
	CKSyncEngineWillSendChangesEventClassOnce.Do(func() {
		CKSyncEngineWillSendChangesEventClass = _CKSyncEngineWillSendChangesEventClass{objc.GetClass("CKSyncEngineWillSendChangesEvent")}
	})
	return CKSyncEngineWillSendChangesEventClass
}

type _CKSyncEngineWillSendChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineWillSendChangesEvent] class.
type ICKSyncEngineWillSendChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about an imminent send of local changes.


// An object that provides information about an imminent send of local changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent
type CKSyncEngineWillSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillSendChangesEventFrom constructs a [CKSyncEngineWillSendChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about an imminent send of local changes.
func CKSyncEngineWillSendChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillSendChangesEvent {
	return CKSyncEngineWillSendChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillSendChangesEventClass) Alloc() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineWillSendChangesEventClass) New() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillSendChangesEvent) Init() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillSendChangesEvent) Autorelease() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillSendChangesEvent creates a new CKSyncEngineWillSendChangesEvent instance.
func NewCKSyncEngineWillSendChangesEvent() CKSyncEngineWillSendChangesEvent {
	return getCKSyncEngineWillSendChangesEventClass().New()
}




