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
	PersistentStoreCoordinatorClass     _PersistentStoreCoordinatorClass
	PersistentStoreCoordinatorClassOnce sync.Once
)

func getPersistentStoreCoordinatorClass() _PersistentStoreCoordinatorClass {
	PersistentStoreCoordinatorClassOnce.Do(func() {
		PersistentStoreCoordinatorClass = _PersistentStoreCoordinatorClass{objc.GetClass("NSPersistentStoreCoordinator")}
	})
	return PersistentStoreCoordinatorClass
}

type _PersistentStoreCoordinatorClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreCoordinator] class.
type IPersistentStoreCoordinator interface {
	objectivec.IObject
	AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) unsafe.Pointer
	AddPersistentStoreWithDescriptionCompletionHandler(storeDescription unsafe.Pointer, block unsafe.Pointer)
	CurrentPersistentHistoryTokenFromStores(stores objc.ID) unsafe.Pointer
	DestroyPersistentStoreAtURLWithTypeOptionsError(url unsafe.Pointer, storeType string, options objc.ID, error_ unsafe.Pointer) bool
	ExecuteRequestWithContextError(request unsafe.Pointer, context unsafe.Pointer, error_ unsafe.Pointer) objc.ID
	FinishDeferredLightweightMigration(error_ unsafe.Pointer) bool
	FinishDeferredLightweightMigrationTask(error_ unsafe.Pointer) bool
	ImportStoreWithIdentifierFromExternalRecordsDirectoryToURLOptionsWithTypeError(storeIdentifier string, externalRecordsURL unsafe.Pointer, destinationURL unsafe.Pointer, options objc.ID, storeType string, error_ unsafe.Pointer) unsafe.Pointer
	ManagedObjectIDForURIRepresentation(url unsafe.Pointer) unsafe.Pointer
	ManagedObjectIDFromUTF8StringLength(utf8string unsafe.Pointer, len uint) unsafe.Pointer
	MetadataForPersistentStore(store unsafe.Pointer) unsafe.Pointer
	MigratePersistentStoreToURLOptionsWithTypeError(store unsafe.Pointer, URL unsafe.Pointer, options objc.ID, storeType string, error_ unsafe.Pointer) unsafe.Pointer
	PerformBlock(block unsafe.Pointer)
	PersistentStoreForURL(URL unsafe.Pointer) unsafe.Pointer
	ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError(destinationURL unsafe.Pointer, destinationOptions objc.ID, sourceURL unsafe.Pointer, sourceOptions objc.ID, storeType string, error_ unsafe.Pointer) bool
	SetMetadataForPersistentStore(metadata unsafe.Pointer, store unsafe.Pointer)
	SetStoresFastSyncDetailsAtURLForPersistentStore(url unsafe.Pointer, store unsafe.Pointer)
	SetURLForPersistentStore(url unsafe.Pointer, store unsafe.Pointer) bool
	SyncWithClientInBackgroundHandlerError(client unsafe.Pointer, flag bool, syncHandler objc.ID, rError unsafe.Pointer) bool
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


// Returns a dictionary containing the metadata stored in the persistent store at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadataForPersistentStore(ofType:at:)
func (pc _PersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLError(storeType string, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("metadataForPersistentStoreOfType:URL:error:"), objc.String(storeType), url, error_)
	return rv
}

// Returns the metadata of a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadataForPersistentStore(ofType:at:options:)
func (pc _PersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("metadataForPersistentStoreOfType:URL:options:error:"), objc.String(storeType), url, options, error_)
	return rv
}

// Registers a persistent store subclass using the specified store type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registerStoreClass(_:forStoreType:)
func (pc _PersistentStoreCoordinatorClass) RegisterStoreClassForStoreType(storeClass objc.Class, storeType string) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("registerStoreClass:forStoreType:"), storeClass, objc.String(storeType))
}

// Sets the metadata for a given store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:forPersistentStoreOfType:at:)
func (pc _PersistentStoreCoordinatorClass) SetMetadataForPersistentStoreOfTypeURLError(metadata unsafe.Pointer, storeType string, url unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("setMetadata:forPersistentStoreOfType:URL:error:"), metadata, objc.String(storeType), url, error_)
	return rv
}

// Updates the metadata of a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:forPersistentStoreOfType:at:options:)
func (pc _PersistentStoreCoordinatorClass) SetMetadataForPersistentStoreOfTypeURLOptionsError(metadata unsafe.Pointer, storeType string, url unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("setMetadata:forPersistentStoreOfType:URL:options:error:"), metadata, objc.String(storeType), url, options, error_)
	return rv
}

// The coordinator’s registered store types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registeredStoreTypes
func (pc _PersistentStoreCoordinatorClass) RegisteredStoreTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("registeredStoreTypes"))
	return rv
}
// Adds a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(ofType:configurationName:at:options:)
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL unsafe.Pointer, options objc.ID, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("addPersistentStoreWithType:configuration:URL:options:error:"), objc.String(storeType), objc.String(configuration), storeURL, options, error_)
	return rv
}

