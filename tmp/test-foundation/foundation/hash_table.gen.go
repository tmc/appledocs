// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var HashTableClass _HashTableClass

func init() {
	HashTableClass = _HashTableClass{objc.GetClass("NSHashTable")}
}

type _HashTableClass struct {
	class objc.Class
}

type HashTable struct {
	objc.ID
}

func HashTableFrom(ptr unsafe.Pointer) HashTable {
	return HashTable{
		ID: objc.ID(ptr),
	}
}




