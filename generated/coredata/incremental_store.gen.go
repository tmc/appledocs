// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [IncrementalStore] class.
var (
	incrementalStoreClass     _IncrementalStoreClass
	incrementalStoreClassOnce sync.Once
)

func getIncrementalStoreClass() _IncrementalStoreClass {
	incrementalStoreClassOnce.Do(func() {
		incrementalStoreClass = _IncrementalStoreClass{objc.GetClass("NSIncrementalStore")}
	})
	return incrementalStoreClass
}

type _IncrementalStoreClass struct {
	class objc.Class
}

// An interface definition for the [IncrementalStore] class.
type IIncrementalStore interface {
	IPersistentStore
	ExecuteRequestWithContextError(request unsafe.Pointer, context unsafe.Pointer, error unsafe.Pointer) objc.ID
	LoadMetadata(error unsafe.Pointer) bool
}

// An abstract superclass defining the API through which Core Data communicates with a store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore
type IncrementalStore struct {
	PersistentStore
}

// IncrementalStoreFrom constructs a [IncrementalStore] from an unsafe.Pointer.
//
// An abstract superclass defining the API through which Core Data communicates with a store.
func IncrementalStoreFrom(ptr unsafe.Pointer) IncrementalStore {
	return IncrementalStore{
		PersistentStore: PersistentStoreFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _IncrementalStoreClass) Alloc() IncrementalStore {
	rv := objc.Send[IncrementalStore](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IncrementalStoreClass) New() IncrementalStore {
	rv := objc.Send[IncrementalStore](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IncrementalStore) Init() IncrementalStore {
	rv := objc.Send[IncrementalStore](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IncrementalStore) Autorelease() IncrementalStore {
	rv := objc.Send[IncrementalStore](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIncrementalStore creates a new IncrementalStore instance.
func NewIncrementalStore() IncrementalStore {
	return getIncrementalStoreClass().New()
}


// Returns a value as appropriate for the given request, or nil if the request cannot be completed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/execute(_:with:)
func (i_ IncrementalStore) ExecuteRequestWithContextError(request unsafe.Pointer, context unsafe.Pointer, error unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("executeRequest:withContext:error:"), request, context, error)
	return rv
}
// Loads the metadata for the store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSIncrementalStore/loadMetadata()
func (i_ IncrementalStore) LoadMetadata(error unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("loadMetadata:"), error)
	return rv
}


