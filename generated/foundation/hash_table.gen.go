// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HashTable] class.
var (
	hashTableClass     _HashTableClass
	hashTableClassOnce sync.Once
)

func getHashTableClass() _HashTableClass {
	hashTableClassOnce.Do(func() {
		hashTableClass = _HashTableClass{objc.GetClass("NSHashTable")}
	})
	return hashTableClass
}

type _HashTableClass struct {
	class objc.Class
}

// An interface definition for the [HashTable] class.
type IHashTable interface {
	objectivec.IObject
}

// A collection similar to a set, but with broader range of available memory semantics.
//
// The hash table is modeled after with the following differences: It can hold weak references to its members. Its members may be copied on input or may use pointer identity for equality and hashing. It can contain arbitrary pointers (its members are not constrained to being objects). You can configure an instance to operate on arbitrary pointers and not just objects, although typically you are encouraged to use the C function API for void * pointers. The object-based API (such as ) will not work for non-object pointers without type-casting. Because of its options, is not a set because it can behave differently (for example, if pointer equality is specified two strings will both be entered). When configuring hash tables, note that only the options listed in guarantee that the rest of the API will work correctly—including copying, archiving, and fast enumeration. While other options are used for certain configurations, such as to hold arbitrary pointers, not all combinations of the options are valid. With some combinations the hash table may not work correctly, or may not even be initialized correctly.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable
type HashTable struct {
	objectivec.Object
}

// HashTableFrom constructs a [HashTable] from an unsafe.Pointer.
//
// A collection similar to a set, but with broader range of available memory semantics.
func HashTableFrom(ptr unsafe.Pointer) HashTable {
	return HashTable{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HashTableClass) Alloc() HashTable {
	rv := objc.Send[HashTable](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HashTableClass) New() HashTable {
	rv := objc.Send[HashTable](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HashTable) Init() HashTable {
	rv := objc.Send[HashTable](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HashTable) Autorelease() HashTable {
	rv := objc.Send[HashTable](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHashTable creates a new HashTable instance.
func NewHashTable() HashTable {
	return getHashTableClass().New()
}




