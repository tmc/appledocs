// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IncrementalStore] class.
var incrementalStoreClass = _IncrementalStoreClass{objc.GetClass("NSIncrementalStore")}

type _IncrementalStoreClass struct {
	class objc.Class
}

// An abstract superclass defining the API through which Core Data communicates with a store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore

type IncrementalStore struct {
	PersistentStore
}

// IncrementalStoreFrom constructs a [IncrementalStore] from an unsafe.Pointer.
//
// An abstract superclass defining the API through which Core Data communicates with a store.
func IncrementalStoreFrom(ptr unsafe.Pointer) IncrementalStore {
	return IncrementalStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

// Returns a value as appropriate for the given request, or nil if the request cannot be completed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/execute(_:with:)
func (i_ IncrementalStore) ExecuteRequestWithContextError(request unsafe.Pointer, context unsafe.Pointer, error unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("executeRequest:withContext:error:"), request, context, error)
	return rv
}
// Loads the metadata for the store. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/loadMetadata()
func (i_ IncrementalStore) LoadMetadata(error unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("loadMetadata:"), error)
	return rv
}


