// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchRecordChangesOperation] class.
var (
	CKFetchRecordChangesOperationClass     _CKFetchRecordChangesOperationClass
	CKFetchRecordChangesOperationClassOnce sync.Once
)

func getCKFetchRecordChangesOperationClass() _CKFetchRecordChangesOperationClass {
	CKFetchRecordChangesOperationClassOnce.Do(func() {
		CKFetchRecordChangesOperationClass = _CKFetchRecordChangesOperationClass{objc.GetClass("CKFetchRecordChangesOperation")}
	})
	return CKFetchRecordChangesOperationClass
}

type _CKFetchRecordChangesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchRecordChangesOperation] class.
type ICKFetchRecordChangesOperation interface {
	ICKDatabaseOperation
	DesiredKeys() string
	SetDesiredKeys(value string)
	FetchRecordChangesCompletionBlock() unsafe.Pointer
	SetFetchRecordChangesCompletionBlock(value unsafe.Pointer)
	MoreComing() bool
	SetMoreComing(value bool)
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	RecordChangedBlock() unsafe.Pointer
	SetRecordChangedBlock(value unsafe.Pointer)
	RecordWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordWithIDWasDeletedBlock(value unsafe.Pointer)
	RecordZoneID() ICKRecordZoneID
	SetRecordZoneID(value ICKRecordZoneID)
	ResultsLimit() int
	SetResultsLimit(value int)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
}

// An operation that reports on the changed and deleted records in the specified record zone.
//
// Use this type of operation object to optimize fetch operations for locally managed sets of records. Specifically, use it when you maintain a local cache of your record data and need to synchronize that cache periodically with the server. To get the most benefit out of a object, you must maintain a local cache of the records from the specified zone. Each time you fetch changes from that zone, the server provides a token that identifies your request. With each subsequent fetch request, you initialize the operation object with the token from the previous request, and the server returns only the records with changes since that request. The blocks you assign to process the fetched records execute serially on an internal queue that the operation manages. Your blocks must be capable of executing on a background thread, so any tasks that require access to the main thread must redirect accordingly. If you assign a completion block to the property of the operation object, the system calls the completion block after the operation executes and returns its results to you. You can use a completion block to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. Any completion block you specify should handle the failure of the operation to complete its task, whether due to an error or an explicit cancellation.


// An operation that reports on the changed and deleted records in the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation
type CKFetchRecordChangesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordChangesOperationFrom constructs a [CKFetchRecordChangesOperation] from an unsafe.Pointer.
//
// An operation that reports on the changed and deleted records in the specified record zone.
func CKFetchRecordChangesOperationFrom(ptr unsafe.Pointer) CKFetchRecordChangesOperation {
	return CKFetchRecordChangesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordChangesOperationClass) Alloc() CKFetchRecordChangesOperation {
	rv := objc.Send[CKFetchRecordChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchRecordChangesOperationClass) New() CKFetchRecordChangesOperation {
	rv := objc.Send[CKFetchRecordChangesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordChangesOperation) Init() CKFetchRecordChangesOperation {
	rv := objc.Send[CKFetchRecordChangesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordChangesOperation) Autorelease() CKFetchRecordChangesOperation {
	rv := objc.Send[CKFetchRecordChangesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordChangesOperation creates a new CKFetchRecordChangesOperation instance.
func NewCKFetchRecordChangesOperation() CKFetchRecordChangesOperation {
	return getCKFetchRecordChangesOperationClass().New()
}



// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/desiredkeys
func (c_ CKFetchRecordChangesOperation) DesiredKeys() string {
	rv := objc.Send[string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/desiredkeys
func (c_ CKFetchRecordChangesOperation) SetDesiredKeys(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), objc.String(value))
}


// The block to execute when the system finishes processing all changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/fetchrecordchangescompletionblock
func (c_ CKFetchRecordChangesOperation) FetchRecordChangesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordChangesCompletionBlock"))
	return rv
}


// The block to execute when the system finishes processing all changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/fetchrecordchangescompletionblock
func (c_ CKFetchRecordChangesOperation) SetFetchRecordChangesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordChangesCompletionBlock:"), value)
}


// A Boolean value that indicates whether more results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/morecoming
func (c_ CKFetchRecordChangesOperation) MoreComing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("moreComing"))
	return rv
}


// A Boolean value that indicates whether more results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/morecoming
func (c_ CKFetchRecordChangesOperation) SetMoreComing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMoreComing:"), value)
}


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/previousserverchangetoken
func (c_ CKFetchRecordChangesOperation) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/previousserverchangetoken
func (c_ CKFetchRecordChangesOperation) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}


// The block to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordchangedblock
func (c_ CKFetchRecordChangesOperation) RecordChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordChangedBlock"))
	return rv
}


// The block to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordchangedblock
func (c_ CKFetchRecordChangesOperation) SetRecordChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangedBlock:"), value)
}


// The block to execute with the ID of a deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordwithidwasdeletedblock
func (c_ CKFetchRecordChangesOperation) RecordWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}


// The block to execute with the ID of a deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordwithidwasdeletedblock
func (c_ CKFetchRecordChangesOperation) SetRecordWithIDWasDeletedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}


// The ID of the record zone with the records you want to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordzoneid
func (c_ CKFetchRecordChangesOperation) RecordZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneID"))
	return rv
}


// The ID of the record zone with the records you want to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/recordzoneid
func (c_ CKFetchRecordChangesOperation) SetRecordZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneID:"), value)
}


// The maximum number of changed records to report with this operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/resultslimit
func (c_ CKFetchRecordChangesOperation) ResultsLimit() int {
	rv := objc.Send[int](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// The maximum number of changed records to report with this operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordchangesoperation/resultslimit
func (c_ CKFetchRecordChangesOperation) SetResultsLimit(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordChangesOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordChangesOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}



