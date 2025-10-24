// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CKSyncEngineSentDatabaseChangesEvent */


/* debug [class_header]: Header for CKSyncEngineSentDatabaseChangesEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineSentDatabaseChangesEvent */
// An interface definition for the [CKSyncEngineSentDatabaseChangesEvent] class.
type ICKSyncEngineSentDatabaseChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineSentDatabaseChangesEvent */
	// properties:
	DeletedZoneIDs() []CKRecordZoneID
	FailedZoneDeletes() foundation.IDictionary
	FailedZoneSaves() []CKSyncEngineFailedZoneSave
	SavedZones() []CKRecordZone
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineSentDatabaseChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineSentDatabaseChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSentDatabaseChangesEventClass) Alloc() CKSyncEngineSentDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineSentDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineSentDatabaseChangesEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineSentDatabaseChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineSentDatabaseChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineSentDatabaseChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineSentDatabaseChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineSentDatabaseChangesEvent */

// The unique identifiers of the deleted record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/deletedZoneIDs
func (c_ CKSyncEngineSentDatabaseChangesEvent) DeletedZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("deletedZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: deletedZoneIDs */


// The unique identifiers of the record zones CloudKit is unable to delete, and the reasons why.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/failedZoneDeletes
func (c_ CKSyncEngineSentDatabaseChangesEvent) FailedZoneDeletes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("failedZoneDeletes"))
	return rv
}/* debug [instance_properties/getter]: failedZoneDeletes */


// The record zones that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/failedZoneSaves
func (c_ CKSyncEngineSentDatabaseChangesEvent) FailedZoneSaves() []CKSyncEngineFailedZoneSave {
	rv := objc.Send[[]CKSyncEngineFailedZoneSave](c_.ID, objc.Sel("failedZoneSaves"))
	return rv
}/* debug [instance_properties/getter]: failedZoneSaves */


// The modified record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentDatabaseChangesEvent/savedZones
func (c_ CKSyncEngineSentDatabaseChangesEvent) SavedZones() []CKRecordZone {
	rv := objc.Send[[]CKRecordZone](c_.ID, objc.Sel("savedZones"))
	return rv
}/* debug [instance_properties/getter]: savedZones */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineSentDatabaseChangesEvent */



