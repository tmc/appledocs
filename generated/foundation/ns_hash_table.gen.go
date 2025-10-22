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
	AddObject(object unsafe.Pointer)
	Count() uint
	AllObjects() unsafe.Pointer
	SetAllObjects(value unsafe.Pointer)
	AnyObject() unsafe.Pointer
	SetAnyObject(value unsafe.Pointer)
	PointerFunctions() NSPointerFunctions
	SetPointerFunctions(value IPointerFunctions)
	SetRepresentation() unsafe.Pointer
	SetSetRepresentation(value unsafe.Pointer)
}

// A collection similar to a set, but with broader range of available memory semantics.
//
// The hash table is modeled after with the following differences: It can hold weak references to its members. Its members may be copied on input or may use pointer identity for equality and hashing. It can contain arbitrary pointers (its members are not constrained to being objects). You can configure an instance to operate on arbitrary pointers and not just objects, although typically you are encouraged to use the C function API for void * pointers. The object-based API (such as ) will not work for non-object pointers without type-casting. Because of its options, is not a set because it can behave differently (for example, if pointer equality is specified two strings will both be entered). When configuring hash tables, note that only the options listed in guarantee that the rest of the API will work correctly—including copying, archiving, and fast enumeration. While other options are used for certain configurations, such as to hold arbitrary pointers, not all combinations of the options are valid. With some combinations the hash table may not work correctly, or may not even be initialized correctly.
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

// Adds a given object to the hash table.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/add(_:)
func (h_ HashTable) AddObject(object unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("addObject:"), object)
}

// The number of elements in the hash table.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTable/count
func (h_ HashTable) Count() uint {
	rv := objc.Send[uint](h_.ID, objc.Sel("count"))
	return rv
}

// The hash table’s members.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/allobjects
func (h_ HashTable) AllObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("allObjects"))
	return rv
}


// SetAllObjects sets the value of the allObjects property.
// The hash table’s members.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/allobjects
func (h_ HashTable) SetAllObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAllObjects:"), value)
}

// One of the objects in the hash table.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/anyobject
func (h_ HashTable) AnyObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("anyObject"))
	return rv
}


// SetAnyObject sets the value of the anyObject property.
// One of the objects in the hash table.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/anyobject
func (h_ HashTable) SetAnyObject(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setAnyObject:"), value)
}

// The pointer functions for the hash table.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (h_ HashTable) PointerFunctions() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](h_.ID, objc.Sel("pointerFunctions"))
	return rv
}


// SetPointerFunctions sets the value of the pointerFunctions property.
// The pointer functions for the hash table.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (h_ HashTable) SetPointerFunctions(value IPointerFunctions) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setPointerFunctions:"), value)
}

// A set that contains the hash table’s members.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/setrepresentation
func (h_ HashTable) SetRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("setRepresentation"))
	return rv
}


// SetSetRepresentation sets the value of the setRepresentation property.
// A set that contains the hash table’s members.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/setrepresentation
func (h_ HashTable) SetSetRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setSetRepresentation:"), value)
}


