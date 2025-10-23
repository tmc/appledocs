// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineDidSendChangesEvent] class.
var (
	CKSyncEngineDidSendChangesEventClass     _CKSyncEngineDidSendChangesEventClass
	CKSyncEngineDidSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidSendChangesEventClass() _CKSyncEngineDidSendChangesEventClass {
	CKSyncEngineDidSendChangesEventClassOnce.Do(func() {
		CKSyncEngineDidSendChangesEventClass = _CKSyncEngineDidSendChangesEventClass{objc.GetClass("CKSyncEngineDidSendChangesEvent")}
	})
	return CKSyncEngineDidSendChangesEventClass
}

type _CKSyncEngineDidSendChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineDidSendChangesEvent] class.
type ICKSyncEngineDidSendChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about a finished send operation.


// An object that provides information about a finished send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent
type CKSyncEngineDidSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidSendChangesEventFrom constructs a [CKSyncEngineDidSendChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a finished send operation.
func CKSyncEngineDidSendChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidSendChangesEvent {
	return CKSyncEngineDidSendChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidSendChangesEventClass) Alloc() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineDidSendChangesEventClass) New() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidSendChangesEvent) Init() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidSendChangesEvent) Autorelease() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidSendChangesEvent creates a new CKSyncEngineDidSendChangesEvent instance.
func NewCKSyncEngineDidSendChangesEvent() CKSyncEngineDidSendChangesEvent {
	return getCKSyncEngineDidSendChangesEventClass().New()
}




