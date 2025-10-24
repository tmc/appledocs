// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchRecordChangesOperation */


/* debug [class_header]: Header for CKFetchRecordChangesOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchRecordChangesOperation */
// An interface definition for the [CKFetchRecordChangesOperation] class.
type ICKFetchRecordChangesOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchRecordChangesOperation */
	// properties:
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	FetchRecordChangesCompletionBlock() unsafe.Pointer
	SetFetchRecordChangesCompletionBlock(value unsafe.Pointer)
	MoreComing() bool
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	RecordChangedBlock() unsafe.Pointer
	SetRecordChangedBlock(value unsafe.Pointer)
	RecordWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordWithIDWasDeletedBlock(value unsafe.Pointer)
	RecordZoneID() ICKRecordZoneID
	SetRecordZoneID(value ICKRecordZoneID)
	ResultsLimit() uint
	SetResultsLimit(value uint)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchRecordChangesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchRecordChangesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordChangesOperationClass) Alloc() CKFetchRecordChangesOperation {
	rv := objc.Send[CKFetchRecordChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchRecordChangesOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchRecordChangesOperation */

// Creates an operation for fetching changes in the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/init(recordZoneID:previousServerChangeToken:)
func NewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken(recordZoneID ICKRecordZoneID, previousServerChangeToken ICKServerChangeToken) CKFetchRecordChangesOperation {
	instance := getCKFetchRecordChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordChangesOperation](instance.ID, objc.Sel("initWithRecordZoneID:previousServerChangeToken:"), recordZoneID, previousServerChangeToken)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchRecordChangesOperationWithRecordZoneIDPreviousServerChangeToken */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchRecordChangesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchRecordChangesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchRecordChangesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchRecordChangesOperation */

// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/desiredKeys
func (c_ CKFetchRecordChangesOperation) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The fields to fetch for the requested records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/desiredKeys
func (c_ CKFetchRecordChangesOperation) SetDesiredKeys(value []string) {
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
}/* debug [instance_properties/setter]: desiredKeys */


// The block to execute when the system finishes processing all changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/fetchRecordChangesCompletionBlock
func (c_ CKFetchRecordChangesOperation) FetchRecordChangesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordChangesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchRecordChangesCompletionBlock */


// The block to execute when the system finishes processing all changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/fetchRecordChangesCompletionBlock
func (c_ CKFetchRecordChangesOperation) SetFetchRecordChangesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordChangesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchRecordChangesCompletionBlock */


// A Boolean value that indicates whether more results are available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/moreComing
func (c_ CKFetchRecordChangesOperation) MoreComing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("moreComing"))
	return rv
}/* debug [instance_properties/getter]: moreComing */


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/previousServerChangeToken
func (c_ CKFetchRecordChangesOperation) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[CKServerChangeToken](c_.ID, objc.Sel("previousServerChangeToken"))
	return rv
}/* debug [instance_properties/getter]: previousServerChangeToken */


// The token that identifies the starting point for retrieving changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/previousServerChangeToken
func (c_ CKFetchRecordChangesOperation) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}/* debug [instance_properties/setter]: previousServerChangeToken */


// The block to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordChangedBlock
func (c_ CKFetchRecordChangesOperation) RecordChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordChangedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordChangedBlock */


// The block to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordChangedBlock
func (c_ CKFetchRecordChangesOperation) SetRecordChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangedBlock:"), value)
}/* debug [instance_properties/setter]: recordChangedBlock */


// The block to execute with the ID of a deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordWithIDWasDeletedBlock
func (c_ CKFetchRecordChangesOperation) RecordWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordWithIDWasDeletedBlock */


// The block to execute with the ID of a deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordWithIDWasDeletedBlock
func (c_ CKFetchRecordChangesOperation) SetRecordWithIDWasDeletedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}/* debug [instance_properties/setter]: recordWithIDWasDeletedBlock */


// The ID of the record zone with the records you want to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordZoneID
func (c_ CKFetchRecordChangesOperation) RecordZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("recordZoneID"))
	return rv
}/* debug [instance_properties/getter]: recordZoneID */


// The ID of the record zone with the records you want to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/recordZoneID
func (c_ CKFetchRecordChangesOperation) SetRecordZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneID:"), value)
}/* debug [instance_properties/setter]: recordZoneID */


// The maximum number of changed records to report with this operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/resultsLimit
func (c_ CKFetchRecordChangesOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of changed records to report with this operation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation/resultsLimit
func (c_ CKFetchRecordChangesOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordChangesOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordChangesOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchRecordChangesOperation */


