// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	ManagedObjectModel() IManagedObjectModel
	Name() string
	SetName(value string)
	NSCoreDataCoreSpotlightExporter() string
	NSPersistentHistoryTrackingKey() string
	PersistentStores() IPersistentStore
	SetPersistentStores(value IPersistentStore)
	NSPersistentStoreDeferredLightweightMigrationOptionKey() string
	NSStoreTypeKey() string
	NSStoreUUIDKey() string
	AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.URL, options objectivec.IObject, error_ unsafe.Pointer) IPersistentStore
	CurrentPersistentHistoryTokenFromStores(stores objectivec.IObject) IPersistentHistoryToken
	FinishDeferredLightweightMigration(error_ unsafe.Pointer) bool
	ManagedObjectIDForURIRepresentation(url foundation.URL) IManagedObjectID
	ManagedObjectIDFromUTF8StringLength(utf8string unsafe.Pointer, len_ uint) IManagedObjectID
	SetMetadataForPersistentStore(metadata foundation.IDictionary, store IPersistentStore)
	URLForPersistentStore(store IPersistentStore) foundation.URL
}

// An object that enables an app’s contexts and the underlying persistent stores to work together.
//
// A managed object context uses a coordinator to facilitate the persistence of its entities in the coordinator’s registered stores. A context can’t function without a coordinator because it relies on the coordinator’s access to the managed object model. The coordinator presents its registered stores as an aggregate, allowing a context to operate on the union of those stores instead of on each individually. A coordinator performs its work on a private queue and executes that work serially. You can use multiple coordinators if the work requires separate queues. Use a coordinator to add or remove persistent stores, change the type or location on-disk of those stores, query the metadata of a specific store, defer a store’s migrations, determine whether two objects originate from the same store, and so on.


// An object that enables an app’s contexts and the underlying persistent stores to work together.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/init(managedObjectModel:)
func NewPersistentStoreCoordinatorWithManagedObjectModel(model IManagedObjectModel) PersistentStoreCoordinator {
	instance := getPersistentStoreCoordinatorClass().Alloc()
	rv := objc.Send[PersistentStoreCoordinator](instance.ID, objc.Sel("initWithManagedObjectModel:"), model)
	rv.Autorelease()
	return rv
}



// Returns the metadata of a specific type of persistent store at the provided location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/metadataForPersistentStore(ofType:at:options:)
func (pc _PersistentStoreCoordinatorClass) MetadataForPersistentStoreOfTypeURLOptionsError(storeType string, url foundation.URL, options objectivec.IObject, error_ unsafe.Pointer) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(pc.class), objc.Sel("metadataForPersistentStoreOfType:URL:options:error:"), objc.String(storeType), url, options, error_)
	return rv
}


// The coordinator’s registered store types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registeredStoreTypes
func (pc _PersistentStoreCoordinatorClass) RegisteredStoreTypes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(pc.class), objc.Sel("registeredStoreTypes"))
	return rv
}

// Adds a specific type of persistent store at the provided location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/addPersistentStore(ofType:configurationName:at:options:)
func (p_ PersistentStoreCoordinator) AddPersistentStoreWithTypeConfigurationURLOptionsError(storeType string, configuration string, storeURL foundation.URL, options objectivec.IObject, error_ unsafe.Pointer) IPersistentStore {
	rv := objc.Send[PersistentStore](p_.ID, objc.Sel("addPersistentStoreWithType:configuration:URL:options:error:"), objc.String(storeType), objc.String(configuration), storeURL, options, error_)
	return rv
}


// Returns a single persistent history token representing all of the specified stores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/currentPersistentHistoryToken(fromStores:)
func (p_ PersistentStoreCoordinator) CurrentPersistentHistoryTokenFromStores(stores objectivec.IObject) IPersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](p_.ID, objc.Sel("currentPersistentHistoryTokenFromStores:"), stores)
	return rv
}


// Executes all remaining tasks of a deferred lightweight migration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/finishDeferredLightweightMigration()
func (p_ PersistentStoreCoordinator) FinishDeferredLightweightMigration(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("finishDeferredLightweightMigration:"), error_)
	return rv
}


// Returns the object identifier for the specified URI representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectID(forURIRepresentation:)
func (p_ PersistentStoreCoordinator) ManagedObjectIDForURIRepresentation(url foundation.URL) IManagedObjectID {
	rv := objc.Send[ManagedObjectID](p_.ID, objc.Sel("managedObjectIDForURIRepresentation:"), url)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectIDFromUTF8String:length:
func (p_ PersistentStoreCoordinator) ManagedObjectIDFromUTF8StringLength(utf8string unsafe.Pointer, len_ uint) IManagedObjectID {
	rv := objc.Send[ManagedObjectID](p_.ID, objc.Sel("managedObjectIDFromUTF8String:length:"), utf8string, len_)
	return rv
}


// Updates the metadata for the specified persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/setMetadata(_:for:)
func (p_ PersistentStoreCoordinator) SetMetadataForPersistentStore(metadata foundation.IDictionary, store IPersistentStore) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetadata:forPersistentStore:"), metadata, store)
}


// Returns the location of the provided persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/url(for:)
func (p_ PersistentStoreCoordinator) URLForPersistentStore(store IPersistentStore) foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("URLForPersistentStore:"), store)
	return rv
}


// The coordinator’s managed object model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/managedObjectModel
func (p_ PersistentStoreCoordinator) ManagedObjectModel() IManagedObjectModel {
	rv := objc.Send[ManagedObjectModel](p_.ID, objc.Sel("managedObjectModel"))
	return rv
}


// The coordinator’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/name
func (p_ PersistentStoreCoordinator) Name() string {
	rv := objc.Send[string](p_.ID, objc.Sel("name"))
	return rv
}


// The coordinator’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/name
func (p_ PersistentStoreCoordinator) SetName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), objc.String(value))
}


// The coordinator’s registered store types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreCoordinator/registeredStoreTypes
func (p_ PersistentStoreCoordinator) RegisteredStoreTypes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("registeredStoreTypes"))
	return rv
}


// The key you use to specify your Core Spotlight delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nscoredatacorespotlightexporter
func (p_ PersistentStoreCoordinator) NSCoreDataCoreSpotlightExporter() string {
	rv := objc.Send[string](p_.ID, objc.Sel("NSCoreDataCoreSpotlightExporter"))
	return rv
}


// The key you use to enable persistent history tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytrackingkey
func (p_ PersistentStoreCoordinator) NSPersistentHistoryTrackingKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("NSPersistentHistoryTrackingKey"))
	return rv
}


// The coordinator’s persistent stores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/persistentstores
func (p_ PersistentStoreCoordinator) PersistentStores() IPersistentStore {
	rv := objc.Send[PersistentStore](p_.ID, objc.Sel("persistentStores"))
	return rv
}


// The coordinator’s persistent stores.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorecoordinator/persistentstores
func (p_ PersistentStoreCoordinator) SetPersistentStores(value IPersistentStore) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPersistentStores:"), value)
}


// The key for enabling deferred lightweight migrations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredeferredlightweightmigrationoptionkey
func (p_ PersistentStoreCoordinator) NSPersistentStoreDeferredLightweightMigrationOptionKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("NSPersistentStoreDeferredLightweightMigrationOptionKey"))
	return rv
}


// A key that identifies the store type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstoretypekey
func (p_ PersistentStoreCoordinator) NSStoreTypeKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("NSStoreTypeKey"))
	return rv
}


// A key that provides the store’s UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstoreuuidkey
func (p_ PersistentStoreCoordinator) NSStoreUUIDKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("NSStoreUUIDKey"))
	return rv
}


