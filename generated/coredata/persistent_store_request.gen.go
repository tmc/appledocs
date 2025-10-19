// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreRequest] class.
var persistentStoreRequestClass = _PersistentStoreRequestClass{objc.GetClass("NSPersistentStoreRequest")}

type _PersistentStoreRequestClass struct {
	class objc.Class
}

// Criteria used to retrieve data from or save data to a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest

type PersistentStoreRequest struct {
	objectivec.Object
}

// PersistentStoreRequestFrom constructs a [PersistentStoreRequest] from an unsafe.Pointer.
//
// Criteria used to retrieve data from or save data to a persistent store.
func PersistentStoreRequestFrom(ptr unsafe.Pointer) PersistentStoreRequest {
	return PersistentStoreRequest{objectivec.Object{objc.ID(ptr)}}
}



