// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CKSyncEngineSentRecordZoneChangesEvent */


/* debug [class_header]: Header for CKSyncEngineSentRecordZoneChangesEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineSentRecordZoneChangesEvent */
// An interface definition for the [CKSyncEngineSentRecordZoneChangesEvent] class.
type ICKSyncEngineSentRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineSentRecordZoneChangesEvent */
	// properties:
	DeletedRecordIDs() []CKRecordID
	FailedRecordDeletes() foundation.IDictionary
	FailedRecordSaves() []CKSyncEngineFailedRecordSave
	SavedRecords() []objc.IObject /* cross-framework: CKRecord */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineSentRecordZoneChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineSentRecordZoneChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSentRecordZoneChangesEventClass) Alloc() CKSyncEngineSentRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineSentRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineSentRecordZoneChangesEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineSentRecordZoneChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineSentRecordZoneChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineSentRecordZoneChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineSentRecordZoneChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineSentRecordZoneChangesEvent */

// The unique identifiers of the deleted records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/deletedRecordIDs
func (c_ CKSyncEngineSentRecordZoneChangesEvent) DeletedRecordIDs() []CKRecordID {
	rv := objc.Send[[]CKRecordID](c_.ID, objc.Sel("deletedRecordIDs"))
	return rv
}/* debug [instance_properties/getter]: deletedRecordIDs */


// The unique identifiers of the records CloudKit is unable to delete, and the reasons why.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/failedRecordDeletes
func (c_ CKSyncEngineSentRecordZoneChangesEvent) FailedRecordDeletes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("failedRecordDeletes"))
	return rv
}/* debug [instance_properties/getter]: failedRecordDeletes */


// The records that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/failedRecordSaves
func (c_ CKSyncEngineSentRecordZoneChangesEvent) FailedRecordSaves() []CKSyncEngineFailedRecordSave {
	rv := objc.Send[[]CKSyncEngineFailedRecordSave](c_.ID, objc.Sel("failedRecordSaves"))
	return rv
}/* debug [instance_properties/getter]: failedRecordSaves */


// The modified records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSentRecordZoneChangesEvent/savedRecords
func (c_ CKSyncEngineSentRecordZoneChangesEvent) SavedRecords() []objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[[]CKRecord](c_.ID, objc.Sel("savedRecords"))
	return rv
}/* debug [instance_properties/getter]: savedRecords */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineSentRecordZoneChangesEvent */



