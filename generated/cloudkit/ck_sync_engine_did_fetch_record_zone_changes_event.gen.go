// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
var (
	CKSyncEngineDidFetchRecordZoneChangesEventClass     _CKSyncEngineDidFetchRecordZoneChangesEventClass
	CKSyncEngineDidFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchRecordZoneChangesEventClass() _CKSyncEngineDidFetchRecordZoneChangesEventClass {
	CKSyncEngineDidFetchRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineDidFetchRecordZoneChangesEventClass = _CKSyncEngineDidFetchRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineDidFetchRecordZoneChangesEvent")}
	})
	return CKSyncEngineDidFetchRecordZoneChangesEventClass
}

type _CKSyncEngineDidFetchRecordZoneChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
type ICKSyncEngineDidFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
	Error() foundation.Error
	ZoneID() CKRecordZoneID
}

// An object that provides information about a finished record zone fetch.


// An object that provides information about a finished record zone fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent

type CKSyncEngineDidFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchRecordZoneChangesEventFrom constructs a [CKSyncEngineDidFetchRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a finished record zone fetch.
func CKSyncEngineDidFetchRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidFetchRecordZoneChangesEvent {
	return CKSyncEngineDidFetchRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineDidFetchRecordZoneChangesEventClass) New() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Init() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchRecordZoneChangesEvent creates a new CKSyncEngineDidFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineDidFetchRecordZoneChangesEvent() CKSyncEngineDidFetchRecordZoneChangesEvent {
	return getCKSyncEngineDidFetchRecordZoneChangesEventClass().New()
}



// An error that describes the cause of a failed fetch operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/error

func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Error() foundation.Error {
	rv := objc.Send[foundation.Error](c_.ID, objc.Sel("error"))
	return rv
}


// The associated record zone’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/zoneID

func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) ZoneID() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}



