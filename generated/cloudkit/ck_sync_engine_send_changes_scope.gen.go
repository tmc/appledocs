// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineSendChangesScope] class.
var (
	CKSyncEngineSendChangesScopeClass     _CKSyncEngineSendChangesScopeClass
	CKSyncEngineSendChangesScopeClassOnce sync.Once
)

func getCKSyncEngineSendChangesScopeClass() _CKSyncEngineSendChangesScopeClass {
	CKSyncEngineSendChangesScopeClassOnce.Do(func() {
		CKSyncEngineSendChangesScopeClass = _CKSyncEngineSendChangesScopeClass{objc.GetClass("CKSyncEngineSendChangesScope")}
	})
	return CKSyncEngineSendChangesScopeClass
}

type _CKSyncEngineSendChangesScopeClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineSendChangesScope] class.
type ICKSyncEngineSendChangesScope interface {
	objectivec.IObject
	ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool
	ContainsRecordID(recordID ICKRecordID) bool
	ExcludedZoneIDs() unsafe.Pointer
	RecordIDs() unsafe.Pointer
	ZoneIDs() unsafe.Pointer
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope
type CKSyncEngineSendChangesScope struct {
	objectivec.Object
}

// CKSyncEngineSendChangesScopeFrom constructs a [CKSyncEngineSendChangesScope] from an unsafe.Pointer.
func CKSyncEngineSendChangesScopeFrom(ptr unsafe.Pointer) CKSyncEngineSendChangesScope {
	return CKSyncEngineSendChangesScope{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineSendChangesScopeClass) Alloc() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineSendChangesScopeClass) New() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineSendChangesScope) Init() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineSendChangesScope) Autorelease() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesScope creates a new CKSyncEngineSendChangesScope instance.
func NewCKSyncEngineSendChangesScope() CKSyncEngineSendChangesScope {
	return getCKSyncEngineSendChangesScopeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineSendChangesScopeWithExcludedZoneIDs(excludedZoneIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), excludedZoneIDs)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithRecordIDs:
func NewCKSyncEngineSendChangesScopeWithRecordIDs(recordIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithRecordIDs:"), recordIDs)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithZoneIDs:
func NewCKSyncEngineSendChangesScopeWithZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineSendChangesScope](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsPendingRecordZoneChange:
func (c_ CKSyncEngineSendChangesScope) ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsPendingRecordZoneChange:"), pendingRecordZoneChange)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsRecordID:
func (c_ CKSyncEngineSendChangesScope) ContainsRecordID(recordID ICKRecordID) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsRecordID:"), recordID)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/excludedZoneIDs
func (c_ CKSyncEngineSendChangesScope) ExcludedZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("excludedZoneIDs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/recordIDs
func (c_ CKSyncEngineSendChangesScope) RecordIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordIDs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/zoneIDs
func (c_ CKSyncEngineSendChangesScope) ZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneIDs"))
	return rv
}


