// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKFetchRecordZonesOperation */


/* debug [class_header]: Header for CKFetchRecordZonesOperation */
// The class instance for the [CKFetchRecordZonesOperation] class.
var (
	CKFetchRecordZonesOperationClass     _CKFetchRecordZonesOperationClass
	CKFetchRecordZonesOperationClassOnce sync.Once
)

func getCKFetchRecordZonesOperationClass() _CKFetchRecordZonesOperationClass {
	CKFetchRecordZonesOperationClassOnce.Do(func() {
		CKFetchRecordZonesOperationClass = _CKFetchRecordZonesOperationClass{objc.GetClass("CKFetchRecordZonesOperation")}
	})
	return CKFetchRecordZonesOperationClass
}

type _CKFetchRecordZonesOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKFetchRecordZonesOperation */
// An interface definition for the [CKFetchRecordZonesOperation] class.
type ICKFetchRecordZonesOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKFetchRecordZonesOperation */
	// properties:
	FetchRecordZonesCompletionBlock() unsafe.Pointer
	SetFetchRecordZonesCompletionBlock(value unsafe.Pointer)
	PerRecordZoneCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerRecordZoneCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)
	FetchRecordZonesResultBlock() objectivec.IObject
	SetFetchRecordZonesResultBlock(value objectivec.IObject)
	PerRecordZoneResultBlock() objectivec.IObject
	SetPerRecordZoneResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKFetchRecordZonesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKFetchRecordZonesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKFetchRecordZonesOperationClass) Alloc() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKFetchRecordZonesOperationClass) New() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKFetchRecordZonesOperation) Init() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKFetchRecordZonesOperation) Autorelease() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZonesOperation creates a new CKFetchRecordZonesOperation instance.
func NewCKFetchRecordZonesOperation() CKFetchRecordZonesOperation {
	return getCKFetchRecordZonesOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKFetchRecordZonesOperation */
// An operation for retrieving record zones from a database.
//
// Use this operation object to fetch record zones so that you can ascertain their capabilities. If you assign a handler to the property of the operation, CloudKit calls it after the operation executes and returns its results. You can use the handler to perform any housekeeping tasks that relate to the operation, but don’t use it to process the results of the operation. The handler you specify should manage any failures, whether due to an error or an explicit cancellation.


// An operation for retrieving record zones from a database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation
type CKFetchRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordZonesOperationFrom constructs a [CKFetchRecordZonesOperation] from an unsafe.Pointer.
//
// An operation for retrieving record zones from a database.
func CKFetchRecordZonesOperationFrom(ptr unsafe.Pointer) CKFetchRecordZonesOperation {
	return CKFetchRecordZonesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKFetchRecordZonesOperation */

// Creates an operation for fetching the specified record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/init(recordZoneIDs:)
func NewCKFetchRecordZonesOperationWithRecordZoneIDs(zoneIDs []CKRecordZoneID) CKFetchRecordZonesOperation {
	instance := getCKFetchRecordZonesOperationClass().Alloc()
	rv := objc.Send[CKFetchRecordZonesOperation](instance.ID, objc.Sel("initWithRecordZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKFetchRecordZonesOperationWithRecordZoneIDs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKFetchRecordZonesOperation */

// Returns an operation for fetching all record zones in the current database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchAllRecordZonesOperation()
func (cc _CKFetchRecordZonesOperationClass) FetchAllRecordZonesOperation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("fetchAllRecordZonesOperation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FetchAllRecordZonesOperation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKFetchRecordZonesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKFetchRecordZonesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKFetchRecordZonesOperation */

// The closure to execute after CloudKit retrieves all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (c_ CKFetchRecordZonesOperation) FetchRecordZonesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fetchRecordZonesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchRecordZonesCompletionBlock */


// The closure to execute after CloudKit retrieves all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (c_ CKFetchRecordZonesOperation) SetFetchRecordZonesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZonesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: fetchRecordZonesCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/perRecordZoneCompletionBlock
func (c_ CKFetchRecordZonesOperation) PerRecordZoneCompletionBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perRecordZoneCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordZoneCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/perRecordZoneCompletionBlock
func (c_ CKFetchRecordZonesOperation) SetPerRecordZoneCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perRecordZoneCompletionBlock */


// The IDs of the record zones to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/recordZoneIDs
func (c_ CKFetchRecordZonesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDs"))
	return rv
}/* debug [instance_properties/getter]: recordZoneIDs */


// The IDs of the record zones to retrieve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/recordZoneIDs
func (c_ CKFetchRecordZonesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
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
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/fetchrecordzonesresultblock
func (c_ CKFetchRecordZonesOperation) FetchRecordZonesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("fetchRecordZonesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: fetchRecordZonesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/fetchrecordzonesresultblock
func (c_ CKFetchRecordZonesOperation) SetFetchRecordZonesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFetchRecordZonesResultBlock:"), value)
}/* debug [instance_properties/setter]: fetchRecordZonesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/perrecordzoneresultblock
func (c_ CKFetchRecordZonesOperation) PerRecordZoneResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perRecordZoneResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordZoneResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonesoperation/perrecordzoneresultblock
func (c_ CKFetchRecordZonesOperation) SetPerRecordZoneResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneResultBlock:"), value)
}/* debug [instance_properties/setter]: perRecordZoneResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordZonesOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKFetchRecordZonesOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKFetchRecordZonesOperation */


