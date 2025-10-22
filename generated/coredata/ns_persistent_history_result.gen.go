// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentHistoryResult] class.
var (
	PersistentHistoryResultClass     _PersistentHistoryResultClass
	PersistentHistoryResultClassOnce sync.Once
)

func getPersistentHistoryResultClass() _PersistentHistoryResultClass {
	PersistentHistoryResultClassOnce.Do(func() {
		PersistentHistoryResultClass = _PersistentHistoryResultClass{objc.GetClass("NSPersistentHistoryResult")}
	})
	return PersistentHistoryResultClass
}

type _PersistentHistoryResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryResult] class.
type IPersistentHistoryResult interface {
	IPersistentStoreResult
	Result() objc.ID
	ResultType() PersistentHistoryResultType
}

// The result of a request to fetch persistent history.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPersistentHistoryResultClass().New()
}


// The result of the history request determined by the persistent history result type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult/result
func (p_ PersistentHistoryResult) Result() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("result"))
	return rv
}

// The type of result that the persistent history change request returns.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult/resultType
func (p_ PersistentHistoryResult) ResultType() PersistentHistoryResultType {
	rv := objc.Send[PersistentHistoryResultType](p_.ID, objc.Sel("resultType"))
	return rv
}



