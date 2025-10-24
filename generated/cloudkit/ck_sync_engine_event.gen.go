// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineEvent */


/* debug [class_header]: Header for CKSyncEngineEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineEvent */
// An interface definition for the [CKSyncEngineEvent] class.
type ICKSyncEngineEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineEvent */
	// properties:
	AccountChangeEvent() ICKSyncEngineAccountChangeEvent
	DidFetchChangesEvent() ICKSyncEngineDidFetchChangesEvent
	DidFetchRecordZoneChangesEvent() ICKSyncEngineDidFetchRecordZoneChangesEvent
	DidSendChangesEvent() ICKSyncEngineDidSendChangesEvent
	FetchedDatabaseChangesEvent() ICKSyncEngineFetchedDatabaseChangesEvent
	FetchedRecordZoneChangesEvent() ICKSyncEngineFetchedRecordZoneChangesEvent
	SentDatabaseChangesEvent() ICKSyncEngineSentDatabaseChangesEvent
	SentRecordZoneChangesEvent() ICKSyncEngineSentRecordZoneChangesEvent
	StateUpdateEvent() ICKSyncEngineStateUpdateEvent
	Type() CKSyncEngineEventType
	WillFetchChangesEvent() ICKSyncEngineWillFetchChangesEvent
	WillFetchRecordZoneChangesEvent() ICKSyncEngineWillFetchRecordZoneChangesEvent
	WillSendChangesEvent() ICKSyncEngineWillSendChangesEvent
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineEventClass) Alloc() CKSyncEngineEvent {
	rv := objc.Send[CKSyncEngineEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineEvent */
// An event that occurs during a sync operation.
//
// All sync operation events descend from this base class, and as such you don’t create instances of it directly. Instead, the sync engine dispatches them to your app’s delegate, periodically, throughout a sync operation. Use the property to determine the event’s proper type, and then use the corresponding convenience property to retrieve a reference to the event that’s downcast to the appropriate subclass. For example, when is set to , use the property to get the downcast reference .


// An event that occurs during a sync operation.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineEvent */

// The event downcast to the subclass that represents a change to the device’s iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/accountChangeEvent
func (c_ CKSyncEngineEvent) AccountChangeEvent() ICKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c_.ID, objc.Sel("accountChangeEvent"))
	return rv
}/* debug [instance_properties/getter]: accountChangeEvent */


// The event downcast to the subclass that represents a completed database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchChangesEvent
func (c_ CKSyncEngineEvent) DidFetchChangesEvent() ICKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c_.ID, objc.Sel("didFetchChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: didFetchChangesEvent */


// The event downcast to the subclass that represents a completed record zone fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didFetchRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) DidFetchRecordZoneChangesEvent() ICKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c_.ID, objc.Sel("didFetchRecordZoneChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: didFetchRecordZoneChangesEvent */


// The event downcast to the subclass that represents a completed send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/didSendChangesEvent
func (c_ CKSyncEngineEvent) DidSendChangesEvent() ICKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c_.ID, objc.Sel("didSendChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: didSendChangesEvent */


// The event downcast to the subclass that represents a set of fetched database changes to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedDatabaseChangesEvent
func (c_ CKSyncEngineEvent) FetchedDatabaseChangesEvent() ICKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](c_.ID, objc.Sel("fetchedDatabaseChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: fetchedDatabaseChangesEvent */


// The event downcast to the subclass that represents a set of fetched record zone changes to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/fetchedRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) FetchedRecordZoneChangesEvent() ICKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](c_.ID, objc.Sel("fetchedRecordZoneChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: fetchedRecordZoneChangesEvent */


// The event downcast to the subclass that represents a sent batch of database changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentDatabaseChangesEvent
func (c_ CKSyncEngineEvent) SentDatabaseChangesEvent() ICKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](c_.ID, objc.Sel("sentDatabaseChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: sentDatabaseChangesEvent */


// The event downcast to the subclass that represents a sent batch of record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/sentRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) SentRecordZoneChangesEvent() ICKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](c_.ID, objc.Sel("sentRecordZoneChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: sentRecordZoneChangesEvent */


// The event downcast to the subclass that represents an update to the sync engine’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/stateUpdateEvent
func (c_ CKSyncEngineEvent) StateUpdateEvent() ICKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c_.ID, objc.Sel("stateUpdateEvent"))
	return rv
}/* debug [instance_properties/getter]: stateUpdateEvent */


// The type of event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/type
func (c_ CKSyncEngineEvent) Type() CKSyncEngineEventType {
	rv := objc.Send[CKSyncEngineEventType](c_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The event downcast to the subclass that represents an imminent database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchChangesEvent
func (c_ CKSyncEngineEvent) WillFetchChangesEvent() ICKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c_.ID, objc.Sel("willFetchChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: willFetchChangesEvent */


// The event downcast to the subclass that represents an imminent fetch of record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willFetchRecordZoneChangesEvent
func (c_ CKSyncEngineEvent) WillFetchRecordZoneChangesEvent() ICKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c_.ID, objc.Sel("willFetchRecordZoneChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: willFetchRecordZoneChangesEvent */


// The event downcast to the subclass that represents an imminent send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineEvent/willSendChangesEvent
func (c_ CKSyncEngineEvent) WillSendChangesEvent() ICKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c_.ID, objc.Sel("willSendChangesEvent"))
	return rv
}/* debug [instance_properties/getter]: willSendChangesEvent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineEvent */



