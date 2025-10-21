// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
	SetOptionForKey(option foundation.IObject, key appkit.string)
	SetValueForPragmaNamed(value foundation.IObject, name appkit.string)
}

// A description object used to create and load a persistent store.
//
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
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/init(url:)
func NewPersistentStoreDescriptionWithURL(url foundation.IURL) PersistentStoreDescription {
	instance := getPersistentStoreDescriptionClass().Alloc()
	rv := objc.Send[PersistentStoreDescription](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Initializes and returns a persistent store description with the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/persistentStoreDescriptionWithURL:
func (pc _PersistentStoreDescriptionClass) PersistentStoreDescriptionWithURL(URL foundation.IURL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("persistentStoreDescriptionWithURL:"), URL)
	return rv
}

// Sets an option on the store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setOption(_:forKey:)
func (p_ PersistentStoreDescription) SetOptionForKey(option foundation.IObject, key appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOption:forKey:"), option, key)
}

// Allows you to set pragmas for the SQLite store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/setValue(_:forPragmaNamed:)
func (p_ PersistentStoreDescription) SetValueForPragmaNamed(value foundation.IObject, name appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setValue:forPragmaNamed:"), value, name)
}

// Options that customize how this store description aligns with a CloudKit database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/cloudKitContainerOptions
func (p_ PersistentStoreDescription) CloudKitContainerOptions() NSPersistentCloudKitContainerOptions {
	rv := objc.Send[NSPersistentCloudKitContainerOptions](p_.ID, objc.Sel("cloudKitContainerOptions"))
	return rv
}


// SetCloudKitContainerOptions sets the value of the cloudKitContainerOptions property.
// Options that customize how this store description aligns with a CloudKit database.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/cloudKitContainerOptions
func (p_ PersistentStoreDescription) SetCloudKitContainerOptions(value NSPersistentCloudKitContainerOptions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCloudKitContainerOptions:"), value)
}

// The name of the configuration used by this store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/configuration
func (p_ PersistentStoreDescription) Configuration() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// The name of the configuration used by this store.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/configuration
func (p_ PersistentStoreDescription) SetConfiguration(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setConfiguration:"), value)
}

// A flag that indicates whether this store will be read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/isReadOnly
func (p_ PersistentStoreDescription) ReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("readOnly"))
	return rv
}


// SetReadOnly sets the value of the readOnly property.
// A flag that indicates whether this store will be read-only.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/isReadOnly
func (p_ PersistentStoreDescription) SetReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReadOnly:"), value)
}

// A dictionary representation of the options set on the associated persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/options
func (p_ PersistentStoreDescription) Options() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("options"))
	return rv
}

// A flag that determines whether the store is added asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldAddStoreAsynchronously
func (p_ PersistentStoreDescription) ShouldAddStoreAsynchronously() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldAddStoreAsynchronously"))
	return rv
}


// SetShouldAddStoreAsynchronously sets the value of the shouldAddStoreAsynchronously property.
// A flag that determines whether the store is added asynchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldAddStoreAsynchronously
func (p_ PersistentStoreDescription) SetShouldAddStoreAsynchronously(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldAddStoreAsynchronously:"), value)
}

// A flag indicating whether a mapping model should be created automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldInferMappingModelAutomatically
func (p_ PersistentStoreDescription) ShouldInferMappingModelAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldInferMappingModelAutomatically"))
	return rv
}


// SetShouldInferMappingModelAutomatically sets the value of the shouldInferMappingModelAutomatically property.
// A flag indicating whether a mapping model should be created automatically.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldInferMappingModelAutomatically
func (p_ PersistentStoreDescription) SetShouldInferMappingModelAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldInferMappingModelAutomatically:"), value)
}

// A flag indicating whether the associated persistent store should be migrated automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldMigrateStoreAutomatically
func (p_ PersistentStoreDescription) ShouldMigrateStoreAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldMigrateStoreAutomatically"))
	return rv
}


// SetShouldMigrateStoreAutomatically sets the value of the shouldMigrateStoreAutomatically property.
// A flag indicating whether the associated persistent store should be migrated automatically.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/shouldMigrateStoreAutomatically
func (p_ PersistentStoreDescription) SetShouldMigrateStoreAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldMigrateStoreAutomatically:"), value)
}

// The SQLite pragmas set for the associated persistent store. (read-only)
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/sqlitePragmas
func (p_ PersistentStoreDescription) SqlitePragmas() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("sqlitePragmas"))
	return rv
}

// The connection timeout for the associated store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/timeout
func (p_ PersistentStoreDescription) Timeout() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](p_.ID, objc.Sel("timeout"))
	return rv
}


// SetTimeout sets the value of the timeout property.
// The connection timeout for the associated store.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/timeout
func (p_ PersistentStoreDescription) SetTimeout(value foundation.ITimeInterval) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeout:"), value)
}

// The URL that the store will use for its location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/url
func (p_ PersistentStoreDescription) URL() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL that the store will use for its location.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreDescription/url
func (p_ PersistentStoreDescription) SetURL(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}

// A flag that indicates whether this store will be read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/isreadonly
func (p_ PersistentStoreDescription) IsReadOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadOnly"))
	return rv
}


// SetIsReadOnly sets the value of the isReadOnly property.
// A flag that indicates whether this store will be read-only.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/isreadonly
func (p_ PersistentStoreDescription) SetIsReadOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadOnly:"), value)
}

// The type of store this description represents.
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/type
func (p_ PersistentStoreDescription) Type() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of store this description represents.

//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstoredescription/type
func (p_ PersistentStoreDescription) SetType(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}


