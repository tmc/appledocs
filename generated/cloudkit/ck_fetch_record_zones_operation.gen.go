// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchRecordZonesOperation] class.
var (
	CKFetchRecordZonesOperationClass     _CKFetchRecordZonesOperationClass
	CKFetchRecordZonesOperationClassOnce sync.Once
)

func getCKFetchRecordZonesOperationClass() _CKFetchRecordZonesOperationClass {
	CKFetchRecordZonesOperationClassOnce.Do(func() {
		CKFetchRecordZonesOperationClass = _CKFetchRecordZonesOperationClass{objc.GetClass("CKFetchRecordZonesOperation")}
	})
	return CKFetchRecordZonesOperationClass
}

type _CKFetchRecordZonesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchRecordZonesOperation] class.
type ICKFetchRecordZonesOperation interface {
	ICKDatabaseOperation
	FetchRecordZonesCompletionBlock() unsafe.Pointer
	SetFetchRecordZonesCompletionBlock(value unsafe.Pointer)
	PerRecordZoneCompletionBlock() unsafe.Pointer
	SetPerRecordZoneCompletionBlock(value unsafe.Pointer)
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)
	FetchRecordZonesResultBlock() unsafe.Pointer
	SetFetchRecordZonesResultBlock(value unsafe.Pointer)
	PerRecordZoneResultBlock() unsafe.Pointer
	SetPerRecordZoneResultBlock(value unsafe.Pointer)
	CompletionBlock() unsafe.Pointer
	SetCompletionBlock(value unsafe.Pointer)
}

// An operation for retrieving record zones from a database.
//
// Use this operation object to fetch record zones so that you can ascertain their capabilities. If you assign a handler to the property of the operation, CloudKit calls it after the operation executes and returns its results. You can use the handler to perform any housekeeping tasks that relate to the operation, but don’t use it to process the results of the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.


// An operation for retrieving record zones from a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation
type CKFetchRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordZonesOperationFrom constructs a [CKFetchRecordZonesOperation] from an unsafe.Pointer.
//
// An operation for retrieving record zones from a database.
func CKFetchRecordZonesOperationFrom(ptr unsafe.Pointer) CKFetchRecordZonesOperation {
	return CKFetchRecordZonesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZonesOperationClass) Alloc() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchRecordZonesOperationClass) New() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordZonesOperation) Init() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordZonesOperation) Autorelease() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZonesOperation creates a new CKFetchRecordZonesOperation instance.
func NewCKFetchRecordZonesOperation() CKFetchRecordZonesOperation {
	return getCKFetchRecordZonesOperationClass().New()
}



// Creates an operation for fetching the specified record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/init(recordZoneIDs:)
func NewCKFetchRecordZonesOperationWithRecordZoneIDs(zoneIDs []CKRecordZoneID) CKFetchRecordZonesOperation {
	instance := getCKFetchRecordZonesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZonesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}



// Returns an operation for fetching all record zones in the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchAllRecordZonesOperation()
func (cc _CKFetchRecordZonesOperationClass) FetchAllRecordZonesOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("fetchAllRecordZonesOperation"))
	return rv
}


// The closure to execute after CloudKit retrieves all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (c_ CKFetchRecordZonesOperation) FetchRecordZonesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZonesCompletionBlock"))
	return rv
}


// The closure to execute after CloudKit retrieves all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (c_ CKFetchRecordZonesOperation) SetFetchRecordZonesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZonesCompletionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/perRecordZoneCompletionBlock
func (c_ CKFetchRecordZonesOperation) PerRecordZoneCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneCompletionBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/perRecordZoneCompletionBlock
func (c_ CKFetchRecordZonesOperation) SetPerRecordZoneCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneCompletionBlock:"), value)
}


// The IDs of the record zones to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/recordZoneIDs
func (c_ CKFetchRecordZonesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDs"))
	return rv
}


// The IDs of the record zones to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/recordZoneIDs
func (c_ CKFetchRecordZonesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneIDs:"), nsArray)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/fetchrecordzonesresultblock
func (c_ CKFetchRecordZonesOperation) FetchRecordZonesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZonesResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/fetchrecordzonesresultblock
func (c_ CKFetchRecordZonesOperation) SetFetchRecordZonesResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZonesResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/perrecordzoneresultblock
func (c_ CKFetchRecordZonesOperation) PerRecordZoneResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/perrecordzoneresultblock
func (c_ CKFetchRecordZonesOperation) SetPerRecordZoneResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneResultBlock:"), value)
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordZonesOperation) CompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionBlock"))
	return rv
}


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordZonesOperation) SetCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}


