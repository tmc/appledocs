// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HashTable] class.
var hashTableClass = _HashTableClass{objc.GetClass("NSHashTable")}

type _HashTableClass struct {
	class objc.Class
}

// An interface definition for the [HashTable] class.
type IHashTable interface {
	objectivec.IObject
}

// A collection similar to a set, but with broader range of available memory semantics. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return hashTableClass.New()
}




