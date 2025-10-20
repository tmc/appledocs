// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchUpdateResult] class.
var (
	batchUpdateResultClass     _BatchUpdateResultClass
	batchUpdateResultClassOnce sync.Once
)

func getBatchUpdateResultClass() _BatchUpdateResultClass {
	batchUpdateResultClassOnce.Do(func() {
		batchUpdateResultClass = _BatchUpdateResultClass{objc.GetClass("NSBatchUpdateResult")}
	})
	return batchUpdateResultClass
}

type _BatchUpdateResultClass struct {
	class objc.Class
}

// An interface definition for the [BatchUpdateResult] class.
type IBatchUpdateResult interface {
	IPersistentStoreResult
}

// The result returned when executing a batch update request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult
type BatchUpdateResult struct {
	PersistentStoreResult
}

// BatchUpdateResultFrom constructs a [BatchUpdateResult] from an unsafe.Pointer.
//
// The result returned when executing a batch update request.
func BatchUpdateResultFrom(ptr unsafe.Pointer) BatchUpdateResult {
	return BatchUpdateResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchUpdateResultClass) Alloc() BatchUpdateResult {
	rv := objc.Send[BatchUpdateResult](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchUpdateResultClass) New() BatchUpdateResult {
	rv := objc.Send[BatchUpdateResult](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchUpdateResult) Init() BatchUpdateResult {
	rv := objc.Send[BatchUpdateResult](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchUpdateResult) Autorelease() BatchUpdateResult {
	rv := objc.Send[BatchUpdateResult](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchUpdateResult creates a new BatchUpdateResult instance.
func NewBatchUpdateResult() BatchUpdateResult {
	return getBatchUpdateResultClass().New()
}


// The type of result that Core Data returns from the request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult/resultType
func (b_ BatchUpdateResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resultType"))
	return rv
}



