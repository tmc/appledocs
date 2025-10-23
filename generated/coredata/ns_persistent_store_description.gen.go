// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreDescription] class.
var (
	PersistentStoreDescriptionClass     _PersistentStoreDescriptionClass
	PersistentStoreDescriptionClassOnce sync.Once
)

func getPersistentStoreDescriptionClass() _PersistentStoreDescriptionClass {
	PersistentStoreDescriptionClassOnce.Do(func() {
		PersistentStoreDescriptionClass = _PersistentStoreDescriptionClass{objc.GetClass("NSPersistentStoreDescription")}
	})
	return PersistentStoreDescriptionClass
}

type _PersistentStoreDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreDescription] class.
type IPersistentStoreDescription interface {
	objectivec.IObject
	// properties:
	Configuration() string /* primitive/slice/pointer. */
	SetConfiguration(value string /* primitive/slice/pointer. */)
	ReadOnly() bool /* primitive/slice/pointer. */
	SetReadOnly(value bool /* primitive/slice/pointer. */)
	Options() foundation.IDictionary /* already interface */
	ShouldAddStoreAsynchronously() bool /* primitive/slice/pointer. */
	SetShouldAddStoreAsynchronously(value bool /* primitive/slice/pointer. */)
	ShouldInferMappingModelAutomatically() bool /* primitive/slice/pointer. */
	SetShouldInferMappingModelAutomatically(value bool /* primitive/slice/pointer. */)
	ShouldMigrateStoreAutomatically() bool /* primitive/slice/pointer. */
	SetShouldMigrateStoreAutomatically(value bool /* primitive/slice/pointer. */)
	SqlitePragmas() foundation.IDictionary /* already interface */
	Timeout() foundation.TimeInterval /* not a class type */
	SetTimeout(value foundation.TimeInterval /* not a class type */)
	CloudKitContainerOptions() IPersistentCloudKitContainerOptions
	SetCloudKitContainerOptions(value IPersistentCloudKitContainerOptions)
	IsReadOnly() bool /* primitive/slice/pointer. */
	SetIsReadOnly(value bool /* primitive/slice/pointer. */)
	Type() string /* primitive/slice/pointer. */
	SetType(value string /* primitive/slice/pointer. */)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
	SetOptionForKey(option objectivec.IObject, key string /* primitive/slice/pointer. */)
	SetValueForPragmaNamed(value objectivec.IObject, name string /* primitive/slice/pointer. */)
}

// A description object used to create and load a persistent store.


// A description object used to create and load a persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription
type PersistentStoreDescription struct {
	objectivec.Object
}

// PersistentStoreDescriptionFrom constructs a [PersistentStoreDescription] from an unsafe.Pointer.
//
// A description object used to create and load a persistent store.
func PersistentStoreDescriptionFrom(ptr unsafe.Pointer) PersistentStoreDescription {
	return PersistentStoreDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreDescriptionClass) Alloc() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreDescriptionClass) New() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreDescription) Init() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreDescription) Autorelease() PersistentStoreDescription {
	rv := objc.Send[PersistentStoreDescription](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreDescription creates a new PersistentStoreDescription instance.
func NewPersistentStoreDescription() PersistentStoreDescription {
	return getPersistentStoreDescriptionClass().New()
}



// Initializes the receiver with a URL for the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/init(url:)
func NewPersistentStoreDescriptionWithURL(url foundation.objc.IObject /* cross-framework URL */) PersistentStoreDescription {
	instance := getPersistentStoreDescriptionClass().Alloc()
	rv := objc.Send[PersistentStoreDescription](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}



// Initializes and returns a persistent store description with the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/persistentStoreDescriptionWithURL:
func (pc _PersistentStoreDescriptionClass) PersistentStoreDescriptionWithURL(URL foundation.objc.IObject /* cross-framework URL */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("persistentStoreDescriptionWithURL:"), URL)
	return rv
}


// Sets an option on the store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setOption(_:forKey:)
func (p_ PersistentStoreDescription) SetOptionForKey(option objectivec.IObject, key string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOption:forKey:"), option, objc.String(key))
}


// Allows you to set pragmas for the SQLite store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setValue(_:forPragmaNamed:)
func (p_ PersistentStoreDescription) SetValueForPragmaNamed(value objectivec.IObject, name string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:forPragmaNamed:"), value, objc.String(name))
}


// The name of the configuration used by this store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/configuration
func (p_ PersistentStoreDescription) Configuration() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("configuration"))
	return rv
}


