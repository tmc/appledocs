// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectModel] class.
var managedObjectModelClass = _ManagedObjectModelClass{objc.GetClass("NSManagedObjectModel")}

type _ManagedObjectModelClass struct {
	class objc.Class
}

// A programmatic representation of the file describing your objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel

type ManagedObjectModel struct {
	objectivec.Object
}

// ManagedObjectModelFrom constructs a [ManagedObjectModel] from an unsafe.Pointer.
//
// A programmatic representation of the file describing your objects.
func ManagedObjectModelFrom(ptr unsafe.Pointer) ManagedObjectModel {
	return ManagedObjectModel{objectivec.Object{objc.ID(ptr)}}
}

// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/isConfiguration(withName:compatibleWithStoreMetadata:)
func (m_ ManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), configuration, metadata)
	return rv
}
// Associates the specified fetch request with the receiver using the given name. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setFetchRequestTemplate(_:forName:)
func (m_ ManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate unsafe.Pointer, name string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplate:forName:"), fetchRequestTemplate, name)
}


