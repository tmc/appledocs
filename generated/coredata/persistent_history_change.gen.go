// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryChange] class.
var persistentHistoryChangeClass = _PersistentHistoryChangeClass{objc.GetClass("NSPersistentHistoryChange")}

type _PersistentHistoryChangeClass struct {
	class objc.Class
}

// A change representing the insertion, update, or deletion of a managed object in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange

type PersistentHistoryChange struct {
	objectivec.Object
}

// PersistentHistoryChangeFrom constructs a [PersistentHistoryChange] from an unsafe.Pointer.
//
// A change representing the insertion, update, or deletion of a managed object in the persistent store.
func PersistentHistoryChangeFrom(ptr unsafe.Pointer) PersistentHistoryChange {
	return PersistentHistoryChange{objectivec.Object{objc.ID(ptr)}}
}



