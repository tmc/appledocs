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
	UbiquitousKeyValueStoreClass     _UbiquitousKeyValueStoreClass
	UbiquitousKeyValueStoreClassOnce sync.Once
)

func getUbiquitousKeyValueStoreClass() _UbiquitousKeyValueStoreClass {
	UbiquitousKeyValueStoreClassOnce.Do(func() {
		UbiquitousKeyValueStoreClass = _UbiquitousKeyValueStoreClass{objc.GetClass("NSUbiquitousKeyValueStore")}
	})
	return UbiquitousKeyValueStoreClass
}

type _UbiquitousKeyValueStoreClass struct {
	class objc.Class
}

// An interface definition for the [UbiquitousKeyValueStore] class.
type IUbiquitousKeyValueStore interface {
	objectivec.IObject
	// properties:
	DictionaryRepresentation() IDictionary /* already interface */
	NSUbiquitousKeyValueStoreChangeReasonKey() IString
	NSUbiquitousKeyValueStoreQuotaViolationChange() int /* primitive/slice/pointer. */
	SetNSUbiquitousKeyValueStoreQuotaViolationChange(value int /* primitive/slice/pointer. */)
	// methods:
	ArrayForKey(aKey IString) IArray
	BoolForKey(aKey IString) bool /* primitive/slice/pointer. */
	DataForKey(aKey IString) IData
	DictionaryForKey(aKey IString) IDictionary /* already interface */
	DoubleForKey(aKey IString) float64 /* primitive/slice/pointer. */
	LongLongForKey(aKey IString) unsafe.Pointer
	ObjectForKey(aKey IString) objc.ID
	RemoveObjectForKey(aKey IString)
	SetDoubleForKey(value float64 /* primitive/slice/pointer. */, aKey IString)
	SetStringForKey(aString IString, aKey IString)
	SetDataForKey(aData IData, aKey IString)
	SetArrayForKey(anArray IArray, aKey IString)
	SetLongLongForKey(value unsafe.Pointer, aKey IString)
	SetBoolForKey(value bool /* primitive/slice/pointer. */, aKey IString)
	SetObjectForKey(anObject objectivec.IObject, aKey IString)
	SetDictionaryForKey(aDictionary IDictionary /* already interface */, aKey IString)
	StringForKey(aKey IString) IString
	Synchronize() bool /* primitive/slice/pointer. */
}

// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices.
//
// Use the iCloud key-value store to make preference, configuration, and app-state data available to every instance of your app on every device connected to a user’s iCloud account. You can store scalar values such as , as well as values containing any of the property list object types: , , , , , and . Changes your app writes to the key-value store object are initially held in memory, then written to disk by the system at appropriate times. If you write to the key-value store object when the user is not signed into an iCloud account, the data is stored locally until the next synchronization opportunity. When the user signs into an iCloud account, the system automatically reconciles your local, on-disk keys and values with those on the iCloud server. Any device running your app, and attached to the same iCloud account, can upload key-value changes to iCloud. To keep track of such changes, register for the notification during app launch. Then, obtain the keys and values from iCloud (which may be newer than those that are local) by calling the method. You need not call the method again during your app’s life cycle, unless your app design requires fast-as-possible upload to iCloud after you change a value. For more information on adopting key-value storage in your app, see in . Avoid using this class for data that is essential to your app’s behavior when offline; instead, store such data directly into the local user defaults database. The total amount of space available in your app’s key-value store, for a given user, is 1 MB. There is a per-key value size limit of 1 MB, and a maximum of 1024 keys. If you attempt to write data that exceeds these quotas, the write attempt fails and no change is made to your iCloud key-value storage. In this scenario, the system posts the notification with a change reason of . The maximum length for key strings for the iCloud key-value store is 64 bytes using UTF8 encoding. Attempting to write a value to a longer key name results in a runtime error. To use this class, you must distribute your app through the App Store or Mac App Store, and you must request the entitlement in your Xcode project. For more on this, see in . This class is not meant to be subclassed.


// An iCloud-based container of key-value pairs you use to share data among instances of your app running on a user’s connected devices.
//
// [Full Topic]
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



// Returns the shared iCloud key-value store object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/default
func (uc _UbiquitousKeyValueStoreClass) DefaultStore() UbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](objc.ID(uc.class), objc.Sel("defaultStore"))
	return rv
}

// Returns the array associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/array(forKey:)
func (u_ UbiquitousKeyValueStore) ArrayForKey(aKey IString) IArray {
	rv := objc.Send[Array](u_.ID, objc.Sel("arrayForKey:"), aKey)
	return rv
}


