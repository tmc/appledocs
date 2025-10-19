// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStore] class.
var persistentStoreClass = _PersistentStoreClass{objc.GetClass("NSPersistentStore")}

type _PersistentStoreClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStore] class.
type IPersistentStore interface {
	objectivec.IObject
}

// The abstract base class for all Core Data persistent stores. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore

type PersistentStore struct {
	objectivec.Object
}

// PersistentStoreFrom constructs a [PersistentStore] from an unsafe.Pointer.
//
// The abstract base class for all Core Data persistent stores.
func PersistentStoreFrom(ptr unsafe.Pointer) PersistentStore {
	return PersistentStore{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreClass) Alloc() PersistentStore {
	rv := objc.Send[PersistentStore](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersistentStoreClass) New() PersistentStore {
	rv := objc.Send[PersistentStore](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStore) Init() PersistentStore {
	rv := objc.Send[PersistentStore](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStore) Autorelease() PersistentStore {
	rv := objc.Send[PersistentStore](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStore creates a new PersistentStore instance.
func NewPersistentStore() PersistentStore {
	return persistentStoreClass.New()
}


// Returns a store initialized with the given arguments. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/init(persistentStoreCoordinator:configurationName:at:options:)
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root unsafe.Pointer, name string, url unsafe.Pointer, options unsafe.Pointer) PersistentStore {
	instance := persistentStoreClass.Alloc()
	rv := objc.Send[PersistentStore](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	rv.Autorelease()
	return rv
}



