// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKQueryOperation */


/* debug [class_header]: Header for CKQueryOperation */
// The class instance for the [CKQueryOperation] class.
var (
	CKQueryOperationClass     _CKQueryOperationClass
	CKQueryOperationClassOnce sync.Once
)

func getCKQueryOperationClass() _CKQueryOperationClass {
	CKQueryOperationClassOnce.Do(func() {
		CKQueryOperationClass = _CKQueryOperationClass{objc.GetClass("CKQueryOperation")}
	})
	return CKQueryOperationClass
}

type _CKQueryOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKQueryOperation */
// An interface definition for the [CKQueryOperation] class.
type ICKQueryOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKQueryOperation */
	// properties:
	Cursor() ICKQueryCursor
	SetCursor(value ICKQueryCursor)
	DesiredKeys() []string
	SetDesiredKeys(value []string)
	Query() objc.IObject /* cross-framework: CKQuery */
	SetQuery(value objc.IObject /* cross-framework: CKQuery */)
	QueryCompletionBlock() unsafe.Pointer
	SetQueryCompletionBlock(value unsafe.Pointer)
	RecordFetchedBlock() unsafe.Pointer
	SetRecordFetchedBlock(value unsafe.Pointer)
	RecordMatchedBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetRecordMatchedBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	ResultsLimit() uint
	SetResultsLimit(value uint)
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
	QueryResultBlock() objectivec.IObject
	SetQueryResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKQueryOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKQueryOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKQueryOperationClass) Alloc() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKQueryOperationClass) New() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQueryOperation) Init() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQueryOperation) Autorelease() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryOperation creates a new CKQueryOperation instance.
func NewCKQueryOperation() CKQueryOperation {
	return getCKQueryOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKQueryOperation */
// An operation for executing queries in a database.
//
// A object is a concrete operation that you can use to execute queries. A query operation applies query parameters to the specified database and record zone, delivering any matching records asynchronously to the handlers that you provide. To perform a new search: Initialize a object with a object that contains the search criteria and sorting information for the records you want. Assign a handler to the property so that you can process the results and execute the operation. If the search yields many records, the operation object may deliver a portion of the total results to your blocks immediately, along with a cursor for obtaining the remaining records. Use the cursor to initialize and execute a separate instance when you’re ready to process the next batch of results. 3. Optionally, configure the results by specifying values for the and properties. 4. Pass the query operation object to the method of the target database to execute the operation. CloudKit restricts queries to the records in a single record zone. For new queries, you specify the zone when you initialize the query operation object. For cursor-based queries, the cursor contains the zone information. To search for records in multiple zones, you must create a separate object for each zone you want to search, although you can initialize each of them with the same object. If you assign a handler to the operation’s property, the operation calls it after it executes and returns any results. Use a handler to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. The handler you provide should manage any failures, whether due to an error or an explicit cancellation.


// An operation for executing queries in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation
type CKQueryOperation struct {
	CKDatabaseOperation
}

// CKQueryOperationFrom constructs a [CKQueryOperation] from an unsafe.Pointer.
//
// An operation for executing queries in a database.
func CKQueryOperationFrom(ptr unsafe.Pointer) CKQueryOperation {
	return CKQueryOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKQueryOperation */

// Creates an operation with additional results from a previous search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(cursor:)
func NewCKQueryOperationWithCursor(cursor ICKQueryCursor) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[CKQueryOperation](instance.ID, objc.Sel("initWithCursor:"), cursor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKQueryOperationWithCursor */


// Creates an operation that searches for records in the specified record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(query:)
func NewCKQueryOperationWithQuery(query objc.IObject /* cross-framework: CKQuery */) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[CKQueryOperation](instance.ID, objc.Sel("initWithQuery:"), query)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKQueryOperationWithQuery */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKQueryOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKQueryOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKQueryOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKQueryOperation */

// The cursor for continuing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/cursor-swift.property
func (c_ CKQueryOperation) Cursor() ICKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("cursor"))
	return rv
}/* debug [instance_properties/getter]: cursor */


// The cursor for continuing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/cursor-swift.property
func (c_ CKQueryOperation) SetCursor(value ICKQueryCursor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCursor:"), value)
}/* debug [instance_properties/setter]: cursor */


// The fields of the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/desiredKeys-4a6vy
func (c_ CKQueryOperation) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The fields of the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/desiredKeys-4a6vy
func (c_ CKQueryOperation) SetDesiredKeys(value []string) {
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


// The query for the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/query
func (c_ CKQueryOperation) Query() objc.IObject /* cross-framework: CKQuery */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("query"))
	return rv
}/* debug [instance_properties/getter]: query */


// The query for the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/query
func (c_ CKQueryOperation) SetQuery(value objc.IObject /* cross-framework: CKQuery */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuery:"), value)
}/* debug [instance_properties/setter]: query */


// The closure to execute after CloudKit retrieves all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
func (c_ CKQueryOperation) QueryCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("queryCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: queryCompletionBlock */


// The closure to execute after CloudKit retrieves all of the records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/queryCompletionBlock
func (c_ CKQueryOperation) SetQueryCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQueryCompletionBlock:"), value)
}/* debug [instance_properties/setter]: queryCompletionBlock */


// The closure to execute when a record becomes available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordFetchedBlock
func (c_ CKQueryOperation) RecordFetchedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordFetchedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordFetchedBlock */


// The closure to execute when a record becomes available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordFetchedBlock
func (c_ CKQueryOperation) SetRecordFetchedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordFetchedBlock:"), value)
}/* debug [instance_properties/setter]: recordFetchedBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordMatchedBlock-7kek0
func (c_ CKQueryOperation) RecordMatchedBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("recordMatchedBlock"))
	return rv
}/* debug [instance_properties/getter]: recordMatchedBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordMatchedBlock-7kek0
func (c_ CKQueryOperation) SetRecordMatchedBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordMatchedBlock:"), value)
}/* debug [instance_properties/setter]: recordMatchedBlock */


// The maximum number of records to return at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/resultsLimit
func (c_ CKQueryOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of records to return at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/resultsLimit
func (c_ CKQueryOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */


// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/zoneID
func (c_ CKQueryOperation) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */


// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/zoneID
func (c_ CKQueryOperation) SetZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}/* debug [instance_properties/setter]: zoneID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/queryresultblock
func (c_ CKQueryOperation) QueryResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("queryResultBlock"))
	return rv
}/* debug [instance_properties/getter]: queryResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/queryresultblock
func (c_ CKQueryOperation) SetQueryResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQueryResultBlock:"), value)
}/* debug [instance_properties/setter]: queryResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKQueryOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKQueryOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKQueryOperation */


