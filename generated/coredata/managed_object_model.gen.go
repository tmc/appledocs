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
	IsConfigurationCompatibleWithStoreMetadata(configuration string, metadata unsafe.Pointer) bool
	SetFetchRequestTemplateForName(fetchRequestTemplate unsafe.Pointer, name string)
}

// A programmatic representation of the file describing your objects.
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


