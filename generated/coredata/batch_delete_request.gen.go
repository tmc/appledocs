// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BatchDeleteRequest] class.
var batchDeleteRequestClass = _BatchDeleteRequestClass{objc.GetClass("NSBatchDeleteRequest")}

type _BatchDeleteRequestClass struct {
	class objc.Class
}

// A request that deletes objects in the SQLite persistent store without loading them into memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSBatchDeleteRequest

type BatchDeleteRequest struct {
	PersistentStoreRequest
}

// BatchDeleteRequestFrom constructs a [BatchDeleteRequest] from an unsafe.Pointer.
//
// A request that deletes objects in the SQLite persistent store without loading them into memory.
func BatchDeleteRequestFrom(ptr unsafe.Pointer) BatchDeleteRequest {
	return BatchDeleteRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}



