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



