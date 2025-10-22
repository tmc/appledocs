// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKSyncEngineFetchedRecordDeletion] class.
var (
	CKSyncEngineFetchedRecordDeletionClass     _CKSyncEngineFetchedRecordDeletionClass
	CKSyncEngineFetchedRecordDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedRecordDeletionClass() _CKSyncEngineFetchedRecordDeletionClass {
	CKSyncEngineFetchedRecordDeletionClassOnce.Do(func() {
		CKSyncEngineFetchedRecordDeletionClass = _CKSyncEngineFetchedRecordDeletionClass{objc.GetClass("CKSyncEngineFetchedRecordDeletion")}
	})
	return CKSyncEngineFetchedRecordDeletionClass
}

type _CKSyncEngineFetchedRecordDeletionClass struct {
	class objc.Class
}

// An interface definition for the [CKSyncEngineFetchedRecordDeletion] class.
type ICKSyncEngineFetchedRecordDeletion interface {
	objectivec.IObject
	RecordID() CKRecordID
	RecordType() unsafe.Pointer
}

// An object that describes the deletion of an individual record.


// An object that describes the deletion of an individual record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion

type CKSyncEngineFetchedRecordDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedRecordDeletionFrom constructs a [CKSyncEngineFetchedRecordDeletion] from an unsafe.Pointer.
//
// An object that describes the deletion of an individual record.
func CKSyncEngineFetchedRecordDeletionFrom(ptr unsafe.Pointer) CKSyncEngineFetchedRecordDeletion {
	return CKSyncEngineFetchedRecordDeletion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineFetchedRecordDeletionClass) Alloc() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKSyncEngineFetchedRecordDeletionClass) New() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineFetchedRecordDeletion) Init() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineFetchedRecordDeletion) Autorelease() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedRecordDeletion creates a new CKSyncEngineFetchedRecordDeletion instance.
func NewCKSyncEngineFetchedRecordDeletion() CKSyncEngineFetchedRecordDeletion {
	return getCKSyncEngineFetchedRecordDeletionClass().New()
}



// The deleted record’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordID

func (c_ CKSyncEngineFetchedRecordDeletion) RecordID() CKRecordID {
	rv := objc.Send[CKRecordID](c_.ID, objc.Sel("recordID"))
	return rv
}


// The record type of the deleted record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordType

func (c_ CKSyncEngineFetchedRecordDeletion) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordType"))
	return rv
}



