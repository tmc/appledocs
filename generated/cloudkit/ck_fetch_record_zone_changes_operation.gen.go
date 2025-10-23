// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchRecordZoneChangesOperation] class.
var (
	CKFetchRecordZoneChangesOperationClass     _CKFetchRecordZoneChangesOperationClass
	CKFetchRecordZoneChangesOperationClassOnce sync.Once
)

func getCKFetchRecordZoneChangesOperationClass() _CKFetchRecordZoneChangesOperationClass {
	CKFetchRecordZoneChangesOperationClassOnce.Do(func() {
		CKFetchRecordZoneChangesOperationClass = _CKFetchRecordZoneChangesOperationClass{objc.GetClass("CKFetchRecordZoneChangesOperation")}
	})
	return CKFetchRecordZoneChangesOperationClass
}

type _CKFetchRecordZoneChangesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchRecordZoneChangesOperation] class.
type ICKFetchRecordZoneChangesOperation interface {
	ICKDatabaseOperation
	ConfigurationsByRecordZoneID() unsafe.Pointer
	SetConfigurationsByRecordZoneID(value unsafe.Pointer)
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	FetchRecordZoneChangesCompletionBlock() unsafe.Pointer
	SetFetchRecordZoneChangesCompletionBlock(value unsafe.Pointer)
	OptionsByRecordZoneID() unsafe.Pointer
	SetOptionsByRecordZoneID(value unsafe.Pointer)
	RecordChangedBlock() unsafe.Pointer
	SetRecordChangedBlock(value unsafe.Pointer)
	RecordWasChangedBlock() unsafe.Pointer
	SetRecordWasChangedBlock(value unsafe.Pointer)
	RecordWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordWithIDWasDeletedBlock(value unsafe.Pointer)
	RecordZoneChangeTokensUpdatedBlock() unsafe.Pointer
	SetRecordZoneChangeTokensUpdatedBlock(value unsafe.Pointer)
	RecordZoneFetchCompletionBlock() unsafe.Pointer
	SetRecordZoneFetchCompletionBlock(value unsafe.Pointer)
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)
	FetchRecordZoneChangesResultBlock() unsafe.Pointer
	SetFetchRecordZoneChangesResultBlock(value unsafe.Pointer)
	RecordZoneFetchResultBlock() unsafe.Pointer
	SetRecordZoneFetchResultBlock(value unsafe.Pointer)
}

// An operation that fetches record zone changes.
//
// Use this operation to fetch record changes in one or more record zones, such as those that occur during record creation, modification, and deletion. You provide a configuration object for each record zone to query for changes. The configuration contains a server change token, which is an opaque pointer to a specific change in the zone’s history. CloudKit returns only the changes that occur after that point. For the first time you fetch a record zone’s changes, or to refetch all changes in a zone’s history, use instead. CloudKit processes the record zones in succession, and returns the changes for each zone in batches. Each batch yields a new change token. If all batches return without error, the operation issues a final change token for that zone. The change tokens conform to and are safe to cache on-disk. This operation’s tokens aren’t compatible with so you should segregate them in your app’s cache. Don’t infer behavior or order from the tokens’ contents. If you create record zones in the private database, fetch all changes the first time the app launches. Cache the results on-device and use to subscribe to future changes. Fetch those changes on receipt of the push notifications the subscription generates. If you use the shared database, subscribe to changes with instead. When a user participates in sharing, CloudKit adds and removes record zones. This means you don’t know in advance which zones exist in the shared database. Use to fetch shared record zones on receipt of the subscription’s push notifications. Then fetch the changes in those zones using this operation. Regardless of which database you use, it’s not necessary to perform fetches each time your app launches, or to schedule fetches at regular intervals. To run the operation, add it to the corresponding database’s operation queue. The operation executes its callbacks on a private serial queue. The following example demonstrates how to create the operation, configure its callbacks, and execute it. For brevity, it omits the delete and operation completion callbacks.


// An operation that fetches record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation
type CKFetchRecordZoneChangesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordZoneChangesOperationFrom constructs a [CKFetchRecordZoneChangesOperation] from an unsafe.Pointer.
//
// An operation that fetches record zone changes.
func CKFetchRecordZoneChangesOperationFrom(ptr unsafe.Pointer) CKFetchRecordZoneChangesOperation {
	return CKFetchRecordZoneChangesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesOperationClass) Alloc() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchRecordZoneChangesOperationClass) New() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordZoneChangesOperation) Init() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordZoneChangesOperation) Autorelease() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZoneChangesOperation creates a new CKFetchRecordZoneChangesOperation instance.
