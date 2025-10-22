// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEnginePendingRecordZoneChange] class.
var (
	CKSyncEnginePendingRecordZoneChangeClass     _CKSyncEnginePendingRecordZoneChangeClass
	CKSyncEnginePendingRecordZoneChangeClassOnce sync.Once
)

func getCKSyncEnginePendingRecordZoneChangeClass() _CKSyncEnginePendingRecordZoneChangeClass {
	CKSyncEnginePendingRecordZoneChangeClassOnce.Do(func() {
		CKSyncEnginePendingRecordZoneChangeClass = _CKSyncEnginePendingRecordZoneChangeClass{objc.GetClass("CKSyncEnginePendingRecordZoneChange")}
	})
	return CKSyncEnginePendingRecordZoneChangeClass
}

type _CKSyncEnginePendingRecordZoneChangeClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEnginePendingRecordZoneChange] class.
type ICKSyncEnginePendingRecordZoneChange interface {
	objectivec.IObject
	RecordID() CKRecordID
	Type() CKSyncEnginePendingRecordZoneChangeType
}

// Describes an unsent record modification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange
type CKSyncEnginePendingRecordZoneChange struct {
	objectivec.Object
}

// CKSyncEnginePendingRecordZoneChangeFrom constructs a [CKSyncEnginePendingRecordZoneChange] from an unsafe.Pointer.
//
// Describes an unsent record modification.
func CKSyncEnginePendingRecordZoneChangeFrom(ptr unsafe.Pointer) CKSyncEnginePendingRecordZoneChange {
	return CKSyncEnginePendingRecordZoneChange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingRecordZoneChangeClass) Alloc() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEnginePendingRecordZoneChangeClass) New() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingRecordZoneChange) Init() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingRecordZoneChange) Autorelease() CKSyncEnginePendingRecordZoneChange {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingRecordZoneChange creates a new CKSyncEnginePendingRecordZoneChange instance.
func NewCKSyncEnginePendingRecordZoneChange() CKSyncEnginePendingRecordZoneChange {
	return getCKSyncEnginePendingRecordZoneChangeClass().New()
}




// Creates a record zone change of the specified type for the given record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/initWithRecordID:type:
func NewCKSyncEnginePendingRecordZoneChangeWithRecordIDType(recordID ICKRecordID, type_ CKSyncEnginePendingRecordZoneChangeType) CKSyncEnginePendingRecordZoneChange {
	instance := getCKSyncEnginePendingRecordZoneChangeClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingRecordZoneChange](instance.ID, objc.Sel("initWithRecordID:type:"), recordID, type_)
	rv.Autorelease()
	return rv
}


// The identifier of the modified record.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/recordID
func (c_ CKSyncEnginePendingRecordZoneChange) RecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}

// The type of change to make.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingRecordZoneChange/type
func (c_ CKSyncEnginePendingRecordZoneChange) Type() CKSyncEnginePendingRecordZoneChangeType {
	rv := objc.Send[CKSyncEnginePendingRecordZoneChangeType](c_.ID, objc.Sel("type"))
	return rv
}


