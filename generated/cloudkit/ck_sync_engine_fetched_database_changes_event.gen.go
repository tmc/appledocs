// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineFetchedDatabaseChangesEvent] class.
var (
	CKSyncEngineFetchedDatabaseChangesEventClass     _CKSyncEngineFetchedDatabaseChangesEventClass
	CKSyncEngineFetchedDatabaseChangesEventClassOnce sync.Once
)

func getCKSyncEngineFetchedDatabaseChangesEventClass() _CKSyncEngineFetchedDatabaseChangesEventClass {
	CKSyncEngineFetchedDatabaseChangesEventClassOnce.Do(func() {
		CKSyncEngineFetchedDatabaseChangesEventClass = _CKSyncEngineFetchedDatabaseChangesEventClass{objc.GetClass("CKSyncEngineFetchedDatabaseChangesEvent")}
	})
	return CKSyncEngineFetchedDatabaseChangesEventClass
}

type _CKSyncEngineFetchedDatabaseChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchedDatabaseChangesEvent] class.
type ICKSyncEngineFetchedDatabaseChangesEvent interface {
	ICKSyncEngineEvent
	// properties:
	// methods:
}

// An object that provides information about fetched database changes.


// An object that provides information about fetched database changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent
type CKSyncEngineFetchedDatabaseChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineFetchedDatabaseChangesEventFrom constructs a [CKSyncEngineFetchedDatabaseChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about fetched database changes.
func CKSyncEngineFetchedDatabaseChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineFetchedDatabaseChangesEvent {
	return CKSyncEngineFetchedDatabaseChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedDatabaseChangesEventClass) Alloc() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchedDatabaseChangesEventClass) New() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedDatabaseChangesEvent) Init() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedDatabaseChangesEvent) Autorelease() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedDatabaseChangesEvent creates a new CKSyncEngineFetchedDatabaseChangesEvent instance.
func NewCKSyncEngineFetchedDatabaseChangesEvent() CKSyncEngineFetchedDatabaseChangesEvent {
	return getCKSyncEngineFetchedDatabaseChangesEventClass().New()
}




