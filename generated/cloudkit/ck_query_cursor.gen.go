// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKQueryCursor */


/* debug [class_header]: Header for CKQueryCursor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKQueryCursor */
// An interface definition for the [CKQueryCursor] class.
type ICKQueryCursor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CKQueryCursor */
	// properties:
	Cursor() ICKQueryCursor
	SetCursor(value ICKQueryCursor)
	DesiredKeys() objectivec.IObject
	SetDesiredKeys(value objectivec.IObject)
	Query() objc.IObject /* cross-framework: CKQuery */
	SetQuery(value objc.IObject /* cross-framework: CKQuery */)
	ResultsLimit() int
	SetResultsLimit(value int)
	ZoneID() ICKRecordZoneID
	SetZoneID(value ICKRecordZoneID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKQueryCursor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKQueryCursor */
// Alloc allocates a new instance without initialization.
func (cc _CKQueryCursorClass) Alloc() CKQueryCursor {
	rv := objc.Send[CKQueryCursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKQueryCursor */
// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
//
// You don’t create instances of this class yourself. When fetching records using a query operation, if the number of results exceeds the limit for the query, CloudKit provides a cursor. Use that cursor to create a new instance of and retrieve the next batch of results for the same query. For information about how to use a object, see .


// An object that marks the stopping point for a query and the starting point for retrieving the remaining results.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKQueryCursor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKQueryCursor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKQueryCursor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKQueryCursor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKQueryCursor */

// The cursor for continuing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/cursor-swift.property
func (c_ CKQueryCursor) Cursor() ICKQueryCursor {
	rv := objc.Send[CKQueryCursor](c_.ID, objc.Sel("cursor"))
	return rv
}/* debug [instance_properties/getter]: cursor */


// The cursor for continuing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/cursor-swift.property
func (c_ CKQueryCursor) SetCursor(value ICKQueryCursor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCursor:"), value)
}/* debug [instance_properties/setter]: cursor */


// The fields of the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c_ CKQueryCursor) DesiredKeys() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("desiredKeys"))
	return rv
}/* debug [instance_properties/getter]: desiredKeys */


// The fields of the records to fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/desiredkeys-7qrse
func (c_ CKQueryCursor) SetDesiredKeys(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDesiredKeys:"), value)
}/* debug [instance_properties/setter]: desiredKeys */


// The query for the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/query
func (c_ CKQueryCursor) Query() objc.IObject /* cross-framework: CKQuery */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("query"))
	return rv
}/* debug [instance_properties/getter]: query */


// The query for the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/query
func (c_ CKQueryCursor) SetQuery(value objc.IObject /* cross-framework: CKQuery */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setQuery:"), value)
}/* debug [instance_properties/setter]: query */


// The maximum number of records to return at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/resultslimit
func (c_ CKQueryCursor) ResultsLimit() int {
	rv := objc.Send[int](c_.ID, objc.Sel("resultsLimit"))
	return rv
}/* debug [instance_properties/getter]: resultsLimit */


// The maximum number of records to return at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/resultslimit
func (c_ CKQueryCursor) SetResultsLimit(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResultsLimit:"), value)
}/* debug [instance_properties/setter]: resultsLimit */


// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQueryCursor) ZoneID() ICKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}/* debug [instance_properties/getter]: zoneID */


// The ID of the record zone that contains the records to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckqueryoperation/zoneid
func (c_ CKQueryCursor) SetZoneID(value ICKRecordZoneID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setZoneID:"), value)
}/* debug [instance_properties/setter]: zoneID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKQueryCursor */



