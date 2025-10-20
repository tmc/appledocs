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

// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices.
//
// Use the iCloud key-value store to make preference, configuration, and app-state data available to every instance of your app on every device connected to a user’s iCloud account. You can store scalar values such as , as well as values containing any of the property list object types: , , , , , and . Changes your app writes to the key-value store object are initially held in memory, then written to disk by the system at appropriate times. If you write to the key-value store object when the user is not signed into an iCloud account, the data is stored locally until the next synchronization opportunity. When the user signs into an iCloud account, the system automatically reconciles your local, on-disk keys and values with those on the iCloud server. Any device running your app, and attached to the same iCloud account, can upload key-value changes to iCloud. To keep track of such changes, register for the notification during app launch. Then, obtain the keys and values from iCloud (which may be newer than those that are local) by calling the method. You need not call the method again during your app’s life cycle, unless your app design requires fast-as-possible upload to iCloud after you change a value. For more information on adopting key-value storage in your app, see in . Avoid using this class for data that is essential to your app’s behavior when offline; instead, store such data directly into the local user defaults database. The total amount of space available in your app’s key-value store, for a given user, is 1 MB. There is a per-key value size limit of 1 MB, and a maximum of 1024 keys. If you attempt to write data that exceeds these quotas, the write attempt fails and no change is made to your iCloud key-value storage. In this scenario, the system posts the notification with a change reason of . The maximum length for key strings for the iCloud key-value store is 64 bytes using UTF8 encoding. Attempting to write a value to a longer key name results in a runtime error. To use this class, you must distribute your app through the App Store or Mac App Store, and you must request the entitlement in your Xcode project. For more on this, see in . This class is not meant to be subclassed.
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




