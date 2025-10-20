// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineFetchedRecordZoneChangesEvent] class.
var (
	CKSyncEngineFetchedRecordZoneChangesEventClass     _CKSyncEngineFetchedRecordZoneChangesEventClass
	CKSyncEngineFetchedRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineFetchedRecordZoneChangesEventClass() _CKSyncEngineFetchedRecordZoneChangesEventClass {
	CKSyncEngineFetchedRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineFetchedRecordZoneChangesEventClass = _CKSyncEngineFetchedRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineFetchedRecordZoneChangesEvent")}
	})
	return CKSyncEngineFetchedRecordZoneChangesEventClass
}

type _CKSyncEngineFetchedRecordZoneChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchedRecordZoneChangesEvent] class.
type ICKSyncEngineFetchedRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about fetched record zone changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent
type CKSyncEngineFetchedRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineFetchedRecordZoneChangesEventFrom constructs a [CKSyncEngineFetchedRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about fetched record zone changes.
func CKSyncEngineFetchedRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineFetchedRecordZoneChangesEvent {
	return CKSyncEngineFetchedRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedRecordZoneChangesEventClass) Alloc() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchedRecordZoneChangesEventClass) New() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedRecordZoneChangesEvent) Init() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedRecordZoneChangesEvent) Autorelease() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedRecordZoneChangesEvent creates a new CKSyncEngineFetchedRecordZoneChangesEvent instance.
func NewCKSyncEngineFetchedRecordZoneChangesEvent() CKSyncEngineFetchedRecordZoneChangesEvent {
	return getCKSyncEngineFetchedRecordZoneChangesEventClass().New()
}




