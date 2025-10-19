// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreResult] class.
var persistentStoreResultClass = _PersistentStoreResultClass{objc.GetClass("NSPersistentStoreResult")}

type _PersistentStoreResultClass struct {
	class objc.Class
}

// The abstract base class for results returned from a persistent store coordinator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreResult

type PersistentStoreResult struct {
	objectivec.Object
}

// PersistentStoreResultFrom constructs a [PersistentStoreResult] from an unsafe.Pointer.
//
// The abstract base class for results returned from a persistent store coordinator.
func PersistentStoreResultFrom(ptr unsafe.Pointer) PersistentStoreResult {
	return PersistentStoreResult{objectivec.Object{objc.ID(ptr)}}
}



