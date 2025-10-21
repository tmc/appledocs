// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
var (
	CKSyncEngineWillFetchRecordZoneChangesEventClass     _CKSyncEngineWillFetchRecordZoneChangesEventClass
	CKSyncEngineWillFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchRecordZoneChangesEventClass() _CKSyncEngineWillFetchRecordZoneChangesEventClass {
	CKSyncEngineWillFetchRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineWillFetchRecordZoneChangesEventClass = _CKSyncEngineWillFetchRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineWillFetchRecordZoneChangesEvent")}
	})
	return CKSyncEngineWillFetchRecordZoneChangesEventClass
}

type _CKSyncEngineWillFetchRecordZoneChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
type ICKSyncEngineWillFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about an imminent fetch of changes in a record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent
type CKSyncEngineWillFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchRecordZoneChangesEventFrom constructs a [CKSyncEngineWillFetchRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about an imminent fetch of changes in a record zone.
func CKSyncEngineWillFetchRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillFetchRecordZoneChangesEvent {
	return CKSyncEngineWillFetchRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineWillFetchRecordZoneChangesEventClass) New() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) Init() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchRecordZoneChangesEvent creates a new CKSyncEngineWillFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineWillFetchRecordZoneChangesEvent() CKSyncEngineWillFetchRecordZoneChangesEvent {
	return getCKSyncEngineWillFetchRecordZoneChangesEventClass().New()
}


// The associated record zone’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent/zoneID
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) ZoneID() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}