// Adds a persistent store using the provided description.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(with:completionHandler:)
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithDescriptionCompletionHandler(storeDescription unsafe.Pointer, block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addPersistentStoreWithDescription:completionHandler:"), storeDescription, block)
}

// Returns a single persistent history token representing all of the specified stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/currentPersistentHistoryToken(fromStores:)
func (p_ PersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores(stores objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentPersistentHistoryTokenFromStores:"), stores)
	return rv
}

// Deletes a specific type of persistent store at the provided location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/destroyPersistentStore(at:ofType:options:)
func (p_ PersistentStoreCoordinator) DestroyPersistentStoreAtURLWithTypeOptionsError(url unsafe.Pointer, storeType string, options objc.ID, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("destroyPersistentStoreAtURL:withType:options:error:"), url, objc.String(storeType), options, error_)
	return rv
}

// Executes the specified request on each of the coordinator’s persistent stores.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/execute(_:with:)
func (p_ PersistentStoreCoordinator) ExecuteRequestWithContextError(request unsafe.Pointer, context unsafe.Pointer, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("executeRequest:withContext:error:"), request, context, error_)
	return rv
}

// Executes all remaining tasks of a deferred lightweight migration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigration()
func (p_ PersistentStoreCoordinator) FinishDeferredLightweightMigration(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finishDeferredLightweightMigration:"), error_)
	return rv
}

// Executes a single pending task of a deferred lightweight migration.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigrationTask()
func (p_ PersistentStoreCoordinator) FinishDeferredLightweightMigrationTask(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finishDeferredLightweightMigrationTask:"), error_)
	return rv
}

// Creates and populates a store with the external records found at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/importStore(withIdentifier:fromExternalRecordsDirectoryAt:to:options:ofType:)
func (p_ PersistentStoreCoordinator) ImportStoreWithIdentifierFromExternalRecordsDirectoryToURLOptionsWithTypeError(storeIdentifier string, externalRecordsURL unsafe.Pointer, destinationURL unsafe.Pointer, options objc.ID, storeType string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("importStoreWithIdentifier:fromExternalRecordsDirectory:toURL:options:withType:error:"), objc.String(storeIdentifier), externalRecordsURL, destinationURL, options, objc.String(storeType), error_)
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

// Returns the metadata of the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadata(for:)
func (p_ PersistentStoreCoordinator) MetadataForPersistentStore(store unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("metadataForPersistentStore:"), store)
	return rv
}

// Changes the location and, if necessary, the store type of the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/migratePersistentStore(_:to:options:withType:)
func (p_ PersistentStoreCoordinator) MigratePersistentStoreToURLOptionsWithTypeError(store unsafe.Pointer, URL unsafe.Pointer, options objc.ID, storeType string, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("migratePersistentStore:toURL:options:withType:error:"), store, URL, options, objc.String(storeType), error_)
	return rv
}

// Executes the provided closure asynchronously on the coordinator’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/perform(_:)-7jqb
func (p_ PersistentStoreCoordinator) PerformBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performBlock:"), block)
}

// Returns the persistent store for the specified file URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/persistentStore(for:)
func (p_ PersistentStoreCoordinator) PersistentStoreForURL(URL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("persistentStoreForURL:"), URL)
	return rv
}

// Replaces one persistent store with another.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/replacePersistentStore(at:destinationOptions:withPersistentStoreFrom:sourceOptions:ofType:)
func (p_ PersistentStoreCoordinator) ReplacePersistentStoreAtURLDestinationOptionsWithPersistentStoreFromURLSourceOptionsStoreTypeError(destinationURL unsafe.Pointer, destinationOptions objc.ID, sourceURL unsafe.Pointer, sourceOptions objc.ID, storeType string, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("replacePersistentStoreAtURL:destinationOptions:withPersistentStoreFromURL:sourceOptions:storeType:error:"), destinationURL, destinationOptions, sourceURL, sourceOptions, objc.String(storeType), error_)
	return rv
}

// Updates the metadata for the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:for:)
func (p_ PersistentStoreCoordinator) SetMetadataForPersistentStore(metadata unsafe.Pointer, store unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetadata:forPersistentStore:"), metadata, store)
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setStoresFastSyncDetailsAtURL:forPersistentStore:
func (p_ PersistentStoreCoordinator) SetStoresFastSyncDetailsAtURLForPersistentStore(url unsafe.Pointer, store unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStoresFastSyncDetailsAtURL:forPersistentStore:"), url, store)
}

// Changes the location of the specified persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setURL(_:for:)
func (p_ PersistentStoreCoordinator) SetURLForPersistentStore(url unsafe.Pointer, store unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setURL:forPersistentStore:"), url, store)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/syncWithClient:inBackground:handler:error:
func (p_ PersistentStoreCoordinator) SyncWithClientInBackgroundHandlerError(client unsafe.Pointer, flag bool, syncHandler objc.ID, rError unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("syncWithClient:inBackground:handler:error:"), client, flag, syncHandler, rError)
	return rv
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
// The coordinator’s registered store types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registeredStoreTypes
func (p_ PersistentStoreCoordinator) RegisteredStoreTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("registeredStoreTypes"))
	return rv
}


