// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStore] class.
var (
	PersistentStoreClass     _PersistentStoreClass
	PersistentStoreClassOnce sync.Once
)

func getPersistentStoreClass() _PersistentStoreClass {
	PersistentStoreClassOnce.Do(func() {
		PersistentStoreClass = _PersistentStoreClass{objc.GetClass("NSPersistentStore")}
	})
	return PersistentStoreClass
}

type _PersistentStoreClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStore] class.
type IPersistentStore interface {
	objectivec.IObject
	LoadMetadata(error_ unsafe.Pointer) bool
	ReadOnly() bool
	SetReadOnly(value bool)
	Metadata() unsafe.Pointer
	SetMetadata(value unsafe.Pointer)
	Options() objc.ID
	PersistentStoreCoordinator() NSPersistentStoreCoordinator
	Type() string
	ConfigurationName() string
	SetConfigurationName(value string)
	CoreSpotlightExporter() NSCoreDataCoreSpotlightDelegate
	SetCoreSpotlightExporter(value ICoreDataCoreSpotlightDelegate)
	Identifier() string
	SetIdentifier(value string)
	IsReadOnly() bool
	SetIsReadOnly(value bool)
	Url() foundation.URL
	SetUrl(value foundation.IURL)
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
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string, url foundation.IURL, options objectivec.IObject) PersistentStore {
	instance := getPersistentStoreClass().Alloc()
	rv := objc.Send[PersistentStore](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	rv.Autorelease()
	return rv
}


// Returns the metadata from the persistent store at the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadataForPersistentStore(with:)
func (pc _PersistentStoreClass) MetadataForPersistentStoreWithURLError(url foundation.IURL, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("metadataForPersistentStoreWithURL:error:"), url, error_)
	return rv
}

// Returns the migration manager class for this store class.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/migrationManagerClass()
func (pc _PersistentStoreClass) MigrationManagerClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(pc.class), objc.Sel("migrationManagerClass"))
	return rv
}

// Sets the metadata for the store at a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/setMetadata(_:forPersistentStoreAt:)
func (pc _PersistentStoreClass) SetMetadataForPersistentStoreWithURLError(metadata unsafe.Pointer, url foundation.IURL, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("setMetadata:forPersistentStoreWithURL:error:"), metadata, url, error_)
	return rv
}

// Instructs the persistent store to load its metadata.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/loadMetadata()
func (p_ PersistentStore) LoadMetadata(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("loadMetadata:"), error_)
	return rv
}

// A Boolean value that indicates whether the persistent store is read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/isReadOnly
func (p_ PersistentStore) ReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readOnly"))
	return rv
}


// SetReadOnly sets the value of the readOnly property.
// A Boolean value that indicates whether the persistent store is read-only.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/isReadOnly
func (p_ PersistentStore) SetReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadOnly:"), value)
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

// The options that Core Data uses to create the store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/options
func (p_ PersistentStore) Options() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("options"))
	return rv
}

// The persistent store coordinator that loads the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/persistentStoreCoordinator
func (p_ PersistentStore) PersistentStoreCoordinator() NSPersistentStoreCoordinator {
	rv := objc.Send[NSPersistentStoreCoordinator](p_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}

// The type string of the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/type
func (p_ PersistentStore) Type() string {
	rv := objc.Send[string](p_.ID, objc.Sel("type"))
	return rv
}

// The name of the managed object model configuration that creates the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/configurationname
func (p_ PersistentStore) ConfigurationName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("configurationName"))
	return rv
}


// SetConfigurationName sets the value of the configurationName property.
// The name of the managed object model configuration that creates the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/configurationname
func (p_ PersistentStore) SetConfigurationName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfigurationName:"), objc.String(value))
}

// The spotlight exporter associated with this persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/corespotlightexporter
func (p_ PersistentStore) CoreSpotlightExporter() NSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](p_.ID, objc.Sel("coreSpotlightExporter"))
	return rv
}


// SetCoreSpotlightExporter sets the value of the coreSpotlightExporter property.
// The spotlight exporter associated with this persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/corespotlightexporter
func (p_ PersistentStore) SetCoreSpotlightExporter(value ICoreDataCoreSpotlightDelegate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCoreSpotlightExporter:"), value)
}

// The unique identifier for the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (p_ PersistentStore) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The unique identifier for the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (p_ PersistentStore) SetIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether the persistent store is read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/isreadonly
func (p_ PersistentStore) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// SetIsReadOnly sets the value of the isReadOnly property.
// A Boolean value that indicates whether the persistent store is read-only.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/isreadonly
func (p_ PersistentStore) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}

// The URL for the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/url
func (p_ PersistentStore) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL for the persistent store.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/url
func (p_ PersistentStore) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}


