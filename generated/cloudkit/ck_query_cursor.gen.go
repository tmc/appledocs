// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKQueryCursor] class.
var (
	CKQueryCursorClass     _CKQueryCursorClass
	CKQueryCursorClassOnce sync.Once
)

func getCKQueryCursorClass() _CKQueryCursorClass {
	CKQueryCursorClassOnce.Do(func() {
		CKQueryCursorClass = _CKQueryCursorClass{objc.GetClass("CKQueryCursor")}
	})
	return CKQueryCursorClass
}

type _CKQueryCursorClass struct {
	class objc.Class
}

// An interface definition for the [CKQueryCursor] class.
type ICKQueryCursor interface {
	objectivec.IObject
}

// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
//
// You don’t create instances of this class yourself. When fetching records using a query operation, if the number of results exceeds the limit for the query, CloudKit provides a cursor. Use that cursor to create a new instance of and retrieve the next batch of results for the same query. For information about how to use a object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKQueryOperation/Cursor-swift.class
type CKQueryCursor struct {
	objectivec.Object
}

// CKQueryCursorFrom constructs a [CKQueryCursor] from an unsafe.Pointer.
//
// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
func CKQueryCursorFrom(ptr unsafe.Pointer) CKQueryCursor {
	return CKQueryCursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKQueryCursorClass) Alloc() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKQueryCursorClass) New() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKQueryCursor) Init() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKQueryCursor) Autorelease() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKQueryCursor creates a new CKQueryCursor instance.
func NewCKQueryCursor() CKQueryCursor {
	return getCKQueryCursorClass().New()
}


// The cursor for continuing the search.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/cursor-swift.property
func (c_ CKQueryCursor) Cursor() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("cursor"))
	return rv
}


// SetCursor sets the value of the cursor property.
// The cursor for continuing the search.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/cursor-swift.property
func (c_ CKQueryCursor) SetCursor(value ICKQueryCursor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCursor:"), value)
}

// The fields of the records to fetch.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c_ CKQueryCursor) DesiredKeys() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("desiredKeys"))
	return rv
}


// SetDesiredKeys sets the value of the desiredKeys property.
// The fields of the records to fetch.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c_ CKQueryCursor) SetDesiredKeys(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}

// The query for the search.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/query
func (c_ CKQueryCursor) Query() CKQuery {
	rv := objc.Send[CKQuery](c_.ID, objc.Sel("query"))
	return rv
}


// SetQuery sets the value of the query property.
// The query for the search.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/query
func (c_ CKQueryCursor) SetQuery(value ICKQuery) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuery:"), value)
}

// The maximum number of records to return at one time.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/resultslimit
func (c_ CKQueryCursor) ResultsLimit() int {
	rv := objc.Send[int](c_.ID, objc.Sel("resultsLimit"))
	return rv
}


// SetResultsLimit sets the value of the resultsLimit property.
// The maximum number of records to return at one time.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/resultslimit
func (c_ CKQueryCursor) SetResultsLimit(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}

// The ID of the record zone that contains the records to search.
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQueryCursor) ZoneID() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}


// SetZoneID sets the value of the zoneID property.
// The ID of the record zone that contains the records to search.

//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQueryCursor) SetZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}



