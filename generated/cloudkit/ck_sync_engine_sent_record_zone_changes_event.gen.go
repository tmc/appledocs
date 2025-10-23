// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineSentRecordZoneChangesEvent] class.
var (
	CKSyncEngineSentRecordZoneChangesEventClass     _CKSyncEngineSentRecordZoneChangesEventClass
	CKSyncEngineSentRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineSentRecordZoneChangesEventClass() _CKSyncEngineSentRecordZoneChangesEventClass {
	CKSyncEngineSentRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineSentRecordZoneChangesEventClass = _CKSyncEngineSentRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineSentRecordZoneChangesEvent")}
	})
	return CKSyncEngineSentRecordZoneChangesEventClass
}

type _CKSyncEngineSentRecordZoneChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineSentRecordZoneChangesEvent] class.
type ICKSyncEngineSentRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about a sent batch of record zone changes.


// An object that provides information about a sent batch of record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent
type CKSyncEngineSentRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineSentRecordZoneChangesEventFrom constructs a [CKSyncEngineSentRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a sent batch of record zone changes.
func CKSyncEngineSentRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineSentRecordZoneChangesEvent {
	return CKSyncEngineSentRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSentRecordZoneChangesEventClass) Alloc() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineSentRecordZoneChangesEventClass) New() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSentRecordZoneChangesEvent) Init() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSentRecordZoneChangesEvent) Autorelease() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSentRecordZoneChangesEvent creates a new CKSyncEngineSentRecordZoneChangesEvent instance.
func NewCKSyncEngineSentRecordZoneChangesEvent() CKSyncEngineSentRecordZoneChangesEvent {
	return getCKSyncEngineSentRecordZoneChangesEventClass().New()
}




