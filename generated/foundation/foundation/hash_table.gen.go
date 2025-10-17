// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HashTable] class.
var HashTableClass objc.Class

func init() {
	HashTableClass = objc.GetClass("NSHashTable")
}

type HashTable struct {
	objc.ID
}

func HashTableFrom(ptr unsafe.Pointer) HashTable {
	return HashTable{
		ID: objc.ID(ptr),
	}
}




