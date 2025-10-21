// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An operation that modifies one or more records.
//
// After modifying the fields of a record, use this operation to save those changes to a database. You also use this operation to delete records permanently from a database. If you’re saving a record that contains a reference to another record, set the reference’s to indicate if the target record’s deletion should cascade to the saved record. This helps avoid orphaned records in explicit record hierarchies. When creating two new records that have a reference between them, use the same operation to save both records at the same time. During a save operation, CloudKit requires that the target record of the reference, if set, exists in the database or is part of the same operation; all other reference fields are exempt from this requirement. When you save records, the value in the property determines how to proceed when CloudKit detects conflicts. Because records can change between the time you fetch them and the time you save them, the save policy determines whether new changes overwrite existing changes. By default, the operation reports an error when there’s a newer version on the server. You can change the default setting to permit your changes to overwrite the server values wholly or partially. The handlers you assign to monitor progress of the operation execute serially on an internal queue that the operation manages. Your handlers must be capable of executing on a background thread, so any tasks that require access to the main thread must redirect accordingly. If you assign a completion handler to the property of the operation, CloudKit calls it after the operation executes and returns the results. Use the completion handler to perform any housekeeping tasks for the operation, but don’t use it to process the results of the operation. The completion handler you provide should manage any failures of the operation, whether due to an error or an explicit cancellation.
//
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


//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordDeleteBlock-7gaqj
func (c_ CKModifyRecordsOperation) PerRecordDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordDeleteBlock"))
	return rv
}


// SetPerRecordDeleteBlock sets the value of the perRecordDeleteBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordDeleteBlock-7gaqj
func (c_ CKModifyRecordsOperation) SetPerRecordDeleteBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordDeleteBlock:"), value)
}



