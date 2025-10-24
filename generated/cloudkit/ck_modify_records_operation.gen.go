// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKModifyRecordsOperation */


/* debug [class_header]: Header for CKModifyRecordsOperation */
// The class instance for the [CKModifyRecordsOperation] class.
var (
	CKModifyRecordsOperationClass     _CKModifyRecordsOperationClass
	CKModifyRecordsOperationClassOnce sync.Once
)

func getCKModifyRecordsOperationClass() _CKModifyRecordsOperationClass {
	CKModifyRecordsOperationClassOnce.Do(func() {
		CKModifyRecordsOperationClass = _CKModifyRecordsOperationClass{objc.GetClass("CKModifyRecordsOperation")}
	})
	return CKModifyRecordsOperationClass
}

type _CKModifyRecordsOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKModifyRecordsOperation */
// An interface definition for the [CKModifyRecordsOperation] class.
type ICKModifyRecordsOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKModifyRecordsOperation */
	// properties:
	ClientChangeTokenData() objc.IObject /* cross-framework: NSData */
	SetClientChangeTokenData(value objc.IObject /* cross-framework: NSData */)
	Atomic() bool
	SetAtomic(value bool)
	ModifyRecordsCompletionBlock() unsafe.Pointer
	SetModifyRecordsCompletionBlock(value unsafe.Pointer)
	PerRecordCompletionBlock() unsafe.Pointer
	SetPerRecordCompletionBlock(value unsafe.Pointer)
	PerRecordDeleteBlock() func(unsafe.Pointer, unsafe.Pointer)
	SetPerRecordDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer))
	PerRecordProgressBlock() unsafe.Pointer
	SetPerRecordProgressBlock(value unsafe.Pointer)
	PerRecordSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerRecordSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	RecordIDsToDelete() []CKRecordID
	SetRecordIDsToDelete(value []CKRecordID)
	RecordsToSave() []objc.IObject /* cross-framework: CKRecord */
	SetRecordsToSave(value []objc.IObject /* cross-framework: CKRecord */)
	SavePolicy() CKRecordSavePolicy
	SetSavePolicy(value CKRecordSavePolicy)
	IsAtomic() bool
	SetIsAtomic(value bool)
	ModifyRecordsResultBlock() objectivec.IObject
	SetModifyRecordsResultBlock(value objectivec.IObject)
	Action() objectivec.IObject
	SetAction(value objectivec.IObject)
	Parent() ICKReference
	SetParent(value ICKReference)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKModifyRecordsOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKModifyRecordsOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKModifyRecordsOperationClass) Alloc() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKModifyRecordsOperationClass) New() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKModifyRecordsOperation) Init() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKModifyRecordsOperation) Autorelease() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifyRecordsOperation creates a new CKModifyRecordsOperation instance.
func NewCKModifyRecordsOperation() CKModifyRecordsOperation {
	return getCKModifyRecordsOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKModifyRecordsOperation */
// An operation that modifies one or more records.
//
// After modifying the fields of a record, use this operation to save those changes to a database. You also use this operation to delete records permanently from a database. If you’re saving a record that contains a reference to another record, set the reference’s to indicate if the target record’s deletion should cascade to the saved record. This helps avoid orphaned records in explicit record hierarchies. When creating two new records that have a reference between them, use the same operation to save both records at the same time. During a save operation, CloudKit requires that the target record of the reference, if set, exists in the database or is part of the same operation; all other reference fields are exempt from this requirement. When you save records, the value in the property determines how to proceed when CloudKit detects conflicts. Because records can change between the time you fetch them and the time you save them, the save policy determines whether new changes overwrite existing changes. By default, the operation reports an error when there’s a newer version on the server. You can change the default setting to permit your changes to overwrite the server values wholly or partially. The handlers you assign to monitor progress of the operation execute serially on an internal queue that the operation manages. Your handlers must be capable of executing on a background thread, so any tasks that require access to the main thread must redirect accordingly. If you assign a completion handler to the property of the operation, CloudKit calls it after the operation executes and returns the results. Use the completion handler to perform any housekeeping tasks for the operation, but don’t use it to process the results of the operation. The completion handler you provide should manage any failures of the operation, whether due to an error or an explicit cancellation.


// An operation that modifies one or more records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation
type CKModifyRecordsOperation struct {
	CKDatabaseOperation
}

// CKModifyRecordsOperationFrom constructs a [CKModifyRecordsOperation] from an unsafe.Pointer.
//
// An operation that modifies one or more records.
func CKModifyRecordsOperationFrom(ptr unsafe.Pointer) CKModifyRecordsOperation {
	return CKModifyRecordsOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKModifyRecordsOperation */

// Creates an operation for modifying the specified records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/initWithRecordsToSave:recordIDsToDelete:
func NewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete(records []objc.IObject /* cross-framework: CKRecord */, recordIDs []CKRecordID) CKModifyRecordsOperation {
	instance := getCKModifyRecordsOperationClass().Alloc()
	rv := objc.Send[CKModifyRecordsOperation](instance.ID, objc.Sel("initWithRecordsToSave:recordIDsToDelete:"), records, recordIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKModifyRecordsOperationWithRecordsToSaveRecordIDsToDelete */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKModifyRecordsOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKModifyRecordsOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKModifyRecordsOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKModifyRecordsOperation */

// A token that tracks local changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/clientChangeTokenData
func (c_ CKModifyRecordsOperation) ClientChangeTokenData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("clientChangeTokenData"))
	return rv
}/* debug [instance_properties/getter]: clientChangeTokenData */


// A token that tracks local changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/clientChangeTokenData
func (c_ CKModifyRecordsOperation) SetClientChangeTokenData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClientChangeTokenData:"), value)
}/* debug [instance_properties/setter]: clientChangeTokenData */


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/isAtomic
func (c_ CKModifyRecordsOperation) Atomic() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("atomic"))
	return rv
}/* debug [instance_properties/getter]: atomic */


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/isAtomic
func (c_ CKModifyRecordsOperation) SetAtomic(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAtomic:"), value)
}/* debug [instance_properties/setter]: atomic */


