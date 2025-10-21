// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchRecordsOperation] class.
var (
	CKFetchRecordsOperationClass     _CKFetchRecordsOperationClass
	CKFetchRecordsOperationClassOnce sync.Once
)

func getCKFetchRecordsOperationClass() _CKFetchRecordsOperationClass {
	CKFetchRecordsOperationClassOnce.Do(func() {
		CKFetchRecordsOperationClass = _CKFetchRecordsOperationClass{objc.GetClass("CKFetchRecordsOperation")}
	})
	return CKFetchRecordsOperationClass
}

type _CKFetchRecordsOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchRecordsOperation] class.
type ICKFetchRecordsOperation interface {
	ICKDatabaseOperation
}

// An operation for retrieving records from a database.
//
// Use this operation to retrieve the entire contents of each record or only a subset of its contained values. As records become available, the operation object reports progress about the state of the operation to several different blocks, which you can use to process the results. Fetching records is a common use of CloudKit, even if your app doesn’t cache record IDs locally. For example, when you fetch a record related to the current record through a object, you use the ID in the reference to perform the fetch. The handlers you assign to process the fetched records execute serially on an internal queue that the fetch operation manages. Your handlers must be capable of executing on a background thread, so any tasks that require access to the main thread must redirect accordingly. In addition to data records, a fetch records operation can fetch the current user record. The method returns a specially configured operation object that retrieves the current user record. That record is a standard object that has no content initially. You can add data to the user record and save it as necessary. Don’t store sensitive personal information, such as passwords, in the user record because other users of your app can access the discoverable user record in a public database. If you must store sensitive information about a user, do so in a separate record that is accessible only to that user. If you assign a closure to the property of the operation object, CloudKit calls it after the operation executes and returns its results. Use a closure to perform any housekeeping tasks for the operation, but don’t use it to process the results of the operation. The closure you specify should handle the failure of the operation to complete its task, whether due to an error or an explicit cancellation.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation
type CKFetchRecordsOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordsOperationFrom constructs a [CKFetchRecordsOperation] from an unsafe.Pointer.
//
// An operation for retrieving records from a database.
func CKFetchRecordsOperationFrom(ptr unsafe.Pointer) CKFetchRecordsOperation {
	return CKFetchRecordsOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordsOperationClass) Alloc() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchRecordsOperationClass) New() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordsOperation) Init() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordsOperation) Autorelease() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordsOperation creates a new CKFetchRecordsOperation instance.
func NewCKFetchRecordsOperation() CKFetchRecordsOperation {
	return getCKFetchRecordsOperationClass().New()
}




// Creates a fetch operation for retrieving the records with the specified IDs.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/init(recordIDs:)
func NewCKFetchRecordsOperationWithRecordIDs(recordIDs unsafe.Pointer) CKFetchRecordsOperation {
	instance := getCKFetchRecordsOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordsOperation](instance.ID, objc.Sel("initWithRecordIDs:"), recordIDs)
	rv.Autorelease()
	return rv
}


// Returns a fetch operation for retrieving the current user record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/fetchCurrentUserRecordOperation()
func (cc _CKFetchRecordsOperationClass) FetchCurrentUserRecordOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("fetchCurrentUserRecordOperation"))
	return rv
}

// The fields of the records to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/desiredKeys-34l1l
func (c_ CKFetchRecordsOperation) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The fields of the records to fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/desiredKeys-34l1l
func (c_ CKFetchRecordsOperation) SetDesiredKeys(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), nsArray)
}

// The record IDs of the records to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/recordIDs
func (c_ CKFetchRecordsOperation) RecordIDs() []CKRecordID {
	rv := objc.Send[[]CKRecordID](c_.ID, objc.Sel("recordIDs"))
	return rv
}


// SetRecordIDs sets the value of the recordIDs property.
// The record IDs of the records to fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/recordIDs
func (c_ CKFetchRecordsOperation) SetRecordIDs(value []CKRecordID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordIDs:"), nsArray)
}

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/fetchrecordsresultblock
func (c_ CKFetchRecordsOperation) FetchRecordsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordsResultBlock"))
	return rv
}


// SetFetchRecordsResultBlock sets the value of the fetchRecordsResultBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/fetchrecordsresultblock
func (c_ CKFetchRecordsOperation) SetFetchRecordsResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordsResultBlock:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/perrecordresultblock
func (c_ CKFetchRecordsOperation) PerRecordResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordResultBlock"))
	return rv
}


// SetPerRecordResultBlock sets the value of the perRecordResultBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/perrecordresultblock
func (c_ CKFetchRecordsOperation) SetPerRecordResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordResultBlock:"), value)
}

// The block to execute after the operation’s main task is completed.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordsOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// SetCompletionBlock sets the value of the completionBlock property.
// The block to execute after the operation’s main task is completed.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordsOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}


