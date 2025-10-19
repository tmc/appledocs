// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentHistoryResult] class.
var persistentHistoryResultClass = _PersistentHistoryResultClass{objc.GetClass("NSPersistentHistoryResult")}

type _PersistentHistoryResultClass struct {
	class objc.Class
}

// The result of a request to fetch persistent history. [Full Topic]
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



