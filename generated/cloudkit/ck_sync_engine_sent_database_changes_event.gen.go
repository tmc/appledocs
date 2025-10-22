// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEngineSentDatabaseChangesEvent] class.
var (
	CKSyncEngineSentDatabaseChangesEventClass     _CKSyncEngineSentDatabaseChangesEventClass
	CKSyncEngineSentDatabaseChangesEventClassOnce sync.Once
)

func getCKSyncEngineSentDatabaseChangesEventClass() _CKSyncEngineSentDatabaseChangesEventClass {
	CKSyncEngineSentDatabaseChangesEventClassOnce.Do(func() {
		CKSyncEngineSentDatabaseChangesEventClass = _CKSyncEngineSentDatabaseChangesEventClass{objc.GetClass("CKSyncEngineSentDatabaseChangesEvent")}
	})
	return CKSyncEngineSentDatabaseChangesEventClass
}

type _CKSyncEngineSentDatabaseChangesEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineSentDatabaseChangesEvent] class.
type ICKSyncEngineSentDatabaseChangesEvent interface {
	ICKSyncEngineEvent
}

// An object that provides information about a sent batch of database changes.


// An object that provides information about a sent batch of database changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent

type CKSyncEngineSentDatabaseChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineSentDatabaseChangesEventFrom constructs a [CKSyncEngineSentDatabaseChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a sent batch of database changes.
func CKSyncEngineSentDatabaseChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineSentDatabaseChangesEvent {
	return CKSyncEngineSentDatabaseChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSentDatabaseChangesEventClass) Alloc() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineSentDatabaseChangesEventClass) New() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSentDatabaseChangesEvent) Init() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSentDatabaseChangesEvent) Autorelease() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSentDatabaseChangesEvent creates a new CKSyncEngineSentDatabaseChangesEvent instance.
func NewCKSyncEngineSentDatabaseChangesEvent() CKSyncEngineSentDatabaseChangesEvent {
	return getCKSyncEngineSentDatabaseChangesEventClass().New()
}




