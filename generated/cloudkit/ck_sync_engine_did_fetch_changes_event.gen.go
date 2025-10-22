// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineDidFetchChangesEvent] class.
var (
	CKSyncEngineDidFetchChangesEventClass     _CKSyncEngineDidFetchChangesEventClass
	CKSyncEngineDidFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchChangesEventClass() _CKSyncEngineDidFetchChangesEventClass {
	CKSyncEngineDidFetchChangesEventClassOnce.Do(func() {
		CKSyncEngineDidFetchChangesEventClass = _CKSyncEngineDidFetchChangesEventClass{objc.GetClass("CKSyncEngineDidFetchChangesEvent")}
	})
	return CKSyncEngineDidFetchChangesEventClass
}

type _CKSyncEngineDidFetchChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineDidFetchChangesEvent] class.
type ICKSyncEngineDidFetchChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that represents a completed database fetch.


// An object that represents a completed database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent

type CKSyncEngineDidFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchChangesEventFrom constructs a [CKSyncEngineDidFetchChangesEvent] from an unsafe.Pointer.
//
// An object that represents a completed database fetch.
func CKSyncEngineDidFetchChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidFetchChangesEvent {
	return CKSyncEngineDidFetchChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidFetchChangesEventClass) Alloc() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineDidFetchChangesEventClass) New() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidFetchChangesEvent) Init() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidFetchChangesEvent) Autorelease() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchChangesEvent creates a new CKSyncEngineDidFetchChangesEvent instance.
func NewCKSyncEngineDidFetchChangesEvent() CKSyncEngineDidFetchChangesEvent {
	return getCKSyncEngineDidFetchChangesEventClass().New()
}




