// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CKQueryOperation] class.
type ICKQueryOperation interface {
	ICKDatabaseOperation
}

// An operation for executing queries in a database.
//
// A object is a concrete operation that you can use to execute queries. A query operation applies query parameters to the specified database and record zone, delivering any matching records asynchronously to the handlers that you provide. To perform a new search: Initialize a object with a object that contains the search criteria and sorting information for the records you want. Assign a handler to the property so that you can process the results and execute the operation. If the search yields many records, the operation object may deliver a portion of the total results to your blocks immediately, along with a cursor for obtaining the remaining records. Use the cursor to initialize and execute a separate instance when you’re ready to process the next batch of results. 3. Optionally, configure the results by specifying values for the and properties. 4. Pass the query operation object to the method of the target database to execute the operation. CloudKit restricts queries to the records in a single record zone. For new queries, you specify the zone when you initialize the query operation object. For cursor-based queries, the cursor contains the zone information. To search for records in multiple zones, you must create a separate object for each zone you want to search, although you can initialize each of them with the same object. If you assign a handler to the operation’s property, the operation calls it after it executes and returns any results. Use a handler to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. The handler you provide should manage any failures, whether due to an error or an explicit cancellation.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKQueryOperationClass) Alloc() CKQueryOperation {
	rv := objc.Send[CKQueryOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an operation with additional results from a previous search.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(cursor:)
func NewCKQueryOperationWithCursor(cursor unsafe.Pointer) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[CKQueryOperation](instance.ID, objc.Sel("initWithCursor:"), cursor)
	rv.Autorelease()
	return rv
}



// Creates an operation that searches for records in the specified record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/init(query:)
func NewCKQueryOperationWithQuery(query unsafe.Pointer) CKQueryOperation {
	instance := getCKQueryOperationClass().Alloc()
	rv := objc.Send[CKQueryOperation](instance.ID, objc.Sel("initWithQuery:"), query)
	rv.Autorelease()
	return rv
}


// The cursor for continuing the search.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/cursor-swift.property
func (c_ CKQueryOperation) Cursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cursor"))
	return rv
}


// SetCursor sets the value of the cursor property.
// The cursor for continuing the search.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/cursor-swift.property
func (c_ CKQueryOperation) SetCursor(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCursor:"), value)
}

// The fields of the records to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/desiredKeys-4a6vy
func (c_ CKQueryOperation) DesiredKeys() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The fields of the records to fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/desiredKeys-4a6vy
func (c_ CKQueryOperation) SetDesiredKeys(value []string) {
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

// The query for the search.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/query
func (c_ CKQueryOperation) Query() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("query"))
	return rv
}


// SetQuery sets the value of the query property.
// The query for the search.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/query
func (c_ CKQueryOperation) SetQuery(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuery:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordMatchedBlock-7kek0
func (c_ CKQueryOperation) RecordMatchedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordMatchedBlock"))
	return rv
}


// SetRecordMatchedBlock sets the value of the recordMatchedBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/recordMatchedBlock-7kek0
func (c_ CKQueryOperation) SetRecordMatchedBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordMatchedBlock:"), value)
}

// The maximum number of records to return at one time.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/resultsLimit
func (c_ CKQueryOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of records to return at one time.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/resultsLimit
func (c_ CKQueryOperation) SetResultsLimit(value uint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}

// The ID of the record zone that contains the records to search.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/zoneID
func (c_ CKQueryOperation) ZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneID"))
	return rv
}


// SetZoneID sets the value of the zoneID property.
// The ID of the record zone that contains the records to search.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/zoneID
func (c_ CKQueryOperation) SetZoneID(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}


