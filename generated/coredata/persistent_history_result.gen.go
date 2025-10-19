// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentHistoryResult] class.
var persistentHistoryResultClass = _PersistentHistoryResultClass{objc.GetClass("NSPersistentHistoryResult")}

type _PersistentHistoryResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryResult] class.
type IPersistentHistoryResult interface {
	IPersistentStoreResult
}

// The result of a request to fetch persistent history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult

type PersistentHistoryResult struct {
	PersistentStoreResult
}

// PersistentHistoryResultFrom constructs a [PersistentHistoryResult] from an unsafe.Pointer.
//
// The result of a request to fetch persistent history.
func PersistentHistoryResultFrom(ptr unsafe.Pointer) PersistentHistoryResult {
	return PersistentHistoryResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PersistentHistoryResultClass) Alloc() PersistentHistoryResult {
	rv := objc.Send[PersistentHistoryResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersistentHistoryResultClass) New() PersistentHistoryResult {
	rv := objc.Send[PersistentHistoryResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentHistoryResult) Init() PersistentHistoryResult {
	rv := objc.Send[PersistentHistoryResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentHistoryResult) Autorelease() PersistentHistoryResult {
	rv := objc.Send[PersistentHistoryResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentHistoryResult creates a new PersistentHistoryResult instance.
func NewPersistentHistoryResult() PersistentHistoryResult {
	return persistentHistoryResultClass.New()
}




