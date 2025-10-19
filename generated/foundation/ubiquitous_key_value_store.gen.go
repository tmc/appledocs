// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UbiquitousKeyValueStore] class.
var (
	ubiquitousKeyValueStoreClass     _UbiquitousKeyValueStoreClass
	ubiquitousKeyValueStoreClassOnce sync.Once
)

func getUbiquitousKeyValueStoreClass() _UbiquitousKeyValueStoreClass {
	ubiquitousKeyValueStoreClassOnce.Do(func() {
		ubiquitousKeyValueStoreClass = _UbiquitousKeyValueStoreClass{objc.GetClass("NSUbiquitousKeyValueStore")}
	})
	return ubiquitousKeyValueStoreClass
}

type _UbiquitousKeyValueStoreClass struct {
	class objc.Class
}

// An interface definition for the [UbiquitousKeyValueStore] class.
type IUbiquitousKeyValueStore interface {
	objectivec.IObject
}

// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore
type UbiquitousKeyValueStore struct {
	objectivec.Object
}

// UbiquitousKeyValueStoreFrom constructs a [UbiquitousKeyValueStore] from an unsafe.Pointer.
//
// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices.
func UbiquitousKeyValueStoreFrom(ptr unsafe.Pointer) UbiquitousKeyValueStore {
	return UbiquitousKeyValueStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UbiquitousKeyValueStoreClass) Alloc() UbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UbiquitousKeyValueStoreClass) New() UbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UbiquitousKeyValueStore) Init() UbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UbiquitousKeyValueStore) Autorelease() UbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUbiquitousKeyValueStore creates a new UbiquitousKeyValueStore instance.
func NewUbiquitousKeyValueStore() UbiquitousKeyValueStore {
	return getUbiquitousKeyValueStoreClass().New()
}




