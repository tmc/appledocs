// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchedZoneDeletion */


/* debug [class_header]: Header for CKSyncEngineFetchedZoneDeletion */
// The class instance for the [CKSyncEngineFetchedZoneDeletion] class.
var (
	CKSyncEngineFetchedZoneDeletionClass     _CKSyncEngineFetchedZoneDeletionClass
	CKSyncEngineFetchedZoneDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedZoneDeletionClass() _CKSyncEngineFetchedZoneDeletionClass {
	CKSyncEngineFetchedZoneDeletionClassOnce.Do(func() {
		CKSyncEngineFetchedZoneDeletionClass = _CKSyncEngineFetchedZoneDeletionClass{objc.GetClass("CKSyncEngineFetchedZoneDeletion")}
	})
	return CKSyncEngineFetchedZoneDeletionClass
}

type _CKSyncEngineFetchedZoneDeletionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchedZoneDeletion */
// An interface definition for the [CKSyncEngineFetchedZoneDeletion] class.
type ICKSyncEngineFetchedZoneDeletion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchedZoneDeletion */
	// properties:
	Reason() CKSyncEngineZoneDeletionReason
	ZoneID() ICKRecordZoneID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchedZoneDeletion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchedZoneDeletion */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedZoneDeletionClass) Alloc() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFetchedZoneDeletionClass) New() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedZoneDeletion) Init() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedZoneDeletion) Autorelease() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedZoneDeletion creates a new CKSyncEngineFetchedZoneDeletion instance.
func NewCKSyncEngineFetchedZoneDeletion() CKSyncEngineFetchedZoneDeletion {
	return getCKSyncEngineFetchedZoneDeletionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchedZoneDeletion */
// An object that describes the deletion of a record zone.


// An object that describes the deletion of a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion
type CKSyncEngineFetchedZoneDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedZoneDeletionFrom constructs a [CKSyncEngineFetchedZoneDeletion] from an unsafe.Pointer.
//
// An object that describes the deletion of a record zone.
func CKSyncEngineFetchedZoneDeletionFrom(ptr unsafe.Pointer) CKSyncEngineFetchedZoneDeletion {
	return CKSyncEngineFetchedZoneDeletion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchedZoneDeletion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchedZoneDeletion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchedZoneDeletion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchedZoneDeletion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchedZoneDeletion */

// The reason for the deletion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion/reason
func (c_ CKSyncEngineFetchedZoneDeletion) Reason() CKSyncEngineZoneDeletionReason {
	rv := objc.Send[CKSyncEngineZoneDeletionReason](c_.ID, objc.Sel("reason"))
	return rv
}/* debug [instance_properties/getter]: reason */


// The identifier of the deleted record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion/zoneID
func (c_ CKSyncEngineFetchedZoneDeletion) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchedZoneDeletion */



