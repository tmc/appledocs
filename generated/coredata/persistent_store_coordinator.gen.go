// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreCoordinator] class.
var (
	persistentStoreCoordinatorClass     _PersistentStoreCoordinatorClass
	persistentStoreCoordinatorClassOnce sync.Once
)

func getPersistentStoreCoordinatorClass() _PersistentStoreCoordinatorClass {
	persistentStoreCoordinatorClassOnce.Do(func() {
		persistentStoreCoordinatorClass = _PersistentStoreCoordinatorClass{objc.GetClass("NSPersistentStoreCoordinator")}
	})
	return persistentStoreCoordinatorClass
}

type _PersistentStoreCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreCoordinator] class.
type IPersistentStoreCoordinator interface {
	objectivec.IObject
	AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL unsafe.Pointer, options objc.ID, error unsafe.Pointer) unsafe.Pointer
	CurrentPersistentHistoryTokenFromStores(stores objc.ID) unsafe.Pointer
	FinishDeferredLightweightMigration(error unsafe.Pointer) bool
	ImportStoreWithIdentifierFromExternalRecordsDirectoryToURLOptionsWithTypeError(storeIdentifier string, externalRecordsURL unsafe.Pointer, destinationURL unsafe.Pointer, options objc.ID, storeType string, error unsafe.Pointer) unsafe.Pointer
	ManagedObjectIDForURIRepresentation(url unsafe.Pointer) unsafe.Pointer
	ManagedObjectIDFromUTF8StringLength(utf8string unsafe.Pointer, len uint) unsafe.Pointer
	SetMetadataForPersistentStore(metadata unsafe.Pointer, store unsafe.Pointer)
	URLForPersistentStore(store unsafe.Pointer) unsafe.Pointer
}

// An object that enables an app’s contexts and the underlying persistent stores to work together.
//
// A managed object context uses a coordinator to facilitate the persistence of its entities in the coordinator’s registered stores. A context can’t function without a coordinator because it relies on the coordinator’s access to the managed object model. The coordinator presents its registered stores as an aggregate, allowing a context to operate on the union of those stores instead of on each individually. A coordinator performs its work on a private queue and executes that work serially. You can use multiple coordinators if the work requires separate queues. Use a coordinator to add or remove persistent stores, change the type or location on-disk of those stores, query the metadata of a specific store, defer a store’s migrations, determine whether two objects originate from the same store, and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator
type PersistentStoreCoordinator struct {
	objectivec.Object
}

// PersistentStoreCoordinatorFrom constructs a [PersistentStoreCoordinator] from an unsafe.Pointer.
//
// An object that enables an app’s contexts and the underlying persistent stores to work together.
func PersistentStoreCoordinatorFrom(ptr unsafe.Pointer) PersistentStoreCoordinator {
	return PersistentStoreCoordinator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreCoordinatorClass) Alloc() PersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreCoordinatorClass) New() PersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreCoordinator) Init() PersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreCoordinator) Autorelease() PersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreCoordinator creates a new PersistentStoreCoordinator instance.
func NewPersistentStoreCoordinator() PersistentStoreCoordinator {
	return getPersistentStoreCoordinatorClass().New()
}


// Creates a persistent store coordinator with the specified managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/init(managedObjectModel:)
func NewPersistentStoreCoordinatorWithManagedObjectModel(model unsafe.Pointer) PersistentStoreCoordinator {
	instance := getPersistentStoreCoordinatorClass().Alloc()
	rv := objc.Send[PersistentStoreCoordinator](instance.ID, objc.Sel("initWithManagedObjectModel:"), model)
	rv.Autorelease()
	return rv
}


// Returns the metadata of a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadataForPersistentStore(ofType:at:options:)
func (pc _PersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url unsafe.Pointer, options objc.ID, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("metadataForPersistentStoreOfType:URL:options:error:"), objc.String(storeType), url, options, error)
	return rv
}

// Adds a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(ofType:configurationName:at:options:)
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL unsafe.Pointer, options objc.ID, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addPersistentStoreWithType:configuration:URL:options:error:"), objc.String(storeType), objc.String(configuration), storeURL, options, error)
	return rv
}

// Returns a single persistent history token representing all of the specified stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/currentPersistentHistoryToken(fromStores:)
func (p_ PersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores(stores objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentPersistentHistoryTokenFromStores:"), stores)
	return rv
}

// Executes all remaining tasks of a deferred lightweight migration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigration()
func (p_ PersistentStoreCoordinator) FinishDeferredLightweightMigration(error unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finishDeferredLightweightMigration:"), error)
	return rv
}

// Creates and populates a store with the external records found at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/importStore(withIdentifier:fromExternalRecordsDirectoryAt:to:options:ofType:)
func (p_ PersistentStoreCoordinator) ImportStoreWithIdentifierFromExternalRecordsDirectoryToURLOptionsWithTypeError(storeIdentifier string, externalRecordsURL unsafe.Pointer, destinationURL unsafe.Pointer, options objc.ID, storeType string, error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("importStoreWithIdentifier:fromExternalRecordsDirectory:toURL:options:withType:error:"), objc.String(storeIdentifier), externalRecordsURL, destinationURL, options, objc.String(storeType), error)
	return rv
}

// Returns the object identifier for the specified URI representation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectID(forURIRepresentation:)
func (p_ PersistentStoreCoordinator) ManagedObjectIDForURIRepresentation(url unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("managedObjectIDForURIRepresentation:"), url)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectIDFromUTF8String:length:
func (p_ PersistentStoreCoordinator) ManagedObjectIDFromUTF8StringLength(utf8string unsafe.Pointer, len uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("managedObjectIDFromUTF8String:length:"), utf8string, len)
	return rv
}

// Updates the metadata for the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:for:)
func (p_ PersistentStoreCoordinator) SetMetadataForPersistentStore(metadata unsafe.Pointer, store unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetadata:forPersistentStore:"), metadata, store)
}

// Returns the location of the provided persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/url(for:)
func (p_ PersistentStoreCoordinator) URLForPersistentStore(store unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("URLForPersistentStore:"), store)
	return rv
}

// The coordinator’s managed object model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectModel
func (p_ PersistentStoreCoordinator) ManagedObjectModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("managedObjectModel"))
	return rv
}

// The coordinator’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/name
func (p_ PersistentStoreCoordinator) Name() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The coordinator’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/name
func (p_ PersistentStoreCoordinator) SetName(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}

