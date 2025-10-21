// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CKModifyRecordZonesOperation] class.
type ICKModifyRecordZonesOperation interface {
	ICKDatabaseOperation
}

// An operation that modifies one or more record zones.
//
// After you create one or more record zones, use this operation to save those zones to the database. You can also use the operation to delete record zones and their records. If you assign a handler to the property of the operation, CloudKit calls the handler after the operation executes and returns its results. Use the handler to perform housekeeping tasks for the operation, but don’t use it to process the results of the operation. The handler you provide should manage any failures of the operation, whether due to an error or an explicit cancellation.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CKModifyRecordZonesOperationClass) Alloc() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates an operation for modifying the specified record zones.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/initWithRecordZonesToSave:recordZoneIDsToDelete:
func NewCKModifyRecordZonesOperationWithRecordZonesToSaveRecordZoneIDsToDelete(recordZonesToSave unsafe.Pointer, recordZoneIDsToDelete unsafe.Pointer) CKModifyRecordZonesOperation {
	instance := getCKModifyRecordZonesOperationClass().Alloc()
	rv := objc.Send[CKModifyRecordZonesOperation](instance.ID, objc.Sel("initWithRecordZonesToSave:recordZoneIDsToDelete:"), recordZonesToSave, recordZoneIDsToDelete)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneDeleteBlock-27i5g
func (c_ CKModifyRecordZonesOperation) PerRecordZoneDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneDeleteBlock"))
	return rv
}


// SetPerRecordZoneDeleteBlock sets the value of the perRecordZoneDeleteBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneDeleteBlock-27i5g
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneDeleteBlock:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneSaveBlock-3txst
func (c_ CKModifyRecordZonesOperation) PerRecordZoneSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("perRecordZoneSaveBlock"))
	return rv
}


// SetPerRecordZoneSaveBlock sets the value of the perRecordZoneSaveBlock property.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/perRecordZoneSaveBlock-3txst
func (c_ CKModifyRecordZonesOperation) SetPerRecordZoneSaveBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerRecordZoneSaveBlock:"), value)
}

// The IDs of the record zones to delete permanently from the database.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZoneIDsToDelete
func (c_ CKModifyRecordZonesOperation) RecordZoneIDsToDelete() []CKRecordZoneID {
	rv := objc.Send[[]CKRecordZoneID](c_.ID, objc.Sel("recordZoneIDsToDelete"))
	return rv
}


// SetRecordZoneIDsToDelete sets the value of the recordZoneIDsToDelete property.
// The IDs of the record zones to delete permanently from the database.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZoneIDsToDelete
func (c_ CKModifyRecordZonesOperation) SetRecordZoneIDsToDelete(value []CKRecordZoneID) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZoneIDsToDelete:"), nsArray)
}

// The record zones to save to the database.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZonesToSave
func (c_ CKModifyRecordZonesOperation) RecordZonesToSave() []CKRecordZone {
	rv := objc.Send[[]CKRecordZone](c_.ID, objc.Sel("recordZonesToSave"))
	return rv
}


// SetRecordZonesToSave sets the value of the recordZonesToSave property.
// The record zones to save to the database.

//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZonesToSave
func (c_ CKModifyRecordZonesOperation) SetRecordZonesToSave(value []CKRecordZone) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordZonesToSave:"), nsArray)
}


