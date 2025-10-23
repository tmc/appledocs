// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKQuery] class.
var (
	CKQueryClass     _CKQueryClass
	CKQueryClassOnce sync.Once
)

func getCKQueryClass() _CKQueryClass {
	CKQueryClassOnce.Do(func() {
		CKQueryClass = _CKQueryClass{objc.GetClass("CKQuery")}
	})
	return CKQueryClass
}

type _CKQueryClass struct {
	class objc.Class
}

// An interface definition for the [CKQuery] class.
type ICKQuery interface {
	objectivec.IObject
	Predicate() foundation.Predicate
	RecordType() unsafe.Pointer
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
}

// A query that describes the criteria to apply when searching for records in a database.
//
// You create a query as the first step in the search process. The query stores the search parameters, including the type of records to search, the match criteria (predicate) to apply, and the sort parameters to apply to the results. Then you use the query to initialize an instance of , which you execute to generate the results. Always designate a record type and predicate when you create a query object. The record type narrows the scope of the search to one type of record, and the predicate defines the conditions for matching records of that type. Predicates usually compare one or more fields of a record to constant values, but you can create predicates that return all records of a specific type or perform more nuanced searches. Because you can’t change the record type and predicate after initialization, you can use the same query to initialize multiple instances of , each of which targets a different database or record zone.


// A query that describes the criteria to apply when searching for records in a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery
type CKQuery struct {
	objectivec.Object
}

// CKQueryFrom constructs a [CKQuery] from an unsafe.Pointer.
//
// A query that describes the criteria to apply when searching for records in a database.
func CKQueryFrom(ptr unsafe.Pointer) CKQuery {
	return CKQuery{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKQueryClass) Alloc() CKQuery {
	rv := objc.Send[CKQuery](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKQueryClass) New() CKQuery {
	rv := objc.Send[CKQuery](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQuery) Init() CKQuery {
	rv := objc.Send[CKQuery](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQuery) Autorelease() CKQuery {
	rv := objc.Send[CKQuery](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQuery creates a new CKQuery instance.
func NewCKQuery() CKQuery {
	return getCKQueryClass().New()
}



// Creates an operation group from a serialized instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/init(coder:)
func NewCKQueryWithCoder(aDecoder foundation.ICoder) CKQuery {
	instance := getCKQueryClass().Alloc()
	rv := objc.Send[CKQuery](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


// Creates a query with the specified record type and predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/initWithRecordType:predicate:
func NewCKQueryWithRecordTypePredicate(recordType unsafe.Pointer, predicate foundation.IPredicate) CKQuery {
	instance := getCKQueryClass().Alloc()
	rv := objc.Send[CKQuery](instance.ID, objc.Sel("initWithRecordType:predicate:"), recordType, predicate)
	rv.Autorelease()
	return rv
}



// The predicate to use for matching records.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/predicate
func (c_ CKQuery) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](c_.ID, objc.Sel("predicate"))
	return rv
}


// The record type to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/recordType-3clp
func (c_ CKQuery) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordType"))
	return rv
}


// The sort descriptors for organizing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/sortDescriptors
func (c_ CKQuery) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Send[[]foundation.SortDescriptor](c_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// The sort descriptors for organizing the query’s results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQuery/sortDescriptors
func (c_ CKQuery) SetSortDescriptors(value []foundation.ISortDescriptor) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}


