// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchedDatabaseChangesEvent */


/* debug [class_header]: Header for CKSyncEngineFetchedDatabaseChangesEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchedDatabaseChangesEvent */
// An interface definition for the [CKSyncEngineFetchedDatabaseChangesEvent] class.
type ICKSyncEngineFetchedDatabaseChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchedDatabaseChangesEvent */
	// properties:
	Deletions() []CKSyncEngineFetchedZoneDeletion
	Modifications() []CKRecordZone
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchedDatabaseChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchedDatabaseChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedDatabaseChangesEventClass) Alloc() CKSyncEngineFetchedDatabaseChangesEvent {
	rv := objc.Send[CKSyncEngineFetchedDatabaseChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchedDatabaseChangesEvent */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchedDatabaseChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchedDatabaseChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchedDatabaseChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchedDatabaseChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchedDatabaseChangesEvent */

// The deleted record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent/deletions
func (c_ CKSyncEngineFetchedDatabaseChangesEvent) Deletions() []CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[[]CKSyncEngineFetchedZoneDeletion](c_.ID, objc.Sel("deletions"))
	return rv
}/* debug [instance_properties/getter]: deletions */


// The modified record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedDatabaseChangesEvent/modifications
func (c_ CKSyncEngineFetchedDatabaseChangesEvent) Modifications() []CKRecordZone {
	rv := objc.Send[[]CKRecordZone](c_.ID, objc.Sel("modifications"))
	return rv
}/* debug [instance_properties/getter]: modifications */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchedDatabaseChangesEvent */



