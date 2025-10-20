// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchInsertResult] class.
var (
	BatchInsertResultClass     _BatchInsertResultClass
	BatchInsertResultClassOnce sync.Once
)

func getBatchInsertResultClass() _BatchInsertResultClass {
	BatchInsertResultClassOnce.Do(func() {
		BatchInsertResultClass = _BatchInsertResultClass{objc.GetClass("NSBatchInsertResult")}
	})
	return BatchInsertResultClass
}

type _BatchInsertResultClass struct {
	class objc.Class
}

// An interface definition for the [BatchInsertResult] class.
type IBatchInsertResult interface {
	IPersistentStoreResult
}

// The result that Core Data returns when executing a batch-insertion request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult
type BatchInsertResult struct {
	PersistentStoreResult
}

// BatchInsertResultFrom constructs a [BatchInsertResult] from an unsafe.Pointer.
//
// The result that Core Data returns when executing a batch-insertion request.
func BatchInsertResultFrom(ptr unsafe.Pointer) BatchInsertResult {
	return BatchInsertResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchInsertResultClass) Alloc() BatchInsertResult {
	rv := objc.Send[BatchInsertResult](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchInsertResultClass) New() BatchInsertResult {
	rv := objc.Send[BatchInsertResult](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchInsertResult) Init() BatchInsertResult {
	rv := objc.Send[BatchInsertResult](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchInsertResult) Autorelease() BatchInsertResult {
	rv := objc.Send[BatchInsertResult](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchInsertResult creates a new BatchInsertResult instance.
func NewBatchInsertResult() BatchInsertResult {
	return getBatchInsertResultClass().New()
}


// The result of a batch-insertion request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult/result
func (b_ BatchInsertResult) Result() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("result"))
	return rv
}

// The type of result that Core Data returns from this request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult/resultType
func (b_ BatchInsertResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resultType"))
	return rv
}



