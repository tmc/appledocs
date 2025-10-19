// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectID] class.
var managedObjectIDClass = _ManagedObjectIDClass{objc.GetClass("NSManagedObjectID")}

type _ManagedObjectIDClass struct {
	class objc.Class
}

// A compact, universal identifier for a managed object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID

type ManagedObjectID struct {
	objectivec.Object
}

// ManagedObjectIDFrom constructs a [ManagedObjectID] from an unsafe.Pointer.
//
// A compact, universal identifier for a managed object.
func ManagedObjectIDFrom(ptr unsafe.Pointer) ManagedObjectID {
	return ManagedObjectID{objectivec.Object{objc.ID(ptr)}}
}

// Returns a URI that provides an archiveable reference to the object for the object ID. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectID/uriRepresentation()
func (m_ ManagedObjectID) URIRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("URIRepresentation"))
	return rv
}


