// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineFetchedZoneDeletion] class.
var (
	CKSyncEngineFetchedZoneDeletionClass     _CKSyncEngineFetchedZoneDeletionClass
	CKSyncEngineFetchedZoneDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedZoneDeletionClass() _CKSyncEngineFetchedZoneDeletionClass {
	CKSyncEngineFetchedZoneDeletionClassOnce.Do(func() {
		CKSyncEngineFetchedZoneDeletionClass = _CKSyncEngineFetchedZoneDeletionClass{objc.GetClass("CKSyncEngineFetchedZoneDeletion")}
	})
	return CKSyncEngineFetchedZoneDeletionClass
}

type _CKSyncEngineFetchedZoneDeletionClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchedZoneDeletion] class.
type ICKSyncEngineFetchedZoneDeletion interface {
	objectivec.IObject
}

// An object that describes the deletion of a record zone.
//
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion
type CKSyncEngineFetchedZoneDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedZoneDeletionFrom constructs a [CKSyncEngineFetchedZoneDeletion] from an unsafe.Pointer.
//
// An object that describes the deletion of a record zone.
func CKSyncEngineFetchedZoneDeletionFrom(ptr unsafe.Pointer) CKSyncEngineFetchedZoneDeletion {
	return CKSyncEngineFetchedZoneDeletion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedZoneDeletionClass) Alloc() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchedZoneDeletionClass) New() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedZoneDeletion) Init() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedZoneDeletion) Autorelease() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedZoneDeletion creates a new CKSyncEngineFetchedZoneDeletion instance.
func NewCKSyncEngineFetchedZoneDeletion() CKSyncEngineFetchedZoneDeletion {
	return getCKSyncEngineFetchedZoneDeletionClass().New()
}




