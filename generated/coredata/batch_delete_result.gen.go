// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchDeleteResult] class.
var (
	batchDeleteResultClass     _BatchDeleteResultClass
	batchDeleteResultClassOnce sync.Once
)

func getBatchDeleteResultClass() _BatchDeleteResultClass {
	batchDeleteResultClassOnce.Do(func() {
		batchDeleteResultClass = _BatchDeleteResultClass{objc.GetClass("NSBatchDeleteResult")}
	})
	return batchDeleteResultClass
}

type _BatchDeleteResultClass struct {
	class objc.Class
}

// An interface definition for the [BatchDeleteResult] class.
type IBatchDeleteResult interface {
	IPersistentStoreResult
}

// An object that describes the result of a batch delete request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult
type BatchDeleteResult struct {
	PersistentStoreResult
}

// BatchDeleteResultFrom constructs a [BatchDeleteResult] from an unsafe.Pointer.
//
// An object that describes the result of a batch delete request.
func BatchDeleteResultFrom(ptr unsafe.Pointer) BatchDeleteResult {
	return BatchDeleteResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BatchDeleteResultClass) Alloc() BatchDeleteResult {
	rv := objc.Send[BatchDeleteResult](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BatchDeleteResultClass) New() BatchDeleteResult {
	rv := objc.Send[BatchDeleteResult](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BatchDeleteResult) Init() BatchDeleteResult {
	rv := objc.Send[BatchDeleteResult](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BatchDeleteResult) Autorelease() BatchDeleteResult {
	rv := objc.Send[BatchDeleteResult](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBatchDeleteResult creates a new BatchDeleteResult instance.
func NewBatchDeleteResult() BatchDeleteResult {
	return getBatchDeleteResultClass().New()
}


// The value the request returns after it executes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult/result
func (b_ BatchDeleteResult) Result() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("result"))
	return rv
}

// The data type of the request’s result value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult/resultType
func (b_ BatchDeleteResult) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("resultType"))
	return rv
}