func NewCKFetchRecordZoneChangesOperation() CKFetchRecordZoneChangesOperation {
	return getCKFetchRecordZoneChangesOperationClass().New()
}



// Creates an operation for fetching record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/initWithRecordZoneIDs:configurationsByRecordZoneID:
func NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsConfigurationsByRecordZoneID(recordZoneIDs []CKRecordZoneID, configurationsByRecordZoneID unsafe.Pointer) CKFetchRecordZoneChangesOperation {
	instance := getCKFetchRecordZoneChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZoneChangesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:configurationsByRecordZoneID:"), recordZoneIDs, configurationsByRecordZoneID)
	rv.Autorelease()
	return rv
}


// Creates an operation for fetching record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/init(recordZoneIDs:optionsByRecordZoneID:)
func NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsOptionsByRecordZoneID(recordZoneIDs []CKRecordZoneID, optionsByRecordZoneID unsafe.Pointer) CKFetchRecordZoneChangesOperation {
	instance := getCKFetchRecordZoneChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZoneChangesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:optionsByRecordZoneID:"), recordZoneIDs, optionsByRecordZoneID)
	rv.Autorelease()
	return rv
}



// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/configurationsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) ConfigurationsByRecordZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("configurationsByRecordZoneID"))
	return rv
}


// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/configurationsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) SetConfigurationsByRecordZoneID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfigurationsByRecordZoneID:"), value)
}


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchAllChanges
func (c_ CKFetchRecordZoneChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fetchAllChanges"))
	return rv
}


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchAllChanges
func (c_ CKFetchRecordZoneChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAllChanges:"), value)
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchRecordZoneChangesCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) FetchRecordZoneChangesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZoneChangesCompletionBlock"))
	return rv
}


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchRecordZoneChangesCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) SetFetchRecordZoneChangesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZoneChangesCompletionBlock:"), value)
}


// Configuration options for each record zone that the operation retrieves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/optionsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) OptionsByRecordZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("optionsByRecordZoneID"))
	return rv
}


// Configuration options for each record zone that the operation retrieves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/optionsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) SetOptionsByRecordZoneID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptionsByRecordZoneID:"), value)
}


// The closure to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordChangedBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordChangedBlock"))
	return rv
}


// The closure to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordChangedBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangedBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWasChangedBlock-90bon
func (c_ CKFetchRecordZoneChangesOperation) RecordWasChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordWasChangedBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWasChangedBlock-90bon
func (c_ CKFetchRecordZoneChangesOperation) SetRecordWasChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWasChangedBlock:"), value)
}


// The block to execute when a record no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWithIDWasDeletedBlock-912xy
func (c_ CKFetchRecordZoneChangesOperation) RecordWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}


// The block to execute when a record no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWithIDWasDeletedBlock-912xy
func (c_ CKFetchRecordZoneChangesOperation) SetRecordWithIDWasDeletedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}


// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneChangeTokensUpdatedBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneChangeTokensUpdatedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneChangeTokensUpdatedBlock"))
	return rv
}


// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneChangeTokensUpdatedBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneChangeTokensUpdatedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneChangeTokensUpdatedBlock:"), value)
}


// The closure to execute when a record zone’s fetch finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneFetchCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneFetchCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneFetchCompletionBlock"))
	return rv
}


// The closure to execute when a record zone’s fetch finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneFetchCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneFetchCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneFetchCompletionBlock:"), value)
}


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneIDs
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDs"))
	return rv
}


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneIDs
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
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
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchrecordzonechangesresultblock
func (c_ CKFetchRecordZoneChangesOperation) FetchRecordZoneChangesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZoneChangesResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchrecordzonechangesresultblock
func (c_ CKFetchRecordZoneChangesOperation) SetFetchRecordZoneChangesResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZoneChangesResultBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzonefetchresultblock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneFetchResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneFetchResultBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzonefetchresultblock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneFetchResultBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneFetchResultBlock:"), value)
}


