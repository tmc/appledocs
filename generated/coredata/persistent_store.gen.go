// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStore] class.
var (
	persistentStoreClass     _PersistentStoreClass
	persistentStoreClassOnce sync.Once
)

func getPersistentStoreClass() _PersistentStoreClass {
	persistentStoreClassOnce.Do(func() {
		persistentStoreClass = _PersistentStoreClass{objc.GetClass("NSPersistentStore")}
	})
	return persistentStoreClass
}

type _PersistentStoreClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStore] class.
type IPersistentStore interface {
	objectivec.IObject
}

// The abstract base class for all Core Data persistent stores.
//
// Core Data provides four store types—SQLite, Binary, XML, and In-Memory (the XML store is not available on iOS); these are described in Persistent Store Features. Core Data also provides subclasses of that you can use to define your own store types: and . The Binary and XML stores are examples of atomic stores that inherit functionality from .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPersistentStoreClass().New()
}


// Returns a store initialized with the given arguments.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/init(persistentStoreCoordinator:configurationName:at:options:)
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root unsafe.Pointer, name string, url unsafe.Pointer, options unsafe.Pointer) PersistentStore {
	instance := getPersistentStoreClass().Alloc()
	rv := objc.Send[PersistentStore](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	rv.Autorelease()
	return rv
}


// The metadata for the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadata
func (p_ PersistentStore) Metadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("metadata"))
	return rv
}

// SetMetadata sets the value of the metadata property.
// The metadata for the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadata
func (p_ PersistentStore) SetMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetadata:"), value)
}
// The type string of the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/type
func (p_ PersistentStore) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}

