// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineRecordZoneChangeBatch] class.
var (
	CKSyncEngineRecordZoneChangeBatchClass     _CKSyncEngineRecordZoneChangeBatchClass
	CKSyncEngineRecordZoneChangeBatchClassOnce sync.Once
)

func getCKSyncEngineRecordZoneChangeBatchClass() _CKSyncEngineRecordZoneChangeBatchClass {
	CKSyncEngineRecordZoneChangeBatchClassOnce.Do(func() {
		CKSyncEngineRecordZoneChangeBatchClass = _CKSyncEngineRecordZoneChangeBatchClass{objc.GetClass("CKSyncEngineRecordZoneChangeBatch")}
	})
	return CKSyncEngineRecordZoneChangeBatchClass
}

type _CKSyncEngineRecordZoneChangeBatchClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineRecordZoneChangeBatch] class.
type ICKSyncEngineRecordZoneChangeBatch interface {
	objectivec.IObject
}

// An object that contains the record changes for a single send operation.


// An object that contains the record changes for a single send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch
type CKSyncEngineRecordZoneChangeBatch struct {
	objectivec.Object
}

// CKSyncEngineRecordZoneChangeBatchFrom constructs a [CKSyncEngineRecordZoneChangeBatch] from an unsafe.Pointer.
//
// An object that contains the record changes for a single send operation.
func CKSyncEngineRecordZoneChangeBatchFrom(ptr unsafe.Pointer) CKSyncEngineRecordZoneChangeBatch {
	return CKSyncEngineRecordZoneChangeBatch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineRecordZoneChangeBatchClass) Alloc() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineRecordZoneChangeBatchClass) New() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineRecordZoneChangeBatch) Init() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineRecordZoneChangeBatch) Autorelease() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineRecordZoneChangeBatch creates a new CKSyncEngineRecordZoneChangeBatch instance.
func NewCKSyncEngineRecordZoneChangeBatch() CKSyncEngineRecordZoneChangeBatch {
	return getCKSyncEngineRecordZoneChangeBatchClass().New()
}




