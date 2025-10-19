// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AtomicStore] class.
var atomicStoreClass = _AtomicStoreClass{objc.GetClass("NSAtomicStore")}

type _AtomicStoreClass struct {
	class objc.Class
}

// An abstract superclass that you subclass to create a Core Data atomic store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStore

type AtomicStore struct {
	PersistentStore
}

// AtomicStoreFrom constructs a [AtomicStore] from an unsafe.Pointer.
//
// An abstract superclass that you subclass to create a Core Data atomic store.
func AtomicStoreFrom(ptr unsafe.Pointer) AtomicStore {
	return AtomicStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

// Saves the cache nodes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStore/save()
func (a_ AtomicStore) Save(error unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("save:"), error)
	return rv
}


