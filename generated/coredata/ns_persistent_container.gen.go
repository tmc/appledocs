// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentContainer] class.
var (
	PersistentContainerClass     _PersistentContainerClass
	PersistentContainerClassOnce sync.Once
)

func getPersistentContainerClass() _PersistentContainerClass {
	PersistentContainerClassOnce.Do(func() {
		PersistentContainerClass = _PersistentContainerClass{objc.GetClass("NSPersistentContainer")}
	})
	return PersistentContainerClass
}

type _PersistentContainerClass struct {
	class objc.Class
}

// An interface definition for the [PersistentContainer] class.
type IPersistentContainer interface {
	objectivec.IObject
	LoadPersistentStoresWithCompletionHandler(block unsafe.Pointer)
	NewBackgroundContext() unsafe.Pointer
	PerformBackgroundTask(block unsafe.Pointer)
}

// A container that encapsulates the Core Data stack in your app.
//
// NSPersistentContainer simplifies the creation and management of the Core Data stack by handling the creation of the managed object model ( ), persistent store coordinator ( ), and the managed object context ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer
type PersistentContainer struct {
	objectivec.Object
}

// PersistentContainerFrom constructs a [PersistentContainer] from an unsafe.Pointer.
//
// A container that encapsulates the Core Data stack in your app.
func PersistentContainerFrom(ptr unsafe.Pointer) PersistentContainer {
	return PersistentContainer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentContainerClass) Alloc() PersistentContainer {
	rv := objc.Send[PersistentContainer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentContainerClass) New() PersistentContainer {
	rv := objc.Send[PersistentContainer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentContainer) Init() PersistentContainer {
	rv := objc.Send[PersistentContainer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentContainer) Autorelease() PersistentContainer {
	rv := objc.Send[PersistentContainer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentContainer creates a new PersistentContainer instance.
func NewPersistentContainer() PersistentContainer {
	return getPersistentContainerClass().New()
}


// Creates a container with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:)
func NewPersistentContainerWithName(name string) PersistentContainer {
	instance := getPersistentContainerClass().Alloc()
	rv := objc.Send[PersistentContainer](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	rv.Autorelease()
	return rv
}

// Create a container with the specified name and managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/init(name:managedObjectModel:)
func NewPersistentContainerWithNameManagedObjectModel(name string, model unsafe.Pointer) PersistentContainer {
	instance := getPersistentContainerClass().Alloc()
	rv := objc.Send[PersistentContainer](instance.ID, objc.Sel("initWithName:managedObjectModel:"), objc.String(name), model)
	rv.Autorelease()
	return rv
}


// Returns the location of the directory that contains the persistent stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/defaultDirectoryURL()
func (pc _PersistentContainerClass) DefaultDirectoryURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("defaultDirectoryURL"))
	return rv
}

// Initializes a new persistent container using the provided name for the container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentContainerWithName:
func (pc _PersistentContainerClass) PersistentContainerWithName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("persistentContainerWithName:"), objc.String(name))
	return rv
}

// Initializes a new persistent container using the provided name and managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentContainerWithName:managedObjectModel:
func (pc _PersistentContainerClass) PersistentContainerWithNameManagedObjectModel(name string, model unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("persistentContainerWithName:managedObjectModel:"), objc.String(name), model)
	return rv
}

// Loads the persistent stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/loadPersistentStores(completionHandler:)
func (p_ PersistentContainer) LoadPersistentStoresWithCompletionHandler(block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("loadPersistentStoresWithCompletionHandler:"), block)
}

// Returns a new managed object context that executes on a private queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/newBackgroundContext()
func (p_ PersistentContainer) NewBackgroundContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("newBackgroundContext"))
	return rv
}

// Executes a closure on a private queue using an ephemeral managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/performBackgroundTask(_:)-39sch
func (p_ PersistentContainer) PerformBackgroundTask(block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performBackgroundTask:"), block)
}

// The container’s managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/managedObjectModel
func (p_ PersistentContainer) ManagedObjectModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("managedObjectModel"))
	return rv
}

// The container’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/name
func (p_ PersistentContainer) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("name"))
	return rv
}

// The container’s persistent store coordinator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentStoreCoordinator
func (p_ PersistentContainer) PersistentStoreCoordinator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The descriptions of the container’s persistent stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentStoreDescriptions
func (p_ PersistentContainer) PersistentStoreDescriptions() []PersistentStoreDescription {
	rv := objc.Send[[]PersistentStoreDescription](p_.ID, objc.Sel("persistentStoreDescriptions"))
	return rv
}


// SetPersistentStoreDescriptions sets the value of the persistentStoreDescriptions property.
// The descriptions of the container’s persistent stores.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/persistentStoreDescriptions
func (p_ PersistentContainer) SetPersistentStoreDescriptions(value []PersistentStoreDescription) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setPersistentStoreDescriptions:"), nsArray)
}
// The main queue’s managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentContainer/viewContext
func (p_ PersistentContainer) ViewContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("viewContext"))
	return rv
}


