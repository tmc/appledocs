// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEnginePendingZoneDelete] class.
var (
	CKSyncEnginePendingZoneDeleteClass     _CKSyncEnginePendingZoneDeleteClass
	CKSyncEnginePendingZoneDeleteClassOnce sync.Once
)

func getCKSyncEnginePendingZoneDeleteClass() _CKSyncEnginePendingZoneDeleteClass {
	CKSyncEnginePendingZoneDeleteClassOnce.Do(func() {
		CKSyncEnginePendingZoneDeleteClass = _CKSyncEnginePendingZoneDeleteClass{objc.GetClass("CKSyncEnginePendingZoneDelete")}
	})
	return CKSyncEnginePendingZoneDeleteClass
}

type _CKSyncEnginePendingZoneDeleteClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEnginePendingZoneDelete] class.
type ICKSyncEnginePendingZoneDelete interface {
	ICKSyncEnginePendingDatabaseChange
}

// An object that describes an unsent record zone deletion.


// An object that describes an unsent record zone deletion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete

type CKSyncEnginePendingZoneDelete struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneDeleteFrom constructs a [CKSyncEnginePendingZoneDelete] from an unsafe.Pointer.
//
// An object that describes an unsent record zone deletion.
func CKSyncEnginePendingZoneDeleteFrom(ptr unsafe.Pointer) CKSyncEnginePendingZoneDelete {
	return CKSyncEnginePendingZoneDelete{
		CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingZoneDeleteClass) Alloc() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEnginePendingZoneDeleteClass) New() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingZoneDelete) Init() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingZoneDelete) Autorelease() CKSyncEnginePendingZoneDelete {
	rv := objc.Send[CKSyncEnginePendingZoneDelete](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneDelete creates a new CKSyncEnginePendingZoneDelete instance.
func NewCKSyncEnginePendingZoneDelete() CKSyncEnginePendingZoneDelete {
	return getCKSyncEnginePendingZoneDeleteClass().New()
}




// Creates a pending zone delete for the specified record zone identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneDelete/initWithZoneID:

func NewCKSyncEnginePendingZoneDeleteWithZoneID(zoneID ICKRecordZoneID) CKSyncEnginePendingZoneDelete {
	instance := getCKSyncEnginePendingZoneDeleteClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingZoneDelete](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	rv.Autorelease()
	return rv
}