// Returns the Boolean value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/bool(forKey:)
func (u_ UbiquitousKeyValueStore) BoolForKey(aKey IString) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("boolForKey:"), aKey)
	return rv
}


// Returns the data object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/data(forKey:)
func (u_ UbiquitousKeyValueStore) DataForKey(aKey IString) IData {
	rv := objc.Send[Data](u_.ID, objc.Sel("dataForKey:"), aKey)
	return rv
}


// Returns the dictionary object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/dictionary(forKey:)
func (u_ UbiquitousKeyValueStore) DictionaryForKey(aKey IString) IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("dictionaryForKey:"), aKey)
	return rv
}


// Returns the double value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/double(forKey:)
func (u_ UbiquitousKeyValueStore) DoubleForKey(aKey IString) float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](u_.ID, objc.Sel("doubleForKey:"), aKey)
	return rv
}


// Returns the value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/longLong(forKey:)
func (u_ UbiquitousKeyValueStore) LongLongForKey(aKey IString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("longLongForKey:"), aKey)
	return rv
}


// Returns the object associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/object(forKey:)
func (u_ UbiquitousKeyValueStore) ObjectForKey(aKey IString) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("objectForKey:"), aKey)
	return rv
}


// Removes the value associated with the specified key from the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/removeObject(forKey:)
func (u_ UbiquitousKeyValueStore) RemoveObjectForKey(aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeObjectForKey:"), aKey)
}


// Sets a double value for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-1xml0
func (u_ UbiquitousKeyValueStore) SetDoubleForKey(value float64 /* primitive/slice/pointer. */, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDouble:forKey:"), value, aKey)
}


// Sets a string object for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-2rlp
func (u_ UbiquitousKeyValueStore) SetStringForKey(aString IString, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setString:forKey:"), aString, aKey)
}


// Sets a data object for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-3ga7z
func (u_ UbiquitousKeyValueStore) SetDataForKey(aData IData, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setData:forKey:"), aData, aKey)
}


// Sets an array object for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-40a8f
func (u_ UbiquitousKeyValueStore) SetArrayForKey(anArray IArray, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setArray:forKey:"), anArray, aKey)
}


// Sets a value for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-7tt20
func (u_ UbiquitousKeyValueStore) SetLongLongForKey(value unsafe.Pointer, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLongLong:forKey:"), value, aKey)
}


// Sets a Boolean value for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-8o8mq
func (u_ UbiquitousKeyValueStore) SetBoolForKey(value bool /* primitive/slice/pointer. */, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setBool:forKey:"), value, aKey)
}


// Sets an object for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-9e3de
func (u_ UbiquitousKeyValueStore) SetObjectForKey(anObject objectivec.IObject, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setObject:forKey:"), anObject, aKey)
}


// Sets a dictionary object for the specified key in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-9vmlm
func (u_ UbiquitousKeyValueStore) SetDictionaryForKey(aDictionary IDictionary /* already interface */, aKey IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDictionary:forKey:"), aDictionary, aKey)
}


// Returns the string associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/string(forKey:)
func (u_ UbiquitousKeyValueStore) StringForKey(aKey IString) IString {
	rv := objc.Send[String](u_.ID, objc.Sel("stringForKey:"), aKey)
	return rv
}


// Explicitly synchronizes in-memory keys and values with those stored on disk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/synchronize()
func (u_ UbiquitousKeyValueStore) Synchronize() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("synchronize"))
	return rv
}


// Returns the shared iCloud key-value store object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/default
func (u_ UbiquitousKeyValueStore) DefaultStore() IUbiquitousKeyValueStore {
	rv := objc.Send[UbiquitousKeyValueStore](u_.ID, objc.Sel("defaultStore"))
	return rv
}


// A dictionary containing all of the key-value pairs in the key-value store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/dictionaryRepresentation
func (u_ UbiquitousKeyValueStore) DictionaryRepresentation() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](u_.ID, objc.Sel("dictionaryRepresentation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorechangereasonkey
func (u_ UbiquitousKeyValueStore) NSUbiquitousKeyValueStoreChangeReasonKey() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUbiquitousKeyValueStoreChangeReasonKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorequotaviolationchange
func (u_ UbiquitousKeyValueStore) NSUbiquitousKeyValueStoreQuotaViolationChange() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUbiquitousKeyValueStoreQuotaViolationChange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorequotaviolationchange
func (u_ UbiquitousKeyValueStore) SetNSUbiquitousKeyValueStoreQuotaViolationChange(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNSUbiquitousKeyValueStoreQuotaViolationChange:"), value)
}



