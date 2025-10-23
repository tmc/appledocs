// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AtomicStore] class.
var (
	AtomicStoreClass     _AtomicStoreClass
	AtomicStoreClassOnce sync.Once
)

func getAtomicStoreClass() _AtomicStoreClass {
	AtomicStoreClassOnce.Do(func() {
		AtomicStoreClass = _AtomicStoreClass{objc.GetClass("NSAtomicStore")}
	})
	return AtomicStoreClass
}

type _AtomicStoreClass struct {
	class objc.Class
}

// An interface definition for the [AtomicStore] class.
type IAtomicStore interface {
	IPersistentStore
	Identifier() string
	SetIdentifier(value string)
	Metadata() string
	SetMetadata(value string)
	Type() string
	SetType(value string)
	NSStoreTypeKey() string
	NSStoreUUIDKey() string
	Save(error_ unsafe.Pointer) bool
}

// An abstract superclass that you subclass to create a Core Data atomic store.
//
// Use an atomic store to handle data sets that can be expressed in memory. The atomic store API favors simplicity over performance. This class provides default implementations of some utility methods. Create a custom atomic store subclass when you have a custom file format that you want to integrate with a Core Data app. When you create a subclass, override the following methods: Also override the following properties and methods of , from which the atomic store class inherits: provides a default dictionary of metadata. This dictionary contains the store type and identifier ( and ) as well as store versioning information. Subclasses must ensure that the metadata is saved along with the store data.


// An abstract superclass that you subclass to create a Core Data atomic store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStore
type AtomicStore struct {
	PersistentStore
}

// AtomicStoreFrom constructs a [AtomicStore] from an unsafe.Pointer.
//
// An abstract superclass that you subclass to create a Core Data atomic store.
func AtomicStoreFrom(ptr unsafe.Pointer) AtomicStore {
	return AtomicStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AtomicStoreClass) Alloc() AtomicStore {
	rv := objc.Send[AtomicStore](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AtomicStoreClass) New() AtomicStore {
	rv := objc.Send[AtomicStore](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AtomicStore) Init() AtomicStore {
	rv := objc.Send[AtomicStore](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AtomicStore) Autorelease() AtomicStore {
	rv := objc.Send[AtomicStore](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAtomicStore creates a new AtomicStore instance.
func NewAtomicStore() AtomicStore {
	return getAtomicStoreClass().New()
}



// Saves the cache nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStore/save()
func (a_ AtomicStore) Save(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("save:"), error_)
	return rv
}


// The unique identifier for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (a_ AtomicStore) Identifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/identifier
func (a_ AtomicStore) SetIdentifier(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The metadata for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/metadata
func (a_ AtomicStore) Metadata() string {
	rv := objc.Send[string](a_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata for the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/metadata
func (a_ AtomicStore) SetMetadata(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), objc.String(value))
}


// The type string of the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/type
func (a_ AtomicStore) Type() string {
	rv := objc.Send[string](a_.ID, objc.Sel("type"))
	return rv
}


// The type string of the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstore/type
func (a_ AtomicStore) SetType(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), objc.String(value))
}


// A key that identifies the store type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstoretypekey
func (a_ AtomicStore) NSStoreTypeKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("NSStoreTypeKey"))
	return rv
}


// A key that provides the store’s UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsstoreuuidkey
func (a_ AtomicStore) NSStoreUUIDKey() string {
	rv := objc.Send[string](a_.ID, objc.Sel("NSStoreUUIDKey"))
	return rv
}