// The name of the configuration used by this store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/configuration
func (p_ PersistentStoreDescription) SetConfiguration(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguration:"), objc.String(value))
}


// A flag that indicates whether this store will be read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/isReadOnly
func (p_ PersistentStoreDescription) ReadOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("readOnly"))
	return rv
}


// A flag that indicates whether this store will be read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/isReadOnly
func (p_ PersistentStoreDescription) SetReadOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadOnly:"), value)
}


// A dictionary representation of the options set on the associated persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/options
func (p_ PersistentStoreDescription) Options() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("options"))
	return rv
}


// A flag that determines whether the store is added asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldAddStoreAsynchronously
func (p_ PersistentStoreDescription) ShouldAddStoreAsynchronously() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldAddStoreAsynchronously"))
	return rv
}


// A flag that determines whether the store is added asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldAddStoreAsynchronously
func (p_ PersistentStoreDescription) SetShouldAddStoreAsynchronously(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldAddStoreAsynchronously:"), value)
}


// A flag indicating whether a mapping model should be created automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldInferMappingModelAutomatically
func (p_ PersistentStoreDescription) ShouldInferMappingModelAutomatically() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldInferMappingModelAutomatically"))
	return rv
}


// A flag indicating whether a mapping model should be created automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldInferMappingModelAutomatically
func (p_ PersistentStoreDescription) SetShouldInferMappingModelAutomatically(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldInferMappingModelAutomatically:"), value)
}


// A flag indicating whether the associated persistent store should be migrated automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldMigrateStoreAutomatically
func (p_ PersistentStoreDescription) ShouldMigrateStoreAutomatically() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldMigrateStoreAutomatically"))
	return rv
}


// A flag indicating whether the associated persistent store should be migrated automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldMigrateStoreAutomatically
func (p_ PersistentStoreDescription) SetShouldMigrateStoreAutomatically(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldMigrateStoreAutomatically:"), value)
}


// The SQLite pragmas set for the associated persistent store. (read-only)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/sqlitePragmas
func (p_ PersistentStoreDescription) SqlitePragmas() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("sqlitePragmas"))
	return rv
}


// The connection timeout for the associated store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/timeout
func (p_ PersistentStoreDescription) Timeout() foundation.TimeInterval /* not a class type */ {
	rv := objc.Send[foundation.TimeInterval](p_.ID, objc.Sel("timeout"))
	return rv
}


// The connection timeout for the associated store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/timeout
func (p_ PersistentStoreDescription) SetTimeout(value foundation.TimeInterval /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeout:"), value)
}


// Options that customize how this store description aligns with a CloudKit database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/cloudkitcontaineroptions
func (p_ PersistentStoreDescription) CloudKitContainerOptions() IPersistentCloudKitContainerOptions {
	rv := objc.Send[PersistentCloudKitContainerOptions](p_.ID, objc.Sel("cloudKitContainerOptions"))
	return rv
}


// Options that customize how this store description aligns with a CloudKit database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/cloudkitcontaineroptions
func (p_ PersistentStoreDescription) SetCloudKitContainerOptions(value IPersistentCloudKitContainerOptions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCloudKitContainerOptions:"), value)
}


// A flag that indicates whether this store will be read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/isreadonly
func (p_ PersistentStoreDescription) IsReadOnly() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// A flag that indicates whether this store will be read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/isreadonly
func (p_ PersistentStoreDescription) SetIsReadOnly(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}


// The type of store this description represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/type
func (p_ PersistentStoreDescription) Type() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("type"))
	return rv
}


// The type of store this description represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/type
func (p_ PersistentStoreDescription) SetType(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), objc.String(value))
}


// The URL that the store will use for its location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/url
func (p_ PersistentStoreDescription) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// The URL that the store will use for its location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/url
func (p_ PersistentStoreDescription) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}


