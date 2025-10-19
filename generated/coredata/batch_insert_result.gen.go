// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchInsertResult] class.
var batchInsertResultClass = _BatchInsertResultClass{objc.GetClass("NSBatchInsertResult")}

type _BatchInsertResultClass struct {
	class objc.Class
}

// An interface definition for the [BatchInsertResult] class.
type IBatchInsertResult interface {
	IPersistentStoreResult
}

// The result that Core Data returns when executing a batch-insertion request. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return batchInsertResultClass.New()
}




