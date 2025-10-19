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



