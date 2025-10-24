// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchRecordZoneChangesOperation */


/* debug [class_header]: Header for CKFetchRecordZoneChangesOperation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchRecordZoneChangesOperation */
// An interface definition for the [CKFetchRecordZoneChangesOperation] class.
type ICKFetchRecordZoneChangesOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchRecordZoneChangesOperation */
	// properties:
	ConfigurationsByRecordZoneID() foundation.IDictionary
	SetConfigurationsByRecordZoneID(value foundation.IDictionary)
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	FetchRecordZoneChangesCompletionBlock() unsafe.Pointer
	SetFetchRecordZoneChangesCompletionBlock(value unsafe.Pointer)
	OptionsByRecordZoneID() foundation.IDictionary
	SetOptionsByRecordZoneID(value foundation.IDictionary)
	RecordChangedBlock() unsafe.Pointer
	SetRecordChangedBlock(value unsafe.Pointer)
	RecordWasChangedBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetRecordWasChangedBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	RecordWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordWithIDWasDeletedBlock(value unsafe.Pointer)
	RecordZoneChangeTokensUpdatedBlock() unsafe.Pointer
	SetRecordZoneChangeTokensUpdatedBlock(value unsafe.Pointer)
	RecordZoneFetchCompletionBlock() unsafe.Pointer
	SetRecordZoneFetchCompletionBlock(value unsafe.Pointer)
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)
	FetchRecordZoneChangesResultBlock() objectivec.IObject
	SetFetchRecordZoneChangesResultBlock(value objectivec.IObject)
	RecordZoneFetchResultBlock() objectivec.IObject
	SetRecordZoneFetchResultBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchRecordZoneChangesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchRecordZoneChangesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZoneChangesOperationClass) Alloc() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchRecordZoneChangesOperation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchRecordZoneChangesOperation */

// Creates an operation for fetching record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/initWithRecordZoneIDs:configurationsByRecordZoneID:
func NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsConfigurationsByRecordZoneID(recordZoneIDs []CKRecordZoneID, configurationsByRecordZoneID foundation.IDictionary) CKFetchRecordZoneChangesOperation {
	instance := getCKFetchRecordZoneChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZoneChangesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:configurationsByRecordZoneID:"), recordZoneIDs, configurationsByRecordZoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsConfigurationsByRecordZoneID */


// Creates an operation for fetching record zone changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/init(recordZoneIDs:optionsByRecordZoneID:)
func NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsOptionsByRecordZoneID(recordZoneIDs []CKRecordZoneID, optionsByRecordZoneID foundation.IDictionary) CKFetchRecordZoneChangesOperation {
	instance := getCKFetchRecordZoneChangesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZoneChangesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:optionsByRecordZoneID:"), recordZoneIDs, optionsByRecordZoneID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchRecordZoneChangesOperationWithRecordZoneIDsOptionsByRecordZoneID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchRecordZoneChangesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchRecordZoneChangesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchRecordZoneChangesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchRecordZoneChangesOperation */

// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/configurationsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) ConfigurationsByRecordZoneID() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("configurationsByRecordZoneID"))
	return rv
}/* debug [instance_properties/getter]: configurationsByRecordZoneID */


// A dictionary of configurations for fetching change operations by zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/configurationsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) SetConfigurationsByRecordZoneID(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfigurationsByRecordZoneID:"), value)
}/* debug [instance_properties/setter]: configurationsByRecordZoneID */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchAllChanges
func (c_ CKFetchRecordZoneChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("fetchAllChanges"))
	return rv
}/* debug [instance_properties/getter]: fetchAllChanges */


// A Boolean value that indicates whether to send repeated requests to the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchAllChanges
func (c_ CKFetchRecordZoneChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchAllChanges:"), value)
}/* debug [instance_properties/setter]: fetchAllChanges */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchRecordZoneChangesCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) FetchRecordZoneChangesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZoneChangesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchRecordZoneChangesCompletionBlock */


// The closure to execute when the operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchRecordZoneChangesCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) SetFetchRecordZoneChangesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZoneChangesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchRecordZoneChangesCompletionBlock */


// Configuration options for each record zone that the operation retrieves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/optionsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) OptionsByRecordZoneID() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("optionsByRecordZoneID"))
	return rv
}/* debug [instance_properties/getter]: optionsByRecordZoneID */


// Configuration options for each record zone that the operation retrieves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/optionsByRecordZoneID
func (c_ CKFetchRecordZoneChangesOperation) SetOptionsByRecordZoneID(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptionsByRecordZoneID:"), value)
}/* debug [instance_properties/setter]: optionsByRecordZoneID */


// The closure to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordChangedBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordChangedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordChangedBlock */


// The closure to execute with the contents of a changed record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordChangedBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordChangedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordChangedBlock:"), value)
}/* debug [instance_properties/setter]: recordChangedBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWasChangedBlock-90bon
func (c_ CKFetchRecordZoneChangesOperation) RecordWasChangedBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("recordWasChangedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordWasChangedBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWasChangedBlock-90bon
func (c_ CKFetchRecordZoneChangesOperation) SetRecordWasChangedBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWasChangedBlock:"), value)
}/* debug [instance_properties/setter]: recordWasChangedBlock */


// The block to execute when a record no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWithIDWasDeletedBlock-912xy
func (c_ CKFetchRecordZoneChangesOperation) RecordWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordWithIDWasDeletedBlock */


// The block to execute when a record no longer exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordWithIDWasDeletedBlock-912xy
func (c_ CKFetchRecordZoneChangesOperation) SetRecordWithIDWasDeletedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}/* debug [instance_properties/setter]: recordWithIDWasDeletedBlock */


// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneChangeTokensUpdatedBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneChangeTokensUpdatedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneChangeTokensUpdatedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneChangeTokensUpdatedBlock */


// The closure to execute when the change token updates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneChangeTokensUpdatedBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneChangeTokensUpdatedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneChangeTokensUpdatedBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneChangeTokensUpdatedBlock */


// The closure to execute when a record zone’s fetch finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneFetchCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneFetchCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordZoneFetchCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneFetchCompletionBlock */


// The closure to execute when a record zone’s fetch finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneFetchCompletionBlock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneFetchCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneFetchCompletionBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneFetchCompletionBlock */


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneIDs
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: recordZoneIDs */


// The IDs of the record zones that contain the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneIDs
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
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
}/* debug [instance_properties/setter]: recordZoneIDs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchrecordzonechangesresultblock
func (c_ CKFetchRecordZoneChangesOperation) FetchRecordZoneChangesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchRecordZoneChangesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchRecordZoneChangesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/fetchrecordzonechangesresultblock
func (c_ CKFetchRecordZoneChangesOperation) SetFetchRecordZoneChangesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZoneChangesResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchRecordZoneChangesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzonefetchresultblock
func (c_ CKFetchRecordZoneChangesOperation) RecordZoneFetchResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("recordZoneFetchResultBlock"))
	return rv
}/* debug [instance_properties/getter]: recordZoneFetchResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordzonefetchresultblock
func (c_ CKFetchRecordZoneChangesOperation) SetRecordZoneFetchResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneFetchResultBlock:"), value)
}/* debug [instance_properties/setter]: recordZoneFetchResultBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchRecordZoneChangesOperation */


