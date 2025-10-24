// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CKModifyRecordsOperation] class.
type ICKModifyRecordsOperation interface {
	ICKDatabaseOperation
	// properties:
	SavePolicy() unsafe.Pointer
	SetSavePolicy(value unsafe.Pointer)
	ClientChangeTokenData() objc.IObject /* cross-framework: Data */
	SetClientChangeTokenData(value objc.IObject /* cross-framework: Data */)
	IsAtomic() bool
	SetIsAtomic(value bool)
	ModifyRecordsCompletionBlock() unsafe.Pointer
	SetModifyRecordsCompletionBlock(value unsafe.Pointer)
	ModifyRecordsResultBlock() unsafe.Pointer
	SetModifyRecordsResultBlock(value unsafe.Pointer)
	PerRecordCompletionBlock() unsafe.Pointer
	SetPerRecordCompletionBlock(value unsafe.Pointer)
	PerRecordDeleteBlock() unsafe.Pointer
	SetPerRecordDeleteBlock(value unsafe.Pointer)
	PerRecordProgressBlock() unsafe.Pointer
	SetPerRecordProgressBlock(value unsafe.Pointer)
	PerRecordSaveBlock() unsafe.Pointer
	SetPerRecordSaveBlock(value unsafe.Pointer)
	RecordIDsToDelete() objc.IObject /* cross-framework: CKRecordID */
	SetRecordIDsToDelete(value objc.IObject /* cross-framework: CKRecordID */)
	RecordsToSave() ICKRecord
	SetRecordsToSave(value ICKRecord)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	Parent() ICKReference
	SetParent(value ICKReference)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CKModifyRecordsOperationClass) Alloc() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The policy to use when saving changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/savePolicy
func (c_ CKModifyRecordsOperation) SavePolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("savePolicy"))
	return rv
}


// The policy to use when saving changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/savePolicy
func (c_ CKModifyRecordsOperation) SetSavePolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSavePolicy:"), value)
}


// A token that tracks local changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/clientchangetokendata
func (c_ CKModifyRecordsOperation) ClientChangeTokenData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("clientChangeTokenData"))
	return rv
}


// A token that tracks local changes to records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/clientchangetokendata
func (c_ CKModifyRecordsOperation) SetClientChangeTokenData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClientChangeTokenData:"), value)
}


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/isatomic
func (c_ CKModifyRecordsOperation) IsAtomic() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAtomic"))
	return rv
}


// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/isatomic
func (c_ CKModifyRecordsOperation) SetIsAtomic(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsAtomic:"), value)
}


// The closure to execute after CloudKit modifies all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordscompletionblock
func (c_ CKModifyRecordsOperation) ModifyRecordsCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordsCompletionBlock"))
	return rv
}


// The closure to execute after CloudKit modifies all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordscompletionblock
func (c_ CKModifyRecordsOperation) SetModifyRecordsCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordsCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordsresultblock
func (c_ CKModifyRecordsOperation) ModifyRecordsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordsResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/modifyrecordsresultblock
func (c_ CKModifyRecordsOperation) SetModifyRecordsResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordsResultBlock:"), value)
}


// The closure to execute when CloudKit saves a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordcompletionblock
func (c_ CKModifyRecordsOperation) PerRecordCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordCompletionBlock"))
	return rv
}


// The closure to execute when CloudKit saves a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordcompletionblock
func (c_ CKModifyRecordsOperation) SetPerRecordCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecorddeleteblock-9czoo
func (c_ CKModifyRecordsOperation) PerRecordDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordDeleteBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecorddeleteblock-9czoo
func (c_ CKModifyRecordsOperation) SetPerRecordDeleteBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordDeleteBlock:"), value)
}


// The closure to execute with progress information for individual records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordprogressblock
func (c_ CKModifyRecordsOperation) PerRecordProgressBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordProgressBlock"))
	return rv
}


// The closure to execute with progress information for individual records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordprogressblock
func (c_ CKModifyRecordsOperation) SetPerRecordProgressBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordProgressBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordsaveblock-7yq9d
func (c_ CKModifyRecordsOperation) PerRecordSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordSaveBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordsaveblock-7yq9d
func (c_ CKModifyRecordsOperation) SetPerRecordSaveBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordSaveBlock:"), value)
}


// The IDs of the records to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/recordidstodelete
func (c_ CKModifyRecordsOperation) RecordIDsToDelete() objc.IObject /* cross-framework: CKRecordID */ {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordIDsToDelete"))
	return rv
}


// The IDs of the records to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/recordidstodelete
func (c_ CKModifyRecordsOperation) SetRecordIDsToDelete(value objc.IObject /* cross-framework: CKRecordID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordIDsToDelete:"), value)
}


// The records to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/recordstosave
func (c_ CKModifyRecordsOperation) RecordsToSave() ICKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("recordsToSave"))
	return rv
}


// The records to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/recordstosave
func (c_ CKModifyRecordsOperation) SetRecordsToSave(value ICKRecord) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordsToSave:"), value)
}


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKModifyRecordsOperation) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("action"))
	return rv
}


// The ownership behavior for the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/reference/action-swift.property
func (c_ CKModifyRecordsOperation) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKModifyRecordsOperation) Parent() ICKReference {
	rv := objc.Send[CKReference](c_.ID, objc.Sel("parent"))
	return rv
}


// A reference to the record’s parent record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecord/parent
func (c_ CKModifyRecordsOperation) SetParent(value ICKReference) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParent:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordsOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordsOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}



