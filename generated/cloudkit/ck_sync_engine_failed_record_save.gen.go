// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineFailedRecordSave] class.
var (
	CKSyncEngineFailedRecordSaveClass     _CKSyncEngineFailedRecordSaveClass
	CKSyncEngineFailedRecordSaveClassOnce sync.Once
)

func getCKSyncEngineFailedRecordSaveClass() _CKSyncEngineFailedRecordSaveClass {
	CKSyncEngineFailedRecordSaveClassOnce.Do(func() {
		CKSyncEngineFailedRecordSaveClass = _CKSyncEngineFailedRecordSaveClass{objc.GetClass("CKSyncEngineFailedRecordSave")}
	})
	return CKSyncEngineFailedRecordSaveClass
}

type _CKSyncEngineFailedRecordSaveClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFailedRecordSave] class.
type ICKSyncEngineFailedRecordSave interface {
	objectivec.IObject
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	Record() ICKRecord
	// methods:
}

// A type that describes an unsuccessful attempt to modify an individual record.


// A type that describes an unsuccessful attempt to modify an individual record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave
type CKSyncEngineFailedRecordSave struct {
	objectivec.Object
}

// CKSyncEngineFailedRecordSaveFrom constructs a [CKSyncEngineFailedRecordSave] from an unsafe.Pointer.
//
// A type that describes an unsuccessful attempt to modify an individual record.
func CKSyncEngineFailedRecordSaveFrom(ptr unsafe.Pointer) CKSyncEngineFailedRecordSave {
	return CKSyncEngineFailedRecordSave{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFailedRecordSaveClass) Alloc() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFailedRecordSaveClass) New() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFailedRecordSave) Init() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFailedRecordSave) Autorelease() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedRecordSave creates a new CKSyncEngineFailedRecordSave instance.
func NewCKSyncEngineFailedRecordSave() CKSyncEngineFailedRecordSave {
	return getCKSyncEngineFailedRecordSaveClass().New()
}



// A error that describes the reason for the unsuccessful attempt to modify the associated record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/error
func (c_ CKSyncEngineFailedRecordSave) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](c_.ID, objc.Sel("error"))
	return rv
}


// The record that CloudKit is unable to modify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/record
func (c_ CKSyncEngineFailedRecordSave) Record() ICKRecord {
	rv := objc.Send[CKRecord](c_.ID, objc.Sel("record"))
	return rv
}



