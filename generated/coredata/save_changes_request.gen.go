// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SaveChangesRequest] class.
var saveChangesRequestClass = _SaveChangesRequestClass{objc.GetClass("NSSaveChangesRequest")}

type _SaveChangesRequestClass struct {
	class objc.Class
}

// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest

type SaveChangesRequest struct {
	PersistentStoreRequest
}

// SaveChangesRequestFrom constructs a [SaveChangesRequest] from an unsafe.Pointer.
//
// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context.
func SaveChangesRequestFrom(ptr unsafe.Pointer) SaveChangesRequest {
	return SaveChangesRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}



