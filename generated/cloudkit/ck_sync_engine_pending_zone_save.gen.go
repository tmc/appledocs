// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CKSyncEnginePendingZoneSave] class.
var (
	CKSyncEnginePendingZoneSaveClass     _CKSyncEnginePendingZoneSaveClass
	CKSyncEnginePendingZoneSaveClassOnce sync.Once
)

func getCKSyncEnginePendingZoneSaveClass() _CKSyncEnginePendingZoneSaveClass {
	CKSyncEnginePendingZoneSaveClassOnce.Do(func() {
		CKSyncEnginePendingZoneSaveClass = _CKSyncEnginePendingZoneSaveClass{objc.GetClass("CKSyncEnginePendingZoneSave")}
	})
	return CKSyncEnginePendingZoneSaveClass
}

type _CKSyncEnginePendingZoneSaveClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEnginePendingZoneSave] class.
type ICKSyncEnginePendingZoneSave interface {
	ICKSyncEnginePendingDatabaseChange
}

// An object that describes an unsent record zone modification.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave
type CKSyncEnginePendingZoneSave struct {
	CKSyncEnginePendingDatabaseChange
}

// CKSyncEnginePendingZoneSaveFrom constructs a [CKSyncEnginePendingZoneSave] from an unsafe.Pointer.
//
// An object that describes an unsent record zone modification.
func CKSyncEnginePendingZoneSaveFrom(ptr unsafe.Pointer) CKSyncEnginePendingZoneSave {
	return CKSyncEnginePendingZoneSave{
		CKSyncEnginePendingDatabaseChange: CKSyncEnginePendingDatabaseChangeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEnginePendingZoneSaveClass) Alloc() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEnginePendingZoneSaveClass) New() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEnginePendingZoneSave) Init() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEnginePendingZoneSave) Autorelease() CKSyncEnginePendingZoneSave {
	rv := objc.Send[CKSyncEnginePendingZoneSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEnginePendingZoneSave creates a new CKSyncEnginePendingZoneSave instance.
func NewCKSyncEnginePendingZoneSave() CKSyncEnginePendingZoneSave {
	return getCKSyncEnginePendingZoneSaveClass().New()
}




// Creates a pending zone save for the specified record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/initWithZone:
func NewCKSyncEnginePendingZoneSaveWithZone(zone ICKRecordZone) CKSyncEnginePendingZoneSave {
	instance := getCKSyncEnginePendingZoneSaveClass().Alloc()
	rv := objc.Send[CKSyncEnginePendingZoneSave](instance.ID, objc.Sel("initWithZone:"), zone)
	rv.Autorelease()
	return rv
}


// The record zone to save.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEnginePendingZoneSave/zone
func (c_ CKSyncEnginePendingZoneSave) Zone() CKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("zone"))
	return rv
}


