// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineAccountChangeEvent */


/* debug [class_header]: Header for CKSyncEngineAccountChangeEvent */
// The class instance for the [CKSyncEngineAccountChangeEvent] class.
var (
	CKSyncEngineAccountChangeEventClass     _CKSyncEngineAccountChangeEventClass
	CKSyncEngineAccountChangeEventClassOnce sync.Once
)

func getCKSyncEngineAccountChangeEventClass() _CKSyncEngineAccountChangeEventClass {
	CKSyncEngineAccountChangeEventClassOnce.Do(func() {
		CKSyncEngineAccountChangeEventClass = _CKSyncEngineAccountChangeEventClass{objc.GetClass("CKSyncEngineAccountChangeEvent")}
	})
	return CKSyncEngineAccountChangeEventClass
}

type _CKSyncEngineAccountChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineAccountChangeEvent */
// An interface definition for the [CKSyncEngineAccountChangeEvent] class.
type ICKSyncEngineAccountChangeEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineAccountChangeEvent */
	// properties:
	ChangeType() CKSyncEngineAccountChangeType
	CurrentUser() ICKRecordID
	PreviousUser() ICKRecordID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineAccountChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineAccountChangeEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineAccountChangeEventClass) Alloc() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineAccountChangeEventClass) New() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineAccountChangeEvent) Init() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineAccountChangeEvent) Autorelease() CKSyncEngineAccountChangeEvent {
	rv := objc.Send[CKSyncEngineAccountChangeEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineAccountChangeEvent creates a new CKSyncEngineAccountChangeEvent instance.
func NewCKSyncEngineAccountChangeEvent() CKSyncEngineAccountChangeEvent {
	return getCKSyncEngineAccountChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineAccountChangeEvent */
// An event that provides information about a change to the device’s iCloud account.


// An event that provides information about a change to the device’s iCloud account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent
type CKSyncEngineAccountChangeEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineAccountChangeEventFrom constructs a [CKSyncEngineAccountChangeEvent] from an unsafe.Pointer.
//
// An event that provides information about a change to the device’s iCloud account.
func CKSyncEngineAccountChangeEventFrom(ptr unsafe.Pointer) CKSyncEngineAccountChangeEvent {
	return CKSyncEngineAccountChangeEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineAccountChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineAccountChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineAccountChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineAccountChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineAccountChangeEvent */

// The iCloud account’s change type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/changeType
func (c_ CKSyncEngineAccountChangeEvent) ChangeType() CKSyncEngineAccountChangeType {
	rv := objc.Send[CKSyncEngineAccountChangeType](c_.ID, objc.Sel("changeType"))
	return rv
}/* debug [instance_properties/getter]: changeType */


// The current iCloud account’s record identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/currentUser
func (c_ CKSyncEngineAccountChangeEvent) CurrentUser() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("currentUser"))
	return rv
}/* debug [instance_properties/getter]: currentUser */


// The previous iCloud account’s record identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineAccountChangeEvent/previousUser
func (c_ CKSyncEngineAccountChangeEvent) PreviousUser() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("previousUser"))
	return rv
}/* debug [instance_properties/getter]: previousUser */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineAccountChangeEvent */



