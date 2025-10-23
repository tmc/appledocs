// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineFetchChangesScope] class.
var (
	CKSyncEngineFetchChangesScopeClass     _CKSyncEngineFetchChangesScopeClass
	CKSyncEngineFetchChangesScopeClassOnce sync.Once
)

func getCKSyncEngineFetchChangesScopeClass() _CKSyncEngineFetchChangesScopeClass {
	CKSyncEngineFetchChangesScopeClassOnce.Do(func() {
		CKSyncEngineFetchChangesScopeClass = _CKSyncEngineFetchChangesScopeClass{objc.GetClass("CKSyncEngineFetchChangesScope")}
	})
	return CKSyncEngineFetchChangesScopeClass
}

type _CKSyncEngineFetchChangesScopeClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchChangesScope] class.
type ICKSyncEngineFetchChangesScope interface {
	objectivec.IObject
	// properties:
	ExcludedZoneIDs() unsafe.Pointer
	ZoneIDs() unsafe.Pointer
	// methods:
	ContainsZoneID(zoneID ICKRecordZoneID) bool /* primitive/slice/pointer. */
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope
type CKSyncEngineFetchChangesScope struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesScopeFrom constructs a [CKSyncEngineFetchChangesScope] from an unsafe.Pointer.
func CKSyncEngineFetchChangesScopeFrom(ptr unsafe.Pointer) CKSyncEngineFetchChangesScope {
	return CKSyncEngineFetchChangesScope{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchChangesScopeClass) Alloc() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchChangesScopeClass) New() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchChangesScope) Init() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchChangesScope) Autorelease() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesScope creates a new CKSyncEngineFetchChangesScope instance.
func NewCKSyncEngineFetchChangesScope() CKSyncEngineFetchChangesScope {
	return getCKSyncEngineFetchChangesScopeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithExcludedZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineFetchChangesScope](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithZoneIDs(zoneIDs unsafe.Pointer) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[CKSyncEngineFetchChangesScope](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/containsZoneID:
func (c_ CKSyncEngineFetchChangesScope) ContainsZoneID(zoneID ICKRecordZoneID) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("containsZoneID:"), zoneID)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/excludedZoneIDs
func (c_ CKSyncEngineFetchChangesScope) ExcludedZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("excludedZoneIDs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/zoneIDs
func (c_ CKSyncEngineFetchChangesScope) ZoneIDs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("zoneIDs"))
	return rv
}


