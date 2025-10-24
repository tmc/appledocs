// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineFetchedRecordDeletion */


/* debug [class_header]: Header for CKSyncEngineFetchedRecordDeletion */
// The class instance for the [CKSyncEngineFetchedRecordDeletion] class.
var (
	CKSyncEngineFetchedRecordDeletionClass     _CKSyncEngineFetchedRecordDeletionClass
	CKSyncEngineFetchedRecordDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedRecordDeletionClass() _CKSyncEngineFetchedRecordDeletionClass {
	CKSyncEngineFetchedRecordDeletionClassOnce.Do(func() {
		CKSyncEngineFetchedRecordDeletionClass = _CKSyncEngineFetchedRecordDeletionClass{objc.GetClass("CKSyncEngineFetchedRecordDeletion")}
	})
	return CKSyncEngineFetchedRecordDeletionClass
}

type _CKSyncEngineFetchedRecordDeletionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineFetchedRecordDeletion */
// An interface definition for the [CKSyncEngineFetchedRecordDeletion] class.
type ICKSyncEngineFetchedRecordDeletion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineFetchedRecordDeletion */
	// properties:
	RecordID() ICKRecordID
	RecordType() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineFetchedRecordDeletion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineFetchedRecordDeletion */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedRecordDeletionClass) Alloc() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineFetchedRecordDeletionClass) New() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedRecordDeletion) Init() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedRecordDeletion) Autorelease() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedRecordDeletion creates a new CKSyncEngineFetchedRecordDeletion instance.
func NewCKSyncEngineFetchedRecordDeletion() CKSyncEngineFetchedRecordDeletion {
	return getCKSyncEngineFetchedRecordDeletionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineFetchedRecordDeletion */
// An object that describes the deletion of an individual record.


// An object that describes the deletion of an individual record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion
type CKSyncEngineFetchedRecordDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedRecordDeletionFrom constructs a [CKSyncEngineFetchedRecordDeletion] from an unsafe.Pointer.
//
// An object that describes the deletion of an individual record.
func CKSyncEngineFetchedRecordDeletionFrom(ptr unsafe.Pointer) CKSyncEngineFetchedRecordDeletion {
	return CKSyncEngineFetchedRecordDeletion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineFetchedRecordDeletion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineFetchedRecordDeletion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineFetchedRecordDeletion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineFetchedRecordDeletion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineFetchedRecordDeletion */

// The deleted record’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordID
func (c_ CKSyncEngineFetchedRecordDeletion) RecordID() ICKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}/* debug [instance_properties/getter]: recordID */


// The record type of the deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordType
func (c_ CKSyncEngineFetchedRecordDeletion) RecordType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordType"))
	return rv
}/* debug [instance_properties/getter]: recordType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineFetchedRecordDeletion */



