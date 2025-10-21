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
}

// An operation that reports on the changed and deleted records in the specified record zone.
//
// Use this type of operation object to optimize fetch operations for locally managed sets of records. Specifically, use it when you maintain a local cache of your record data and need to synchronize that cache periodically with the server. To get the most benefit out of a object, you must maintain a local cache of the records from the specified zone. Each time you fetch changes from that zone, the server provides a token that identifies your request. With each subsequent fetch request, you initialize the operation object with the token from the previous request, and the server returns only the records with changes since that request. The blocks you assign to process the fetched records execute serially on an internal queue that the operation manages. Your blocks must be capable of executing on a background thread, so any tasks that require access to the main thread must redirect accordingly. If you assign a completion block to the property of the operation object, the system calls the completion block after the operation executes and returns its results to you. You can use a completion block to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. Any completion block you specify should handle the failure of the operation to complete its task, whether due to an error or an explicit cancellation.
//
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




// Creates an operation for fetching changes in the specified record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/init(recordZoneID:previousServerChangeToken:)
func NewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken(recordZoneID unsafe.Pointer, previousServerChangeToken unsafe.Pointer) CKFetchRecordChangesOperation {
	instance := getCKFetchRecordChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordChangesOperation](instance.ID, objc.Sel("initWithRecordZoneID:previousServerChangeToken:"), recordZoneID, previousServerChangeToken)
	rv.Autorelease()
	return rv
}


// The fields to fetch for the requested records.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/desiredKeys
func (c_ CKFetchRecordChangesOperation) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The fields to fetch for the requested records.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/desiredKeys
func (c_ CKFetchRecordChangesOperation) SetDesiredKeys(value []string) {
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

// A Boolean value that indicates whether more results are available.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/moreComing
func (c_ CKFetchRecordChangesOperation) MoreComing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("moreComing"))
	return rv
}

// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/previousServerChangeToken
func (c_ CKFetchRecordChangesOperation) PreviousServerChangeToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}


// SetPreviousServerChangeToken sets the value of the previousServerChangeToken property.
// The token that identifies the starting point for retrieving changes.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/previousServerChangeToken
func (c_ CKFetchRecordChangesOperation) SetPreviousServerChangeToken(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}

// The ID of the record zone with the records you want to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordZoneID
func (c_ CKFetchRecordChangesOperation) RecordZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneID"))
	return rv
}


// SetRecordZoneID sets the value of the recordZoneID property.
// The ID of the record zone with the records you want to fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordZoneID
func (c_ CKFetchRecordChangesOperation) SetRecordZoneID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneID:"), value)
}

// The maximum number of changed records to report with this operation object.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/resultsLimit
func (c_ CKFetchRecordChangesOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of changed records to report with this operation object.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/resultsLimit
func (c_ CKFetchRecordChangesOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}


