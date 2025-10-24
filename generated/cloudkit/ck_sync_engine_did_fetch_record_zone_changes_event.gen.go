// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
)

/* debug [class.gen.go]: Generating class CKSyncEngineDidFetchRecordZoneChangesEvent */


/* debug [class_header]: Header for CKSyncEngineDidFetchRecordZoneChangesEvent */
// The class instance for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
var (
	CKSyncEngineDidFetchRecordZoneChangesEventClass     _CKSyncEngineDidFetchRecordZoneChangesEventClass
	CKSyncEngineDidFetchRecordZoneChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchRecordZoneChangesEventClass() _CKSyncEngineDidFetchRecordZoneChangesEventClass {
	CKSyncEngineDidFetchRecordZoneChangesEventClassOnce.Do(func() {
		CKSyncEngineDidFetchRecordZoneChangesEventClass = _CKSyncEngineDidFetchRecordZoneChangesEventClass{objc.GetClass("CKSyncEngineDidFetchRecordZoneChangesEvent")}
	})
	return CKSyncEngineDidFetchRecordZoneChangesEventClass
}

type _CKSyncEngineDidFetchRecordZoneChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineDidFetchRecordZoneChangesEvent */
// An interface definition for the [CKSyncEngineDidFetchRecordZoneChangesEvent] class.
type ICKSyncEngineDidFetchRecordZoneChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineDidFetchRecordZoneChangesEvent */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	ZoneID() ICKRecordZoneID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineDidFetchRecordZoneChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineDidFetchRecordZoneChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidFetchRecordZoneChangesEventClass) Alloc() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineDidFetchRecordZoneChangesEventClass) New() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Init() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Autorelease() CKSyncEngineDidFetchRecordZoneChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchRecordZoneChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchRecordZoneChangesEvent creates a new CKSyncEngineDidFetchRecordZoneChangesEvent instance.
func NewCKSyncEngineDidFetchRecordZoneChangesEvent() CKSyncEngineDidFetchRecordZoneChangesEvent {
	return getCKSyncEngineDidFetchRecordZoneChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineDidFetchRecordZoneChangesEvent */
// An object that provides information about a finished record zone fetch.


// An object that provides information about a finished record zone fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent
type CKSyncEngineDidFetchRecordZoneChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchRecordZoneChangesEventFrom constructs a [CKSyncEngineDidFetchRecordZoneChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a finished record zone fetch.
func CKSyncEngineDidFetchRecordZoneChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidFetchRecordZoneChangesEvent {
	return CKSyncEngineDidFetchRecordZoneChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineDidFetchRecordZoneChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineDidFetchRecordZoneChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineDidFetchRecordZoneChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineDidFetchRecordZoneChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineDidFetchRecordZoneChangesEvent */

// An error that describes the cause of a failed fetch operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/error
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](c_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The associated record zone’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchRecordZoneChangesEvent/zoneID
func (c_ CKSyncEngineDidFetchRecordZoneChangesEvent) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineDidFetchRecordZoneChangesEvent */



