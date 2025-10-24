// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineWillFetchRecordZoneChangesEvent */


/* debug [class_header]: Header for CKSyncEngineWillFetchRecordZoneChangesEvent */
// The class instance for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
var (
	CKSyncEngineWillFetchRecordZoneChangesEventClass     _CKSyncEngineWillFetchRecordZoneChangesEventClass
	CKSyncEngineWillFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchRecordZoneChangesEventClass() _CKSyncEngineWillFetchRecordZoneChangesEventClass {
	CKSyncEngineWillFetchRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineWillFetchRecordZoneChangesEventClass = _CKSyncEngineWillFetchRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineWillFetchRecordZoneChangesEvent")}
	})
	return CKSyncEngineWillFetchRecordZoneChangesEventClass
}

type _CKSyncEngineWillFetchRecordZoneChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineWillFetchRecordZoneChangesEvent */
// An interface definition for the [CKSyncEngineWillFetchRecordZoneChangesEvent] class.
type ICKSyncEngineWillFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineWillFetchRecordZoneChangesEvent */
	// properties:
	ZoneID() ICKRecordZoneID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineWillFetchRecordZoneChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineWillFetchRecordZoneChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineWillFetchRecordZoneChangesEventClass) New() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) Init() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineWillFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchRecordZoneChangesEvent creates a new CKSyncEngineWillFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineWillFetchRecordZoneChangesEvent() CKSyncEngineWillFetchRecordZoneChangesEvent {
	return getCKSyncEngineWillFetchRecordZoneChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineWillFetchRecordZoneChangesEvent */
// An object that provides information about an imminent fetch of changes in a record zone.


// An object that provides information about an imminent fetch of changes in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent
type CKSyncEngineWillFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchRecordZoneChangesEventFrom constructs a [CKSyncEngineWillFetchRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about an imminent fetch of changes in a record zone.
func CKSyncEngineWillFetchRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillFetchRecordZoneChangesEvent {
	return CKSyncEngineWillFetchRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineWillFetchRecordZoneChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineWillFetchRecordZoneChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineWillFetchRecordZoneChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineWillFetchRecordZoneChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineWillFetchRecordZoneChangesEvent */

// The associated record zone’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchRecordZoneChangesEvent/zoneID
func (c_ CKSyncEngineWillFetchRecordZoneChangesEvent) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineWillFetchRecordZoneChangesEvent */



