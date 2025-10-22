// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEnginePendingDatabaseChange] class.
var (
	CKSyncEnginePendingDatabaseChangeClass     _CKSyncEnginePendingDatabaseChangeClass
	CKSyncEnginePendingDatabaseChangeClassOnce sync.Once
)

func getCKSyncEnginePendingDatabaseChangeClass() _CKSyncEnginePendingDatabaseChangeClass {
	CKSyncEnginePendingDatabaseChangeClassOnce.Do(func() {
		CKSyncEnginePendingDatabaseChangeClass = _CKSyncEnginePendingDatabaseChangeClass{objc.GetClass("CKSyncEnginePendingDatabaseChange")}
	})
	return CKSyncEnginePendingDatabaseChangeClass
}

type _CKSyncEnginePendingDatabaseChangeClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEnginePendingDatabaseChange] class.
type ICKSyncEnginePendingDatabaseChange interface {
	objectivec.IObject
	Type() CKSyncEnginePendingDatabaseChangeType
	ZoneID() CKRecordZoneID
}

// An object that describes an unsent database modification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange
type CKSyncEnginePendingDatabaseChange struct {
	objectivec.Object
}

// CKSyncEnginePendingDatabaseChangeFrom constructs a [CKSyncEnginePendingDatabaseChange] from an unsafe.Pointer.
//
// An object that describes an unsent database modification.
func CKSyncEnginePendingDatabaseChangeFrom(ptr unsafe.Pointer) CKSyncEnginePendingDatabaseChange {
	return CKSyncEnginePendingDatabaseChange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingDatabaseChangeClass) Alloc() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEnginePendingDatabaseChangeClass) New() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingDatabaseChange) Init() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingDatabaseChange) Autorelease() CKSyncEnginePendingDatabaseChange {
	rv := objc.Send[CKSyncEnginePendingDatabaseChange](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingDatabaseChange creates a new CKSyncEnginePendingDatabaseChange instance.
func NewCKSyncEnginePendingDatabaseChange() CKSyncEnginePendingDatabaseChange {
	return getCKSyncEnginePendingDatabaseChangeClass().New()
}


// The type of database change.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/type
func (c_ CKSyncEnginePendingDatabaseChange) Type() CKSyncEnginePendingDatabaseChangeType {
	rv := objc.Send[CKSyncEnginePendingDatabaseChangeType](c_.ID, objc.Sel("type"))
	return rv
}

// The identifier of the record zone to change.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingDatabaseChange/zoneID
func (c_ CKSyncEnginePendingDatabaseChange) ZoneID() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c_.ID, objc.Sel("zoneID"))
	return rv
}



