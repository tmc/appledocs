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



