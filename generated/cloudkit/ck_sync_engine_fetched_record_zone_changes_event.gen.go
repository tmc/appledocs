// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchedRecordZoneChangesEvent */


/* debug [class_header]: Header for CKSyncEngineFetchedRecordZoneChangesEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchedRecordZoneChangesEvent */
// An interface definition for the [CKSyncEngineFetchedRecordZoneChangesEvent] class.
type ICKSyncEngineFetchedRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchedRecordZoneChangesEvent */
	// properties:
	Deletions() []CKSyncEngineFetchedRecordDeletion
	Modifications() []objc.IObject /* cross-framework: CKRecord */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchedRecordZoneChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchedRecordZoneChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedRecordZoneChangesEventClass) Alloc() CKSyncEngineFetchedRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchedRecordZoneChangesEvent */
// An object that provides information about fetched record zone changes.


// An object that provides information about fetched record zone changes.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchedRecordZoneChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchedRecordZoneChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchedRecordZoneChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchedRecordZoneChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchedRecordZoneChangesEvent */

// The deleted records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent/deletions
func (c_ CKSyncEngineFetchedRecordZoneChangesEvent) Deletions() []CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[[]CKSyncEngineFetchedRecordDeletion](c_.ID, objc.Sel("deletions"))
	return rv
}/* debug [instance_properties/getter]: deletions */


// The modified records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordZoneChangesEvent/modifications
func (c_ CKSyncEngineFetchedRecordZoneChangesEvent) Modifications() []objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[[]CKRecord](c_.ID, objc.Sel("modifications"))
	return rv
}/* debug [instance_properties/getter]: modifications */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchedRecordZoneChangesEvent */



