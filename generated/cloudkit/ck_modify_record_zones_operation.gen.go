// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKModifyRecordZonesOperation */


/* debug [class_header]: Header for CKModifyRecordZonesOperation */
// The class instance for the [CKModifyRecordZonesOperation] class.
var (
	CKModifyRecordZonesOperationClass     _CKModifyRecordZonesOperationClass
	CKModifyRecordZonesOperationClassOnce sync.Once
)

func getCKModifyRecordZonesOperationClass() _CKModifyRecordZonesOperationClass {
	CKModifyRecordZonesOperationClassOnce.Do(func() {
		CKModifyRecordZonesOperationClass = _CKModifyRecordZonesOperationClass{objc.GetClass("CKModifyRecordZonesOperation")}
	})
	return CKModifyRecordZonesOperationClass
}

type _CKModifyRecordZonesOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKModifyRecordZonesOperation */
// An interface definition for the [CKModifyRecordZonesOperation] class.
type ICKModifyRecordZonesOperation interface {
	ICKDatabaseOperation
	
/* debug [class_interface_properties]: Properties for CKModifyRecordZonesOperation */
	// properties:
	ModifyRecordZonesCompletionBlock() unsafe.Pointer
	SetModifyRecordZonesCompletionBlock(value unsafe.Pointer)
	PerRecordZoneDeleteBlock() func(unsafe.Pointer, unsafe.Pointer)
	SetPerRecordZoneDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer))
	PerRecordZoneSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	SetPerRecordZoneSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer))
	RecordZoneIDsToDelete() []CKRecordZoneID
	SetRecordZoneIDsToDelete(value []CKRecordZoneID)
	RecordZonesToSave() []CKRecordZone
	SetRecordZonesToSave(value []CKRecordZone)
	ModifyRecordZonesResultBlock() objectivec.IObject
	SetModifyRecordZonesResultBlock(value objectivec.IObject)
	CompletionBlock() objectivec.IObject
	SetCompletionBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKModifyRecordZonesOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKModifyRecordZonesOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKModifyRecordZonesOperationClass) Alloc() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKModifyRecordZonesOperationClass) New() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKModifyRecordZonesOperation) Init() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKModifyRecordZonesOperation) Autorelease() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifyRecordZonesOperation creates a new CKModifyRecordZonesOperation instance.
func NewCKModifyRecordZonesOperation() CKModifyRecordZonesOperation {
	return getCKModifyRecordZonesOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKModifyRecordZonesOperation */
// An operation that modifies one or more record zones.
//
// After you create one or more record zones, use this operation to save those zones to the database. You can also use the operation to delete record zones and their records. If you assign a handler to the property of the operation, CloudKit calls the handler after the operation executes and returns its results. Use the handler to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. The handler you provide should manage any failures of the operation, whether due to an error or an explicit cancellation.


// An operation that modifies one or more record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation
type CKModifyRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKModifyRecordZonesOperationFrom constructs a [CKModifyRecordZonesOperation] from an unsafe.Pointer.
//
// An operation that modifies one or more record zones.
func CKModifyRecordZonesOperationFrom(ptr unsafe.Pointer) CKModifyRecordZonesOperation {
	return CKModifyRecordZonesOperation{
		CKDatabaseOperation: CKDatabaseOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKModifyRecordZonesOperation */

// Creates an operation for modifying the specified record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/initWithRecordZonesToSave:recordZoneIDsToDelete:
func NewCKModifyRecordZonesOperationWithRecordZonesToSaveRecordZoneIDsToDelete(recordZonesToSave []CKRecordZone, recordZoneIDsToDelete []CKRecordZoneID) CKModifyRecordZonesOperation {
	instance := getCKModifyRecordZonesOperationClass().Alloc()
	rv := objc.Send[CKModifyRecordZonesOperation](instance.ID, objc.Sel("initWithRecordZonesToSave:recordZoneIDsToDelete:"), recordZonesToSave, recordZoneIDsToDelete)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKModifyRecordZonesOperationWithRecordZonesToSaveRecordZoneIDsToDelete */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKModifyRecordZonesOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKModifyRecordZonesOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKModifyRecordZonesOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKModifyRecordZonesOperation */

// The closure to execute after CloudKit modifies all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/modifyRecordZonesCompletionBlock
func (c_ CKModifyRecordZonesOperation) ModifyRecordZonesCompletionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("modifyRecordZonesCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: modifyRecordZonesCompletionBlock */


// The closure to execute after CloudKit modifies all of the record zones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/modifyRecordZonesCompletionBlock
func (c_ CKModifyRecordZonesOperation) SetModifyRecordZonesCompletionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordZonesCompletionBlock:"), value)
}/* debug [instance_properties/setter]: modifyRecordZonesCompletionBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneDeleteBlock-27i5g
func (c_ CKModifyRecordZonesOperation) PerRecordZoneDeleteBlock() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perRecordZoneDeleteBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordZoneDeleteBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneDeleteBlock-27i5g
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneDeleteBlock:"), value)
}/* debug [instance_properties/setter]: perRecordZoneDeleteBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneSaveBlock-3txst
func (c_ CKModifyRecordZonesOperation) PerRecordZoneSaveBlock() func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perRecordZoneSaveBlock"))
	return rv
}/* debug [instance_properties/getter]: perRecordZoneSaveBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneSaveBlock-3txst
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneSaveBlock(value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneSaveBlock:"), value)
}/* debug [instance_properties/setter]: perRecordZoneSaveBlock */


// The IDs of the record zones to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZoneIDsToDelete
func (c_ CKModifyRecordZonesOperation) RecordZoneIDsToDelete() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDsToDelete"))
	return rv
}/* debug [instance_properties/getter]: recordZoneIDsToDelete */


// The IDs of the record zones to delete permanently from the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZoneIDsToDelete
func (c_ CKModifyRecordZonesOperation) SetRecordZoneIDsToDelete(value []CKRecordZoneID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneIDsToDelete:"), nsArray)
}/* debug [instance_properties/setter]: recordZoneIDsToDelete */


// The record zones to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZonesToSave
func (c_ CKModifyRecordZonesOperation) RecordZonesToSave() []CKRecordZone {
	rv := objc.Send[[]CKRecordZone](c_.ID, objc.Sel("recordZonesToSave"))
	return rv
}/* debug [instance_properties/getter]: recordZonesToSave */


// The record zones to save to the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZonesToSave
func (c_ CKModifyRecordZonesOperation) SetRecordZonesToSave(value []CKRecordZone) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZonesToSave:"), nsArray)
}/* debug [instance_properties/setter]: recordZonesToSave */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonesresultblock
func (c_ CKModifyRecordZonesOperation) ModifyRecordZonesResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("modifyRecordZonesResultBlock"))
	return rv
}/* debug [instance_properties/getter]: modifyRecordZonesResultBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonesresultblock
func (c_ CKModifyRecordZonesOperation) SetModifyRecordZonesResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setModifyRecordZonesResultBlock:"), value)
}/* debug [instance_properties/setter]: modifyRecordZonesResultBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordZonesOperation) CompletionBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("completionBlock"))
	return rv
}/* debug [instance_properties/getter]: completionBlock */


// The block to execute after the operation’s main task is completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
func (c_ CKModifyRecordZonesOperation) SetCompletionBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionBlock:"), value)
}/* debug [instance_properties/setter]: completionBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKModifyRecordZonesOperation */


