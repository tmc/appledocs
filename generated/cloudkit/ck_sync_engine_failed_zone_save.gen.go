// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineFailedZoneSave] class.
var (
	CKSyncEngineFailedZoneSaveClass     _CKSyncEngineFailedZoneSaveClass
	CKSyncEngineFailedZoneSaveClassOnce sync.Once
)

func getCKSyncEngineFailedZoneSaveClass() _CKSyncEngineFailedZoneSaveClass {
	CKSyncEngineFailedZoneSaveClassOnce.Do(func() {
		CKSyncEngineFailedZoneSaveClass = _CKSyncEngineFailedZoneSaveClass{objc.GetClass("CKSyncEngineFailedZoneSave")}
	})
	return CKSyncEngineFailedZoneSaveClass
}

type _CKSyncEngineFailedZoneSaveClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFailedZoneSave] class.
type ICKSyncEngineFailedZoneSave interface {
	objectivec.IObject
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	RecordZone() ICKRecordZone
	// methods:
}

// An object that describes an unsuccessful attempt to modify a single record zone.


// An object that describes an unsuccessful attempt to modify a single record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave
type CKSyncEngineFailedZoneSave struct {
	objectivec.Object
}

// CKSyncEngineFailedZoneSaveFrom constructs a [CKSyncEngineFailedZoneSave] from an unsafe.Pointer.
//
// An object that describes an unsuccessful attempt to modify a single record zone.
func CKSyncEngineFailedZoneSaveFrom(ptr unsafe.Pointer) CKSyncEngineFailedZoneSave {
	return CKSyncEngineFailedZoneSave{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFailedZoneSaveClass) Alloc() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFailedZoneSaveClass) New() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFailedZoneSave) Init() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFailedZoneSave) Autorelease() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedZoneSave creates a new CKSyncEngineFailedZoneSave instance.
func NewCKSyncEngineFailedZoneSave() CKSyncEngineFailedZoneSave {
	return getCKSyncEngineFailedZoneSaveClass().New()
}



// A error that describes the reason for the unsuccessful attempt to modify the associated record zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/error
func (c_ CKSyncEngineFailedZoneSave) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](c_.ID, objc.Sel("error"))
	return rv
}


// The record zone that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/recordZone
func (c_ CKSyncEngineFailedZoneSave) RecordZone() ICKRecordZone {
	rv := objc.Send[CKRecordZone](c_.ID, objc.Sel("recordZone"))
	return rv
}



