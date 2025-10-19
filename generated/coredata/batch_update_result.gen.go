// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchUpdateResult] class.
var batchUpdateResultClass = _BatchUpdateResultClass{objc.GetClass("NSBatchUpdateResult")}

type _BatchUpdateResultClass struct {
	class objc.Class
}

// An interface definition for the [BatchUpdateResult] class.
type IBatchUpdateResult interface {
	IPersistentStoreResult
}

// The result returned when executing a batch update request. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return batchUpdateResultClass.New()
}




