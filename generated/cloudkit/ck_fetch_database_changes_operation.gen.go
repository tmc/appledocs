// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKFetchDatabaseChangesOperation] class.
var (
	CKFetchDatabaseChangesOperationClass     _CKFetchDatabaseChangesOperationClass
	CKFetchDatabaseChangesOperationClassOnce sync.Once
)

func getCKFetchDatabaseChangesOperationClass() _CKFetchDatabaseChangesOperationClass {
	CKFetchDatabaseChangesOperationClassOnce.Do(func() {
		CKFetchDatabaseChangesOperationClass = _CKFetchDatabaseChangesOperationClass{objc.GetClass("CKFetchDatabaseChangesOperation")}
	})
	return CKFetchDatabaseChangesOperationClass
}

type _CKFetchDatabaseChangesOperationClass struct {
	class objc.Class
}

// An interface definition for the [CKFetchDatabaseChangesOperation] class.
type ICKFetchDatabaseChangesOperation interface {
	ICKDatabaseOperation
}

// An operation that fetches database changes.
//
// Use this operation to fetch all record zone changes in a database. This includes new record zones, changed zones — including deleted or purged zones — and zones that contain record changes. When you create the operation, you provide a server change token, which is an opaque token that represents a specific point in the database’s history. CloudKit returns only the changes that occur after that point. For your app’s first fetch, or to refetch every change in the database’s history, use instead. The operation yields new change tokens during its execution, and issues a final change token when it completes without error. The change tokens conform to and are safe to cache on-disk. This operation’s tokens aren’t compatible with so you should segregate them in your cache. Don’t infer any behavior or order from the tokens’ contents. When your app launches for the first time, use this operation to fetch all the database’s changes. Cache the results on-device and use to subscribe to future changes. Fetch those changes on receipt of the push notifications the subscription generates. It’s not necessary to perform a fetch each time your app launches, or to schedule fetches at regular intervals. The operation calls for each zone that contains record changes. It also calls it for new and modified record zones. Store the IDs that CloudKit provides to this callback. Use those IDs with to fetch the corresponding changes. There are similar callbacks for deleted and purged record zones. To run the operation, add it to the corresponding database’s operation queue. The operation executes its callbacks on a private serial queue. The following example shows how to create the operation, configure its callbacks, and execute it. For brevity, it omits the delete and purge callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation
type CKFetchDatabaseChangesOperation struct {
	CKDatabaseOperation
}

// CKFetchDatabaseChangesOperationFrom constructs a [CKFetchDatabaseChangesOperation] from an unsafe.Pointer.
//
// An operation that fetches database changes.
func CKFetchDatabaseChangesOperationFrom(ptr unsafe.Pointer) CKFetchDatabaseChangesOperation {
	return CKFetchDatabaseChangesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKFetchDatabaseChangesOperationClass) Alloc() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKFetchDatabaseChangesOperationClass) New() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchDatabaseChangesOperation) Init() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchDatabaseChangesOperation) Autorelease() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchDatabaseChangesOperation creates a new CKFetchDatabaseChangesOperation instance.
func NewCKFetchDatabaseChangesOperation() CKFetchDatabaseChangesOperation {
	return getCKFetchDatabaseChangesOperationClass().New()
}


// Creates an operation for fetching database changes.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/init(previousServerChangeToken:)
func NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken(previousServerChangeToken unsafe.Pointer) CKFetchDatabaseChangesOperation {
	instance := getCKFetchDatabaseChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchDatabaseChangesOperation](instance.ID, objc.Sel("initWithPreviousServerChangeToken:"), previousServerChangeToken)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchAllChanges
func (c_ CKFetchDatabaseChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fetchAllChanges"))
	return rv
}


// SetFetchAllChanges sets the value of the fetchAllChanges property.
// A Boolean value that indicates whether to send repeated requests to the server.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchAllChanges
func (c_ CKFetchDatabaseChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAllChanges:"), value)
}
// The server change token.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/previousServerChangeToken
func (c_ CKFetchDatabaseChangesOperation) PreviousServerChangeToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}


// SetPreviousServerChangeToken sets the value of the previousServerChangeToken property.
// The server change token.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/previousServerChangeToken
func (c_ CKFetchDatabaseChangesOperation) SetPreviousServerChangeToken(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}
// The maximum number of results that the operation fetches.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/resultsLimit
func (c_ CKFetchDatabaseChangesOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of results that the operation fetches.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/resultsLimit
func (c_ CKFetchDatabaseChangesOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}

