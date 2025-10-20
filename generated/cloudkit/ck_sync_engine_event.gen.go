// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineEvent] class.
var (
	CKSyncEngineEventClass     _CKSyncEngineEventClass
	CKSyncEngineEventClassOnce sync.Once
)

func getCKSyncEngineEventClass() _CKSyncEngineEventClass {
	CKSyncEngineEventClassOnce.Do(func() {
		CKSyncEngineEventClass = _CKSyncEngineEventClass{objc.GetClass("CKSyncEngineEvent")}
	})
	return CKSyncEngineEventClass
}

type _CKSyncEngineEventClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineEvent] class.
type ICKSyncEngineEvent interface {
	objectivec.IObject
}

// An event that occurs during a sync operation.
//
// All sync operation events descend from this base class, and as such you don’t create instances of it directly. Instead, the sync engine dispatches them to your app’s delegate, periodically, throughout a sync operation. Use the property to determine the event’s proper type, and then use the corresponding convenience property to retrieve a reference to the event that’s downcast to the appropriate subclass. For example, when is set to , use the property to get the downcast reference .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent
type CKSyncEngineEvent struct {
	objectivec.Object
}

// CKSyncEngineEventFrom constructs a [CKSyncEngineEvent] from an unsafe.Pointer.
//
// An event that occurs during a sync operation.
func CKSyncEngineEventFrom(ptr unsafe.Pointer) CKSyncEngineEvent {
	return CKSyncEngineEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineEventClass) Alloc() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineEventClass) New() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineEvent) Init() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineEvent) Autorelease() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineEvent creates a new CKSyncEngineEvent instance.
func NewCKSyncEngineEvent() CKSyncEngineEvent {
	return getCKSyncEngineEventClass().New()
}


// The event downcast to the subclass that represents a change to the device’s iCloud account.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/accountChangeEvent
func (c_ CKSyncEngineEvent) AccountChangeEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("accountChangeEvent"))
	return rv
}

// The event downcast to the subclass that represents a completed database fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchChangesEvent
func (c_ CKSyncEngineEvent) DidFetchChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("didFetchChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a completed record zone fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) DidFetchRecordZoneChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("didFetchRecordZoneChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a completed send operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didSendChangesEvent
func (c_ CKSyncEngineEvent) DidSendChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("didSendChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a set of fetched database changes to process.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedDatabaseChangesEvent
func (c_ CKSyncEngineEvent) FetchedDatabaseChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchedDatabaseChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a set of fetched record zone changes to process.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) FetchedRecordZoneChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchedRecordZoneChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a sent batch of database changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentDatabaseChangesEvent
func (c_ CKSyncEngineEvent) SentDatabaseChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sentDatabaseChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents a sent batch of record zone changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) SentRecordZoneChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sentRecordZoneChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents an update to the sync engine’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/stateUpdateEvent
func (c_ CKSyncEngineEvent) StateUpdateEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("stateUpdateEvent"))
	return rv
}

// The type of event.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/type
func (c_ CKSyncEngineEvent) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}

// The event downcast to the subclass that represents an imminent database fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchChangesEvent
func (c_ CKSyncEngineEvent) WillFetchChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("willFetchChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents an imminent fetch of record zone changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) WillFetchRecordZoneChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("willFetchRecordZoneChangesEvent"))
	return rv
}

// The event downcast to the subclass that represents an imminent send operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willSendChangesEvent
func (c_ CKSyncEngineEvent) WillSendChangesEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("willSendChangesEvent"))
	return rv
}



