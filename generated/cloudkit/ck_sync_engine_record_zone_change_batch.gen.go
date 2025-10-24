// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKSyncEngineRecordZoneChangeBatch */


/* debug [class_header]: Header for CKSyncEngineRecordZoneChangeBatch */
// The class instance for the [CKSyncEngineRecordZoneChangeBatch] class.
var (
	CKSyncEngineRecordZoneChangeBatchClass     _CKSyncEngineRecordZoneChangeBatchClass
	CKSyncEngineRecordZoneChangeBatchClassOnce sync.Once
)

func getCKSyncEngineRecordZoneChangeBatchClass() _CKSyncEngineRecordZoneChangeBatchClass {
	CKSyncEngineRecordZoneChangeBatchClassOnce.Do(func() {
		CKSyncEngineRecordZoneChangeBatchClass = _CKSyncEngineRecordZoneChangeBatchClass{objc.GetClass("CKSyncEngineRecordZoneChangeBatch")}
	})
	return CKSyncEngineRecordZoneChangeBatchClass
}

type _CKSyncEngineRecordZoneChangeBatchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineRecordZoneChangeBatch */
// An interface definition for the [CKSyncEngineRecordZoneChangeBatch] class.
type ICKSyncEngineRecordZoneChangeBatch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKSyncEngineRecordZoneChangeBatch */
	// properties:
	AtomicByZone() bool
	SetAtomicByZone(value bool)
	RecordIDsToDelete() []CKRecordID
	RecordsToSave() []objc.IObject /* cross-framework: CKRecord */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineRecordZoneChangeBatch */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineRecordZoneChangeBatch */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineRecordZoneChangeBatchClass) Alloc() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineRecordZoneChangeBatchClass) New() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineRecordZoneChangeBatch) Init() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineRecordZoneChangeBatch) Autorelease() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineRecordZoneChangeBatch creates a new CKSyncEngineRecordZoneChangeBatch instance.
func NewCKSyncEngineRecordZoneChangeBatch() CKSyncEngineRecordZoneChangeBatch {
	return getCKSyncEngineRecordZoneChangeBatchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineRecordZoneChangeBatch */
// An object that contains the record changes for a single send operation.


// An object that contains the record changes for a single send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch
type CKSyncEngineRecordZoneChangeBatch struct {
	objectivec.Object
}

// CKSyncEngineRecordZoneChangeBatchFrom constructs a [CKSyncEngineRecordZoneChangeBatch] from an unsafe.Pointer.
//
// An object that contains the record changes for a single send operation.
func CKSyncEngineRecordZoneChangeBatchFrom(ptr unsafe.Pointer) CKSyncEngineRecordZoneChangeBatch {
	return CKSyncEngineRecordZoneChangeBatch{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineRecordZoneChangeBatch */

// Creates a batch of records to modify using the provided record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/initWithPendingChanges:recordProvider:
func NewCKSyncEngineRecordZoneChangeBatchWithPendingChangesRecordProvider(pendingChanges []CKSyncEnginePendingRecordZoneChange, recordProvider unsafe.Pointer) CKSyncEngineRecordZoneChangeBatch {
	instance := getCKSyncEngineRecordZoneChangeBatchClass().Alloc()
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](instance.ID, objc.Sel("initWithPendingChanges:recordProvider:"), pendingChanges, recordProvider)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineRecordZoneChangeBatchWithPendingChangesRecordProvider */


// Creates a batch of records to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/initWithRecordsToSave:recordIDsToDelete:atomicByZone:
func NewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone(recordsToSave []objc.IObject /* cross-framework: CKRecord */, recordIDsToDelete []CKRecordID, atomicByZone bool) CKSyncEngineRecordZoneChangeBatch {
	instance := getCKSyncEngineRecordZoneChangeBatchClass().Alloc()
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](instance.ID, objc.Sel("initWithRecordsToSave:recordIDsToDelete:atomicByZone:"), recordsToSave, recordIDsToDelete, atomicByZone)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineRecordZoneChangeBatch */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineRecordZoneChangeBatch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineRecordZoneChangeBatch */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineRecordZoneChangeBatch */

// A Boolean value that determines whether CloudKit modifies records atomically by record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/atomicByZone
func (c_ CKSyncEngineRecordZoneChangeBatch) AtomicByZone() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("atomicByZone"))
	return rv
}/* debug [instance_properties/getter]: atomicByZone */


// A Boolean value that determines whether CloudKit modifies records atomically by record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/atomicByZone
func (c_ CKSyncEngineRecordZoneChangeBatch) SetAtomicByZone(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAtomicByZone:"), value)
}/* debug [instance_properties/setter]: atomicByZone */


// The unique identifiers of the records to delete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/recordIDsToDelete
func (c_ CKSyncEngineRecordZoneChangeBatch) RecordIDsToDelete() []CKRecordID {
	rv := objc.Send[[]CKRecordID](c_.ID, objc.Sel("recordIDsToDelete"))
	return rv
}/* debug [instance_properties/getter]: recordIDsToDelete */


// The records to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/recordsToSave
func (c_ CKSyncEngineRecordZoneChangeBatch) RecordsToSave() []objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[[]CKRecord](c_.ID, objc.Sel("recordsToSave"))
	return rv
}/* debug [instance_properties/getter]: recordsToSave */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineRecordZoneChangeBatch */


