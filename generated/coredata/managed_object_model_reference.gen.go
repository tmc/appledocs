// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectModelReference] class.
var managedObjectModelReferenceClass = _ManagedObjectModelReferenceClass{objc.GetClass("NSManagedObjectModelReference")}

type _ManagedObjectModelReferenceClass struct {
	class objc.Class
}

// An object that describes a specific version of an object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModelReference

type ManagedObjectModelReference struct {
	objectivec.Object
}

// ManagedObjectModelReferenceFrom constructs a [ManagedObjectModelReference] from an unsafe.Pointer.
//
// An object that describes a specific version of an object model.
func ManagedObjectModelReferenceFrom(ptr unsafe.Pointer) ManagedObjectModelReference {
	return ManagedObjectModelReference{objectivec.Object{objc.ID(ptr)}}
}



