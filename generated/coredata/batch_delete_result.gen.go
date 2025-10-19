// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchDeleteResult] class.
var batchDeleteResultClass = _BatchDeleteResultClass{objc.GetClass("NSBatchDeleteResult")}

type _BatchDeleteResultClass struct {
	class objc.Class
}

// An object that describes the result of a batch delete request. [Full Topic]
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



