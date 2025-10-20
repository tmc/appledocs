// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ManagedObjectModel] class.
var (
	managedObjectModelClass     _ManagedObjectModelClass
	managedObjectModelClassOnce sync.Once
)

func getManagedObjectModelClass() _ManagedObjectModelClass {
	managedObjectModelClassOnce.Do(func() {
		managedObjectModelClass = _ManagedObjectModelClass{objc.GetClass("NSManagedObjectModel")}
	})
	return managedObjectModelClass
}

type _ManagedObjectModelClass struct {
	class objc.Class
}

// An interface definition for the [ManagedObjectModel] class.
type IManagedObjectModel interface {
	objectivec.IObject
	FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables unsafe.Pointer) unsafe.Pointer
	FetchRequestTemplateForName(name string) unsafe.Pointer
	IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool
	SetFetchRequestTemplateForName(fetchRequestTemplate unsafe.Pointer, name string)
}

// A programmatic representation of the file describing your objects.
//
// The model contains one or more objects representing the entities in the schema. Each object has property description objects (instances of subclasses of ) that represent the properties (or fields) of the entity in the schema. The Core Data framework uses this description in several ways: Constraining UI creation in Interface Builder Validating attribute and relationship values at runtime Mapping between your managed objects and a database or file-based schema for object persistence A managed object model maintains a mapping between each of its entity objects and a corresponding managed object class for use with the persistent storage mechanisms in the Core Data framework. You can determine the entity for a particular managed object with the method. You typically create managed object models using the data modeling tool in Xcode, but it’s possible to build a model programmatically if needed.
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

// Alloc allocates a new instance without initialization.
func (mc _ManagedObjectModelClass) Alloc() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ManagedObjectModelClass) New() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ManagedObjectModel) Init() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ManagedObjectModel) Autorelease() ManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewManagedObjectModel creates a new ManagedObjectModel instance.
func NewManagedObjectModel() ManagedObjectModel {
	return getManagedObjectModelClass().New()
}


// Returns a copy of the fetch request template with the variables substituted by values from the substitutions dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestFromTemplate(withName:substitutionVariables:)
func (m_ ManagedObjectModel) FetchRequestFromTemplateWithNameSubstitutionVariables(name string, variables unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fetchRequestFromTemplateWithName:substitutionVariables:"), objc.String(name), variables)
	return rv
}

// Returns the fetch request with a specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/fetchRequestTemplate(forName:)
func (m_ ManagedObjectModel) FetchRequestTemplateForName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fetchRequestTemplateForName:"), objc.String(name))
	return rv
}

// Returns a Boolean value that indicates whether a given configuration in the model is compatible with given metadata from a persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/isConfiguration(withName:compatibleWithStoreMetadata:)
func (m_ ManagedObjectModel) IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isConfiguration:compatibleWithStoreMetadata:"), objc.String(configuration), metadata)
	return rv
}

// Associates the specified fetch request with the receiver using the given name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSManagedObjectModel/setFetchRequestTemplate(_:forName:)
func (m_ ManagedObjectModel) SetFetchRequestTemplateForName(fetchRequestTemplate unsafe.Pointer, name string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFetchRequestTemplate:forName:"), fetchRequestTemplate, objc.String(name))
}



