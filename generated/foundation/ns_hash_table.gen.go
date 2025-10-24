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
	HashTableClass     _HashTableClass
	HashTableClassOnce sync.Once
)

func getHashTableClass() _HashTableClass {
	HashTableClassOnce.Do(func() {
		HashTableClass = _HashTableClass{objc.GetClass("NSHashTable")}
	})
	return HashTableClass
}

type _HashTableClass struct {
	class objc.Class
}

// An interface definition for the [HashTable] class.
type IHashTable interface {
	objectivec.IObject
	// properties:
	AllObjects() []objc.ID
	AnyObject() unsafe.Pointer
	Count() uint
	PointerFunctions() IPointerFunctions
	SetRepresentation() unsafe.Pointer
	// methods:
	AddObject(object unsafe.Pointer)
	ContainsObject(anObject unsafe.Pointer) bool
	IntersectHashTable(other unsafe.Pointer)
	IntersectsHashTable(other unsafe.Pointer) bool
	IsEqualToHashTable(other unsafe.Pointer) bool
	IsSubsetOfHashTable(other unsafe.Pointer) bool
	Member(object unsafe.Pointer) unsafe.Pointer
	MinusHashTable(other unsafe.Pointer)
	ObjectEnumerator() unsafe.Pointer
	RemoveObject(object unsafe.Pointer)
	RemoveAllObjects()
	UnionHashTable(other unsafe.Pointer)
}

// A collection similar to a set, but with broader range of available memory semantics.
//
// The hash table is modeled after with the following differences: It can hold weak references to its members. Its members may be copied on input or may use pointer identity for equality and hashing. It can contain arbitrary pointers (its members are not constrained to being objects). You can configure an instance to operate on arbitrary pointers and not just objects, although typically you are encouraged to use the C function API for void * pointers. The object-based API (such as ) will not work for non-object pointers without type-casting. Because of its options, is not a set because it can behave differently (for example, if pointer equality is specified two strings will both be entered). When configuring hash tables, note that only the options listed in guarantee that the rest of the API will work correctly—including copying, archiving, and fast enumeration. While other options are used for certain configurations, such as to hold arbitrary pointers, not all combinations of the options are valid. With some combinations the hash table may not work correctly, or may not even be initialized correctly.


// A collection similar to a set, but with broader range of available memory semantics.
//
// [Full Topic]
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



// Returns a hash table with given pointer functions options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:)
func NewHashTableWithOptions(options PointerFunctionsOptions) HashTable {
	rv := objc.Send[HashTable](objc.ID(getHashTableClass().class), objc.Sel("hashTableWithOptions:"), options)
	return rv
}


// Returns a hash table initialized with the given attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:capacity:)
func NewHashTableWithOptionsCapacity(options PointerFunctionsOptions, initialCapacity uint) HashTable {
	instance := getHashTableClass().Alloc()
	rv := objc.Send[HashTable](instance.ID, objc.Sel("initWithOptions:capacity:"), options, initialCapacity)
	rv.Autorelease()
	return rv
}


// Returns a hash table initialized with the given functions and capacity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/init(pointerFunctions:capacity:)
func NewHashTableWithPointerFunctionsCapacity(functions IPointerFunctions, initialCapacity uint) HashTable {
	instance := getHashTableClass().Alloc()
	rv := objc.Send[HashTable](instance.ID, objc.Sel("initWithPointerFunctions:capacity:"), functions, initialCapacity)
	rv.Autorelease()
	return rv
}



// Returns a new hash table for storing weak references to its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/hashTableWithWeakObjects
func (hc _HashTableClass) HashTableWithWeakObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(hc.class), objc.Sel("hashTableWithWeakObjects"))
	return rv
}


// Returns a hash table with given pointer functions options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/init(options:)
func (hc _HashTableClass) HashTableWithOptions(options PointerFunctionsOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("hashTableWithOptions:"), options)
	return rv
}


// Returns a new hash table for storing weak references to its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/weakObjects()
func (hc _HashTableClass) WeakObjectsHashTable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("weakObjectsHashTable"))
	return rv
}


// Adds a given object to the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/add(_:)
func (h_ HashTable) AddObject(object unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addObject:"), object)
}


// Returns a Boolean value that indicates whether the hash table contains a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/contains(_:)
func (h_ HashTable) ContainsObject(anObject unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("containsObject:"), anObject)
	return rv
}


// Removes from the receiving hash table each element that isn’t a member of another given hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/intersect(_:)
func (h_ HashTable) IntersectHashTable(other unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("intersectHashTable:"), other)
}


// Returns a Boolean value that indicates whether a given hash table intersects with the receiving hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/intersects(_:)
func (h_ HashTable) IntersectsHashTable(other unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("intersectsHashTable:"), other)
	return rv
}


// Returns a Boolean value that indicates whether a given hash table is equal to the receiving hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/isEqual(to:)
func (h_ HashTable) IsEqualToHashTable(other unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isEqualToHashTable:"), other)
	return rv
}


// Returns a Boolean value that indicates whether every element in the receiving hash table is also present in another given hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/isSubset(of:)
func (h_ HashTable) IsSubsetOfHashTable(other unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isSubsetOfHashTable:"), other)
	return rv
}


// Determines whether the hash table contains a given object, and returns that object if it is present
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/member(_:)
func (h_ HashTable) Member(object unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("member:"), object)
	return rv
}


// Removes each element in another given hash table from the receiving hash table, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/minus(_:)
func (h_ HashTable) MinusHashTable(other unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("minusHashTable:"), other)
}


// Returns an enumerator object that lets you access each object in the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/objectEnumerator()
func (h_ HashTable) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("objectEnumerator"))
	return rv
}


// Removes a given object from the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/remove(_:)
func (h_ HashTable) RemoveObject(object unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeObject:"), object)
}


// Removes all objects from the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/removeAllObjects()
func (h_ HashTable) RemoveAllObjects() {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeAllObjects"))
}


// Adds each element in another given hash table to the receiving hash table, if not present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/union(_:)
func (h_ HashTable) UnionHashTable(other unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("unionHashTable:"), other)
}


// The hash table’s members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/allObjects
func (h_ HashTable) AllObjects() []objc.ID {
	rv := objc.Send[[]objc.ID](h_.ID, objc.Sel("allObjects"))
	return rv
}


// One of the objects in the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/anyObject
func (h_ HashTable) AnyObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("anyObject"))
	return rv
}


// The number of elements in the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/count
func (h_ HashTable) Count() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("count"))
	return rv
}


// The pointer functions for the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/pointerFunctions
func (h_ HashTable) PointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](h_.ID, objc.Sel("pointerFunctions"))
	return rv
}


// A set that contains the hash table’s members.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/setRepresentation
func (h_ HashTable) SetRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("setRepresentation"))
	return rv
}


