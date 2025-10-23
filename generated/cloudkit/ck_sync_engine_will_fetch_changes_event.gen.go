// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineWillFetchChangesEvent] class.
var (
	CKSyncEngineWillFetchChangesEventClass     _CKSyncEngineWillFetchChangesEventClass
	CKSyncEngineWillFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchChangesEventClass() _CKSyncEngineWillFetchChangesEventClass {
	CKSyncEngineWillFetchChangesEventClassOnce.Do(func() {
		CKSyncEngineWillFetchChangesEventClass = _CKSyncEngineWillFetchChangesEventClass{objc.GetClass("CKSyncEngineWillFetchChangesEvent")}
	})
	return CKSyncEngineWillFetchChangesEventClass
}

type _CKSyncEngineWillFetchChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineWillFetchChangesEvent] class.
type ICKSyncEngineWillFetchChangesEvent interface {
	ICKSyncEngineEvent
	// properties:
	// methods:
}

// An object that represents an imminent database fetch.


// An object that represents an imminent database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent
type CKSyncEngineWillFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchChangesEventFrom constructs a [CKSyncEngineWillFetchChangesEvent] from an unsafe.Pointer.
//
// An object that represents an imminent database fetch.
func CKSyncEngineWillFetchChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillFetchChangesEvent {
	return CKSyncEngineWillFetchChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillFetchChangesEventClass) Alloc() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineWillFetchChangesEventClass) New() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillFetchChangesEvent) Init() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillFetchChangesEvent) Autorelease() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchChangesEvent creates a new CKSyncEngineWillFetchChangesEvent instance.
func NewCKSyncEngineWillFetchChangesEvent() CKSyncEngineWillFetchChangesEvent {
	return getCKSyncEngineWillFetchChangesEventClass().New()
}




