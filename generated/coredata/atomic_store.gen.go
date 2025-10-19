// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AtomicStore] class.
var (
	atomicStoreClass     _AtomicStoreClass
	atomicStoreClassOnce sync.Once
)

func getAtomicStoreClass() _AtomicStoreClass {
	atomicStoreClassOnce.Do(func() {
		atomicStoreClass = _AtomicStoreClass{objc.GetClass("NSAtomicStore")}
	})
	return atomicStoreClass
}

type _AtomicStoreClass struct {
	class objc.Class
}

// An interface definition for the [AtomicStore] class.
type IAtomicStore interface {
	IPersistentStore
	Save(error unsafe.Pointer) bool
}

// An abstract superclass that you subclass to create a Core Data atomic store. [Full Topic]
//
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


// Saves the cache nodes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSAtomicStore/save()
func (a_ AtomicStore) Save(error unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("save:"), error)
	return rv
}


