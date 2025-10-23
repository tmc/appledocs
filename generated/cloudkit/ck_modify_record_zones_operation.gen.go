// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKModifyRecordZonesOperation] class.
var (
	CKModifyRecordZonesOperationClass     _CKModifyRecordZonesOperationClass
	CKModifyRecordZonesOperationClassOnce sync.Once
)

func getCKModifyRecordZonesOperationClass() _CKModifyRecordZonesOperationClass {
	CKModifyRecordZonesOperationClassOnce.Do(func() {
		CKModifyRecordZonesOperationClass = _CKModifyRecordZonesOperationClass{objc.GetClass("CKModifyRecordZonesOperation")}
	})
	return CKModifyRecordZonesOperationClass
}

type _CKModifyRecordZonesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKModifyRecordZonesOperation] class.
type ICKModifyRecordZonesOperation interface {
	ICKDatabaseOperation
	ModifyRecordZonesCompletionBlock() unsafe.Pointer
	SetModifyRecordZonesCompletionBlock(value unsafe.Pointer)
	ModifyRecordZonesResultBlock() unsafe.Pointer
	SetModifyRecordZonesResultBlock(value unsafe.Pointer)
	PerRecordZoneDeleteBlock() unsafe.Pointer
	SetPerRecordZoneDeleteBlock(value unsafe.Pointer)
	PerRecordZoneSaveBlock() unsafe.Pointer
	SetPerRecordZoneSaveBlock(value unsafe.Pointer)
	RecordZoneIDsToDelete() ICKRecordZoneID
	SetRecordZoneIDsToDelete(value ICKRecordZoneID)
	RecordZonesToSave() ICKRecordZone
	SetRecordZonesToSave(value ICKRecordZone)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
}

// An operation that modifies one or more record zones.
//
// After you create one or more record zones, use this operation to save those zones to the database. You can also use the operation to delete record zones and their records. If you assign a handler to the property of the operation, CloudKit calls the handler after the operation executes and returns its results. Use the handler to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. The handler you provide should manage any failures of the operation, whether due to an error or an explicit cancellation.


// An operation that modifies one or more record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation
type CKModifyRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKModifyRecordZonesOperationFrom constructs a [CKModifyRecordZonesOperation] from an unsafe.Pointer.
//
// An operation that modifies one or more record zones.
func CKModifyRecordZonesOperationFrom(ptr unsafe.Pointer) CKModifyRecordZonesOperation {
	return CKModifyRecordZonesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKModifyRecordZonesOperationClass) Alloc() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKModifyRecordZonesOperationClass) New() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKModifyRecordZonesOperation) Init() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKModifyRecordZonesOperation) Autorelease() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifyRecordZonesOperation creates a new CKModifyRecordZonesOperation instance.
func NewCKModifyRecordZonesOperation() CKModifyRecordZonesOperation {
	return getCKModifyRecordZonesOperationClass().New()
}



// The closure to execute after CloudKit modifies all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonescompletionblock
func (c_ CKModifyRecordZonesOperation) ModifyRecordZonesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordZonesCompletionBlock"))
	return rv
}


// The closure to execute after CloudKit modifies all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonescompletionblock
func (c_ CKModifyRecordZonesOperation) SetModifyRecordZonesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordZonesCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonesresultblock
func (c_ CKModifyRecordZonesOperation) ModifyRecordZonesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordZonesResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonesresultblock
func (c_ CKModifyRecordZonesOperation) SetModifyRecordZonesResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordZonesResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonedeleteblock-6c82y
func (c_ CKModifyRecordZonesOperation) PerRecordZoneDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneDeleteBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonedeleteblock-6c82y
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneDeleteBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonesaveblock-1m45y
func (c_ CKModifyRecordZonesOperation) PerRecordZoneSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneSaveBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonesaveblock-1m45y
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneSaveBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneSaveBlock:"), value)
}


// The IDs of the record zones to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/recordzoneidstodelete
func (c_ CKModifyRecordZonesOperation) RecordZoneIDsToDelete() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDsToDelete"))
	return rv
}


// The IDs of the record zones to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/recordzoneidstodelete
func (c_ CKModifyRecordZonesOperation) SetRecordZoneIDsToDelete(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneIDsToDelete:"), value)
}


// The record zones to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/recordzonestosave
func (c_ CKModifyRecordZonesOperation) RecordZonesToSave() ICKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("recordZonesToSave"))
	return rv
}


// The record zones to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/recordzonestosave
func (c_ CKModifyRecordZonesOperation) SetRecordZonesToSave(value ICKRecordZone) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZonesToSave:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordZonesOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordZonesOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}



