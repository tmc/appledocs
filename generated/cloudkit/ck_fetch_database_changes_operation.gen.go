// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchDatabaseChangesOperation */


/* debug [class_header]: Header for CKFetchDatabaseChangesOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchDatabaseChangesOperation */
// An interface definition for the [CKFetchDatabaseChangesOperation] class.
type ICKFetchDatabaseChangesOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchDatabaseChangesOperation */
	// properties:
	ChangeTokenUpdatedBlock() unsafe.Pointer
	SetChangeTokenUpdatedBlock(value unsafe.Pointer)
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	FetchDatabaseChangesCompletionBlock() unsafe.Pointer
	SetFetchDatabaseChangesCompletionBlock(value unsafe.Pointer)
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	RecordZoneWithIDChangedBlock() unsafe.Pointer
	SetRecordZoneWithIDChangedBlock(value unsafe.Pointer)
	RecordZoneWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordZoneWithIDWasDeletedBlock(value unsafe.Pointer)
	RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() unsafe.Pointer
	SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value unsafe.Pointer)
	RecordZoneWithIDWasPurgedBlock() unsafe.Pointer
	SetRecordZoneWithIDWasPurgedBlock(value unsafe.Pointer)
	ResultsLimit() uint
	SetResultsLimit(value uint)
	FetchDatabaseChangesResultBlock() objectivec.IObject
	SetFetchDatabaseChangesResultBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchDatabaseChangesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchDatabaseChangesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchDatabaseChangesOperationClass) Alloc() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchDatabaseChangesOperation */
// An operation that fetches database changes.
//
// Use this operation to fetch all record zone changes in a database. This includes new record zones, changed zones — including deleted or purged zones — and zones that contain record changes. When you create the operation, you provide a server change token, which is an opaque token that represents a specific point in the database’s history. CloudKit returns only the changes that occur after that point. For your app’s first fetch, or to refetch every change in the database’s history, use instead. The operation yields new change tokens during its execution, and issues a final change token when it completes without error. The change tokens conform to and are safe to cache on-disk. This operation’s tokens aren’t compatible with so you should segregate them in your cache. Don’t infer any behavior or order from the tokens’ contents. When your app launches for the first time, use this operation to fetch all the database’s changes. Cache the results on-device and use to subscribe to future changes. Fetch those changes on receipt of the push notifications the subscription generates. It’s not necessary to perform a fetch each time your app launches, or to schedule fetches at regular intervals. The operation calls for each zone that contains record changes. It also calls it for new and modified record zones. Store the IDs that CloudKit provides to this callback. Use those IDs with to fetch the corresponding changes. There are similar callbacks for deleted and purged record zones. To run the operation, add it to the corresponding database’s operation queue. The operation executes its callbacks on a private serial queue. The following example shows how to create the operation, configure its callbacks, and execute it. For brevity, it omits the delete and purge callbacks.


// An operation that fetches database changes.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchDatabaseChangesOperation */

// Creates an operation for fetching database changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/init(previousServerChangeToken:)
func NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken(previousServerChangeToken ICKServerChangeToken) CKFetchDatabaseChangesOperation {
	instance := getCKFetchDatabaseChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchDatabaseChangesOperation](instance.ID, objc.Sel("initWithPreviousServerChangeToken:"), previousServerChangeToken)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchDatabaseChangesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchDatabaseChangesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchDatabaseChangesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchDatabaseChangesOperation */

// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/changeTokenUpdatedBlock
func (c_ CKFetchDatabaseChangesOperation) ChangeTokenUpdatedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("changeTokenUpdatedBlock"))
	return rv
}/* debug [instance_properties/getter]: changeTokenUpdatedBlock */


// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/changeTokenUpdatedBlock
func (c_ CKFetchDatabaseChangesOperation) SetChangeTokenUpdatedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChangeTokenUpdatedBlock:"), value)
}/* debug [instance_properties/setter]: changeTokenUpdatedBlock */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchAllChanges
func (c_ CKFetchDatabaseChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fetchAllChanges"))
	return rv
}/* debug [instance_properties/getter]: fetchAllChanges */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchAllChanges
func (c_ CKFetchDatabaseChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAllChanges:"), value)
}/* debug [instance_properties/setter]: fetchAllChanges */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchDatabaseChangesCompletionBlock
func (c_ CKFetchDatabaseChangesOperation) FetchDatabaseChangesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchDatabaseChangesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchDatabaseChangesCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchDatabaseChangesCompletionBlock
func (c_ CKFetchDatabaseChangesOperation) SetFetchDatabaseChangesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchDatabaseChangesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchDatabaseChangesCompletionBlock */


// The server change token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/previousServerChangeToken
func (c_ CKFetchDatabaseChangesOperation) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}/* debug [instance_properties/getter]: previousServerChangeToken */


// The server change token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/previousServerChangeToken
func (c_ CKFetchDatabaseChangesOperation) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}/* debug [instance_properties/setter]: previousServerChangeToken */


// The closure to execute with a single record zone change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDChangedBlock
func (c_ CKFetchDatabaseChangesOperation) RecordZoneWithIDChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneWithIDChangedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneWithIDChangedBlock */


// The closure to execute with a single record zone change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDChangedBlock
func (c_ CKFetchDatabaseChangesOperation) SetRecordZoneWithIDChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneWithIDChangedBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneWithIDChangedBlock */


// The closure to execute when a record zone no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedBlock
func (c_ CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneWithIDWasDeletedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneWithIDWasDeletedBlock */


// The closure to execute when a record zone no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedBlock
func (c_ CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneWithIDWasDeletedBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneWithIDWasDeletedBlock */


// The closure to execute when a user-invoked account reset deletes a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock
func (c_ CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock */


// The closure to execute when a user-invoked account reset deletes a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock
func (c_ CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock */


// The closure to execute when CloudKit purges a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasPurgedBlock
func (c_ CKFetchDatabaseChangesOperation) RecordZoneWithIDWasPurgedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneWithIDWasPurgedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneWithIDWasPurgedBlock */


// The closure to execute when CloudKit purges a record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasPurgedBlock
func (c_ CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasPurgedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneWithIDWasPurgedBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneWithIDWasPurgedBlock */


// The maximum number of results that the operation fetches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/resultsLimit
func (c_ CKFetchDatabaseChangesOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of results that the operation fetches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/resultsLimit
func (c_ CKFetchDatabaseChangesOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/fetchdatabasechangesresultblock
func (c_ CKFetchDatabaseChangesOperation) FetchDatabaseChangesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchDatabaseChangesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchDatabaseChangesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/fetchdatabasechangesresultblock
func (c_ CKFetchDatabaseChangesOperation) SetFetchDatabaseChangesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchDatabaseChangesResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchDatabaseChangesResultBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchDatabaseChangesOperation */


