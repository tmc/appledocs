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
	// properties:
	Metadata() foundation.IDictionary /* already interface */
	SetMetadata(value foundation.IDictionary /* already interface */)
	Type() string /* primitive/slice/pointer. */
	ConfigurationName() string /* primitive/slice/pointer. */
	SetConfigurationName(value string /* primitive/slice/pointer. */)
	CoreSpotlightExporter() ICoreDataCoreSpotlightDelegate
	SetCoreSpotlightExporter(value ICoreDataCoreSpotlightDelegate)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	IsReadOnly() bool /* primitive/slice/pointer. */
	SetIsReadOnly(value bool /* primitive/slice/pointer. */)
	Options() unsafe.Pointer
	SetOptions(value unsafe.Pointer)
	PersistentStoreCoordinator() IPersistentStoreCoordinator
	SetPersistentStoreCoordinator(value IPersistentStoreCoordinator)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

// The abstract base class for all Core Data persistent stores.
//
// Core Data provides four store types—SQLite, Binary, XML, and In-Memory (the XML store is not available on iOS); these are described in Persistent Store Features. Core Data also provides subclasses of that you can use to define your own store types: and . The Binary and XML stores are examples of atomic stores that inherit functionality from .


// The abstract base class for all Core Data persistent stores.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/init(persistentStoreCoordinator:configurationName:at:options:)
func NewPersistentStoreWithPersistentStoreCoordinatorConfigurationNameURLOptions(root IPersistentStoreCoordinator, name string /* primitive/slice/pointer. */, url foundation.objc.IObject /* cross-framework URL */, options objectivec.IObject) PersistentStore {
	instance := getPersistentStoreClass().Alloc()
	rv := objc.Send[PersistentStore](instance.ID, objc.Sel("initWithPersistentStoreCoordinator:configurationName:URL:options:"), root, objc.String(name), url, options)
	rv.Autorelease()
	return rv
}



// The metadata for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadata
func (p_ PersistentStore) Metadata() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/metadata
func (p_ PersistentStore) SetMetadata(value foundation.IDictionary /* already interface */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMetadata:"), value)
}


// The type string of the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStore/type
func (p_ PersistentStore) Type() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("type"))
	return rv
}


// The name of the managed object model configuration that creates the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/configurationname
func (p_ PersistentStore) ConfigurationName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("configurationName"))
	return rv
}


// The name of the managed object model configuration that creates the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/configurationname
func (p_ PersistentStore) SetConfigurationName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfigurationName:"), objc.String(value))
}


// The spotlight exporter associated with this persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/corespotlightexporter
func (p_ PersistentStore) CoreSpotlightExporter() ICoreDataCoreSpotlightDelegate {
	rv := objc.Send[CoreDataCoreSpotlightDelegate](p_.ID, objc.Sel("coreSpotlightExporter"))
	return rv
}


// The spotlight exporter associated with this persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/corespotlightexporter
func (p_ PersistentStore) SetCoreSpotlightExporter(value ICoreDataCoreSpotlightDelegate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCoreSpotlightExporter:"), value)
}


// The unique identifier for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (p_ PersistentStore) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (p_ PersistentStore) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// A Boolean value that indicates whether the persistent store is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/isreadonly
func (p_ PersistentStore) IsReadOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// A Boolean value that indicates whether the persistent store is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/isreadonly
func (p_ PersistentStore) SetIsReadOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}


// The options that Core Data uses to create the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/options
func (p_ PersistentStore) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("options"))
	return rv
}


// The options that Core Data uses to create the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/options
func (p_ PersistentStore) SetOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOptions:"), value)
}


// The persistent store coordinator that loads the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/persistentstorecoordinator
func (p_ PersistentStore) PersistentStoreCoordinator() IPersistentStoreCoordinator {
	rv := objc.Send[PersistentStoreCoordinator](p_.ID, objc.Sel("persistentStoreCoordinator"))
	return rv
}


// The persistent store coordinator that loads the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/persistentstorecoordinator
func (p_ PersistentStore) SetPersistentStoreCoordinator(value IPersistentStoreCoordinator) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPersistentStoreCoordinator:"), value)
}


// The URL for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/url
func (p_ PersistentStore) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// The URL for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/url
func (p_ PersistentStore) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}


