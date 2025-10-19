// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentStoreAsynchronousResult] class.
var persistentStoreAsynchronousResultClass = _PersistentStoreAsynchronousResultClass{objc.GetClass("NSPersistentStoreAsynchronousResult")}

type _PersistentStoreAsynchronousResultClass struct {
	class objc.Class
}

// A concrete class used to represent the results of an asynchronous request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult

type PersistentStoreAsynchronousResult struct {
	PersistentStoreResult
}

// PersistentStoreAsynchronousResultFrom constructs a [PersistentStoreAsynchronousResult] from an unsafe.Pointer.
//
// A concrete class used to represent the results of an asynchronous request.
func PersistentStoreAsynchronousResultFrom(ptr unsafe.Pointer) PersistentStoreAsynchronousResult {
	return PersistentStoreAsynchronousResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

// Cancels the asynchronous fetch request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/cancel()
func (p_ PersistentStoreAsynchronousResult) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}