// The closure to execute after CloudKit modifies all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/modifyRecordsCompletionBlock
func (c_ CKModifyRecordsOperation) ModifyRecordsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordsCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: modifyRecordsCompletionBlock */


// The closure to execute after CloudKit modifies all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/modifyRecordsCompletionBlock
func (c_ CKModifyRecordsOperation) SetModifyRecordsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordsCompletionBlock:"), value)
}/* debug [instance_properties/setter]: modifyRecordsCompletionBlock */


// The closure to execute when CloudKit saves a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordCompletionBlock
func (c_ CKModifyRecordsOperation) PerRecordCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordCompletionBlock */


// The closure to execute when CloudKit saves a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordCompletionBlock
func (c_ CKModifyRecordsOperation) SetPerRecordCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perRecordCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordDeleteBlock-7gaqj
func (c_ CKModifyRecordsOperation) PerRecordDeleteBlock() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perRecordDeleteBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordDeleteBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordDeleteBlock-7gaqj
func (c_ CKModifyRecordsOperation) SetPerRecordDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordDeleteBlock:"), value)
}/* debug [instance_properties/setter]: perRecordDeleteBlock */


// The closure to execute with progress information for individual records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordProgressBlock
func (c_ CKModifyRecordsOperation) PerRecordProgressBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordProgressBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordProgressBlock */


// The closure to execute with progress information for individual records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordProgressBlock
func (c_ CKModifyRecordsOperation) SetPerRecordProgressBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordProgressBlock:"), value)
}/* debug [instance_properties/setter]: perRecordProgressBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordSaveBlock-80dn4
func (c_ CKModifyRecordsOperation) PerRecordSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perRecordSaveBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordSaveBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordSaveBlock-80dn4
func (c_ CKModifyRecordsOperation) SetPerRecordSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordSaveBlock:"), value)
}/* debug [instance_properties/setter]: perRecordSaveBlock */


// The IDs of the records to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordIDsToDelete
func (c_ CKModifyRecordsOperation) RecordIDsToDelete() []CKRecordID {
	rv := objc.Send[[]CKRecordID](c_.ID, objc.Sel("recordIDsToDelete"))
	return rv
}/* debug [instance_properties/getter]: recordIDsToDelete */


// The IDs of the records to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordIDsToDelete
func (c_ CKModifyRecordsOperation) SetRecordIDsToDelete(value []CKRecordID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordIDsToDelete:"), nsArray)
}/* debug [instance_properties/setter]: recordIDsToDelete */


// The records to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordsToSave
func (c_ CKModifyRecordsOperation) RecordsToSave() []objc.IObject /* cross-framework: CKRecord */ {
	rv := objc.Send[[]CKRecord](c_.ID, objc.Sel("recordsToSave"))
	return rv
}/* debug [instance_properties/getter]: recordsToSave */


// The records to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordsToSave
func (c_ CKModifyRecordsOperation) SetRecordsToSave(value []objc.IObject /* cross-framework: CKRecord */) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordsToSave:"), nsArray)
}/* debug [instance_properties/setter]: recordsToSave */


// The policy to use when saving changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/savePolicy
func (c_ CKModifyRecordsOperation) SavePolicy() CKRecordSavePolicy {
	rv := objc.Send[CKRecordSavePolicy](c_.ID, objc.Sel("savePolicy"))
	return rv
}/* debug [instance_properties/getter]: savePolicy */


// The policy to use when saving changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/savePolicy
func (c_ CKModifyRecordsOperation) SetSavePolicy(value CKRecordSavePolicy) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSavePolicy:"), value)
}/* debug [instance_properties/setter]: savePolicy */


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/isatomic
func (c_ CKModifyRecordsOperation) IsAtomic() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAtomic"))
	return rv
}/* debug [instance_properties/getter]: isAtomic */


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/isatomic
func (c_ CKModifyRecordsOperation) SetIsAtomic(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAtomic:"), value)
}/* debug [instance_properties/setter]: isAtomic */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordsresultblock
func (c_ CKModifyRecordsOperation) ModifyRecordsResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("modifyRecordsResultBlock"))
	return rv
}/* debug [instance_properties/getter]: modifyRecordsResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordsresultblock
func (c_ CKModifyRecordsOperation) SetModifyRecordsResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordsResultBlock:"), value)
}/* debug [instance_properties/setter]: modifyRecordsResultBlock */


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKModifyRecordsOperation) Action() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKModifyRecordsOperation) SetAction(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKModifyRecordsOperation) Parent() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKModifyRecordsOperation) SetParent(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordsOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordsOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKModifyRecordsOperation */


